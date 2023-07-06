// (c) 2022 Gobalsky Labs Limited
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

package sqlstore_test

import (
	"context"
	"testing"
	"time"

	"code.vegaprotocol.io/vega/datanode/entities"
	"code.vegaprotocol.io/vega/datanode/sqlstore"
	"code.vegaprotocol.io/vega/protos/vega"
	v1 "code.vegaprotocol.io/vega/protos/vega/data/v1"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarkets_Add(t *testing.T) {
	t.Run("Add should insert a valid market record", shouldInsertAValidMarketRecord)
	t.Run("Add should update a valid market record if the block number already exists", shouldUpdateAValidMarketRecord)
}

func TestMarkets_Get(t *testing.T) {
	t.Run("GetByID should return the request market if it exists", getByIDShouldReturnTheRequestedMarketIfItExists)
	t.Run("GetByID should return error if the market does not exist", getByIDShouldReturnErrorIfTheMarketDoesNotExist)
	t.Run("GetAllPaged should not include rejected markets", getAllPagedShouldNotIncludeRejectedMarkets)
	t.Run("GetByTxHash", getByTxHashReturnsMatchingMarkets)
}

func getByIDShouldReturnTheRequestedMarketIfItExists(t *testing.T) {
	bs, md := setupMarketsTest(t)

	ctx, rollback := tempTransaction(t)
	defer rollback()
	block := addTestBlock(t, ctx, bs)

	market := entities.Market{
		ID:       "deadbeef",
		TxHash:   generateTxHash(),
		VegaTime: block.VegaTime,
		State:    entities.MarketStateActive,
	}
	err := md.Upsert(ctx, &market)
	require.NoError(t, err, "Saving market entity to database")

	marketFromDB, err := md.GetByID(ctx, market.ID.String())
	require.NoError(t, err)
	assert.Equal(t, market.ID, marketFromDB.ID)
	assert.Equal(t, market.TxHash, marketFromDB.TxHash)
	assert.Equal(t, market.VegaTime, marketFromDB.VegaTime)
	assert.Equal(t, market.State, marketFromDB.State)
}

func getByTxHashReturnsMatchingMarkets(t *testing.T) {
	bs, md := setupMarketsTest(t)

	ctx, rollback := tempTransaction(t)
	defer rollback()
	block := addTestBlock(t, ctx, bs)

	market := entities.Market{
		ID:       "deadbeef",
		TxHash:   generateTxHash(),
		VegaTime: block.VegaTime,
		State:    entities.MarketStateActive,
	}
	err := md.Upsert(ctx, &market)
	require.NoError(t, err, "Saving market entity to database")

	foundMarkets, err := md.GetByTxHash(ctx, market.TxHash)
	require.NoError(t, err)
	require.Len(t, foundMarkets, 1)
	assert.Equal(t, market.ID, foundMarkets[0].ID)
	assert.Equal(t, market.TxHash, foundMarkets[0].TxHash)
	assert.Equal(t, market.VegaTime, foundMarkets[0].VegaTime)
	assert.Equal(t, market.State, foundMarkets[0].State)
}

func getByIDShouldReturnErrorIfTheMarketDoesNotExist(t *testing.T) {
	bs, md := setupMarketsTest(t)

	ctx, rollback := tempTransaction(t)
	defer rollback()
	block := addTestBlock(t, ctx, bs)

	market := entities.Market{
		ID:       "deadbeef",
		TxHash:   generateTxHash(),
		VegaTime: block.VegaTime,
		State:    entities.MarketStateActive,
	}
	err := md.Upsert(ctx, &market)
	require.NoError(t, err, "Saving market entity to database")

	_, err = md.GetByID(ctx, "not-a-market")
	require.Error(t, err)
}

func getAllPagedShouldNotIncludeRejectedMarkets(t *testing.T) {
	bs, md := setupMarketsTest(t)

	ctx, rollback := tempTransaction(t)
	defer rollback()
	block := addTestBlock(t, ctx, bs)

	market := entities.Market{
		ID:       "deadbeef",
		TxHash:   generateTxHash(),
		VegaTime: block.VegaTime,
		State:    entities.MarketStateActive,
	}
	err := md.Upsert(ctx, &market)
	require.NoError(t, err, "Saving market entity to database")

	rejected := entities.Market{
		ID:       "DEADBAAD",
		TxHash:   generateTxHash(),
		VegaTime: block.VegaTime,
		State:    entities.MarketStateRejected,
	}
	err = md.Upsert(ctx, &rejected)
	require.NoError(t, err, "Saving market entity to database")

	markets, pageInfo, err := md.GetAllPaged(ctx, "", entities.CursorPagination{}, true)
	require.NoError(t, err)
	assert.Len(t, markets, 1)
	assert.Equal(t, market.ID, markets[0].ID)
	assert.Equal(t, market.TxHash, markets[0].TxHash)
	assert.Equal(t, market.VegaTime, markets[0].VegaTime)
	assert.Equal(t, market.State, markets[0].State)
	assert.Equal(t, entities.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
		StartCursor:     market.Cursor().Encode(),
		EndCursor:       market.Cursor().Encode(),
	}, pageInfo)
}

func shouldInsertAValidMarketRecord(t *testing.T) {
	bs, md := setupMarketsTest(t)

	ctx, rollback := tempTransaction(t)
	defer rollback()

	conn := connectionSource.Connection
	var rowCount int

	err := conn.QueryRow(ctx, `select count(*) from markets`).Scan(&rowCount)
	require.NoError(t, err)
	assert.Equal(t, 0, rowCount)

	block := addTestBlock(t, ctx, bs)

	marketProto := getTestMarket(true)
	marketProto.LiquidityMonitoringParameters.TriggeringRatio = "0.3"

	market, err := entities.NewMarketFromProto(marketProto, generateTxHash(), block.VegaTime)
	require.NoError(t, err, "Converting market proto to database entity")

	err = md.Upsert(ctx, market)
	require.NoError(t, err, "Saving market entity to database")
	err = conn.QueryRow(ctx, `select count(*) from markets`).Scan(&rowCount)
	assert.NoError(t, err)
	assert.Equal(t, 1, rowCount)
}

