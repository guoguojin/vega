package settlement_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	bmock "code.vegaprotocol.io/vega/broker/mocks"
	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/settlement"
	"code.vegaprotocol.io/vega/settlement/mocks"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testEngine struct {
	*settlement.Engine
	ctrl      *gomock.Controller
	prod      *mocks.MockProduct
	positions []*mocks.MockMarketPosition
	broker    *bmock.MockBroker
	market    string
}

type posValue struct {
	trader string
	price  *num.Uint // absolute Mark price
	size   int64
}

type marginVal struct {
	events.MarketPosition
	asset, marketID                  string
	margin, general, marginShortFall uint64
}

func TestMarketExpiry(t *testing.T) {
	t.Run("Settle at market expiry - success", testSettleExpiredSuccess)
	t.Run("Settle at market expiry - error", testSettleExpiryFail)
	t.Run("Settle at market expiry with mark price - success", testSettleExpiredSuccessWithMarkPrice)
	t.Run("Settle at market expiry - failure invalid method", testSettleExpiredSuccessErrorInvalidSettlementMethod)
}

func TestMarkToMarket(t *testing.T) {
	t.Run("No settle positions if none were on channel", testMarkToMarketEmpty)
	t.Run("Settle positions are pushed onto the slice channel in order", testMarkToMarketOrdered)
	t.Run("Trade adds new trader to market, no MTM settlement because markPrice is the same", testAddNewTrader)
	// add this test case because we had a runtime panic on the trades map earlier
	t.Run("Trade adds new trader, immediately closing out with themselves", testAddNewTraderSelfTrade)
	t.Run("Test MTM settle when the network is closed out", testMTMNetworkZero)
}

func testSettleExpiredSuccess(t *testing.T) {
	engine := getTestEngine(t)
	defer engine.Finish()
	// these are mark prices, product will provide the actual value
	pr := num.NewUint(1000)
	data := []posValue{ // {{{2
		{
			trader: "trader1",
			price:  pr,
			size:   10,
		},
		{
			trader: "trader2",
			price:  pr,
			size:   -5,
		},
		{
			trader: "trader3",
			price:  pr,
			size:   -5,
		},
	}
	half := num.NewUint(500)
	expect := []*types.Transfer{
		{
			Owner: data[1].trader,
			Amount: &types.FinancialAmount{
				Amount: half,
			},
			Type: types.TransferType_TRANSFER_TYPE_LOSS,
		},
		{
			Owner: data[2].trader,
			Amount: &types.FinancialAmount{
				Amount: half,
			},
			Type: types.TransferType_TRANSFER_TYPE_LOSS,
		},
		{
			Owner: data[0].trader,
			Amount: &types.FinancialAmount{
				Amount: pr,
			},
			Type: types.TransferType_TRANSFER_TYPE_WIN,
		},
	} // }}}
	oraclePrice := num.NewUint(1100)
	settleF := func(price *num.Uint, size int64) (*types.FinancialAmount, error) {
		if size < 0 {
			size *= -1
		}
		amt := num.Zero().Sub(oraclePrice, price)
		amt = amt.Mul(amt, num.NewUint(uint64(size)))
		return &types.FinancialAmount{
			Amount: amt,
		}, nil
	}
	positions := engine.getExpiryPositions(data...)
	for _, d := range data {
		// we expect settle calls for each position
		engine.prod.EXPECT().Settle(d.price, d.size).Times(1).DoAndReturn(settleF)
	}
	// ensure positions are set
	engine.Update(positions)
	// now settle:
	got, err := engine.Settle(time.Now(), num.Zero())
	assert.NoError(t, err)
	assert.Equal(t, len(expect), len(got))
	for i, p := range got {
		e := expect[i]
		assert.Equal(t, e.Type, p.Type)
		assert.Equal(t, e.Amount.Amount, p.Amount.Amount)
	}
}

