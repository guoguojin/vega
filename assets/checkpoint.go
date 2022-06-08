package assets

import (
	"bytes"
	"context"
	"sort"

	checkpoint "code.vegaprotocol.io/protos/vega/checkpoint/v1"
	vgcrypto "code.vegaprotocol.io/vega/libs/crypto"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/types"

	"code.vegaprotocol.io/vega/libs/proto"
)

func (*Service) Name() types.CheckpointName {
	return types.AssetsCheckpoint
}

func (s *Service) Checkpoint() ([]byte, error) {
	t := &checkpoint.Assets{
		Assets: s.getEnabled(),
	}
	return proto.Marshal(t)
}

func (s *Service) Load(_ context.Context, cp []byte) error {
	data := &checkpoint.Assets{}
	if err := proto.Unmarshal(cp, data); err != nil {
		return err
	}
	s.amu.Lock()
	s.pamu.Lock()
	s.pendingAssets = map[string]*Asset{}
	s.pamu.Unlock()
	s.amu.Unlock()
	for _, a := range data.Assets {
		details, _ := types.AssetDetailsFromProto(a.AssetDetails)
		// first check if the asset did get loaded from genesis
		// if yes and the details + IDs are same, then all good
		if existing, ok := s.assets[a.Id]; ok {
			// we know this ID, are the details the same
			// if not, then  there's an error, and we should not overwrite an existing
			// asset, only new ones can be added
			if !bytes.Equal(vgcrypto.Hash([]byte(details.String())), vgcrypto.Hash([]byte(existing.Type().Details.String()))) {
				s.log.Panic("invalid asset loaded from genesis",
					logging.String("id", a.Id),
					logging.String("details-genesis", existing.String()),
					logging.String("details-checkpoint", details.String()))
			}
			continue
		}

		// asset didn't match anything, we need to go through the process to add it.

		id, err := s.NewAsset(a.Id, details)
		if err != nil {
			return err
		}
		pa, _ := s.Get(a.Id)
		if s.isValidator {
			if err := pa.Validate(); err != nil {
				return err
			}
		} else {
			pa.SetValidNonValidator()
		}
		if err := s.Enable(id); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) getEnabled() []*checkpoint.AssetEntry {
	s.amu.RLock()
	keys := make([]string, 0, len(s.assets))
	vals := make(map[string]*checkpoint.AssetEntry, len(s.assets))
	for k, a := range s.assets {
		keys = append(keys, k)
		vals[k] = &checkpoint.AssetEntry{
			Id:           k,
			AssetDetails: a.Type().Details.IntoProto(),
		}
	}
	s.amu.RUnlock()
	if len(keys) == 0 {
		return nil
	}
	ret := make([]*checkpoint.AssetEntry, 0, len(vals))
	sort.Strings(keys)
	for _, k := range keys {
		ret = append(ret, vals[k])
	}
	return ret
}