func setupMarketsTest(t *testing.T) (*sqlstore.Blocks, *sqlstore.Markets) {
	t.Helper()
	bs := sqlstore.NewBlocks(connectionSource)
	md := sqlstore.NewMarkets(connectionSource)
	return bs, md
}

func shouldUpdateAValidMarketRecord(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()

	conn := connectionSource.Connection
	var rowCount int

	t.Run("should have no markets in the database", func(t *testing.T) {
		err := conn.QueryRow(ctx, `select count(*) from markets`).Scan(&rowCount)
		require.NoError(t, err)
		assert.Equal(t, 0, rowCount)
	})

	var block entities.Block
	var marketProto *vega.Market

	t.Run("should insert a valid market record to the database", func(t *testing.T) {
		block = addTestBlock(t, ctx, bs)
		marketProto = getTestMarket(false)

		market, err := entities.NewMarketFromProto(marketProto, generateTxHash(), block.VegaTime)
		require.NoError(t, err, "Converting market proto to database entity")

		err = md.Upsert(ctx, market)
		require.NoError(t, err, "Saving market entity to database")

		var got entities.Market
		err = pgxscan.Get(ctx, conn, &got, `select * from markets where id = $1 and vega_time = $2`, market.ID, market.VegaTime)
		assert.NoError(t, err)
		assert.Equal(t, "TEST_INSTRUMENT", market.InstrumentID)

		assert.Equal(t, marketProto.TradableInstrument, got.TradableInstrument.ToProto())
	})

	marketProto.TradableInstrument.Instrument.Name = "Updated Test Instrument"
	marketProto.TradableInstrument.Instrument.Metadata.Tags = append(marketProto.TradableInstrument.Instrument.Metadata.Tags, "CCC")

	t.Run("should update a valid market record to the database if the block number already exists", func(t *testing.T) {
		market, err := entities.NewMarketFromProto(marketProto, generateTxHash(), block.VegaTime)

		require.NoError(t, err, "Converting market proto to database entity")

		err = md.Upsert(ctx, market)
		require.NoError(t, err, "Saving market entity to database")

		var got entities.Market
		err = pgxscan.Get(ctx, conn, &got, `select * from markets where id = $1 and vega_time = $2`, market.ID, market.VegaTime)
		assert.NoError(t, err)
		assert.Equal(t, "TEST_INSTRUMENT", market.InstrumentID)

		assert.Equal(t, marketProto.TradableInstrument, got.TradableInstrument.ToProto())
	})

	t.Run("should add the updated market record to the database if the block number has changed", func(t *testing.T) {
		newMarketProto := marketProto.DeepClone()
		newMarketProto.TradableInstrument.Instrument.Metadata.Tags = append(newMarketProto.TradableInstrument.Instrument.Metadata.Tags, "DDD")
		newBlock := addTestBlockForTime(t, ctx, bs, time.Now().Add(time.Second))

		market, err := entities.NewMarketFromProto(newMarketProto, generateTxHash(), newBlock.VegaTime)
		require.NoError(t, err, "Converting market proto to database entity")

		err = md.Upsert(ctx, market)
		require.NoError(t, err, "Saving market entity to database")

		err = conn.QueryRow(ctx, `select count(*) from markets`).Scan(&rowCount)
		require.NoError(t, err)
		assert.Equal(t, 2, rowCount)

		var gotFirstBlock, gotSecondBlock entities.Market

		err = pgxscan.Get(ctx, conn, &gotFirstBlock, `select * from markets where id = $1 and vega_time = $2`, market.ID, block.VegaTime)
		assert.NoError(t, err)
		assert.Equal(t, "TEST_INSTRUMENT", market.InstrumentID)

		assert.Equal(t, marketProto.TradableInstrument, gotFirstBlock.TradableInstrument.ToProto())

		err = pgxscan.Get(ctx, conn, &gotSecondBlock, `select * from markets where id = $1 and vega_time = $2`, market.ID, newBlock.VegaTime)
		assert.NoError(t, err)
		assert.Equal(t, "TEST_INSTRUMENT", market.InstrumentID)

		assert.Equal(t, newMarketProto.TradableInstrument, gotSecondBlock.TradableInstrument.ToProto())
	})
}

