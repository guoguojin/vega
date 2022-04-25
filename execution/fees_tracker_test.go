package execution_test

import (
	"bytes"
	"context"
	"testing"

	vgproto "code.vegaprotocol.io/protos/vega"
	snapshotpb "code.vegaprotocol.io/protos/vega/snapshot/v1"
	"code.vegaprotocol.io/vega/execution"
	"code.vegaprotocol.io/vega/libs/proto"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"
	"github.com/stretchr/testify/require"
)

type TestEpochEngine struct {
	target func(context.Context, types.Epoch)
}

func (e *TestEpochEngine) NotifyOnEpoch(f func(context.Context, types.Epoch), _ func(context.Context, types.Epoch)) {
	e.target = f
}

func TestFeesTracker(t *testing.T) {
	epochEngine := &TestEpochEngine{}
	feesTracker := execution.NewFeesTracker(epochEngine)
	epochEngine.target(context.Background(), types.Epoch{Seq: 1})

	partyScores := feesTracker.GetFeePartyScores("does not exist", types.TransferTypeMakerFeeReceive)
	require.Equal(t, 0, len(partyScores))

	key := (&types.PayloadFeeTracker{}).Key()
	hash1, err := feesTracker.GetHash(key)
	require.NoError(t, err)

	// update with a few transfers
	transfersM1 := []*types.Transfer{
		{Owner: "party1", Type: types.TransferTypeMakerFeeReceive, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(100)}},
		{Owner: "party1", Type: types.TransferTypeMakerFeePay, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(200)}},
		{Owner: "party1", Type: types.TransferTypeLiquidityFeeDistribute, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(200)}},
		{Owner: "party1", Type: types.TransferTypeMakerFeeReceive, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(400)}},
		{Owner: "party1", Type: types.TransferTypeMakerFeePay, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(300)}},
		{Owner: "party1", Type: types.TransferTypeLiquidityFeeDistribute, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(600)}},
		{Owner: "party2", Type: types.TransferTypeMakerFeeReceive, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(900)}},
		{Owner: "party2", Type: types.TransferTypeMakerFeePay, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(800)}},
		{Owner: "party2", Type: types.TransferTypeLiquidityFeeDistribute, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(700)}},
		{Owner: "party2", Type: types.TransferTypeMakerFeeReceive, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(600)}},
		{Owner: "party2", Type: types.TransferTypeMakerFeePay, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(200)}},
		{Owner: "party2", Type: types.TransferTypeLiquidityFeeDistribute, Amount: &types.FinancialAmount{Asset: "asset1", Amount: num.NewUint(1000)}},
	}
	feesTracker.UpdateFeesFromTransfers("market1", transfersM1)

	transfersM2 := []*types.Transfer{
		{Owner: "party1", Type: types.TransferTypeMakerFeeReceive, Amount: &types.FinancialAmount{Asset: "asset2", Amount: num.NewUint(150)}},
		{Owner: "party2", Type: types.TransferTypeMakerFeePay, Amount: &types.FinancialAmount{Asset: "asset2", Amount: num.NewUint(150)}},
	}
	feesTracker.UpdateFeesFromTransfers("market2", transfersM2)

	// asset1, types.TransferTypeMakerFeeReceive
	// party1 received 500
	// party2 received 1500
	scores := feesTracker.GetFeePartyScores("market1", types.TransferTypeMakerFeeReceive)
	require.Equal(t, "0.25", scores[0].Score.String())
	require.Equal(t, "party1", scores[0].Party)
	require.Equal(t, "0.75", scores[1].Score.String())
	require.Equal(t, "party2", scores[1].Party)

	// asset1 TransferTypeMakerFeePay
	// party1 paid 500
	// party2 paid 1000
	scores = feesTracker.GetFeePartyScores("market1", types.TransferTypeMakerFeePay)
	require.Equal(t, "0.3333333333333333", scores[0].Score.String())
	require.Equal(t, "party1", scores[0].Party)
	require.Equal(t, "0.6666666666666667", scores[1].Score.String())
	require.Equal(t, "party2", scores[1].Party)

	// asset1 TransferTypeLiquidityFeeDistribute
	// party1 paid 800
	// party2 paid 1700
	scores = feesTracker.GetFeePartyScores("market1", types.TransferTypeLiquidityFeeDistribute)
	require.Equal(t, "0.32", scores[0].Score.String())
	require.Equal(t, "party1", scores[0].Party)
	require.Equal(t, "0.68", scores[1].Score.String())
	require.Equal(t, "party2", scores[1].Party)

	// asset2 TransferTypeMakerFeePay
	scores = feesTracker.GetFeePartyScores("market2", types.TransferTypeMakerFeeReceive)
	require.Equal(t, 1, len(scores))
	require.Equal(t, "1", scores[0].Score.String())
	require.Equal(t, "party1", scores[0].Party)

	// asset2 TransferTypeMakerFeePay
	scores = feesTracker.GetFeePartyScores("market2", types.TransferTypeMakerFeePay)
	require.Equal(t, 1, len(scores))
	require.Equal(t, "1", scores[0].Score.String())
	require.Equal(t, "party2", scores[0].Party)

	// check hash has changed
	hash2, err := feesTracker.GetHash(key)
	require.NoError(t, err)
	require.False(t, bytes.Equal(hash1, hash2))

	state2, _, err := feesTracker.GetState(key)
	require.NoError(t, err)

	epochEngineLoad := &TestEpochEngine{}
	feesTrackerLoad := execution.NewFeesTracker(epochEngineLoad)
	epochEngineLoad.target(context.Background(), types.Epoch{Seq: 1})

	pl := snapshotpb.Payload{}
	require.NoError(t, proto.Unmarshal(state2, &pl))
	feesTrackerLoad.LoadState(context.Background(), types.PayloadFromProto(&pl))

	hash3, err := feesTrackerLoad.GetHash(key)
	require.NoError(t, err)
	require.True(t, bytes.Equal(hash2, hash3))

	state3, _, err := feesTrackerLoad.GetState(key)
	require.NoError(t, err)
	require.True(t, bytes.Equal(state2, state3))

	// check a restored party exist in the restored engine
	scores = feesTrackerLoad.GetFeePartyScores("market2", types.TransferTypeMakerFeeReceive)
	require.Equal(t, 1, len(scores))
	require.Equal(t, "1", scores[0].Score.String())
	require.Equal(t, "party1", scores[0].Party)

	// New epoch should scrub the state an produce a difference hash
	epochEngineLoad.target(context.Background(), types.Epoch{Seq: 2, Action: vgproto.EpochAction_EPOCH_ACTION_START})
	hash4, err := feesTrackerLoad.GetHash(key)
	require.NoError(t, err)
	require.False(t, bytes.Equal(hash3, hash4))
}