func testSettleExpiredSuccessWithMarkPrice(t *testing.T) {
	engine := getTestEngine(t)
	defer engine.Finish()
	// these are mark prices, product will provide the actual value
	pr := num.NewUint(1000)
	data := []posValue{ // {{{2
		{
			trader: "trader1",
			price:  pr,
			size:   10,
		},
		{
			trader: "trader2",
			price:  pr,
			size:   -5,
		},
		{
			trader: "trader3",
			price:  pr,
			size:   -5,
		},
	}
	half := num.NewUint(500)
	expect := []*types.Transfer{
		{
			Owner: data[1].trader,
			Amount: &types.FinancialAmount{
				Amount: half,
			},
			Type: types.TransferType_TRANSFER_TYPE_LOSS,
		},
		{
			Owner: data[2].trader,
			Amount: &types.FinancialAmount{
				Amount: half,
			},
			Type: types.TransferType_TRANSFER_TYPE_LOSS,
		},
		{
			Owner: data[0].trader,
			Amount: &types.FinancialAmount{
				Amount: pr,
			},
			Type: types.TransferType_TRANSFER_TYPE_WIN,
		},
	} // }}}

	// settlement price at markPrice
	markPrice := num.NewUint(1100)
	// set the FinalSettlement to the MarkPrice method
	engine.Engine.Config.FinalSettlement.FinalSettlement = settlement.FinalSettlementMarkPrice

	positions := engine.getExpiryPositions(data...)
	engine.prod.EXPECT().GetAsset().Return("ETH").AnyTimes()
	// ensure positions are set
	engine.Update(positions)
	// now settle:
	got, err := engine.Settle(time.Now(), markPrice)
	assert.NoError(t, err)
	assert.Equal(t, len(expect), len(got))
	for i, p := range got {
		e := expect[i]
		assert.Equal(t, e.Type, p.Type)
		assert.Equal(t, e.Amount.Amount, p.Amount.Amount)
	}
}

func testSettleExpiredSuccessErrorInvalidSettlementMethod(t *testing.T) {
	engine := getTestEngine(t)
	defer engine.Finish()
	// settlement price at markPrice
	markPrice := num.NewUint(1100)
	// set the FinalSettlement to the MarkPrice method
	engine.Engine.Config.FinalSettlement.FinalSettlement = "not a settlement"
	// now settle:
	_, err := engine.Settle(time.Now(), markPrice)
	assert.Error(t, err)
}

func testSettleExpiryFail(t *testing.T) {
	engine := getTestEngine(t)
	defer engine.Finish()
	// these are mark prices, product will provide the actual value
	data := []posValue{
		{
			trader: "trader1",
			price:  num.NewUint(1000),
			size:   10,
		},
	}
	errExp := errors.New("product.Settle error")
	positions := engine.getExpiryPositions(data...)
	engine.prod.EXPECT().Settle(data[0].price, data[0].size).Times(1).Return(nil, errExp)
	engine.Update(positions)
	empty, err := engine.Settle(time.Now(), num.Zero())
	assert.Empty(t, empty)
	assert.Error(t, err)
	assert.Equal(t, errExp, err)
}

func testMarkToMarketEmpty(t *testing.T) {
	markPrice := num.NewUint(10000)
	// there's only 1 trade to test here
	data := posValue{
		price:  markPrice,
		size:   1,
		trader: "test",
	}
	engine := getTestEngine(t)
	defer engine.Finish()
	pos := mocks.NewMockMarketPosition(engine.ctrl)
	pos.EXPECT().Party().AnyTimes().Return(data.trader)
	pos.EXPECT().Size().AnyTimes().Return(data.size)
	pos.EXPECT().Price().AnyTimes().Return(markPrice)
	engine.Update([]events.MarketPosition{pos})
	result := engine.SettleMTM(context.Background(), markPrice, []events.MarketPosition{pos})
	assert.Empty(t, result)
}

func testAddNewTraderSelfTrade(t *testing.T) {
	engine := getTestEngine(t)
	defer engine.Finish()
	markPrice := num.NewUint(1000)
	t1 := testPos{
		price: markPrice.Clone(),
		party: "trader1",
		size:  5,
	}
	init := []events.MarketPosition{
		t1,
		testPos{
			price: markPrice.Clone(),
			party: "trader2",
			size:  -5,
		},
	}
	// let's not change the markPrice
	// just add a trader to the market, buying from an existing trader
	trade := &types.Trade{
		Buyer:  "trader3",
		Seller: "trader3",
		Price:  markPrice.Clone(),
		Size:   1,
	}
	// the first trader is the seller
	// so these are the new positions after the trade
	t1.size -= 1
	positions := []events.MarketPosition{
		t1,
		init[1],
		testPos{
			party: "trader3",
			size:  0,
			price: markPrice.Clone(),
		},
	}
	engine.Update(init)
	engine.AddTrade(trade)
	noTransfers := engine.SettleMTM(context.Background(), markPrice, positions)
	assert.Len(t, noTransfers, 1)
	assert.Nil(t, noTransfers[0].Transfer())
}