func getTestMarket(termInt bool) *vega.Market {
	term := &vega.DataSourceSpec{
		Id:        "",
		CreatedAt: 0,
		UpdatedAt: 0,
		Data: vega.NewDataSourceDefinition(
			vega.DataSourceDefinitionTypeExt,
		).SetOracleConfig(
			&vega.DataSourceSpecConfiguration{
				Signers: nil,
				Filters: nil,
			},
		),
		Status: 0,
	}

	if termInt {
		term = &vega.DataSourceSpec{
			Id:        "",
			CreatedAt: 0,
			UpdatedAt: 0,
			Data: vega.NewDataSourceDefinition(
				vega.DataSourceDefinitionTypeInt,
			).SetTimeTriggerConditionConfig(
				[]*v1.Condition{
					{
						Operator: v1.Condition_OPERATOR_GREATER_THAN,
						Value:    "test-value",
					},
				},
			),
			Status: 0,
		}
	}

	return &vega.Market{
		Id: "DEADBEEF",
		TradableInstrument: &vega.TradableInstrument{
			Instrument: &vega.Instrument{
				Id:   "TEST_INSTRUMENT",
				Code: "TEST",
				Name: "Test Instrument",
				Metadata: &vega.InstrumentMetadata{
					Tags: []string{"AAA", "BBB"},
				},
				Product: &vega.Instrument_Future{
					Future: &vega.Future{
						SettlementAsset: "Test Asset",
						QuoteName:       "Test Quote",
						DataSourceSpecForSettlementData: &vega.DataSourceSpec{
							Id:        "",
							CreatedAt: 0,
							UpdatedAt: 0,
							Data: vega.NewDataSourceDefinition(
								vega.DataSourceDefinitionTypeExt,
							).SetOracleConfig(
								&vega.DataSourceSpecConfiguration{
									Signers: nil,
									Filters: nil,
								},
							),
							Status: 0,
						},
						DataSourceSpecForTradingTermination: term,
						DataSourceSpecBinding: &vega.DataSourceSpecToFutureBinding{
							SettlementDataProperty:     "",
							TradingTerminationProperty: "",
						},
					},
				},
			},
			MarginCalculator: &vega.MarginCalculator{
				ScalingFactors: &vega.ScalingFactors{
					SearchLevel:       0,
					InitialMargin:     0,
					CollateralRelease: 0,
				},
			},
			RiskModel: &vega.TradableInstrument_SimpleRiskModel{
				SimpleRiskModel: &vega.SimpleRiskModel{
					Params: &vega.SimpleModelParams{
						FactorLong:           0,
						FactorShort:          0,
						MaxMoveUp:            0,
						MinMoveDown:          0,
						ProbabilityOfTrading: 0,
					},
				},
			},
		},
		DecimalPlaces: 16,
		Fees: &vega.Fees{
			Factors: &vega.FeeFactors{
				MakerFee:          "",
				InfrastructureFee: "",
				LiquidityFee:      "",
			},
		},
		OpeningAuction: &vega.AuctionDuration{
			Duration: 0,
			Volume:   0,
		},
		PriceMonitoringSettings: &vega.PriceMonitoringSettings{
			Parameters: &vega.PriceMonitoringParameters{
				Triggers: []*vega.PriceMonitoringTrigger{
					{
						Horizon:          0,
						Probability:      "",
						AuctionExtension: 0,
					},
				},
			},
		},
		LiquidityMonitoringParameters: &vega.LiquidityMonitoringParameters{
			TargetStakeParameters: &vega.TargetStakeParameters{
				TimeWindow:    0,
				ScalingFactor: 0,
			},
			TriggeringRatio:  "0",
			AuctionExtension: 0,
		},
		TradingMode: vega.Market_TRADING_MODE_CONTINUOUS,
		State:       vega.Market_STATE_ACTIVE,
		MarketTimestamps: &vega.MarketTimestamps{
			Proposed: 0,
			Pending:  0,
			Open:     0,
			Close:    0,
		},
		PositionDecimalPlaces:   8,
		LpPriceRange:            "0.95",
		LinearSlippageFactor:    "1.23",
		QuadraticSlippageFactor: "5.67",
	}
}

func populateTestMarkets(ctx context.Context, t *testing.T, bs *sqlstore.Blocks, md *sqlstore.Markets, blockTimes map[string]time.Time) {
	t.Helper()

	markets := []entities.Market{
		{
			ID:           entities.MarketID("02a16077"),
			InstrumentID: "AAA",
		},
		{
			ID:           entities.MarketID("44eea1bc"),
			InstrumentID: "BBB",
		},
		{
			ID:           entities.MarketID("65be62cd"),
			InstrumentID: "CCC",
		},
		{
			ID:           entities.MarketID("7a797e0e"),
			InstrumentID: "DDD",
		},
		{
			ID:           entities.MarketID("7bb2356e"),
			InstrumentID: "EEE",
		},
		{
			ID:           entities.MarketID("b7c84b8e"),
			InstrumentID: "FFF",
		},
		{
			ID:           entities.MarketID("c612300d"),
			InstrumentID: "GGG",
		},
		{
			ID:           entities.MarketID("c8744329"),
			InstrumentID: "HHH",
		},
		{
			ID:           entities.MarketID("da8d1803"),
			InstrumentID: "III",
		},
		{
			ID:           entities.MarketID("fb1528a5"),
			InstrumentID: "JJJ",
		},
	}

	source := &testBlockSource{bs, time.Now()}
	for _, market := range markets {
		block := source.getNextBlock(t, ctx)
		market.VegaTime = block.VegaTime
		blockTimes[market.ID.String()] = block.VegaTime
		err := md.Upsert(ctx, &market)
		require.NoError(t, err)
	}
}

func TestMarketsCursorPagination(t *testing.T) {
	t.Run("Should return the market if Market ID is provided", testCursorPaginationReturnsTheSpecifiedMarket)
	t.Run("Should return all markets if no market ID and no cursor is provided", testCursorPaginationReturnsAllMarkets)
	t.Run("Should return the first page when first limit is provided with no after cursor", testCursorPaginationReturnsFirstPage)
	t.Run("Should return the last page when last limit is provided with first before cursor", testCursorPaginationReturnsLastPage)
	t.Run("Should return the page specified by the first limit and after cursor", testCursorPaginationReturnsPageTraversingForward)
	t.Run("Should return the page specified by the last limit and before cursor", testCursorPaginationReturnsPageTraversingBackward)

	t.Run("Should return the market if Market ID is provided - newest first", testCursorPaginationReturnsTheSpecifiedMarketNewestFirst)
	t.Run("Should return all markets if no market ID and no cursor is provided - newest first", testCursorPaginationReturnsAllMarketsNewestFirst)
	t.Run("Should return the first page when first limit is provided with no after cursor - newest first", testCursorPaginationReturnsFirstPageNewestFirst)
	t.Run("Should return the last page when last limit is provided with first before cursor - newest first", testCursorPaginationReturnsLastPageNewestFirst)
	t.Run("Should return the page specified by the first limit and after cursor - newest first", testCursorPaginationReturnsPageTraversingForwardNewestFirst)
	t.Run("Should return the page specified by the last limit and before cursor - newest first", testCursorPaginationReturnsPageTraversingBackwardNewestFirst)
}

