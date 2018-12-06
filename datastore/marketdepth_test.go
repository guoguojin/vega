package datastore

import (
	"fmt"
	"math/rand"
	"testing"
	"vega/msg"
	"github.com/stretchr/testify/assert"
)

func TestOrderBookDepth_All(t *testing.T){

	var marketDepth MarketDepth

	ordersList := []*msg.Order{
		{Side: msg.Side_Buy,Price: 116, Remaining: 100},
		{Side: msg.Side_Buy,Price: 110, Remaining: 100},
		{Side: msg.Side_Buy,Price: 111, Remaining: 100},
		{Side: msg.Side_Buy,Price: 111, Remaining: 100},
		{Side: msg.Side_Buy,Price: 113, Remaining: 100},
		{Side: msg.Side_Buy,Price: 114, Remaining: 100},
		{Side: msg.Side_Buy,Price: 116, Remaining: 100},
	}

	for _, elem := range ordersList {
		marketDepth.Add(elem)
	}

	assert.Equal(t, marketDepth.Buy[0].Price, uint64(116))
	assert.Equal(t, marketDepth.Buy[0].Volume, uint64(200))
	assert.Equal(t, marketDepth.Buy[0].NumberOfOrders, uint64(2))
	assert.Equal(t, marketDepth.Buy[0].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Buy[1].Price, uint64(114))
	assert.Equal(t, marketDepth.Buy[1].Volume, uint64(100))
	assert.Equal(t, marketDepth.Buy[1].NumberOfOrders, uint64(1))
	assert.Equal(t, marketDepth.Buy[1].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Buy[2].Price, uint64(113))
	assert.Equal(t, marketDepth.Buy[2].Volume, uint64(100))
	assert.Equal(t, marketDepth.Buy[2].NumberOfOrders, uint64(1))
	assert.Equal(t, marketDepth.Buy[2].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Buy[3].Price, uint64(111))
	assert.Equal(t, marketDepth.Buy[3].Volume, uint64(200))
	assert.Equal(t, marketDepth.Buy[3].NumberOfOrders, uint64(2))
	assert.Equal(t, marketDepth.Buy[3].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Buy[4].Price, uint64(110))
	assert.Equal(t, marketDepth.Buy[4].Volume, uint64(100))
	assert.Equal(t, marketDepth.Buy[4].NumberOfOrders, uint64(1))
	assert.Equal(t, marketDepth.Buy[4].CumulativeVolume, uint64(0))

	marketDepth.DecreaseByTradedVolume(&msg.Order{Side: msg.Side_Buy,Price: 111, Remaining: 50}, 50)
	marketDepth.DecreaseByTradedVolume(&msg.Order{Side: msg.Side_Buy,Price: 114, Remaining: 80}, 20)
	marketDepth.DecreaseByTradedVolume(&msg.Order{Side: msg.Side_Buy,Price: 113, Remaining: 100}, 100)

	assert.Equal(t, marketDepth.Buy[0].Price, uint64(116))
	assert.Equal(t, marketDepth.Buy[0].Volume, uint64(200))
	assert.Equal(t, marketDepth.Buy[0].NumberOfOrders, uint64(2))
	assert.Equal(t, marketDepth.Buy[0].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Buy[1].Price, uint64(114))
	assert.Equal(t, marketDepth.Buy[1].Volume, uint64(80))
	assert.Equal(t, marketDepth.Buy[1].NumberOfOrders, uint64(1))
	assert.Equal(t, marketDepth.Buy[1].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Buy[2].Price, uint64(111))
	assert.Equal(t, marketDepth.Buy[2].Volume, uint64(150))
	assert.Equal(t, marketDepth.Buy[2].NumberOfOrders, uint64(2))
	assert.Equal(t, marketDepth.Buy[2].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Buy[3].Price, uint64(110))
	assert.Equal(t, marketDepth.Buy[3].Volume, uint64(100))
	assert.Equal(t, marketDepth.Buy[3].NumberOfOrders, uint64(1))
	assert.Equal(t, marketDepth.Buy[3].CumulativeVolume, uint64(0))


	// test sell side
	ordersList = []*msg.Order{
		{Side: msg.Side_Sell,Price: 123, Remaining: 100},
		{Side: msg.Side_Sell,Price: 119, Remaining: 100},
		{Side: msg.Side_Sell,Price: 120, Remaining: 100},
		{Side: msg.Side_Sell,Price: 120, Remaining: 100},
		{Side: msg.Side_Sell,Price: 121, Remaining: 100},
		{Side: msg.Side_Sell,Price: 121, Remaining: 100},
		{Side: msg.Side_Sell,Price: 122, Remaining: 100},
		{Side: msg.Side_Sell,Price: 123, Remaining: 100},
	}

	for _, elem := range ordersList {
		marketDepth.Add(elem)
	}

	assert.Equal(t, marketDepth.Sell[0].Price, uint64(119))
	assert.Equal(t, marketDepth.Sell[0].Volume, uint64(100))
	assert.Equal(t, marketDepth.Sell[0].NumberOfOrders, uint64(1))
	assert.Equal(t, marketDepth.Sell[0].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Sell[1].Price, uint64(120))
	assert.Equal(t, marketDepth.Sell[1].Volume, uint64(200))
	assert.Equal(t, marketDepth.Sell[1].NumberOfOrders, uint64(2))
	assert.Equal(t, marketDepth.Sell[1].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Sell[2].Price, uint64(121))
	assert.Equal(t, marketDepth.Sell[2].Volume, uint64(200))
	assert.Equal(t, marketDepth.Sell[2].NumberOfOrders, uint64(2))
	assert.Equal(t, marketDepth.Sell[2].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Sell[3].Price, uint64(122))
	assert.Equal(t, marketDepth.Sell[3].Volume, uint64(100))
	assert.Equal(t, marketDepth.Sell[3].NumberOfOrders, uint64(1))
	assert.Equal(t, marketDepth.Sell[3].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Sell[4].Price, uint64(123))
	assert.Equal(t, marketDepth.Sell[4].Volume, uint64(200))
	assert.Equal(t, marketDepth.Sell[4].NumberOfOrders, uint64(2))
	assert.Equal(t, marketDepth.Sell[4].CumulativeVolume, uint64(0))

	marketDepth.DecreaseByTradedVolume(&msg.Order{Side: msg.Side_Sell,Price: 119, Remaining: 100}, 50)
	marketDepth.DecreaseByTradedVolume(&msg.Order{Side: msg.Side_Sell,Price: 120, Remaining: 100}, 20)
	marketDepth.DecreaseByTradedVolume(&msg.Order{Side: msg.Side_Sell,Price: 122, Remaining: 100}, 100)

	assert.Equal(t, marketDepth.Sell[0].Price, uint64(119))
	assert.Equal(t, marketDepth.Sell[0].Volume, uint64(50))
	assert.Equal(t, marketDepth.Sell[0].NumberOfOrders, uint64(1))
	assert.Equal(t, marketDepth.Sell[0].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Sell[1].Price, uint64(120))
	assert.Equal(t, marketDepth.Sell[1].Volume, uint64(180))
	assert.Equal(t, marketDepth.Sell[1].NumberOfOrders, uint64(2))
	assert.Equal(t, marketDepth.Sell[1].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Sell[2].Price, uint64(121))
	assert.Equal(t, marketDepth.Sell[2].Volume, uint64(200))
	assert.Equal(t, marketDepth.Sell[2].NumberOfOrders, uint64(2))
	assert.Equal(t, marketDepth.Sell[2].CumulativeVolume, uint64(0))

	assert.Equal(t, marketDepth.Sell[3].Price, uint64(123))
	assert.Equal(t, marketDepth.Sell[3].Volume, uint64(200))
	assert.Equal(t, marketDepth.Sell[3].NumberOfOrders, uint64(2))
	assert.Equal(t, marketDepth.Sell[3].CumulativeVolume, uint64(0))
}