func testAddNewTrader(t *testing.T) {
	engine := getTestEngine(t)
	defer engine.Finish()
	markPrice := num.NewUint(1000)
	t1 := testPos{
		price: markPrice.Clone(),
		party: "trader1",
		size:  5,
	}
	init := []events.MarketPosition{
		t1,
		testPos{
			price: markPrice.Clone(),
			party: "trader2",
			size:  -5,
		},
	}
	// let's not change the markPrice
	// just add a trader to the market, buying from an existing trader
	trade := &types.Trade{
		Buyer:  "trader3",
		Seller: t1.party,
		Price:  markPrice.Clone(),
		Size:   1,
	}
	// the first trader is the seller
	// so these are the new positions after the trade
	t1.size -= 1
	positions := []events.MarketPosition{
		t1,
		init[1],
		testPos{
			party: "trader3",
			size:  1,
			price: markPrice.Clone(),
		},
	}
	engine.Update(init)
	engine.AddTrade(trade)
	noTransfers := engine.SettleMTM(context.Background(), markPrice, positions)
	assert.Len(t, noTransfers, 2)
	for _, v := range noTransfers {
		assert.Nil(t, v.Transfer())
	}
}

// This tests MTM results put losses first, trades tested are Long going longer, short going shorter
// and long going short, short going long, and a third trader who's not trading at all
func testMarkToMarketOrdered(t *testing.T) {
	engine := getTestEngine(t)
	defer engine.Finish()
	pr := num.NewUint(10000)
	positions := []posValue{
		{
			price:  pr,
			size:   1,
			trader: "trader1", // mocks will create 2 traders (long & short)
		},
		{
			price:  pr.Clone(),
			size:   -1,
			trader: "trader2",
		},
	}
	markPrice := pr.Clone()
	markPrice = markPrice.Add(markPrice, num.NewUint(1000))
	neutral := testPos{
		party: "neutral",
		size:  5,
		price: pr.Clone(),
	}
	init := []events.MarketPosition{
		neutral,
		testPos{
			price: neutral.price.Clone(),
			party: "trader1",
			size:  1,
		},
		testPos{
			price: neutral.price.Clone(),
			party: "trader2",
			size:  -1,
		},
	}
	short, long := make([]events.MarketPosition, 0, 3), make([]events.MarketPosition, 0, 3)
	// the SettleMTM data must contain the new mark price already
	neutral.price = markPrice.Clone()
	short = append(short, neutral)
	long = append(long, neutral)
	// we have a long and short trade example
	trades := map[string]*types.Trade{
		"long": {
			Price: markPrice,
			Size:  1,
		},
		// to go short, the trade has to be 2
		"short": {
			Price: markPrice,
			Size:  2,
		},
	}
	// creates trades and event slices we'll be needing later on
	for _, p := range positions {
		if p.size > 0 {
			trades["long"].Buyer = p.trader
			trades["short"].Seller = p.trader
			long = append(long, testPos{
				party: p.trader,
				price: markPrice.Clone(),
				size:  p.size + int64(trades["long"].Size),
			})
			short = append(short, testPos{
				party: p.trader,
				price: markPrice.Clone(),
				size:  p.size - int64(trades["short"].Size),
			})
		} else {
			trades["long"].Seller = p.trader
			trades["short"].Buyer = p.trader
			long = append(long, testPos{
				party: p.trader,
				price: markPrice.Clone(),
				size:  p.size - int64(trades["long"].Size),
			})
			short = append(short, testPos{
				party: p.trader,
				price: markPrice.Clone(),
				size:  p.size + int64(trades["short"].Size),
			})
		}
	}
	updates := map[string][]events.MarketPosition{
		"long":  long,
		"short": short,
	}
	// set up the engine, ready to run the scenario's
	// for each data-set we reset the state in the engine, then we check the MTM is performed
	// correctly
	for k, trade := range trades {
		engine.Update(init)
		engine.AddTrade(trade)
		update := updates[k]
		transfers := engine.SettleMTM(context.Background(), markPrice, update)
		assert.NotEmpty(t, transfers)
		assert.Equal(t, 3, len(transfers))
		// start with losses, end with wins
		assert.Equal(t, types.TransferType_TRANSFER_TYPE_MTM_LOSS, transfers[0].Transfer().Type)
		assert.Equal(t, types.TransferType_TRANSFER_TYPE_MTM_WIN, transfers[len(transfers)-1].Transfer().Type)
		assert.Equal(t, "trader2", transfers[0].Party()) // we expect trader2 to have a loss
	}
}