func testCursorPaginationReturnsTheSpecifiedMarket(t *testing.T) {
	ctx, rollback := tempTransaction(t)
	defer rollback()

	bs, md := setupMarketsTest(t)

	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)
	pagination, err := entities.NewCursorPagination(nil, nil, nil, nil, false)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "c612300d", pagination, true)
	require.NoError(t, err)
	assert.Equal(t, 1, len(got))
	assert.Equal(t, "c612300d", got[0].ID.String())
	assert.Equal(t, "GGG", got[0].InstrumentID)

	mc := entities.MarketCursor{
		VegaTime: blockTimes["c612300d"],
		ID:       "c612300d",
	}

	wantStartCursor := entities.NewCursor(mc.String()).Encode()
	wantEndCursor := entities.NewCursor(mc.String()).Encode()

	assert.Equal(t, entities.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func testCursorPaginationReturnsAllMarkets(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()
	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)

	pagination, err := entities.NewCursorPagination(nil, nil, nil, nil, false)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "", pagination, true)
	require.NoError(t, err)
	assert.Equal(t, 10, len(got))
	assert.Equal(t, "02a16077", got[0].ID.String())
	assert.Equal(t, "fb1528a5", got[9].ID.String())
	assert.Equal(t, "AAA", got[0].InstrumentID)
	assert.Equal(t, "JJJ", got[9].InstrumentID)

	wantStartCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["02a16077"],
			ID:       "02a16077",
		}.String(),
	).Encode()
	wantEndCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["fb1528a5"],
			ID:       "fb1528a5",
		}.String(),
	).Encode()
	assert.Equal(t, entities.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func testCursorPaginationReturnsFirstPage(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()
	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)
	first := int32(3)
	pagination, err := entities.NewCursorPagination(&first, nil, nil, nil, false)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "", pagination, true)
	require.NoError(t, err)

	assert.Equal(t, 3, len(got))
	assert.Equal(t, "02a16077", got[0].ID.String())
	assert.Equal(t, "65be62cd", got[2].ID.String())
	assert.Equal(t, "AAA", got[0].InstrumentID)
	assert.Equal(t, "CCC", got[2].InstrumentID)

	wantStartCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["02a16077"],
			ID:       "02a16077",
		}.String(),
	).Encode()
	wantEndCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["65be62cd"],
			ID:       "65be62cd",
		}.String(),
	).Encode()
	assert.Equal(t, entities.PageInfo{
		HasNextPage:     true,
		HasPreviousPage: false,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func testCursorPaginationReturnsLastPage(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()
	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)
	last := int32(3)
	pagination, err := entities.NewCursorPagination(nil, nil, &last, nil, false)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "", pagination, true)
	require.NoError(t, err)

	assert.Equal(t, 3, len(got))
	assert.Equal(t, "c8744329", got[0].ID.String())
	assert.Equal(t, "fb1528a5", got[2].ID.String())
	assert.Equal(t, "HHH", got[0].InstrumentID)
	assert.Equal(t, "JJJ", got[2].InstrumentID)

	wantStartCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["c8744329"],
			ID:       "c8744329",
		}.String(),
	).Encode()
	wantEndCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["fb1528a5"],
			ID:       "fb1528a5",
		}.String(),
	).Encode()
	assert.Equal(t, entities.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: true,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func testCursorPaginationReturnsPageTraversingForward(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()
	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)
	first := int32(3)
	after := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["65be62cd"],
			ID:       "65be62cd",
		}.String(),
	).Encode()
	pagination, err := entities.NewCursorPagination(&first, &after, nil, nil, false)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "", pagination, true)
	require.NoError(t, err)

	assert.Equal(t, 3, len(got))
	assert.Equal(t, "7a797e0e", got[0].ID.String())
	assert.Equal(t, "b7c84b8e", got[2].ID.String())
	assert.Equal(t, "DDD", got[0].InstrumentID)
	assert.Equal(t, "FFF", got[2].InstrumentID)

	wantStartCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["7a797e0e"],
			ID:       "7a797e0e",
		}.String(),
	).Encode()
	wantEndCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["b7c84b8e"],
			ID:       "b7c84b8e",
		}.String(),
	).Encode()
	assert.Equal(t, entities.PageInfo{
		HasNextPage:     true,
		HasPreviousPage: true,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func testCursorPaginationReturnsPageTraversingBackward(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()
	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)
	last := int32(3)
	before := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["c8744329"],
			ID:       "c8744329",
		}.String(),
	).Encode()
	pagination, err := entities.NewCursorPagination(nil, nil, &last, &before, false)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "", pagination, true)
	require.NoError(t, err)

	assert.Equal(t, 3, len(got))
	assert.Equal(t, "7bb2356e", got[0].ID.String())
	assert.Equal(t, "c612300d", got[2].ID.String())
	assert.Equal(t, "EEE", got[0].InstrumentID)
	assert.Equal(t, "GGG", got[2].InstrumentID)

	wantStartCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["7bb2356e"],
			ID:       "7bb2356e",
		}.String(),
	).Encode()
	wantEndCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["c612300d"],
			ID:       "c612300d",
		}.String(),
	).Encode()
	assert.Equal(t, entities.PageInfo{
		HasNextPage:     true,
		HasPreviousPage: true,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func testCursorPaginationReturnsTheSpecifiedMarketNewestFirst(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()
	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)
	pagination, err := entities.NewCursorPagination(nil, nil, nil, nil, true)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "c612300d", pagination, true)
	require.NoError(t, err)
	assert.Equal(t, 1, len(got))
	assert.Equal(t, "c612300d", got[0].ID.String())
	assert.Equal(t, "GGG", got[0].InstrumentID)

	wantStartCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["c612300d"],
			ID:       "c612300d",
		}.String(),
	).Encode()
	wantEndCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["c612300d"],
			ID:       "c612300d",
		}.String(),
	).Encode()

	assert.Equal(t, entities.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func testCursorPaginationReturnsAllMarketsNewestFirst(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()
	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)

	pagination, err := entities.NewCursorPagination(nil, nil, nil, nil, true)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "", pagination, true)
	require.NoError(t, err)
	assert.Equal(t, 10, len(got))
	assert.Equal(t, "fb1528a5", got[0].ID.String())
	assert.Equal(t, "02a16077", got[9].ID.String())
	assert.Equal(t, "JJJ", got[0].InstrumentID)
	assert.Equal(t, "AAA", got[9].InstrumentID)

	wantStartCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["fb1528a5"],
			ID:       "fb1528a5",
		}.String(),
	).Encode()
	wantEndCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["02a16077"],
			ID:       "02a16077",
		}.String(),
	).Encode()
	assert.Equal(t, entities.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func testCursorPaginationReturnsFirstPageNewestFirst(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()
	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)
	first := int32(3)
	pagination, err := entities.NewCursorPagination(&first, nil, nil, nil, true)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "", pagination, true)
	require.NoError(t, err)

	assert.Equal(t, 3, len(got))
	assert.Equal(t, "fb1528a5", got[0].ID.String())
	assert.Equal(t, "c8744329", got[2].ID.String())
	assert.Equal(t, "JJJ", got[0].InstrumentID)
	assert.Equal(t, "HHH", got[2].InstrumentID)

	wantStartCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["fb1528a5"],
			ID:       "fb1528a5",
		}.String(),
	).Encode()
	wantEndCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["c8744329"],
			ID:       "c8744329",
		}.String(),
	).Encode()
	assert.Equal(t, entities.PageInfo{
		HasNextPage:     true,
		HasPreviousPage: false,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func testCursorPaginationReturnsLastPageNewestFirst(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()
	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)
	last := int32(3)
	pagination, err := entities.NewCursorPagination(nil, nil, &last, nil, true)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "", pagination, true)
	require.NoError(t, err)

	assert.Equal(t, 3, len(got))
	assert.Equal(t, "65be62cd", got[0].ID.String())
	assert.Equal(t, "02a16077", got[2].ID.String())
	assert.Equal(t, "CCC", got[0].InstrumentID)
	assert.Equal(t, "AAA", got[2].InstrumentID)

	wantStartCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["65be62cd"],
			ID:       "65be62cd",
		}.String(),
	).Encode()
	wantEndCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["02a16077"],
			ID:       "02a16077",
		}.String(),
	).Encode()
	assert.Equal(t, entities.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: true,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func testCursorPaginationReturnsPageTraversingForwardNewestFirst(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()
	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)
	first := int32(3)
	after := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["c8744329"],
			ID:       "c8744329",
		}.String(),
	).Encode()
	pagination, err := entities.NewCursorPagination(&first, &after, nil, nil, true)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "", pagination, true)
	require.NoError(t, err)

	assert.Equal(t, 3, len(got))
	assert.Equal(t, "c612300d", got[0].ID.String())
	assert.Equal(t, "7bb2356e", got[2].ID.String())
	assert.Equal(t, "GGG", got[0].InstrumentID)
	assert.Equal(t, "EEE", got[2].InstrumentID)

	wantStartCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["c612300d"],
			ID:       "c612300d",
		}.String(),
	).Encode()
	wantEndCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["7bb2356e"],
			ID:       "7bb2356e",
		}.String(),
	).Encode()
	assert.Equal(t, entities.PageInfo{
		HasNextPage:     true,
		HasPreviousPage: true,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func testCursorPaginationReturnsPageTraversingBackwardNewestFirst(t *testing.T) {
	bs, md := setupMarketsTest(t)
	ctx, rollback := tempTransaction(t)
	defer rollback()
	blockTimes := make(map[string]time.Time)
	populateTestMarkets(ctx, t, bs, md, blockTimes)
	last := int32(3)
	before := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["65be62cd"],
			ID:       "65be62cd",
		}.String(),
	).Encode()
	pagination, err := entities.NewCursorPagination(nil, nil, &last, &before, true)
	require.NoError(t, err)

	got, pageInfo, err := md.GetAllPaged(ctx, "", pagination, true)
	require.NoError(t, err)

	assert.Equal(t, 3, len(got))
	assert.Equal(t, "b7c84b8e", got[0].ID.String())
	assert.Equal(t, "7a797e0e", got[2].ID.String())
	assert.Equal(t, "FFF", got[0].InstrumentID)
	assert.Equal(t, "DDD", got[2].InstrumentID)

	wantStartCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["b7c84b8e"],
			ID:       "b7c84b8e",
		}.String(),
	).Encode()
	wantEndCursor := entities.NewCursor(
		entities.MarketCursor{
			VegaTime: blockTimes["7a797e0e"],
			ID:       "7a797e0e",
		}.String(),
	).Encode()
	assert.Equal(t, entities.PageInfo{
		HasNextPage:     true,
		HasPreviousPage: true,
		StartCursor:     wantStartCursor,
		EndCursor:       wantEndCursor,
	}, pageInfo)
}

