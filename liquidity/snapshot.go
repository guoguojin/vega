package liquidity

import (
	"context"
	"sort"
	"strconv"

	typespb "code.vegaprotocol.io/protos/vega"
	snapshotpb "code.vegaprotocol.io/protos/vega/snapshot/v1"
	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/libs/crypto"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"

	"github.com/golang/protobuf/proto"
)

type SnapshotEngine struct {
	*Engine
	pl     types.Payload
	market string

	// liquidity types
	parametersChanged bool
	hashes            map[string][]byte
	serialised        map[string][]byte
	serialisers       map[string]*proto.Buffer
	serialisersFuncs  map[string]func() ([]byte, bool, error)

	// keys, need to be computed when the engine is
	// instantiated as they are dynamic
	hashKeys                  []string
	parametersKey             string
	partiesLiquidityOrdersKey string
	partiesOrdersKey          string
	pendingProvisionsKey      string
	provisionsKey             string
}

func NewSnapshotEngine(config Config,
	log *logging.Logger,
	broker Broker,
	idGen IDGen,
	riskModel RiskModel,
	priceMonitor PriceMonitor,
	market string,
) *SnapshotEngine {
	se := &SnapshotEngine{
		Engine: NewEngine(config, log, broker, idGen, riskModel, priceMonitor, market),
		pl:     types.Payload{},
		market: market,

		parametersChanged: true,
		// empty so default to nil to force update
		hashes:           map[string][]byte{},
		serialised:       map[string][]byte{},
		serialisers:      map[string]*proto.Buffer{},
		serialisersFuncs: map[string]func() ([]byte, bool, error){},
	}

	// build the keys
	se.buildHashKeys(market)

	// map the serialisations functions
	se.serialisersFuncs[se.parametersKey] = se.serialiseParameters
	se.serialisersFuncs[se.partiesLiquidityOrdersKey] = se.serialisePartiesLiquidityOrders
	se.serialisersFuncs[se.partiesOrdersKey] = se.serialisePartiesOrders
	se.serialisersFuncs[se.pendingProvisionsKey] = se.serialisePendingProvisions
	se.serialisersFuncs[se.provisionsKey] = se.serialiseProvisions

	// inialised some stuff
	for _, v := range se.hashKeys {
		se.serialisers[v] = proto.NewBuffer(nil)
		se.serialisers[v].SetDeterministic(true)
	}

	return se
}

func (e *SnapshotEngine) Namespace() types.SnapshotNamespace {
	return types.LiquiditySnapshot
}

func (e *SnapshotEngine) Keys() []string {
	return e.hashKeys
}

func (e *SnapshotEngine) OnSuppliedStakeToObligationFactorUpdate(v num.Decimal) {
	e.parametersChanged = true
	e.Engine.OnSuppliedStakeToObligationFactorUpdate(v)
}

func (e *SnapshotEngine) OnMaximumLiquidityFeeFactorLevelUpdate(f num.Decimal) {
	e.parametersChanged = true
	e.Engine.OnMaximumLiquidityFeeFactorLevelUpdate(f)
}

func (e *SnapshotEngine) OnMarketLiquidityProvisionShapesMaxSizeUpdate(v int64) error {
	e.parametersChanged = true
	return e.Engine.OnMarketLiquidityProvisionShapesMaxSizeUpdate(v)
}

func (e *SnapshotEngine) GetHash(k string) ([]byte, error) {
	_, hash, err := e.serialise(k)
	return hash, err
}

func (e *SnapshotEngine) GetState(k string) ([]byte, []types.StateProvider, error) {
	state, _, err := e.serialise(k)
	return state, nil, err
}

func (e *SnapshotEngine) LoadState(ctx context.Context, p *types.Payload) ([]types.StateProvider, error) {
	if e.Namespace() != p.Data.Namespace() {
		return nil, types.ErrInvalidSnapshotNamespace
	}
	// see what we're reloading
	switch pl := p.Data.(type) {
	case *types.PayloadLiquidityPendingProvisions:
		return nil, e.loadPendingProvisions(
			ctx, pl.PendingProvisions.GetPendingProvisions())
	case *types.PayloadLiquidityProvisions:
		return nil, e.loadProvisions(
			ctx, pl.Provisions.GetLiquidityProvisions())
	case *types.PayloadLiquidityParameters:
		return nil, e.loadParameters(ctx, pl.Parameters)
	case *types.PayloadLiquidityPartiesOrders:
		return nil, e.loadPartiesOrders(ctx, pl.PartiesOrders.GetOrders())
	case *types.PayloadLiquidityPartiesLiquidityOrders:
		return nil, e.loadPartiesLiquidityOrders(
			ctx, pl.PartiesLiquidityOrders.GetOrders())
	default:
		return nil, types.ErrUnknownSnapshotType
	}
}

