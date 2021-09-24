package collateral

import (
	"context"
	"sort"
	"strings"

	snapshot "code.vegaprotocol.io/protos/vega/snapshot/v1"
	"code.vegaprotocol.io/vega/libs/crypto"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type accState struct {
	accs       map[string]*types.Account
	assets     map[string]types.Asset
	accIDs     []string
	assetIDs   []string
	hashes     map[string][]byte
	updates    map[string]bool
	serialised map[string][]byte
}

var (
	hashKeys = []string{
		"account",
		"asset",
	}

	ErrSnapshotKeyDoesNotExist = errors.New("unknown key for collateral snapshot")
)

func (e *Engine) Name() types.CheckpointName {
	return types.CollateralCheckpoint
}

func (e *Engine) Checkpoint() ([]byte, error) {
	msg := &snapshot.Collateral{
		Balances: e.getSnapshotBalances(),
	}
	ret, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (e *Engine) Load(checkpoint []byte) error {
	msg := snapshot.Collateral{}
	if err := proto.Unmarshal(checkpoint, &msg); err != nil {
		return err
	}
	for _, balance := range msg.Balances {
		ub, _ := num.UintFromString(balance.Balance, 10)
		if balance.Party == systemOwner {
			accID := e.accountID(noMarket, systemOwner, balance.Asset, types.AccountTypeGlobalInsurance)
			if _, err := e.GetAccountByID(accID); err != nil {
				// this account is created when the asset is enabled. If we can't get this account,
				// then the asset is not yet enabled and we have a problem...
				return err
			}
			e.UpdateBalance(context.Background(), accID, ub)
			continue
		}
		accID := e.accountID(noMarket, balance.Party, balance.Asset, types.AccountTypeGeneral)
		if _, err := e.GetAccountByID(accID); err != nil {
			accID, _ = e.CreatePartyGeneralAccount(context.Background(), balance.Party, balance.Asset)
		}
		e.UpdateBalance(context.Background(), accID, ub)
	}
	return nil
}

// get all balances for snapshot
func (e *Engine) getSnapshotBalances() []*snapshot.AssetBalance {
	parties := make([]string, 0, len(e.partiesAccs))
	pbal := make(map[string][]*snapshot.AssetBalance, len(e.partiesAccs))
	entries := 0
	for party, accs := range e.partiesAccs {
		assets := make([]string, 0, len(accs))
		balances := map[string]*num.Uint{}
		for _, acc := range accs {
			switch acc.Type {
			case types.AccountTypeMargin, types.AccountTypeGeneral, types.AccountTypeBond,
				types.AccountTypeInsurance, types.AccountTypeGlobalInsurance:
				assetBal, ok := balances[acc.Asset]
				if !ok {
					assetBal = num.Zero()
					balances[acc.Asset] = assetBal
					assets = append(assets, acc.Asset)
				}
				assetBal.AddSum(acc.Balance)
			case types.AccountTypeSettlement:
				if !acc.Balance.IsZero() {
					e.log.Panic("Settlement balance is not zero",
						logging.String("market-id", acc.MarketID))
				}
			}
		}
		ln := len(assets)
		if ln == 0 {
			continue
		}
		entries += ln
		pbal[party] = make([]*snapshot.AssetBalance, 0, len(assets))
		parties = append(parties, party)
		// sort by asset -> each party will have their balances appended in alphabetic order
		sort.Strings(assets)
		for _, a := range assets {
			bal := balances[a]
			pbal[party] = append(pbal[party], &snapshot.AssetBalance{
				Party:   party,
				Asset:   a,
				Balance: bal.String(),
			})
		}
	}
	ret := make([]*snapshot.AssetBalance, 0, entries)
	sort.Strings(parties)
	for _, party := range parties {
		ret = append(ret, pbal[party]...)
	}
	return ret
}

func (e *Engine) Namespace() string {
	return string(types.CollateralSnapshot)
}

func (e *Engine) Keys() []string {
	return hashKeys
}

func (e *Engine) GetHash(k string) ([]byte, error) {
	return e.state.getHash(k)
}

func (e *Engine) GetState(k string) ([]byte, error) {
	return e.state.getState(k)
}

func (e *Engine) Snapshot() (map[string][]byte, error) {
	r := make(map[string][]byte, len(hashKeys))
	for _, k := range hashKeys {
		state, err := e.state.getState(k)
		if err != nil {
			return nil, err
		}
		r[k] = state
	}
	return r, nil
}

func newAccState() *accState {
	return &accState{
		accs:     map[string]*types.Account{},
		assets:   map[string]types.Asset{},
		accIDs:   []string{},
		assetIDs: []string{},
		updates: map[string]bool{
			"account": false,
			"asset":   false,
		},
		hashes: map[string][]byte{
			"account": nil,
			"asset":   nil,
		},
		serialised: map[string][]byte{
			"account": nil,
			"asset":   nil,
		},
	}
}

func (a *accState) enableAsset(asset types.Asset) {
	a.assets[asset.ID] = asset
	a.assetIDs = append(a.assetIDs, asset.ID)
	sort.Strings(a.assetIDs)
	a.updates["asset"] = true
}

func (a *accState) add(accs ...*types.Account) {
	if len(accs) == 0 {
		return
	}
	ids := make([]string, 0, len(accs))
	for _, acc := range accs {
		if _, ok := a.accs[acc.ID]; !ok {
			ids = append(ids, acc.ID)
		}
		a.accs[acc.ID] = acc.Clone()
	}
	if len(ids) > 0 {
		a.accIDs = append(a.accIDs, ids...)
		sort.Strings(a.accIDs)
	}
	a.updates["account"] = true
}

func (a *accState) delAcc(accs ...*types.Account) {
	if len(accs) == 0 {
		return
	}
	updated := false
	for _, acc := range accs {
		if _, ok := a.accs[acc.ID]; ok {
			updated = true
			delete(a.accs, acc.ID)
			// find ID in slice, this should always be present
			i := sort.Search(len(a.accIDs), func(i int) bool {
				return a.accIDs[i] >= acc.ID
			})
			// just make sure we found a match, this should be optional
			if a.accIDs[i] == acc.ID {
				copy(a.accIDs[i:], a.accIDs[i+1:])
			}
		}
	}
	if updated {
		a.updates["account"] = true
	}
}

func (a *accState) hashAssets() error {
	if !a.updates["asset"] {
		return nil
	}
	data := []byte(strings.Join(a.assetIDs, ""))
	// @TODO populate type to persist && serialise, then save it in the serialised field
	a.hashes["asset"] = crypto.Hash(data)
	a.updates["asset"] = false
	return nil
}

func (a *accState) hashAccounts() error {
	if !a.updates["account"] {
		return nil
	}
	data := make([]byte, 0, len(a.accIDs)*32)
	i := 0
	for _, id := range a.accIDs {
		b := a.accs[id].Balance.Bytes()
		copy(data[i:], b[:])
		i += 32
		// @TODO populate type to persist && serialise, then save it in the serialised field
	}
	a.hashes["account"] = crypto.Hash(data)
	a.updates["account"] = false
	return nil
}

func (a *accState) getState(k string) ([]byte, error) {
	update, exist := a.updates[k]
	if !exist {
		return nil, ErrSnapshotKeyDoesNotExist
	}
	if !update {
		h := a.serialised[k]
		return h, nil
	}
	if k == "asset" {
		if err := a.hashAssets(); err != nil {
			return nil, err
		}
	} else if err := a.hashAccounts(); err != nil {
		return nil, err
	}
	h := a.serialised[k]
	return h, nil
}

func (a *accState) getHash(k string) ([]byte, error) {
	update, exist := a.updates[k]
	if !exist {
		return nil, ErrSnapshotKeyDoesNotExist
	}
	// we have a pending update
	if update {
		// hash whichever one we need to update
		if k == "asset" {
			if err := a.hashAssets(); err != nil {
				return nil, err
			}
		} else if err := a.hashAccounts(); err != nil {
			return nil, err
		}
	}
	// fetch the new hash and return
	h := a.hashes[k]
	return h, nil
}