func TestSuccessorMarkets(t *testing.T) {
	t.Run("should create a market lineage record when a successor market proposal is approved", testMarketLineageCreated)
	t.Run("ListSuccessorMarkets should return the market lineage", testListSuccessorMarkets)
	t.Run("GetMarket should return the market with its parent and successor if they exist", testGetMarketWithParentAndSuccessor)
}

func testMarketLineageCreated(t *testing.T) {
	ctx, rollback := tempTransaction(t)
	defer rollback()

	bs, md := setupMarketsTest(t)
	parentMarket := entities.Market{
		ID:           entities.MarketID("deadbeef01"),
		InstrumentID: "deadbeef01",
	}

	successorMarketA := entities.Market{
		ID:             entities.MarketID("deadbeef02"),
		InstrumentID:   "deadbeef02",
		ParentMarketID: parentMarket.ID,
	}

	successorMarketB := entities.Market{
		ID:             entities.MarketID("deadbeef03"),
		InstrumentID:   "deadbeef03",
		ParentMarketID: successorMarketA.ID,
	}

	conn := connectionSource.Connection
	var rowCount int64

	source := &testBlockSource{bs, time.Now()}
	block := source.getNextBlock(t, ctx)
	t.Run("parent market should create a market lineage record with no parent market id", func(t *testing.T) {
		parentMarket.VegaTime = block.VegaTime
		parentMarket.State = entities.MarketStateProposed
		err := md.Upsert(ctx, &parentMarket)
		require.NoError(t, err)
		err = conn.QueryRow(ctx, `select count(*) from market_lineage where market_id = $1`, parentMarket.ID).Scan(&rowCount)
		require.NoError(t, err)
		assert.Equal(t, int64(0), rowCount)

		block = source.getNextBlock(t, ctx)
		parentMarket.State = entities.MarketStatePending
		parentMarket.TradingMode = entities.MarketTradingModeOpeningAuction
		parentMarket.VegaTime = block.VegaTime
		err = md.Upsert(ctx, &parentMarket)
		require.NoError(t, err)

		block = source.getNextBlock(t, ctx)
		parentMarket.State = entities.MarketStateActive
		parentMarket.TradingMode = entities.MarketTradingModeContinuous
		parentMarket.VegaTime = block.VegaTime
		err = md.Upsert(ctx, &parentMarket)
		require.NoError(t, err)

		var marketID, parentMarketID, rootID entities.MarketID
		err = conn.QueryRow(ctx,
			`select market_id, parent_market_id, root_id from market_lineage where market_id = $1`,
			parentMarket.ID,
		).Scan(&marketID, &parentMarketID, &rootID)
		require.NoError(t, err)
		assert.Equal(t, parentMarket.ID, marketID)
		assert.Equal(t, entities.MarketID(""), parentMarketID)
		assert.Equal(t, parentMarket.ID, rootID)
	})

	block = source.getNextBlock(t, ctx)
	t.Run("successor market should create a market lineage record pointing to the parent market and the root market", func(t *testing.T) {
		successorMarketA.VegaTime = block.VegaTime
		successorMarketA.State = entities.MarketStateProposed
		err := md.Upsert(ctx, &successorMarketA)
		require.NoError(t, err)
		// proposed market successor only, so it should not create a lineage record yet
		err = conn.QueryRow(ctx, `select count(*) from market_lineage where market_id = $1`, successorMarketA.ID).Scan(&rowCount)
		require.NoError(t, err)
		assert.Equal(t, int64(0), rowCount)

		block = source.getNextBlock(t, ctx)
		successorMarketA.State = entities.MarketStatePending
		successorMarketA.TradingMode = entities.MarketTradingModeOpeningAuction
		successorMarketA.VegaTime = block.VegaTime
		err = md.Upsert(ctx, &successorMarketA)
		require.NoError(t, err)

		block = source.getNextBlock(t, ctx)
		successorMarketA.State = entities.MarketStateActive
		successorMarketA.TradingMode = entities.MarketTradingModeContinuous
		successorMarketA.VegaTime = block.VegaTime
		err = md.Upsert(ctx, &successorMarketA)
		require.NoError(t, err)
		// proposed market successor has been accepted and is pending, so we should now have a lineage record pointing to the parent
		var marketID, parentMarketID, rootID entities.MarketID
		err = conn.QueryRow(ctx,
			`select market_id, parent_market_id, root_id from market_lineage where market_id = $1`,
			successorMarketA.ID,
		).Scan(&marketID, &parentMarketID, &rootID)
		require.NoError(t, err)
		assert.Equal(t, successorMarketA.ID, marketID)
		assert.Equal(t, parentMarket.ID, parentMarketID)
		assert.Equal(t, parentMarket.ID, rootID)
	})

	block = source.getNextBlock(t, ctx)
	t.Run("second successor market should create a lineage record pointing to the parent market and the root market", func(t *testing.T) {
		successorMarketB.VegaTime = block.VegaTime
		successorMarketB.State = entities.MarketStateProposed
		err := md.Upsert(ctx, &successorMarketB)
		require.NoError(t, err)
		// proposed market successor only, so it should not create a lineage record yet
		err = conn.QueryRow(ctx, `select count(*) from market_lineage where market_id = $1`, successorMarketB.ID).Scan(&rowCount)
		require.NoError(t, err)
		assert.Equal(t, int64(0), rowCount)

		block = source.getNextBlock(t, ctx)
		successorMarketB.State = entities.MarketStatePending
		successorMarketB.TradingMode = entities.MarketTradingModeOpeningAuction
		successorMarketB.VegaTime = block.VegaTime
		err = md.Upsert(ctx, &successorMarketB)
		require.NoError(t, err)
		// proposed market successor has been accepted and is pending, so we should now have a lineage record pointing to the parent
		block = source.getNextBlock(t, ctx)
		successorMarketB.State = entities.MarketStateActive
		successorMarketB.TradingMode = entities.MarketTradingModeContinuous
		successorMarketB.VegaTime = block.VegaTime
		err = md.Upsert(ctx, &successorMarketB)
		require.NoError(t, err)
		var marketID, parentMarketID, rootID entities.MarketID
		err = conn.QueryRow(ctx,
			`select market_id, parent_market_id, root_id from market_lineage where market_id = $1`,
			successorMarketB.ID,
		).Scan(&marketID, &parentMarketID, &rootID)
		require.NoError(t, err)
		assert.Equal(t, successorMarketB.ID, marketID)
		assert.Equal(t, successorMarketA.ID, parentMarketID)
		assert.Equal(t, parentMarket.ID, rootID)
	})
}