func TestOrderBookDepthBuySide(t *testing.T) {
	// test orderbook depth

	// Scenario:

	// POST few orders to datastore
	// call getMarketDepth and see if order book depth is OK

	// create impacted orders and call PUT on them
	// call getMarketDepth and see if order book depth is OK

	// call DELETE on orders
	// call getMarketDepth and see if order book depth is OK

	//var memStore = NewMemStore([]string{testMarket}, []string{testParty, testPartyA, testPartyB})
	var newOrderStore = NewOrderStore("../tmp/orderstore")
	defer newOrderStore.Close()

	orders := []*msg.Order{
		{
			Id:     fmt.Sprintf("%d", rand.Intn(1000000000000)),
			Side: msg.Side_Buy,
			Market: testMarket,
			Party: testPartyA,
			Price: 111,
			Remaining: 100,
		},
		{
			Id:     fmt.Sprintf("%d", rand.Intn(1000000000000)),
			Side: msg.Side_Buy,
			Market: testMarket,
			Party: testPartyA,
			Price: 112,
			Remaining: 100,
		},
		{
			Id:     fmt.Sprintf("%d", rand.Intn(1000000000000)),
			Side: msg.Side_Buy,
			Market: testMarket,
			Party: testPartyA,
			Price: 112,
			Remaining: 100,
		},
		{
			Id:     fmt.Sprintf("%d", rand.Intn(1000000000000)),
			Side: msg.Side_Buy,
			Market: testMarket,
			Party: testPartyA,
			Price: 113,
			Remaining: 100,
		},
	}

	for idx, _ := range orders {
		newOrderStore.Post(orders[idx])
	}

	newOrderStore.Commit()

	marketDepth, _ := newOrderStore.GetMarketDepth(testMarket)

	assert.Equal(t, uint64(113), marketDepth.Buy[0].Price)
	assert.Equal(t, uint64(100), marketDepth.Buy[0].Volume)
	assert.Equal(t, uint64(1), marketDepth.Buy[0].NumberOfOrders)
	assert.Equal(t, uint64(100), marketDepth.Buy[0].CumulativeVolume)

	assert.Equal(t, uint64(112), marketDepth.Buy[1].Price)
	assert.Equal(t, uint64(200), marketDepth.Buy[1].Volume)
	assert.Equal(t, uint64(2), marketDepth.Buy[1].NumberOfOrders)
	assert.Equal(t, uint64(300), marketDepth.Buy[1].CumulativeVolume)

	assert.Equal(t, uint64(111), marketDepth.Buy[2].Price)
	assert.Equal(t, uint64(100), marketDepth.Buy[2].Volume)
	assert.Equal(t, uint64(1), marketDepth.Buy[2].NumberOfOrders)
	assert.Equal(t, uint64(400), marketDepth.Buy[2].CumulativeVolume)

	ordersUpdate := []*msg.Order{
		{
			Id:     orders[0].Id,
			Side: msg.Side_Buy,
			Market: testMarket,
			Party: testPartyA,
			Price: 111,
			Remaining: 50,
		},
		{
			Id:     orders[2].Id,
			Side: msg.Side_Buy,
			Market: testMarket,
			Party: testPartyA,
			Price: 112,
			Remaining: 50,
		},
		{
			Id:    orders[3].Id,
			Side: msg.Side_Buy,
			Market: testMarket,
			Party: testPartyA,
			Price: 113,
			Remaining: 80,
			Status: msg.Order_Expired,
		},
	}

	for idx, _ := range ordersUpdate {
		newOrderStore.Put(ordersUpdate[idx])
	}

	marketDepth, _ = newOrderStore.GetMarketDepth(testMarket)

	// 113 is removed

	assert.Equal(t, uint64(112), marketDepth.Buy[0].Price)
	assert.Equal(t, uint64(150), marketDepth.Buy[0].Volume)
	assert.Equal(t, uint64(2), marketDepth.Buy[0].NumberOfOrders)
	assert.Equal(t, uint64(150), marketDepth.Buy[0].CumulativeVolume)

	assert.Equal(t, uint64(111), marketDepth.Buy[1].Price)
	assert.Equal(t, uint64(50), marketDepth.Buy[1].Volume)
	assert.Equal(t, uint64(1), marketDepth.Buy[1].NumberOfOrders)
	assert.Equal(t, uint64(200), marketDepth.Buy[1].CumulativeVolume)

	ordersRemove := []*msg.Order{
		{
			Id:     orders[0].Id,
			Side: msg.Side_Buy,
			Market: testMarket,
			Party: testPartyA,
			Price: 111,
			Remaining: 50,
		},
		{
			Id:     orders[1].Id,
			Side: msg.Side_Buy,
			Market: testMarket,
			Party: testPartyA,
			Price: 112,
			Remaining: 100,
		},
	}

	for idx, _ := range ordersRemove {
		newOrderStore.Delete(ordersRemove[idx])
	}

	marketDepth, _ = newOrderStore.GetMarketDepth(testMarket)

	fmt.Printf("marketDepth %+v\n", marketDepth.Buy)

	assert.Equal(t, uint64(112), marketDepth.Buy[0].Price)
	assert.Equal(t, uint64(50), marketDepth.Buy[0].Volume)
	assert.Equal(t, uint64(1), marketDepth.Buy[0].NumberOfOrders)
	assert.Equal(t, uint64(50), marketDepth.Buy[0].CumulativeVolume)

	assert.Equal(t, 1, len(marketDepth.Buy))
}

