package liquidity_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"testing"
	"time"

	snapshotpb "code.vegaprotocol.io/protos/vega/snapshot/v1"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var initialTime = time.Date(2020, 10, 20, 1, 1, 1, 0, time.UTC)

func TestSnapshotRoundTrip(t *testing.T) {
	var (
		party1 = "p1"
		party2 = "p2"
		party3 = "p3"
		market = "market-id"
		ctx    = context.Background()
		e1     = newTestEngine(t, initialTime)
		e2     = newTestEngine(t, initialTime)
		e3     = newTestEngine(t, initialTime)
	)

	e1.broker.EXPECT().Send(gomock.Any()).AnyTimes()
	e2.broker.EXPECT().Send(gomock.Any()).AnyTimes()
	e2.broker.EXPECT().SendBatch(gomock.Any()).AnyTimes()
	e3.broker.EXPECT().Send(gomock.Any()).AnyTimes()
	e3.broker.EXPECT().SendBatch(gomock.Any()).AnyTimes()

	lp1 := &types.LiquidityProvisionSubmission{
		MarketID:         market,
		Fee:              num.MustDecimalFromString("0.01"),
		CommitmentAmount: num.NewUint(1000),
		Buys: []*types.LiquidityOrder{
			{
				Reference:  types.PeggedReferenceMid,
				Offset:     -1,
				Proportion: 1,
			},
		},
		Sells: []*types.LiquidityOrder{
			{
				Reference:  types.PeggedReferenceMid,
				Offset:     1,
				Proportion: 1,
			},
		},
	}
	lp2 := &types.LiquidityProvisionSubmission{
		MarketID:         market,
		Fee:              num.MustDecimalFromString("0.42"),
		CommitmentAmount: num.NewUint(4242),
		Buys: []*types.LiquidityOrder{
			{
				Reference:  types.PeggedReferenceMid,
				Offset:     -10,
				Proportion: 42,
			},
		},
		Sells: []*types.LiquidityOrder{
			{
				Reference:  types.PeggedReferenceMid,
				Offset:     42,
				Proportion: 58,
			},
		},
	}

	require.NoError(t,
		e1.engine.SubmitLiquidityProvision(ctx, lp1, party1, "some-id-1"),
	)
	require.NoError(t,
		e1.engine.SubmitLiquidityProvision(ctx, lp2, party2, "some-id-2"),
	)

	keys := e1.engine.Keys()
	kToH := map[string][]byte{}
	kToS := map[string][]byte{}

	expectedHashes := map[string]string{
		"parameters:market-id":             "d663375fd6843a0807d17b10ad8425a6ba45c8c2dd6339f400c5b2426f900c13",
		"partiesLiquidityOrders:market-id": "0254d8b74441ca3bac8f9b141408502d9b1f297e8ef1054d45775566677a8072",
		"partiesOrders:market-id":          "f9cb31b1c4c8df91f6a348d43978c302c8887336107c265259bc74fdddf00e19",
		"pendingProvisions:market-id":      "6cc4d407a2ea45e37e27993eb6f94134b3f906d080777d94bf99551aa82dc461",
		"provisions:market-id":             "236cf44c1a419b99955aa7662dacfed9bba4d5f391e4606bb00f50676f99b40a",
	}

	for _, key := range keys {
		h, err := e1.engine.GetHash(key)
		assert.NoError(t, err)
		kToH[key] = h
		s, _, err := e1.engine.GetState(key)
		assert.NoError(t, err)
		kToS[key] = s

		// compare hashes to the expected ones
		assert.Equal(t, hex.EncodeToString(h), expectedHashes[key])
	}

	// now we reload the keys / state
	for _, s := range kToS {
		pl := snapshotpb.Payload{}
		assert.NoError(t, proto.Unmarshal(s, &pl))
		_, err := e2.engine.LoadState(ctx, types.PayloadFromProto(&pl))
		assert.NoError(t, err)
	}

	// now ensure both are producing same hashes
	for k, e1h := range kToH {
		e2h, err := e2.engine.GetHash(k)
		assert.NoError(t, err)
		assert.True(t, bytes.Equal(e1h, e2h))
	}

	// now we update the state of e2 to see if hashes changes

	expectedHashes2 := map[string]string{
		"parameters:market-id":             "b5eec91c297baf1f06830350dbcb37d79937561ae605d2304eb12680e443775c",
		"partiesLiquidityOrders:market-id": "4f70bcae9080f9c4879db76c814cb1d1e538501a565a459114bbf7f6db8ccd9b",
		"partiesOrders:market-id":          "f9cb31b1c4c8df91f6a348d43978c302c8887336107c265259bc74fdddf00e19",
		"pendingProvisions:market-id":      "627ef55af7f36bea0d09b0081b85d66531a01df060d8e9447e17049a4e152b12",
		"provisions:market-id":             "89bcde9a9715401764cf27403e3585126c9321a477e54c28d84b74f95fb78189",
	}

	lp3 := &types.LiquidityProvisionSubmission{
		MarketID:         market,
		Fee:              num.MustDecimalFromString("0.2"),
		CommitmentAmount: num.NewUint(5000),
		Buys: []*types.LiquidityOrder{
			{
				Reference:  types.PeggedReferenceMid,
				Offset:     -10,
				Proportion: 42,
			},
		},
		Sells: []*types.LiquidityOrder{
			{
				Reference:  types.PeggedReferenceMid,
				Offset:     42,
				Proportion: 58,
			},
		},
	}

	require.NoError(t,
		e2.engine.SubmitLiquidityProvision(ctx, lp3, party3, "some-id-2"),
	)

	require.NoError(t,
		e2.engine.OnMarketLiquidityProvisionShapesMaxSizeUpdate(4200),
	)

	repriceFN := func(
		order *types.PeggedOrder, side types.Side,
	) (*num.Uint, *types.PeggedOrder, error) {
		return num.NewUint(100), order, nil
	}

	e2.priceMonitor.EXPECT().GetValidPriceRange().
		Return(num.NewWrappedDecimal(num.Zero(), num.DecimalZero()), num.NewWrappedDecimal(num.NewUint(90), num.DecimalFromInt64(110))).
		AnyTimes()

	_, _, err := e2.engine.Update(ctx, num.NewUint(99), num.NewUint(101),
		repriceFN, []*types.Order{
			{
				ID:        "order-id-1",
				Party:     party1,
				MarketID:  market,
				Side:      types.SideBuy,
				Price:     num.NewUint(90),
				Size:      10,
				Remaining: 10,
			},
		},
	)

	require.NoError(t, err)

	for _, key := range keys {
		h, err := e2.engine.GetHash(key)
		assert.NoError(t, err)

		s, _, err := e2.engine.GetState(key)
		assert.NoError(t, err)

		// compare hashes to the expected ones
		assert.Equal(t, hex.EncodeToString(h), expectedHashes2[key], key)

		pl := snapshotpb.Payload{}
		assert.NoError(t, proto.Unmarshal(s, &pl))
		_, err = e3.engine.LoadState(ctx, types.PayloadFromProto(&pl))
		assert.NoError(t, err)
	}

	for _, key := range keys {
		h, err := e3.engine.GetHash(key)
		assert.NoError(t, err)
		// compare hashes to the expected ones
		assert.Equal(t, hex.EncodeToString(h), expectedHashes2[key], key)
	}
}
