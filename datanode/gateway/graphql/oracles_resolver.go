// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.DATANODE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package gql

import (
	"context"

	"code.vegaprotocol.io/vega/libs/ptr"
	v2 "code.vegaprotocol.io/vega/protos/data-node/api/v2"
	vegapb "code.vegaprotocol.io/vega/protos/vega"
	v1 "code.vegaprotocol.io/vega/protos/vega/data/v1"
)

type oracleSpecResolver VegaResolverRoot

func (o *oracleSpecResolver) DataSourceSpec(_ context.Context, obj *vegapb.OracleSpec) (extDss *ExternalDataSourceSpec, _ error) {
	extDss = &ExternalDataSourceSpec{Spec: &DataSourceSpec{Data: &vegapb.DataSourceDefinition{}}}
	if obj.ExternalDataSourceSpec != nil {
		extDss.Spec = resolveDataSourceSpec(obj.ExternalDataSourceSpec.Spec)
	}
	return
}

func (o *oracleSpecResolver) DataConnection(ctx context.Context, obj *vegapb.OracleSpec, pagination *v2.Pagination) (*v2.OracleDataConnection, error) {
	var specID *string
	if ed := obj.ExternalDataSourceSpec; ed != nil && ed.Spec != nil && ed.Spec.Id != "" {
		specID = &ed.Spec.Id
	}

	req := v2.ListOracleDataRequest{
		OracleSpecId: specID,
		Pagination:   pagination,
	}

	resp, err := o.tradingDataClientV2.ListOracleData(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp.OracleData, nil
}

type oracleDataResolver VegaResolverRoot

func (o *oracleDataResolver) ExternalData(_ context.Context, obj *vegapb.OracleData) (ed *ExternalData, _ error) {
	ed = &ExternalData{
		Data: &Data{},
	}

	oed := obj.ExternalData
	if oed == nil || oed.Data == nil {
		return
	}

	ed.Data.Signers = resolveSigners(oed.Data.Signers)
	ed.Data.Data = oed.Data.Data
	ed.Data.MatchedSpecIds = oed.Data.MatchedSpecIds
	ed.Data.BroadcastAt = oed.Data.BroadcastAt

	return
}

func resolveTrigger(obj any) (trigger TriggerKind) {
	switch trig := obj.(type) {
	case *vegapb.EthCallTrigger_TimeTrigger:
		if trig.TimeTrigger != nil {
			init := int64(*trig.TimeTrigger.Initial)
			every := int64(*trig.TimeTrigger.Every)
			until := int64(*trig.TimeTrigger.Until)
			trigger = &EthTimeTrigger{
				Initial: &init,
				Every:   &every,
				Until:   &until,
			}
		}
	}

	return
}

func resolveSigners(obj []*v1.Signer) (signers []*Signer) {
	for i := range obj {
		signers = append(signers, &Signer{Signer: resolveSigner(obj[i].Signer)})
	}
	return
}

func resolveSigner(obj any) (signer SignerKind) {
	switch sig := obj.(type) {
	case *v1.Signer_PubKey:
		signer = &PubKey{Key: &sig.PubKey.Key}
	case *v1.Signer_EthAddress:
		signer = &ETHAddress{Address: &sig.EthAddress.Address}
	}
	return
}

func resolveDataSourceDefinition(d *vegapb.DataSourceDefinition) (ds *vegapb.DataSourceDefinition) {
	ds = &vegapb.DataSourceDefinition{}
	if d == nil || d.SourceType == nil {
		return ds
	}
	data := d.Content()
	if data != nil {
		switch tp := data.(type) {
		case *vegapb.DataSourceSpecConfiguration:
			ds.SourceType = &vegapb.DataSourceDefinition_External{
				External: &vegapb.DataSourceDefinitionExternal{
					SourceType: &vegapb.DataSourceDefinitionExternal_Oracle{
						Oracle: tp,
					},
				},
			}

		case *vegapb.EthCallSpec:
			ds.SourceType = &vegapb.DataSourceDefinition_External{
				External: &vegapb.DataSourceDefinitionExternal{
					SourceType: &vegapb.DataSourceDefinitionExternal_EthOracle{
						EthOracle: tp,
					},
				},
			}

		case *vegapb.DataSourceSpecConfigurationTime:
			ds.SourceType = &vegapb.DataSourceDefinition_Internal{
				Internal: &vegapb.DataSourceDefinitionInternal{
					SourceType: &vegapb.DataSourceDefinitionInternal_Time{
						Time: tp,
					},
				},
			}
		}
	}

	return ds
}

func resolveDataSourceSpec(d *vegapb.DataSourceSpec) (ds *DataSourceSpec) {
	ds = &DataSourceSpec{
		Data: &vegapb.DataSourceDefinition{},
	}
	if d == nil {
		return
	}

	ds.ID = d.GetId()
	ds.CreatedAt = d.CreatedAt
	if d.UpdatedAt != 0 {
		ds.UpdatedAt = ptr.From(d.UpdatedAt)
	}

	switch d.Status {
	case vegapb.DataSourceSpec_STATUS_ACTIVE:
		ds.Status = DataSourceSpecStatusStatusActive
	case vegapb.DataSourceSpec_STATUS_DEACTIVATED:
		ds.Status = DataSourceSpecStatusStatusDeactivated
	}

	if d.Data != nil {
		ds.Data = resolveDataSourceDefinition(d.Data)
	}

	return
}