func TestOrderBookDepthSellSide(t *testing.T) {
	// test orderbook depth

	// Scenario:

	// POST few orders to datastore
	// call getMarketDepth and see if order book depth is OK

	// create impacted orders and call PUT on them
	// call getMarketDepth and see if order book depth is OK

	// call DELETE on orders
	// call getMarketDepth and see if order book depth is OK

	//var memStore = NewMemStore([]string{testMarket}, []string{testParty, testPartyA, testPartyB})
	var newOrderStore = NewOrderStore("../tmp/orderStore")
	defer newOrderStore.Close()

	orders := []*msg.Order{
		{
			Id:     fmt.Sprintf("%d", rand.Intn(1000000000000)),
			Side: msg.Side_Sell,
			Market: testMarket,
			Party: testPartyA,
			Price: 111,
			Remaining: 100,
		},
		{
			Id:     fmt.Sprintf("%d", rand.Intn(1000000000000)),
			Side: msg.Side_Sell,
			Market: testMarket,
			Party: testPartyA,
			Price: 112,
			Remaining: 100,
		},
		{
			Id:     fmt.Sprintf("%d", rand.Intn(1000000000000)),
			Side: msg.Side_Sell,
			Market: testMarket,
			Party: testPartyA,
			Price: 112,
			Remaining: 100,
		},
		{
			Id:     fmt.Sprintf("%d", rand.Intn(1000000000000)),
			Side: msg.Side_Sell,
			Market: testMarket,
			Party: testPartyA,
			Price: 113,
			Remaining: 100,
		},
	}

	for idx, _ := range orders {
		newOrderStore.Post(orders[idx])
	}

	newOrderStore.Commit()

	marketDepth, _ := newOrderStore.GetMarketDepth(testMarket)

	assert.Equal(t, uint64(111), marketDepth.Sell[0].Price)
	assert.Equal(t, uint64(100), marketDepth.Sell[0].Volume)
	assert.Equal(t, uint64(1), marketDepth.Sell[0].NumberOfOrders)
	assert.Equal(t, uint64(100), marketDepth.Sell[0].CumulativeVolume)

	assert.Equal(t, uint64(112), marketDepth.Sell[1].Price)
	assert.Equal(t, uint64(200), marketDepth.Sell[1].Volume)
	assert.Equal(t, uint64(2), marketDepth.Sell[1].NumberOfOrders)
	assert.Equal(t, uint64(300), marketDepth.Sell[1].CumulativeVolume)

	assert.Equal(t, uint64(113), marketDepth.Sell[2].Price)
	assert.Equal(t, uint64(100), marketDepth.Sell[2].Volume)
	assert.Equal(t, uint64(1), marketDepth.Sell[2].NumberOfOrders)
	assert.Equal(t, uint64(400), marketDepth.Sell[2].CumulativeVolume)

	ordersUpdate := []*msg.Order{
		{
			Id:     orders[0].Id,
			Side: msg.Side_Sell,
			Market: testMarket,
			Party: testPartyA,
			Price: 111,
			Remaining: 50,
		},
		{
			Id:     orders[2].Id,
			Side: msg.Side_Sell,
			Market: testMarket,
			Party: testPartyA,
			Price: 112,
			Remaining: 50,
		},
		{
			Id:    orders[3].Id,
			Side: msg.Side_Sell,
			Market: testMarket,
			Party: testPartyA,
			Price: 113,
			Remaining: 80,
			Status: msg.Order_Expired,
		},
	}

	for idx, _ := range ordersUpdate {
		newOrderStore.Put(ordersUpdate[idx])
	}

	marketDepth, _ = newOrderStore.GetMarketDepth(testMarket)

	assert.Equal(t, uint64(111), marketDepth.Sell[0].Price)
	assert.Equal(t, uint64(50), marketDepth.Sell[0].Volume)
	assert.Equal(t, uint64(1), marketDepth.Sell[0].NumberOfOrders)
	assert.Equal(t, uint64(50), marketDepth.Sell[0].CumulativeVolume)

	assert.Equal(t, uint64(112), marketDepth.Sell[1].Price)
	assert.Equal(t, uint64(150), marketDepth.Sell[1].Volume)
	assert.Equal(t, uint64(2), marketDepth.Sell[1].NumberOfOrders)
	assert.Equal(t, uint64(200), marketDepth.Sell[1].CumulativeVolume)

	// 113 is removed
	assert.Equal(t, 2, len(marketDepth.Sell))

	ordersRemove := []*msg.Order{
		{
			Id:     orders[0].Id,
			Side: msg.Side_Sell,
			Market: testMarket,
			Party: testPartyA,
			Price: 111,
			Remaining: 50,
		},
		{
			Id:     orders[1].Id,
			Side: msg.Side_Sell,
			Market: testMarket,
			Party: testPartyA,
			Price: 112,
			Remaining: 100,
		},
	}

	for idx, _ := range ordersRemove {
		newOrderStore.Delete(ordersRemove[idx])
	}

	marketDepth, _ = newOrderStore.GetMarketDepth(testMarket)

	assert.Equal(t, uint64(112), marketDepth.Sell[0].Price)
	assert.Equal(t, uint64(50), marketDepth.Sell[0].Volume)
	assert.Equal(t, uint64(1), marketDepth.Sell[0].NumberOfOrders)
	assert.Equal(t, uint64(50), marketDepth.Sell[0].CumulativeVolume)

	assert.Equal(t, 1, len(marketDepth.Sell))
}