func (e *SnapshotEngine) loadPartiesOrders(
	_ context.Context, orders []*typespb.Order,
) error {
	e.Engine.orders = NewPartiesOrders()
	for _, v := range orders {
		order, err := types.OrderFromProto(v)
		if err != nil {
			return err
		}
		e.Engine.orders.Add(order.Party, order)
	}

	return nil
}

func (e *SnapshotEngine) loadPartiesLiquidityOrders(
	_ context.Context, orders []*typespb.Order,
) error {
	e.Engine.liquidityOrders = NewPartiesOrders()
	for _, v := range orders {
		order, err := types.OrderFromProto(v)
		if err != nil {
			return err
		}
		e.Engine.liquidityOrders.Add(order.Party, order)
	}

	return nil
}

func (e *SnapshotEngine) loadParameters(
	_ context.Context, parameters *snapshotpb.LiquidityParameters,
) error {
	maxShapesSize, err := strconv.ParseInt(parameters.MaxShapeSize, 10, 64)
	if err != nil {
		return err
	}
	if err := e.OnMarketLiquidityProvisionShapesMaxSizeUpdate(maxShapesSize); err != nil {
		return err
	}

	maxFee, err := num.DecimalFromString(parameters.MaxFee)
	if err != nil {
		return err
	}
	e.OnMaximumLiquidityFeeFactorLevelUpdate(maxFee)

	stof, err := num.DecimalFromString(parameters.StakeToObligationFactor)
	if err != nil {
		return err
	}
	e.OnSuppliedStakeToObligationFactorUpdate(stof)

	return nil
}

func (e *SnapshotEngine) loadPendingProvisions(
	_ context.Context, pendingProvisions []string,
) error {
	e.Engine.pendings = NewSnapshotablePendingProvisions()
	for _, v := range pendingProvisions {
		e.Engine.pendings.Add(v)
	}

	return nil
}

func (e *SnapshotEngine) loadProvisions(
	ctx context.Context, provisions []*typespb.LiquidityProvision,
) error {
	e.Engine.provisions = NewSnapshotableProvisionsPerParty()
	evts := make([]events.Event, 0, len(provisions))
	for _, v := range provisions {
		provision, err := types.LiquidityProvisionFromProto(v)
		if err != nil {
			return err
		}
		e.Engine.provisions.Set(v.PartyId, provision)
		evts = append(evts, events.NewLiquidityProvisionEvent(ctx, provision))
	}

	e.broker.SendBatch(evts)

	return nil
}

func (e *SnapshotEngine) serialise(k string) ([]byte, []byte, error) {
	f, ok := e.serialisersFuncs[k]
	if !ok {
		return nil, nil, types.ErrSnapshotKeyDoesNotExist
	}

	buf, changed, err := f()
	if err != nil {
		return nil, nil, err
	}

	if !changed {
		return e.serialised[k], e.hashes[k], nil
	}

	e.serialised[k] = buf
	h := crypto.Hash(buf)
	e.hashes[k] = h

	return buf, h, nil
}

func (e *SnapshotEngine) serialiseParameters() ([]byte, bool, error) {
	var key = e.parametersKey
	if !e.parametersChanged {
		return e.serialised[key], false, nil
	}

	// reset the flag
	e.parametersChanged = false

	payload := &snapshotpb.Payload{
		Data: &snapshotpb.Payload_LiquidityParameters{
			LiquidityParameters: &snapshotpb.LiquidityParameters{
				MaxFee:                  e.Engine.maxFee.String(),
				MaxShapeSize:            strconv.Itoa(int(e.Engine.maxShapesSize)),
				StakeToObligationFactor: e.Engine.stakeToObligationFactor.String(),
				MarketId:                e.market,
			},
		},
	}

	buf := e.serialisers[key]
	buf.Reset()
	err := buf.Marshal(payload)
	if err != nil {
		return nil, false, err
	}

	return buf.Bytes(), true, nil
}

