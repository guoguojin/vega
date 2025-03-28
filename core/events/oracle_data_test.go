// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.VEGA file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package events_test

import (
	"context"
	"testing"

	"code.vegaprotocol.io/vega/core/events"
	"code.vegaprotocol.io/vega/core/types"
	vegapb "code.vegaprotocol.io/vega/protos/vega"
	datapb "code.vegaprotocol.io/vega/protos/vega/data/v1"
	"github.com/stretchr/testify/assert"
)

func TestOracleDataDeepClone(t *testing.T) {
	ctx := context.Background()
	pubKeys := []*types.Signer{
		types.CreateSignerFromString("PK1", types.DataSignerTypePubKey),
		types.CreateSignerFromString("PK2", types.DataSignerTypePubKey),
		types.CreateSignerFromString("PK3", types.DataSignerTypePubKey),
	}

	od := datapb.ExternalData{
		Data: &datapb.Data{
			Signers: types.SignersIntoProto(pubKeys),
			Data: []*datapb.Property{
				{
					Name:  "Name",
					Value: "Value",
				},
			},
			MatchedSpecIds: []string{
				"MS1", "MS2",
			},
			BroadcastAt: 10000,
		},
	}

	odEvent := events.NewOracleDataEvent(ctx, vegapb.OracleData{ExternalData: &od})

	od2 := odEvent.OracleData()

	// Change the original values
	pk1 := types.CreateSignerFromString("Changed1", types.DataSignerTypePubKey)
	pk2 := types.CreateSignerFromString("Changed2", types.DataSignerTypePubKey)
	pk3 := types.CreateSignerFromString("Changed3", types.DataSignerTypePubKey)

	od.Data.Signers[0] = pk1.IntoProto()
	od.Data.Signers[1] = pk2.IntoProto()
	od.Data.Signers[2] = pk3.IntoProto()
	od.Data.Data[0].Name = "Changed"
	od.Data.Data[0].Value = "Changed"
	od.Data.MatchedSpecIds[0] = "Changed1"
	od.Data.MatchedSpecIds[1] = "Changed2"
	od.Data.BroadcastAt = 999

	// Check things have changed
	assert.NotEqual(t, od.Data.Signers[0], od2.ExternalData.Data.Signers[0])

	assert.NotEqual(t, od.Data.Signers[1], od2.ExternalData.Data.Signers[1])
	assert.NotEqual(t, od.Data.Signers[2], od2.ExternalData.Data.Signers[2])
	assert.NotEqual(t, od.Data.Data[0].Name, od2.ExternalData.Data.Data[0].Name)
	assert.NotEqual(t, od.Data.Data[0].Value, od2.ExternalData.Data.Data[0].Value)
	assert.NotEqual(t, od.Data.MatchedSpecIds[0], od2.ExternalData.Data.MatchedSpecIds[0])
	assert.NotEqual(t, od.Data.MatchedSpecIds[1], od2.ExternalData.Data.MatchedSpecIds[1])
	assert.NotEqual(t, od.Data.BroadcastAt, od2.ExternalData.Data.BroadcastAt)
}