func testMTMNetworkZero(t *testing.T) {
	t.Skip("not implemented yet")
	engine := getTestEngine(t)
	defer engine.Finish()
	markPrice := num.NewUint(1000)
	init := []events.MarketPosition{
		testPos{
			price: markPrice.Clone(),
			party: "trader1",
			size:  5,
		},
		testPos{
			price: markPrice.Clone(),
			party: "trader2",
			size:  -5,
		},
		testPos{
			price: markPrice.Clone(),
			party: "trader3",
			size:  10,
		},
		testPos{
			price: markPrice.Clone(),
			party: "trader4",
			size:  -10,
		},
	}
	// initialise the engine with the positions above
	engine.Update(init)
	// assume trader 4 is distressed, network has to trade and buy 10
	// ensure the network loses in this scenario: the price has gone up
	cPrice := num.Sum(markPrice, num.NewUint(1))
	trade := &types.Trade{
		Buyer:  types.NetworkParty,
		Seller: "trader1",
		Size:   5, // trader 1 only has 5 on the book, let's pretend we can close him our
		Price:  cPrice.Clone(),
	}
	engine.AddTrade(trade)
	engine.AddTrade(&types.Trade{
		Buyer:  types.NetworkParty,
		Seller: "trader3",
		Size:   2,
		Price:  cPrice.Clone(),
	})
	engine.AddTrade(&types.Trade{
		Buyer:  types.NetworkParty,
		Seller: "trader2",
		Size:   3,
		Price:  cPrice.Clone(), // trader 2 is going from -5 to -8
	})
	// the new positions of the traders who have traded with the network...
	positions := []events.MarketPosition{
		testPos{
			party: "trader1", // trader 1 was 5 long, sold 5 to network, so closed out
			price: markPrice.Clone(),
			size:  0,
		},
		testPos{
			party: "trader3",
			size:  8, // long 10, sold 2
			price: markPrice.Clone(),
		},
		testPos{
			party: "trader2",
			size:  -8,
			price: markPrice.Clone(), // trader 2 was -5, shorted an additional 3 => -8
		},
	}
	// new markprice is cPrice
	noTransfers := engine.SettleMTM(context.Background(), cPrice, positions)
	assert.Len(t, noTransfers, 3)
	hasNetwork := false
	for i, v := range noTransfers {
		assert.NotNil(t, v.Transfer())
		if v.Party() == types.NetworkParty {
			// network hás to lose
			require.Equal(t, types.TransferType_TRANSFER_TYPE_MTM_LOSS, v.Transfer().Type)
			// network loss should be at the start of the slice
			require.Equal(t, 0, i)
			hasNetwork = true
		}
	}
	require.True(t, hasNetwork)
}

// {{{
func (te *testEngine) getExpiryPositions(positions ...posValue) []events.MarketPosition {
	te.positions = make([]*mocks.MockMarketPosition, 0, len(positions))
	mpSlice := make([]events.MarketPosition, 0, len(positions))
	for _, p := range positions {
		pos := mocks.NewMockMarketPosition(te.ctrl)
		// these values should only be obtained once, and assigned internally
		pos.EXPECT().Party().MinTimes(1).AnyTimes().Return(p.trader)
		pos.EXPECT().Size().MinTimes(1).AnyTimes().Return(p.size)
		pos.EXPECT().Price().Times(1).Return(p.price)
		te.positions = append(te.positions, pos)
		mpSlice = append(mpSlice, pos)
	}
	return mpSlice
}