func testListSuccessorMarkets(t *testing.T) {
	ctx, rollback := tempTransaction(t)
	defer rollback()
	md, markets, proposals := setupSuccessorMarkets(t, ctx)

	successors := []entities.SuccessorMarket{
		{
			Market: markets[5],
			Proposals: []*entities.Proposal{
				&proposals[1],
				&proposals[2],
			},
		},
		{
			Market: markets[6],
			Proposals: []*entities.Proposal{
				&proposals[3],
				&proposals[4],
			},
		},
		{
			Market: markets[8],
		},
	}

	t.Run("should list the full history if children only is false", func(t *testing.T) {
		got, _, err := md.ListSuccessorMarkets(ctx, "deadbeef02", true, entities.CursorPagination{})
		require.NoError(t, err)
		want := successors[:]
		assert.Equal(t, want, got)
	})

	t.Run("should list only the successor markets if children only is true", func(t *testing.T) {
		got, _, err := md.ListSuccessorMarkets(ctx, "deadbeef02", false, entities.CursorPagination{})
		require.NoError(t, err)
		want := successors[1:]

		assert.Equal(t, want, got)
	})

	t.Run("should paginate results if pagination is provided", func(t *testing.T) {
		first := int32(2)
		after := entities.NewCursor(
			entities.MarketCursor{
				VegaTime: markets[5].VegaTime,
				ID:       markets[5].ID,
			}.String(),
		).Encode()
		pagination, err := entities.NewCursorPagination(&first, &after, nil, nil, false)
		require.NoError(t, err)
		got, pageInfo, err := md.ListSuccessorMarkets(ctx, "deadbeef01", true, pagination)
		require.NoError(t, err)
		want := successors[1:]

		assert.Equal(t, want, got, "paged successor markets do not match")
		wantStartCursor := entities.NewCursor(
			entities.MarketCursor{
				VegaTime: markets[6].VegaTime,
				ID:       markets[6].ID,
			}.String(),
		).Encode()
		wantEndCursor := entities.NewCursor(
			entities.MarketCursor{
				VegaTime: markets[8].VegaTime,
				ID:       markets[8].ID,
			}.String(),
		).Encode()
		assert.Equal(t, entities.PageInfo{
			HasNextPage:     false,
			HasPreviousPage: true,
			StartCursor:     wantStartCursor,
			EndCursor:       wantEndCursor,
		}, pageInfo)
	})
}

