package datastore

import (
	"testing"
	"vega/msg"
	"github.com/stretchr/testify/assert"
)

//
//import (
//	"testing"
//	"vega/log"
//	"vega/msg"
//	"github.com/stretchr/testify/assert"
//	"vega/filters"
//	"fmt"
//)
//
//const testMarket = "market"
//const testParty = "party"
//const testPartyA = "partyA"
//const testPartyB = "partyB"
//
//// this runs just once as first
//func init() {
//	log.InitConsoleLogger(log.DebugLevel)
//}
//
//func TestNewMemStore_ReturnsNewMemStoreInstance(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	assert.NotNil(t, memStore)
//}
//
//func TestNewMemStore_ReturnsNewTradeStoreInstance(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newTradeStore = NewTradeStore(&memStore)
//	assert.NotNil(t, newTradeStore)
//}
//
//func TestNewMemStore_ReturnsNewOrderStoreInstance(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	assert.NotNil(t, newOrderStore)
//}
//
func TestMemStore_PostAndGetNewOrder(t *testing.T) {
	newOrderStore := NewOrderStore("./orderStore")
	defer newOrderStore.Close()

	var order = &msg.Order{
		Id:     "45305210ff7a9bb9450b1833cc10368a",
		Market: "testMarket",
		Party:  "testParty",
	}

	err := newOrderStore.Post(order)
	assert.Nil(t, err)

	o, err := newOrderStore.GetByMarketAndId("testMarket", order.Id)
	assert.Nil(t, err)
	assert.Equal(t, order, o)
}
//
//func TestMemStore_PostDuplicateOrder(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//
//	var order = Order{
//		Order: msg.Order{
//			Id:     "45305210ff7a9bb9450b1833cc10368a",
//			Market: testMarket,
//			Party:  testParty,
//		},
//	}
//
//	err := newOrderStore.Post(order)
//	assert.Nil(t, err)
//	err = newOrderStore.Post(order)
//	assert.Error(t, err, "order exists in store")
//}
//
//func TestMemStore_PostOrderToNoneExistentMarket(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	var order = Order{
//		Order: msg.Order{
//			Id:     "45305210ff7a9bb9450b1833cc10368a",
//			Market: "GBP/EUR19",
//		},
//	}
//	err := newOrderStore.Post(order)
//	assert.Error(t, err, "market does not exist")
//}
//
//func TestMemStore_PostPutAndGetExistingOrder(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//
//	var order = Order{
//		Order: msg.Order{
//			Id:     "c471bdd5f381aa3654d98f4591eaa968",
//			Market: testMarket,
//			Party:  testParty,
//			Price:  100,
//			Size:   1,
//			Status: msg.Order_Active,
//		},
//	}
//
//	err := newOrderStore.Post(order)
//	assert.Nil(t, err)
//
//	o, err := newOrderStore.GetByMarketAndId(testMarket, order.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, uint64(100), o.Price)
//	assert.Equal(t, uint64(1), o.Size)
//
//	order.Price = 1000
//	order.Size = 5
//	order.Status = msg.Order_Cancelled
//
//	err = newOrderStore.Put(order)
//	assert.Nil(t, err)
//
//	o, err = newOrderStore.GetByMarketAndId(testMarket, order.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, order, o)
//	assert.Equal(t, uint64(1000), o.Price)
//	assert.Equal(t, uint64(5), o.Size)
//	assert.Equal(t, msg.Order_Cancelled, o.Status)
//}
//
//func TestMemStore_PutNoneExistentOrder(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	var order = Order{
//		Order: msg.Order{
//			Id:     "45305210ff7a9bb9450b1833cc10368a",
//			Market: testMarket,
//			Party:  testParty,
//		},
//	}
//	err := newOrderStore.Put(order)
//	assert.Error(t, err, "order not found in store")
//}
//
//func TestMemStore_PutOrderToNoneExistentMarket(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	var order = Order{
//		Order: msg.Order{
//			Id:     "45305210ff7a9bb9450b1833cc10368a",
//			Market: "GBP/EUR19",
//			Party:  testParty,
//		},
//	}
//	err := newOrderStore.Put(order)
//	assert.Error(t, err, "market does not exist")
//}
//
//func TestMemStore_PostAndDeleteOrder(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//
//	var order = Order{
//		Order: msg.Order{
//			Id:     "45305210ff7a9bb9450b1833cc10368a",
//			Market: testMarket,
//			Party:  testParty,
//		},
//	}
//
//	err := newOrderStore.Post(order)
//	assert.Nil(t, err)
//
//	o, err := newOrderStore.GetByMarketAndId(testMarket, order.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, order, o)
//
//	err = newOrderStore.Delete(o)
//	assert.Nil(t, err)
//
//	o, err = newOrderStore.GetByMarketAndId(testMarket, order.Id)
//	assert.Error(t, err)
//}
//
//func TestMemStore_DeleteOrderFromNoneExistentMarket(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	var order = Order{
//		Order: msg.Order{
//			Id:     "45305210ff7a9bb9450b1833cc10368a",
//			Market: "GBP/EUR19",
//			Party:  testParty,
//		},
//	}
//	err := newOrderStore.Delete(order)
//	assert.Error(t, err, "market does not exist")
//}
//
//func TestMemStore_GetOrderForNoneExistentMarket(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testPartyA, testPartyB})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	_, err := newOrderStore.GetByMarketAndId("UNKNOWN", "ID")
//	assert.Error(t, err, "market does not exist")
//}
//
//func TestMemStore_PostAndGetTrade(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testPartyA, testPartyB})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	var newTradeStore = NewTradeStore(&memStore)
//
//	var trade = Trade{
//		Trade:             msg.Trade{Market: testMarket, Buyer: testPartyA, Seller: testPartyB},
//		AggressiveOrderId: "d41d8cd98f00b204e9800998ecf8427e",
//		PassiveOrderId:    "d41d8cd98f00b204e9800998ecf9999e",
//	}
//
//	var passiveOrder = Order{
//		Order: msg.Order{
//			Id:     "d41d8cd98f00b204e9800998ecf9999e",
//			Market: testMarket,
//			Party:  testPartyB,
//		},
//	}
//
//	var aggressiveOrder = Order{
//		Order: msg.Order{
//			Id:     "d41d8cd98f00b204e9800998ecf8427e",
//			Market: testMarket,
//			Party:  testPartyA,
//		},
//	}
//
//	err := newOrderStore.Post(passiveOrder)
//	assert.Nil(t, err)
//
//	err = newOrderStore.Post(aggressiveOrder)
//	assert.Nil(t, err)
//
//	err = newTradeStore.Post(trade)
//	assert.Nil(t, err)
//
//	tr, err := newTradeStore.GetByMarketAndId(testMarket, trade.Id)
//	assert.Equal(t, trade, tr)
//}
//
//func TestMemStore_PutAndDeleteTrade(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testPartyA, testPartyB})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	var newTradeStore = NewTradeStore(&memStore)
//
//	var passiveOrder = Order{
//		Order: msg.Order{Id: "d41d8cd98f00b204e9800998ecf9999e", Market: testMarket, Party: testPartyA},
//	}
//
//	var aggressiveOrder = Order{
//		Order: msg.Order{Id: "d41d8cd98f00b204e9800998ecf8427e", Market: testMarket, Party: testPartyB},
//	}
//
//	var trade = Trade{
//		Trade:             msg.Trade{Market: testMarket, Buyer: testPartyA, Seller: testPartyB},
//		AggressiveOrderId: "d41d8cd98f00b204e9800998ecf8427e",
//		PassiveOrderId:    "d41d8cd98f00b204e9800998ecf9999e",
//	}
//
//	err := newOrderStore.Post(passiveOrder)
//	assert.Nil(t, err)
//
//	err = newOrderStore.Post(aggressiveOrder)
//	assert.Nil(t, err)
//
//	err = newTradeStore.Post(trade)
//	assert.Nil(t, err)
//
//	tr, err := newTradeStore.GetByMarketAndId(testMarket, trade.Id)
//	assert.Equal(t, trade, tr)
//
//	err = newTradeStore.Delete(tr)
//	assert.Nil(t, err)
//
//	tr, err = newTradeStore.GetByMarketAndId(testMarket, trade.Id)
//	assert.Error(t, err)
//}
//
//func TestMemStore_PostTradeOrderNotFound(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newTradeStore = NewTradeStore(&memStore)
//	trade := Trade{
//		Trade: msg.Trade{
//			Id:     "one",
//			Market: testMarket,
//		},
//		AggressiveOrderId: "mystery",
//		PassiveOrderId:    "d41d8cd98f00b204e9800998ecf9999e",
//	}
//	err := newTradeStore.Post(trade)
//	assert.Error(t, err)
//}
//
////func TestMemStore_PostAndFindByOrderId(t *testing.T) {
////	var memStore = NewMemStore([]string{testMarket})
////	var newOrderStore = NewOrderStore(&memStore)
////	var newTradeStore = NewTradeStore(&memStore)
////
////	trade1 := Trade{
////		Trade: msg.Trade{
////			Id:     "one",
////			Market: testMarket,
////		},
////		OrderId: "d41d8cd98f00b204e9800998ecf8427e",
////	}
////	trade2 := Trade{
////		Trade: msg.Trade{
////			Id:     "two",
////			Market: testMarket,
////		},
////		OrderId: "d41d8cd98f00b204e9800998ecf8427e",
////	}
////	order := Order{
////		Order: msg.Order{
////			Id:     "d41d8cd98f00b204e9800998ecf8427e",
////			Market: testMarket,
////		},
////	}
////
////	err := newOrderStore.Post(order)
////	assert.Nil(t, err)
////
////	err = newTradeStore.Post(trade1)
////	assert.Nil(t, err)
////
////	err = newTradeStore.Post(trade2)
////	assert.Nil(t, err)
////
////	trades, err := newTradeStore.GetByOrderId(testMarket, order.Id, GetParams{Limit: 12345} )
////	assert.Nil(t, err)
////
////	assert.Equal(t, 2, len(trades))
////	assert.Equal(t, "one", trades[0].Id)
////	assert.Equal(t, "two", trades[1].Id)
////}
//
//func TestMemStore_GetAllOrdersForMarket(t *testing.T) {
//
//	var tests = []struct {
//		inMarkets      []string
//		inOrders       []Order
//		inLimit        uint64
//		inMarket       string
//		outOrdersCount int
//	}{
//		{
//			inMarkets: []string{"testMarket1", "marketZ"},
//			inOrders: []Order{
//				{
//					Order: msg.Order{
//						Id:     "d41d8cd98f00b204e9800998ecf8427e",
//						Market: "testMarket1",
//						Party:  testParty,
//					},
//				},
//				{
//					Order: msg.Order{
//						Id:     "ad2dc275947362c45893bbeb30fc3098",
//						Market: "marketZ",
//						Party:  testParty,
//					},
//				},
//				{
//					Order: msg.Order{
//						Id:     "4e8e41367997cfe705d62ea80592cbcc",
//						Market: "testMarket1",
//						Party:  testParty,
//					},
//				},
//			},
//			inLimit:        5000,
//			inMarket:       "testMarket1",
//			outOrdersCount: 2,
//		},
//		{
//			inMarkets: []string{testMarket, "marketABC"},
//			inOrders: []Order{
//				{
//					Order: msg.Order{
//						Id:     "d41d8cd98f00b204e9800998ecf8427e",
//						Market: testMarket,
//						Party:  testParty,
//					},
//				},
//				{
//					Order: msg.Order{
//						Id:     "ad2dc275947362c45893bbeb30fc3098",
//						Market: "marketABC",
//						Party:  testParty,
//					},
//				},
//				{
//					Order: msg.Order{
//						Id:     "4e8e41367997cfe705d62ea80592cbcc",
//						Market: testMarket,
//						Party:  testParty,
//					},
//				},
//			},
//			inLimit:        5000,
//			inMarket:       "marketABC",
//			outOrdersCount: 1,
//		},
//		{
//			inMarkets: []string{testMarket},
//			inOrders: []Order{
//				{
//					Order: msg.Order{
//						Id:     "d41d8cd98f00b204e9800998ecf8427e",
//						Market: testMarket,
//						Party:  testParty,
//					},
//				},
//				{
//					Order: msg.Order{
//						Id:     "ad2dc275947362c45893bbeb30fc3098",
//						Market: testMarket,
//						Party:  testParty,
//					},
//				},
//				{
//					Order: msg.Order{
//						Id:     "4e8e41367997cfe705d62ea80592cbcc",
//						Market: testMarket,
//						Party:  testParty,
//					},
//				},
//			},
//			inLimit:        2,
//			inMarket:       testMarket,
//			outOrdersCount: 2,
//		},
//	}
//	for testIdx, tt := range tests[:3] {
//		fmt.Printf("TEST NUMBER #%d\n", testIdx)
//		var memStore = NewMemStore(tt.inMarkets, []string{testParty})
//		var newOrderStore = NewOrderStore(&memStore)
//
//		for _, order := range tt.inOrders {
//			err := newOrderStore.Post(order)
//			assert.Nil(t, err)
//		}
//
//		filters := &filters.OrderQueryFilters{}
//		filters.Last = &tt.inLimit
//		orders, err := newOrderStore.GetByMarket(tt.inMarket, filters)
//		assert.Nil(t, err)
//		assert.Equal(t, tt.outOrdersCount, len(orders))
//		newOrderStore.Close()
//	}
//}
//
//func TestMemStore_GetAllOrdersForNoneExistentMarket(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	o, err := newOrderStore.GetByMarket("UNKNOWN", nil)
//	assert.Error(t, err, "market does not exist")
//	assert.Nil(t, o)
//}
//
//func TestMemStore_GetAllTradesForMarket(t *testing.T) {
//	otherMarket := "another"
//	var memStore = NewMemStore([]string{testMarket, otherMarket}, []string{testPartyA, testPartyB})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	var newTradeStore = NewTradeStore(&memStore)
//
//	orderIdA := "d41d8cd98f00b204e9800998ecf8427e"
//	orderIdB := "d41d8cd98f00b204e9800998ecf9999e"
//
//	orderA := Order{
//		Order: msg.Order{
//			Id:     orderIdA,
//			Market: testMarket,
//			Party:  testPartyA,
//		},
//	}
//
//	orderB := Order{
//		Order: msg.Order{
//			Id:     orderIdB,
//			Market: testMarket,
//			Party:  testPartyB,
//		},
//	}
//
//	tradeA := Trade{
//		Trade: msg.Trade{
//			Id:     "c0e8490aa4b1d0071ae8f01cdf45c6aa",
//			Price:  1000,
//			Market: testMarket,
//			Buyer:  testPartyA,
//			Seller: testPartyB,
//		},
//		PassiveOrderId:    orderIdA,
//		AggressiveOrderId: orderIdB,
//	}
//	tradeB := Trade{
//		Trade: msg.Trade{
//			Id:     "d41d8cd98fsb204e9800998ecf8427e",
//			Price:  2000,
//			Market: testMarket,
//			Buyer:  testPartyA,
//			Seller: testPartyB,
//		},
//		PassiveOrderId:    orderIdA,
//		AggressiveOrderId: orderIdB,
//	}
//
//	err := newOrderStore.Post(orderA)
//	assert.Nil(t, err)
//
//	err = newOrderStore.Post(orderB)
//	assert.Nil(t, err)
//
//	err = newTradeStore.Post(tradeA)
//	assert.Nil(t, err)
//	err = newTradeStore.Post(tradeB)
//	assert.Nil(t, err)
//
//	last := uint64(10000)
//	filters := &filters.TradeQueryFilters{}
//	filters.Last = &last
//
//	trades, err := newTradeStore.GetByMarket(testMarket, filters)
//
//	assert.Nil(t, err)
//	assert.Equal(t, 2, len(trades))
//}
//
//func TestMemStore_GetAllTradesForNoneExistentMarket(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newTradeStore = NewTradeStore(&memStore)
//	o, err := newTradeStore.GetByMarket("UNKNOWN", nil)
//	assert.Error(t, err, "market does not exist")
//	assert.Nil(t, o)
//}
//
//func TestMemStore_PutTrade(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testPartyA, testPartyB})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	var newTradeStore = NewTradeStore(&memStore)
//
//	passiveOrderId := "d41d8cd98f00b204e9800998ecf9999e"
//	passiveOrder := Order{
//		Order: msg.Order{
//			Id:     passiveOrderId,
//			Market: testMarket,
//			Party:  testPartyA,
//		},
//	}
//
//	aggressiveOrderId := "d41d8cd98f00b204e9800998ecf8427e"
//	aggressiveOrder := Order{
//		Order: msg.Order{
//			Id:     aggressiveOrderId,
//			Market: testMarket,
//			Party:  testPartyB,
//		},
//	}
//
//	tradeId := "c0e8490aa4b1d0071ae8f01cdf45c6aa"
//	trade := Trade{
//		Trade: msg.Trade{
//			Id:     tradeId,
//			Price:  1000,
//			Market: testMarket,
//			Buyer:  testPartyA,
//			Seller: testPartyB,
//		},
//		PassiveOrderId:    passiveOrderId,
//		AggressiveOrderId: aggressiveOrderId,
//	}
//
//	err := newOrderStore.Post(aggressiveOrder)
//	assert.Nil(t, err)
//
//	err = newOrderStore.Post(passiveOrder)
//	assert.Nil(t, err)
//
//	err = newTradeStore.Post(trade)
//	assert.Nil(t, err)
//
//	tradeOut, err := newTradeStore.GetByMarketAndId(testMarket, tradeId)
//	assert.Nil(t, err)
//	assert.Equal(t, uint64(1000), tradeOut.Price)
//
//	trade = Trade{
//		Trade: msg.Trade{
//			Id:     tradeId,
//			Price:  9000,
//			Market: testMarket,
//			Buyer:  testPartyA,
//			Seller: testPartyB,
//		},
//		PassiveOrderId:    passiveOrderId,
//		AggressiveOrderId: aggressiveOrderId,
//	}
//
//	err = newTradeStore.Put(trade)
//	assert.Nil(t, err)
//
//	tradeOut, err = newTradeStore.GetByMarketAndId(testMarket, tradeId)
//	assert.Nil(t, err)
//	assert.Equal(t, uint64(9000), tradeOut.Price)
//}
//
//func TestMemStore_PutGetAndDeleteTradeForNoneExistentMarket(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newTradeStore = NewTradeStore(&memStore)
//
//	trade := Trade{
//		Trade: msg.Trade{
//			Id:     "A",
//			Price:  9000,
//			Market: "UNKNOWN",
//		},
//		AggressiveOrderId: "Y",
//		PassiveOrderId:    "Z",
//	}
//
//	err := newTradeStore.Put(trade)
//	assert.Error(t, err, "market does not exist")
//
//	_, err = newTradeStore.GetByMarketAndId("UNKNOWN", "ID")
//	assert.Error(t, err, "market does not exist")
//
//	err = newTradeStore.Delete(trade)
//	assert.Error(t, err, "market does not exist")
//
//}
//
//func TestMemOrder_ToString(t *testing.T) {
//	orderId := "d41d8cd98f00b204e9800998ecf8427e"
//	order := Order{
//		Order: msg.Order{
//			Id:     orderId,
//			Market: testMarket,
//			Party:  testParty,
//		},
//	}
//	memOrder := memOrder{
//		order: order,
//	}
//	assert.Equal(t, "memOrder::order-id=d41d8cd98f00b204e9800998ecf8427e", memOrder.String())
//}
//
//func TestMemTrade_ToString(t *testing.T) {
//	tradeId := "c0e8490aa4b1d0071ae8f01cdf45c6aa"
//	passiveOrderId := "d41d8cd98f00b204e9800998ecf8427e"
//	aggressiveOrderId := "d41d8cd98f00b204e9800998ecf9999e"
//	trade := Trade{
//		Trade: msg.Trade{
//			Id:     tradeId,
//			Price:  9000,
//			Market: testMarket,
//		},
//		AggressiveOrderId: aggressiveOrderId,
//		PassiveOrderId:    passiveOrderId,
//	}
//	memTrade := memTrade{
//		trade: trade,
//	}
//	assert.Equal(t, "memTrade::trade-id=c0e8490aa4b1d0071ae8f01cdf45c6aa", memTrade.String())
//}
//
//func TestMemOrderStore_Parties(t *testing.T) {
//	// test when store is added they are added to parties map
//	var memStore = NewMemStore([]string{testMarket}, []string{testPartyA, testPartyB})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	var newTradeStore = NewTradeStore(&memStore)
//
//	passiveOrder := Order{
//		Order: msg.Order{
//			Id:        "d41d8cd98f00b204e9800998ecf9999e",
//			Market:    testMarket,
//			Party:     testPartyA,
//			Remaining: 0,
//		},
//	}
//
//	aggressiveOrder := Order{
//		Order: msg.Order{
//			Id:        "d41d8cd98f00b204e9800998ecf8427e",
//			Market:    testMarket,
//			Party:     testPartyB,
//			Remaining: 100,
//		},
//	}
//
//	trade := Trade{
//		Trade: msg.Trade{
//			Id:        "trade-id",
//			Price:     9000,
//			Market:    testMarket,
//			Buyer:     testPartyA,
//			Seller:    testPartyB,
//			Aggressor: msg.Side_Buy,
//		},
//		AggressiveOrderId: aggressiveOrder.Order.Id,
//		PassiveOrderId:    passiveOrder.Order.Id,
//	}
//
//	err := newOrderStore.Post(passiveOrder)
//	assert.Nil(t, err)
//
//	err = newOrderStore.Post(aggressiveOrder)
//	assert.Nil(t, err)
//
//	err = newTradeStore.Post(trade)
//
//	ordersAtPartyA, err := newOrderStore.GetByParty(testPartyA, nil)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(ordersAtPartyA))
//
//	ordersAtPartyB, err := newOrderStore.GetByParty(testPartyB, nil)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(ordersAtPartyB))
//
//	orderAtPartyA, err := newOrderStore.GetByPartyAndId(testPartyA, passiveOrder.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, passiveOrder, orderAtPartyA)
//
//	orderAtPartyB, err := newOrderStore.GetByPartyAndId(testPartyB, aggressiveOrder.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, aggressiveOrder, orderAtPartyB)
//
//	tradesAtPartyA, err := newTradeStore.GetByParty(testPartyA, nil)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(tradesAtPartyA))
//
//	tradesAtPartyB, err := newTradeStore.GetByParty(testPartyB, nil)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(tradesAtPartyB))
//
//	// update order, parties should also be updated as its a pointer
//	updatedAggressiveOrder := Order{
//		Order: msg.Order{
//			Id:        "d41d8cd98f00b204e9800998ecf8427e",
//			Market:    testMarket,
//			Party:     testPartyB,
//			Remaining: 0,
//		},
//	}
//
//	err = newOrderStore.Put(updatedAggressiveOrder)
//	assert.Nil(t, err)
//	orderAtPartyB, err = newOrderStore.GetByPartyAndId(testPartyB, aggressiveOrder.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, updatedAggressiveOrder, orderAtPartyB)
//
//	// delete trade from trade store, parties should be updated
//	err = newTradeStore.Delete(trade)
//	assert.Nil(t, err)
//
//	tradesAtPartyA, err = newTradeStore.GetByParty(testPartyA, nil)
//	assert.Nil(t, err)
//	assert.Equal(t, 0, len(tradesAtPartyA))
//
//	tradesAtPartyB, err = newTradeStore.GetByParty(testPartyB, nil)
//	assert.Nil(t, err)
//	assert.Equal(t, 0, len(tradesAtPartyB))
//
//	// delete order from trade store, parties should be updated
//	err = newOrderStore.Delete(passiveOrder)
//	assert.Nil(t, err)
//
//	ordersAtPartyA, err = newOrderStore.GetByParty(testPartyA, nil)
//	assert.Nil(t, err)
//	assert.Equal(t, 0, len(ordersAtPartyA))
//
//	// delete order from trade store, parties should be updated
//	err = newOrderStore.Delete(aggressiveOrder)
//	assert.Nil(t, err)
//
//	ordersAtPartyB, err = newOrderStore.GetByParty(testPartyB, nil)
//	assert.Nil(t, err)
//	assert.Equal(t, 0, len(ordersAtPartyB))
//}
//
//func TestAddPartiesOnTheFly(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	var newTradeStore = NewTradeStore(&memStore)
//
//	passiveOrder := Order{
//		Order: msg.Order{
//			Id:        "d41d8cd98f00b204e9800998ecf9999e",
//			Market:    testMarket,
//			Party:     testPartyA,
//			Remaining: 0,
//		},
//	}
//
//	aggressiveOrder := Order{
//		Order: msg.Order{
//			Id:        "d41d8cd98f00b204e9800998ecf8427e",
//			Market:    testMarket,
//			Party:     testPartyB,
//			Remaining: 100,
//		},
//	}
//
//	trade := Trade{
//		Trade: msg.Trade{
//			Id:        "trade-id",
//			Price:     9000,
//			Market:    testMarket,
//			Buyer:     testPartyA,
//			Seller:    testPartyB,
//			Aggressor: msg.Side_Buy,
//		},
//		AggressiveOrderId: aggressiveOrder.Order.Id,
//		PassiveOrderId:    passiveOrder.Order.Id,
//	}
//
//	err := newOrderStore.Post(passiveOrder)
//	assert.Nil(t, err)
//
//	err = newOrderStore.Post(aggressiveOrder)
//	assert.Nil(t, err)
//
//	err = newTradeStore.Post(trade)
//
//	orderAtPartyA, err := newOrderStore.GetByPartyAndId(testPartyA, passiveOrder.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, passiveOrder, orderAtPartyA)
//
//	orderAtPartyB, err := newOrderStore.GetByPartyAndId(testPartyB, aggressiveOrder.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, aggressiveOrder, orderAtPartyB)
//
//	tradesAtPartyA, err := newTradeStore.GetByParty(testPartyA, nil)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(tradesAtPartyA))
//
//	tradesAtPartyB, err := newTradeStore.GetByParty(testPartyB, nil)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(tradesAtPartyB))
//}
//
//func TestNewOrderStore_Filtering(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//
//	order1 := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf9999a",
//			Market:     testMarket,
//			Party:      testPartyA,
//			Side:       msg.Side_Sell,
//			Price:      100,
//			Size:       1000,
//			Remaining:  0,
//			Type:       msg.Order_GTC,
//			Timestamp:  0,
//			Status:     msg.Order_Active,
//		},
//	}
//
//	order2 := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf8427b",
//			Market:     testMarket,
//			Party:      testPartyB,
//			Side:       msg.Side_Buy,
//			Price:      110,
//			Size:       900,
//			Remaining:  0,
//			Type:       msg.Order_GTC,
//			Timestamp:  0,
//			Status:     msg.Order_Active,
//		},
//	}
//
//	order3 := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf8427c",
//			Market:     testMarket,
//			Party:      testPartyA,
//			Side:       msg.Side_Buy,
//			Price:      1000,
//			Size:       1000,
//			Remaining:  1000,
//			Type:       msg.Order_GTC,
//			Timestamp:  1,
//			Status:     msg.Order_Cancelled,
//		},
//	}
//
//	order4 := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf8427d",
//			Market:     testMarket,
//			Party:      testPartyA,
//			Side:       msg.Side_Sell,
//			Price:      100,
//			Size:       100,
//			Remaining:  100,
//			Type:       msg.Order_GTC,
//			Timestamp:  1,
//			Status:     msg.Order_Active,
//		},
//	}
//
//	err := newOrderStore.Post(order1)
//	assert.Nil(t, err)
//	err = newOrderStore.Post(order2)
//	assert.Nil(t, err)
//	err = newOrderStore.Post(order3)
//	assert.Nil(t, err)
//	err = newOrderStore.Post(order4)
//	assert.Nil(t, err)
//
//	orderFilters := &filters.OrderQueryFilters{
//		MarketFilter: &filters.QueryFilter{Eq: testMarket},
//		PartyFilter:  &filters.QueryFilter{Eq: testPartyA},
//	}
//	orders, err := newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(orders))
//
//	// get all orders
//	orderFilters = &filters.OrderQueryFilters{
//		MarketFilter: &filters.QueryFilter{Eq: testMarket},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 4, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		PriceFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{Lower: uint64(150), Upper: uint64(1150)}, Kind: "uint64"},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		PriceFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{Lower: uint64(99), Upper: uint64(200)}, Kind: "uint64"},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		RemainingFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{Lower: uint64(1), Upper: uint64(10000)}, Kind: "uint64"},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 2, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		RemainingFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{Lower: uint64(0), Upper: uint64(10000)}, Kind: "uint64"},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 4, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		SizeFilter: &filters.QueryFilter{Eq: uint64(900)},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		SizeFilter: &filters.QueryFilter{Neq: uint64(900)},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		TypeFilter: &filters.QueryFilter{Eq: msg.Order_GTC},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 4, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		TypeFilter: &filters.QueryFilter{Neq: msg.Order_GTC},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 0, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		PriceFilter: &filters.QueryFilter{Eq: uint64(1000)},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		TypeFilter:  &filters.QueryFilter{Neq: msg.Order_GTC},
//		PriceFilter: &filters.QueryFilter{Eq: uint64(1000)},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 0, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		TimestampFilter: &filters.QueryFilter{Eq: uint64(1)},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 2, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		TimestampFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{uint64(1), uint64(10)}, Kind: "uint64"},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 2, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		TimestampFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{uint64(5), uint64(10)}, Kind: "uint64"},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 0, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		TimestampFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{uint64(0), uint64(10)}, Kind: "uint64"},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 4, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		StatusFilter: &filters.QueryFilter{Eq: msg.Order_Active},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		StatusFilter: &filters.QueryFilter{Eq: msg.Order_Cancelled},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		StatusFilter: &filters.QueryFilter{Eq: msg.Order_Expired},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 0, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		StatusFilter: &filters.QueryFilter{Neq: msg.Order_Expired},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 4, len(orders))
//
//	orderFilters = &filters.OrderQueryFilters{
//		IdFilter: &filters.QueryFilter{ Eq: "d41d8cd98f00b204e9800998ecf8427c"},
//	}
//	orders, err = newOrderStore.GetByMarket(testMarket, orderFilters)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(orders))
//}
//
//func TestNewTradeStore_Filtering(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//	var newTradeStore = NewTradeStore(&memStore)
//
//	order1p := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf9999a",
//			Market:     testMarket,
//			Party:      testPartyA,
//			Side:       msg.Side_Sell,
//			Price:      100,
//			Size:       1000,
//			Remaining:  0,
//			Type:       msg.Order_GTC,
//			Timestamp:  0,
//			Status:     msg.Order_Active,
//		},
//	}
//
//	order1a := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf8427b",
//			Market:     testMarket,
//			Party:      testPartyB,
//			Side:       msg.Side_Buy,
//			Price:      100,
//			Size:       1000,
//			Remaining:  0,
//			Type:       msg.Order_GTC,
//			Timestamp:  0,
//			Status:     msg.Order_Active,
//		},
//	}
//
//	trade1 := Trade{
//		Trade: msg.Trade{
//			Id:        "trade-id-1",
//			Price:     9000,
//			Size:      1000,
//			Market:    testMarket,
//			Buyer:     testPartyA,
//			Seller:    testPartyB,
//			Aggressor: msg.Side_Buy,
//		},
//		AggressiveOrderId: order1a.Order.Id,
//		PassiveOrderId:    order1p.Order.Id,
//	}
//
//	err := newOrderStore.Post(order1a)
//	assert.Nil(t, err)
//	err = newOrderStore.Post(order1p)
//	assert.Nil(t, err)
//	err = newTradeStore.Post(trade1)
//	assert.Nil(t, err)
//
//	order2p := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf9999c",
//			Market:     testMarket,
//			Party:      testPartyA,
//			Side:       msg.Side_Sell,
//			Price:      100,
//			Size:       1000,
//			Remaining:  0,
//			Type:       msg.Order_GTC,
//			Timestamp:  0,
//			Status:     msg.Order_Active,
//		},
//	}
//
//	order2a := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf8427d",
//			Market:     testMarket,
//			Party:      testPartyB,
//			Side:       msg.Side_Buy,
//			Price:      100,
//			Size:       1000,
//			Remaining:  0,
//			Type:       msg.Order_GTC,
//			Timestamp:  0,
//			Status:     msg.Order_Active,
//		},
//	}
//
//	trade2 := Trade{
//		Trade: msg.Trade{
//			Id:        "trade-id-2",
//			Price:     100,
//			Size:      140,
//			Market:    testMarket,
//			Buyer:     testPartyA,
//			Seller:    testPartyB,
//			Aggressor: msg.Side_Buy,
//		},
//		AggressiveOrderId: order2a.Order.Id,
//		PassiveOrderId:    order2p.Order.Id,
//	}
//
//	err = newOrderStore.Post(order2a)
//	assert.Nil(t, err)
//	err = newOrderStore.Post(order2p)
//	assert.Nil(t, err)
//	err = newTradeStore.Post(trade2)
//	assert.Nil(t, err)
//
//	order3p := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf9999e",
//			Market:     testMarket,
//			Party:      testPartyB,
//			Side:       msg.Side_Buy,
//			Price:      110,
//			Size:       900,
//			Remaining:  0,
//			Type:       msg.Order_GTC,
//			Timestamp:  0,
//			Status:     msg.Order_Active,
//		},
//	}
//
//	order3a := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf8427f",
//			Market:     testMarket,
//			Party:      testPartyB,
//			Side:       msg.Side_Buy,
//			Price:      110,
//			Size:       900,
//			Remaining:  0,
//			Type:       msg.Order_GTC,
//			Timestamp:  0,
//			Status:     msg.Order_Active,
//		},
//	}
//
//	trade3 := Trade{
//		Trade: msg.Trade{
//			Id:        "trade-id-3",
//			Price:     110,
//			Size:      1050,
//			Market:    testMarket,
//			Buyer:     testPartyA,
//			Seller:    testPartyB,
//			Aggressor: msg.Side_Buy,
//		},
//		AggressiveOrderId: order3a.Order.Id,
//		PassiveOrderId:    order3p.Order.Id,
//	}
//
//	err = newOrderStore.Post(order3a)
//	assert.Nil(t, err)
//	err = newOrderStore.Post(order3p)
//	assert.Nil(t, err)
//	err = newTradeStore.Post(trade3)
//	assert.Nil(t, err)
//
//	order4a := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf8427g",
//			Market:     testMarket,
//			Party:      testPartyA,
//			Side:       msg.Side_Buy,
//			Price:      1000,
//			Size:       1000,
//			Remaining:  1000,
//			Type:       msg.Order_GTC,
//			Timestamp:  1,
//			Status:     msg.Order_Cancelled,
//		},
//	}
//
//	order4p := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf8427h",
//			Market:     testMarket,
//			Party:      testPartyA,
//			Side:       msg.Side_Sell,
//			Price:      100,
//			Size:       100,
//			Remaining:  100,
//			Type:       msg.Order_GTC,
//			Timestamp:  1,
//			Status:     msg.Order_Active,
//		},
//	}
//
//	trade4 := Trade{
//		Trade: msg.Trade{
//			Id:        "trade-id-4",
//			Price:     100,
//			Size:      100,
//			Market:    testMarket,
//			Buyer:     testPartyB,
//			Seller:    testPartyA,
//			Aggressor: msg.Side_Sell,
//			Timestamp: 1,
//		},
//		AggressiveOrderId: order4a.Order.Id,
//		PassiveOrderId:    order4p.Order.Id,
//	}
//
//	err = newOrderStore.Post(order4a)
//	assert.Nil(t, err)
//	err = newOrderStore.Post(order4p)
//	assert.Nil(t, err)
//	err = newTradeStore.Post(trade4)
//	assert.Nil(t, err)
//
//	getTradeParams := &filters.TradeQueryFilters{
//		MarketFilter: &filters.QueryFilter{Eq: testMarket},
//	}
//	trades, err := newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 4, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		PriceFilter: &filters.QueryFilter{Eq: uint64(9000)},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		PriceFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{uint64(0), uint64(150)}, Kind: "uint64"},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SizeFilter: &filters.QueryFilter{Neq: uint64(10000)},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 4, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SizeFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{uint64(0), uint64(150)}, Kind: "uint64"},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 2, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SizeFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{uint64(150), uint64(1500)}, Kind: "uint64"},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 2, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SizeFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{uint64(1020), uint64(1500)}, Kind: "uint64"},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SizeFilter: &filters.QueryFilter{Eq: uint64(1500)},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 0, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SizeFilter: &filters.QueryFilter{Eq: uint64(1050)},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		BuyerFilter: &filters.QueryFilter{Eq: testPartyA},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		BuyerFilter: &filters.QueryFilter{Eq: testPartyB},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		BuyerFilter: &filters.QueryFilter{Neq: testPartyA},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		BuyerFilter: &filters.QueryFilter{Neq: testPartyB},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SellerFilter: &filters.QueryFilter{Eq: testPartyA},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SellerFilter: &filters.QueryFilter{Eq: testPartyB},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SellerFilter: &filters.QueryFilter{Neq: testPartyA},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SellerFilter: &filters.QueryFilter{Neq: testPartyB},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SellerFilter:    &filters.QueryFilter{Eq: testPartyB},
//		AggressorFilter: &filters.QueryFilter{Eq: msg.Side_Buy},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		SellerFilter:    &filters.QueryFilter{Eq: testPartyB},
//		AggressorFilter: &filters.QueryFilter{Eq: msg.Side_Sell},
//		Operator:        filters.QueryFilterOperatorOr,
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		AggressorFilter: &filters.QueryFilter{Eq: msg.Side_Sell},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		TimestampFilter: &filters.QueryFilter{Eq: uint64(0)},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		TimestampFilter: &filters.QueryFilter{Neq: uint64(0)},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		TimestampFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{uint64(0), uint64(0)}, Kind: "uint64"},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 3, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		TimestampFilter: &filters.QueryFilter{FilterRange: &filters.QueryFilterRange{uint64(1), uint64(1000)}, Kind: "uint64"},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//
//	getTradeParams = &filters.TradeQueryFilters{
//		IdFilter: &filters.QueryFilter{ Eq: "trade-id-3"},
//	}
//	trades, err = newTradeStore.GetByMarket(testMarket, getTradeParams)
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(trades))
//}
//
//
//func TestMemStore_GetOrderByReference(t *testing.T) {
//	var memStore = NewMemStore([]string{testMarket}, []string{testParty})
//	var newOrderStore = NewOrderStore(&memStore)
//	defer newOrderStore.Close()
//
//	order := Order{
//		Order: msg.Order{
//			Id:         "d41d8cd98f00b204e9800998ecf8427b",
//			Market:     testMarket,
//			Party:      testPartyA,
//			Side:       msg.Side_Buy,
//			Price:      100,
//			Size:       1000,
//			Remaining:  0,
//			Type:       msg.Order_GTC,
//			Timestamp:  0,
//			Status:     msg.Order_Active,
//			Reference:  "123123-34334343-1231231",
//		},
//	}
//
//	err := newOrderStore.Post(order)
//	assert.Nil(t, err)
//
//	fetchedOrder, err := newOrderStore.GetByPartyAndReference(testPartyA, "123123-34334343-1231231")
//	assert.Nil(t, err)
//
//	assert.Equal(t, order, fetchedOrder)
//}