func (te *testEngine) getMockMarketPositions(data []posValue) ([]settlement.MarketPosition, []events.MarketPosition) {
	raw, evts := make([]settlement.MarketPosition, 0, len(data)), make([]events.MarketPosition, 0, len(data))
	for _, pos := range data {
		mock := mocks.NewMockMarketPosition(te.ctrl)
		mock.EXPECT().Party().MinTimes(1).Return(pos.trader)
		mock.EXPECT().Size().MinTimes(1).Return(pos.size)
		mock.EXPECT().Price().MinTimes(1).Return(pos.price)
		raw = append(raw, mock)
		evts = append(evts, mock)
	}
	return raw, evts
}

func TestConcurrent(t *testing.T) {
	const N = 10

	engine := getTestEngine(t)
	defer engine.Finish()
	engine.prod.EXPECT().Settle(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(price *num.Uint, size int64) (*types.FinancialAmount, error) {
		return &types.FinancialAmount{Amount: num.Zero()}, nil
	})

	cfg := engine.Config
	cfg.Level.Level = logging.DebugLevel
	engine.ReloadConf(cfg)
	cfg.Level.Level = logging.InfoLevel
	engine.ReloadConf(cfg)

	var wg sync.WaitGroup

	now := time.Now()
	wg.Add(N * 3)
	for i := 0; i < N; i++ {
		data := []posValue{
			{
				trader: "testtrader1",
				price:  num.NewUint(1234),
				size:   100,
			},
			{
				trader: "testtrader2",
				price:  num.NewUint(1235),
				size:   0,
			},
		}
		raw, evts := engine.getMockMarketPositions(data)
		// margin evt
		marginEvts := make([]events.Margin, 0, len(raw))
		for _, pe := range raw {
			marginEvts = append(marginEvts, marginVal{
				MarketPosition: pe,
			})
		}

		go func() {
			defer wg.Done()
			// Update requires posMu
			engine.Update(evts)
		}()
		go func() {
			defer wg.Done()
			// RemoveDistressed requires posMu and closedMu
			engine.RemoveDistressed(context.Background(), marginEvts)
		}()
		go func() {
			defer wg.Done()
			// Settle requires posMu
			_, err := engine.Settle(now, num.Zero())
			assert.NoError(t, err)
		}()
	}

	wg.Wait()
}

// Finish - call finish on controller, remove test state (positions)
func (te *testEngine) Finish() {
	te.ctrl.Finish()
	te.positions = nil
}

// Quick mock implementation of the events.MarketPosition interface
type testPos struct {
	party           string
	size, buy, sell int64
	price           *num.Uint
	vwBuy, vwSell   uint64
}

func (t testPos) Party() string {
	return t.party
}

func (t testPos) Size() int64 {
	return t.size
}

func (t testPos) Buy() int64 {
	return t.buy
}

func (t testPos) Sell() int64 {
	return t.sell
}

func (t testPos) Price() *num.Uint {
	if t.price == nil {
		return num.Zero()
	}
	return t.price
}

func (t testPos) VWBuy() *num.Uint {
	return num.NewUint(t.vwBuy)
}

func (t testPos) VWSell() *num.Uint {
	return num.NewUint(t.vwSell)
}

func (t testPos) ClearPotentials() {}

func getTestEngine(t *testing.T) *testEngine {
	ctrl := gomock.NewController(t)
	conf := settlement.NewDefaultConfig()
	prod := mocks.NewMockProduct(ctrl)
	broker := bmock.NewMockBroker(ctrl)
	broker.EXPECT().SendBatch(gomock.Any()).AnyTimes()
	market := "BTC/DEC19"
	prod.EXPECT().GetAsset().AnyTimes().Do(func() string { return "BTC" })
	return &testEngine{
		Engine:    settlement.New(logging.NewTestLogger(), conf, prod, market, broker),
		ctrl:      ctrl,
		prod:      prod,
		broker:    broker,
		positions: nil,
		market:    market,
	}
} // }}}

func (m marginVal) Asset() string {
	return m.asset
}

func (m marginVal) MarketID() string {
	return m.marketID
}

func (m marginVal) MarginBalance() *num.Uint {
	return num.NewUint(m.margin)
}

func (m marginVal) GeneralBalance() *num.Uint {
	return num.NewUint(m.general)
}

func (m marginVal) BondBalance() *num.Uint {
	return num.Zero()
}

func (m marginVal) MarginShortFall() *num.Uint {
	return num.NewUint(m.marginShortFall)
}

//  vim: set ts=4 sw=4 tw=0 foldlevel=1 foldmethod=marker noet :