func testGetMarketWithParentAndSuccessor(t *testing.T) {
	ctx, rollback := tempTransaction(t)
	defer rollback()

	md, _, _ := setupSuccessorMarkets(t, ctx)

	t.Run("should return successor market id only if the first market in a succession line", func(t *testing.T) {
		got, err := md.GetByID(ctx, "deadbeef01")
		require.NoError(t, err)
		assert.Equal(t, "", got.ParentMarketID.String())
		assert.Equal(t, "deadbeef02", got.SuccessorMarketID.String())
	})

	t.Run("should return parent and successor market id if the market is within a succession line", func(t *testing.T) {
		got, err := md.GetByID(ctx, "deadbeef02")
		require.NoError(t, err)
		assert.Equal(t, "deadbeef01", got.ParentMarketID.String())
		assert.Equal(t, "deadbeef03", got.SuccessorMarketID.String())
	})

	t.Run("should return parent market id only if the last market in a succession line", func(t *testing.T) {
		got, err := md.GetByID(ctx, "deadbeef03")
		require.NoError(t, err)
		assert.Equal(t, "deadbeef02", got.ParentMarketID.String())
		assert.Equal(t, "", got.SuccessorMarketID.String())
	})
}

func setupSuccessorMarkets(t *testing.T, ctx context.Context) (*sqlstore.Markets, []entities.Market, []entities.Proposal) {
	t.Helper()

	bs, md := setupMarketsTest(t)
	ps := sqlstore.NewProposals(connectionSource)
	ts := sqlstore.NewParties(connectionSource)

	parentMarket := entities.Market{
		ID:           entities.MarketID("deadbeef01"),
		InstrumentID: "deadbeef01",
		TradableInstrument: entities.TradableInstrument{
			TradableInstrument: &vega.TradableInstrument{},
		},
	}

	successorMarketA := entities.Market{
		ID:           entities.MarketID("deadbeef02"),
		InstrumentID: "deadbeef02",
		TradableInstrument: entities.TradableInstrument{
			TradableInstrument: &vega.TradableInstrument{},
		},
		ParentMarketID: parentMarket.ID,
	}

	parentMarket.SuccessorMarketID = successorMarketA.ID

	successorMarketB := entities.Market{
		ID:           entities.MarketID("deadbeef03"),
		InstrumentID: "deadbeef03",
		TradableInstrument: entities.TradableInstrument{
			TradableInstrument: &vega.TradableInstrument{},
		},
		ParentMarketID: successorMarketA.ID,
	}

	successorMarketA.SuccessorMarketID = successorMarketB.ID

	source := &testBlockSource{bs, time.Now()}

	block := source.getNextBlock(t, ctx)

	pt1 := addTestParty(t, ctx, ts, block)
	pt2 := addTestParty(t, ctx, ts, block)

	proposals := []struct {
		id        string
		party     entities.Party
		reference string
		block     entities.Block
		state     entities.ProposalState
		rationale entities.ProposalRationale
		terms     entities.ProposalTerms
		reason    entities.ProposalError
	}{
		{
			id:        "deadbeef01",
			party:     pt1,
			reference: "deadbeef01",
			block:     source.getNextBlock(t, ctx),
			state:     entities.ProposalStateEnacted,
			rationale: entities.ProposalRationale{ProposalRationale: &vega.ProposalRationale{Title: "myurl1.com", Description: "mydescription1"}},
			terms:     entities.ProposalTerms{ProposalTerms: &vega.ProposalTerms{Change: &vega.ProposalTerms_NewMarket{NewMarket: &vega.NewMarket{}}}},
			reason:    entities.ProposalErrorUnspecified,
		},
		{
			id:        "deadbeef02",
			party:     pt1,
			reference: "deadbeef02",
			block:     source.getNextBlock(t, ctx),
			state:     entities.ProposalStateEnacted,
			rationale: entities.ProposalRationale{ProposalRationale: &vega.ProposalRationale{Title: "myurl1.com", Description: "mydescription1"}},
			terms: entities.ProposalTerms{ProposalTerms: &vega.ProposalTerms{Change: &vega.ProposalTerms_NewMarket{NewMarket: &vega.NewMarket{
				Changes: &vega.NewMarketConfiguration{
					Successor: &vega.SuccessorConfiguration{
						ParentMarketId:        "deadbeef01",
						InsurancePoolFraction: "1.0",
					},
				},
			}}}},
			reason: entities.ProposalErrorUnspecified,
		},
		{
			id:        "deadbeefaa",
			party:     pt2,
			reference: "deadbeefaa",
			block:     source.getNextBlock(t, ctx),
			state:     entities.ProposalStateEnacted,
			rationale: entities.ProposalRationale{ProposalRationale: &vega.ProposalRationale{Title: "myurl1.com", Description: "mydescription1"}},
			terms: entities.ProposalTerms{ProposalTerms: &vega.ProposalTerms{Change: &vega.ProposalTerms_NewMarket{NewMarket: &vega.NewMarket{
				Changes: &vega.NewMarketConfiguration{
					Successor: &vega.SuccessorConfiguration{
						ParentMarketId:        "deadbeef01",
						InsurancePoolFraction: "1.0",
					},
				},
			}}}},
			reason: entities.ProposalErrorParticipationThresholdNotReached,
		},
		{
			id:        "deadbeef03",
			party:     pt1,
			reference: "deadbeef03",
			block:     source.getNextBlock(t, ctx),
			state:     entities.ProposalStateEnacted,
			rationale: entities.ProposalRationale{ProposalRationale: &vega.ProposalRationale{Title: "myurl1.com", Description: "mydescription1"}},
			terms: entities.ProposalTerms{ProposalTerms: &vega.ProposalTerms{Change: &vega.ProposalTerms_NewMarket{NewMarket: &vega.NewMarket{
				Changes: &vega.NewMarketConfiguration{
					Successor: &vega.SuccessorConfiguration{
						ParentMarketId:        "deadbeef02",
						InsurancePoolFraction: "1.0",
					},
				},
			}}}},
			reason: entities.ProposalErrorUnspecified,
		},
		{
			id:        "deadbeefbb",
			party:     pt2,
			reference: "deadbeefbb",
			block:     source.getNextBlock(t, ctx),
			state:     entities.ProposalStateEnacted,
			rationale: entities.ProposalRationale{ProposalRationale: &vega.ProposalRationale{Title: "myurl1.com", Description: "mydescription1"}},
			terms: entities.ProposalTerms{ProposalTerms: &vega.ProposalTerms{Change: &vega.ProposalTerms_NewMarket{NewMarket: &vega.NewMarket{
				Changes: &vega.NewMarketConfiguration{
					Successor: &vega.SuccessorConfiguration{
						ParentMarketId:        "deadbeef02",
						InsurancePoolFraction: "1.0",
					},
				},
			}}}},
			reason: entities.ProposalErrorParticipationThresholdNotReached,
		},
	}

	props := []entities.Proposal{}
	for _, p := range proposals {
		p := addTestProposal(t, ctx, ps, p.id, p.party, p.reference, p.block, p.state,
			p.rationale, p.terms, p.reason)

		props = append(props, p)
	}

	markets := []struct {
		market      entities.Market
		state       entities.MarketState
		tradingMode entities.MarketTradingMode
	}{
		{
			market:      parentMarket,
			state:       entities.MarketStateProposed,
			tradingMode: entities.MarketTradingModeUnspecified,
		},
		{
			market:      parentMarket,
			state:       entities.MarketStatePending,
			tradingMode: entities.MarketTradingModeOpeningAuction,
		},
		{
			market:      parentMarket,
			state:       entities.MarketStateActive,
			tradingMode: entities.MarketTradingModeContinuous,
		},
		{
			market:      successorMarketA,
			state:       entities.MarketStateProposed,
			tradingMode: entities.MarketTradingModeUnspecified,
		},
		{
			market:      successorMarketA,
			state:       entities.MarketStatePending,
			tradingMode: entities.MarketTradingModeOpeningAuction,
		},
		{
			market:      parentMarket,
			state:       entities.MarketStateSettled,
			tradingMode: entities.MarketTradingModeNoTrading,
		},
		{
			market:      successorMarketA,
			state:       entities.MarketStateActive,
			tradingMode: entities.MarketTradingModeContinuous,
		},
		{
			market:      successorMarketB,
			state:       entities.MarketStateProposed,
			tradingMode: entities.MarketTradingModeUnspecified,
		},
		{
			market:      successorMarketB,
			state:       entities.MarketStatePending,
			tradingMode: entities.MarketTradingModeOpeningAuction,
		},
		{
			market:      successorMarketB,
			state:       entities.MarketStateActive,
			tradingMode: entities.MarketTradingModeContinuous,
		},
	}

	entries := make([]entities.Market, 0, len(markets))

	for _, u := range markets {
		block := source.getNextBlock(t, ctx)
		u.market.VegaTime = block.VegaTime
		u.market.State = u.state
		u.market.TradingMode = u.tradingMode
		err := md.Upsert(ctx, &u.market)
		entries = append(entries, u.market)
		require.NoError(t, err)
	}

	return md, entries, props
}