func (e *SnapshotEngine) serialisePartiesLiquidityOrders() ([]byte, bool, error) {
	var key = e.partiesLiquidityOrdersKey
	if !e.Engine.liquidityOrders.HasUpdates() {
		return e.serialised[key], false, nil
	}

	e.Engine.liquidityOrders.ResetUpdated()

	pborders := []*typespb.Order{}
	for _, orders := range e.Engine.liquidityOrders.m {
		for _, order := range orders {
			pborders = append(pborders, order.IntoProto())
		}
	}
	sort.SliceStable(pborders, func(i, j int) bool { return pborders[i].Id < pborders[j].Id })

	payload := &snapshotpb.Payload{
		Data: &snapshotpb.Payload_LiquidityPartiesLiquidityOrders{
			LiquidityPartiesLiquidityOrders: &snapshotpb.LiquidityPartiesLiquidityOrders{
				MarketId: e.market,
				Orders:   pborders,
			},
		},
	}

	buf := e.serialisers[key]
	buf.Reset()
	err := buf.Marshal(payload)
	if err != nil {
		return nil, false, err
	}

	return buf.Bytes(), true, nil
}

func (e *SnapshotEngine) serialisePartiesOrders() ([]byte, bool, error) {
	var key = e.partiesOrdersKey
	if !e.Engine.orders.HasUpdates() {
		return e.serialised[key], false, nil
	}

	e.Engine.orders.ResetUpdated()

	pborders := []*typespb.Order{}
	for _, orders := range e.Engine.orders.m {
		for _, order := range orders {
			pborders = append(pborders, order.IntoProto())
		}
	}
	sort.SliceStable(pborders, func(i, j int) bool { return pborders[i].Id < pborders[j].Id })

	payload := &snapshotpb.Payload{
		Data: &snapshotpb.Payload_LiquidityPartiesOrders{
			LiquidityPartiesOrders: &snapshotpb.LiquidityPartiesOrders{
				MarketId: e.market,
				Orders:   pborders,
			},
		},
	}

	buf := e.serialisers[key]
	buf.Reset()
	err := buf.Marshal(payload)
	if err != nil {
		return nil, false, err
	}

	return buf.Bytes(), true, nil
}

func (e *SnapshotEngine) serialisePendingProvisions() ([]byte, bool, error) {
	var key = e.pendingProvisionsKey
	if !e.Engine.pendings.HasUpdates() {
		return e.serialised[key], false, nil
	}

	e.Engine.pendings.ResetUpdated()

	pbpendings := make([]string, 0, len(e.Engine.pendings.m))
	for k := range e.Engine.pendings.m {
		pbpendings = append(pbpendings, k)
	}
	sort.Strings(pbpendings)

	payload := &snapshotpb.Payload{
		Data: &snapshotpb.Payload_LiquidityPendingProvisions{
			LiquidityPendingProvisions: &snapshotpb.LiquidityPendingProvisions{
				MarketId:          e.market,
				PendingProvisions: pbpendings,
			},
		},
	}

	buf := e.serialisers[key]
	buf.Reset()
	err := buf.Marshal(payload)
	if err != nil {
		return nil, false, err
	}

	return buf.Bytes(), true, nil
}

func (e *SnapshotEngine) serialiseProvisions() ([]byte, bool, error) {
	var key = e.provisionsKey
	if !e.Engine.provisions.HasUpdates() {
		return e.serialised[key], false, nil
	}

	e.Engine.provisions.ResetUpdated()

	// these are sorted already, only a convertion to proto is needed
	lps := e.Engine.provisions.Slice()
	pblps := make([]*typespb.LiquidityProvision, 0, len(lps))
	for _, v := range lps {
		pblps = append(pblps, v.IntoProto())
	}

	payload := &snapshotpb.Payload{
		Data: &snapshotpb.Payload_LiquidityProvisions{
			LiquidityProvisions: &snapshotpb.LiquidityProvisions{
				MarketId:            e.market,
				LiquidityProvisions: pblps,
			},
		},
	}

	buf := e.serialisers[key]
	buf.Reset()
	err := buf.Marshal(payload)
	if err != nil {
		return nil, false, err
	}

	return buf.Bytes(), true, nil
}

func (e *SnapshotEngine) buildHashKeys(market string) {
	e.parametersKey = (&types.PayloadLiquidityParameters{Parameters: &snapshotpb.LiquidityParameters{MarketId: market}}).Key()
	e.partiesLiquidityOrdersKey = (&types.PayloadLiquidityPartiesLiquidityOrders{}).Key()
	e.partiesOrdersKey = (&types.PayloadLiquidityPartiesOrders{}).Key()
	e.pendingProvisionsKey = (&types.PayloadLiquidityPendingProvisions{}).Key()
	e.provisionsKey = (&types.PayloadLiquidityProvisions{}).Key()

	e.hashKeys = append([]string{}, e.parametersKey, e.partiesLiquidityOrdersKey,
		e.partiesOrdersKey, e.pendingProvisionsKey, e.provisionsKey)
}
