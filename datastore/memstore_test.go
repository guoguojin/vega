package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"vega/proto"
	"context"
)

const testMarket = "market"

func TestNewMemStore_ReturnsNewMemStoreInstance(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	assert.NotNil(t, memStore)
}

func TestNewTradeStore_ReturnsNewTradeStoreInstance(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newTradeStore = NewTradeStore(&memStore)
	assert.NotNil(t, newTradeStore)
}

func TestNewOrderStore_ReturnsNewOrderStoreInstance(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	assert.NotNil(t, newOrderStore)
}

func TestMemOrderStore_PutAndGetNewOrder(t *testing.T) {
	var ctx = context.Background()
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)

	var order = &Order{
		Order: msg.Order{
			Id:     "45305210ff7a9bb9450b1833cc10368a",
			Market: testMarket,
		},
	}

	err := newOrderStore.Put(ctx, order)
	assert.Nil(t, err)

	o, err := newOrderStore.Get(ctx, testMarket, order.Id)
	assert.Nil(t, err)
	assert.Equal(t, order, o)
}

func TestMemOrderStore_PutAndGetExistingOrder(t *testing.T) {
	var ctx = context.Background()
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)

	var order = &Order{
		Order: msg.Order{
			Id:     "c471bdd5f381aa3654d98f4591eaa968",
			Market: testMarket,
			Party:  "tester",
			Price:  100,
			Size:   1,
		},
	}

	err := newOrderStore.Put(ctx, order)
	assert.Nil(t, err)

	o, err := newOrderStore.Get(ctx, testMarket, order.Id)
	assert.Nil(t, err)
	assert.Equal(t, uint64(100), o.Price)
	assert.Equal(t, uint64(1), o.Size)

	order.Price = 1000
	order.Size = 5

	err = newOrderStore.Put(ctx, order)
	assert.Nil(t, err)

	o, err = newOrderStore.Get(ctx, testMarket, order.Id)
	assert.Nil(t, err)
	assert.Equal(t, order, o)
	assert.Equal(t, uint64(1000), o.Price)
	assert.Equal(t, uint64(5), o.Size)
}

func TestMemOrderStore_PutAndDeleteOrder(t *testing.T) {
	var ctx = context.Background()
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)

	var order = &Order{
		Order: msg.Order{
			Id:     "45305210ff7a9bb9450b1833cc10368a",
			Market: testMarket,
		},

	}

	err := newOrderStore.Put(ctx, order)
	assert.Nil(t, err)

	o, err := newOrderStore.Get(ctx, testMarket, order.Id)
	assert.Nil(t, err)
	assert.Equal(t, order, o)

	err = newOrderStore.Delete(ctx, o)
	assert.Nil(t, err)

	o, err = newOrderStore.Get(ctx, testMarket, order.Id)
	assert.Nil(t, o)
}

func TestMemOrderStore_PutAndGetTrade(t *testing.T) {
	var ctx = context.Background()
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var newTradeStore = NewTradeStore(&memStore)

	var trade = &Trade{
		Trade: msg.Trade{Market:  testMarket, },
		OrderID: "d41d8cd98f00b204e9800998ecf8427e",
	}

	var order = &Order{
		Order: msg.Order{
			Id:     "d41d8cd98f00b204e9800998ecf8427e",
			Market: testMarket,
		},
	}

	err := newOrderStore.Put(ctx, order)
	assert.Nil(t, err)

	err = newTradeStore.Put(ctx, trade)
	assert.Nil(t, err)

	tr, err := newTradeStore.Get(ctx, testMarket, trade.Id)
	assert.Equal(t, trade, tr)
}

func TestMemOrderStore_PutAndDeleteTrade(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var newTradeStore = NewTradeStore(&memStore)
	var ctx = context.Background()

	var order = &Order{
		Order: msg.Order{ Id: "d41d8cd98f00b204e9800998ecf8427e", Market: testMarket },
	}
	var trade = &Trade{
		OrderID: "d41d8cd98f00b204e9800998ecf8427e",
		Trade: msg.Trade{ Market:  testMarket, },
	}

	err := newOrderStore.Put(ctx, order)
	assert.Nil(t, err)

	err = newTradeStore.Put(ctx, trade)
	assert.Nil(t, err)

	tr, err := newTradeStore.Get(ctx, testMarket, trade.Id)
	assert.Equal(t, trade, tr)

	err = newTradeStore.Delete(ctx, tr)
	assert.Nil(t, err)

	tr, err = newTradeStore.Get(ctx, testMarket, trade.Id)
	assert.Nil(t, tr)
}

func TestMemOrderStore_PutTradeOrderNotFound(t *testing.T) {
	var ctx = context.Background()
	var memStore = NewMemStore([]string{testMarket})
	var newTradeStore = NewTradeStore(&memStore)
	trade := &Trade{
		Trade: msg.Trade{
			Id:      "one",
			Market:  testMarket,
		},
		OrderID: "mystery",
	}
	err := newTradeStore.Put(ctx, trade)
	assert.Error(t, err)
}

func TestMemOrderStore_PutAndFindByOrderId(t *testing.T) {
	var ctx = context.Background()
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var newTradeStore = NewTradeStore(&memStore)

	trade1 := &Trade{
		Trade: msg.Trade{
			Id:      "one",
			Market:  testMarket,
		},
		OrderID: "d41d8cd98f00b204e9800998ecf8427e",
	}
	trade2 := &Trade{
		Trade: msg.Trade{
			Id:      "two",
			Market:  testMarket,
		},
		OrderID: "d41d8cd98f00b204e9800998ecf8427e",
	}
	order := &Order{
		Order: msg.Order{
			Id:     "d41d8cd98f00b204e9800998ecf8427e",
			Market: testMarket,
		},
	}

	err := newOrderStore.Put(ctx, order)
	assert.Nil(t, err)

	err = newTradeStore.Put(ctx, trade1)
	assert.Nil(t, err)

	err = newTradeStore.Put(ctx, trade2)
	assert.Nil(t, err)

	trades, err := newTradeStore.GetByOrderID(ctx, testMarket, order.Id)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(trades))
	assert.Equal(t, "one", trades[0].Id)
	assert.Equal(t, "two", trades[1].Id)
}



func TestMemOrderStore_GetAllOrdersForMarket(t *testing.T) {
	otherMarket := "another"
	var ctx = context.Background()
	var memStore = NewMemStore([]string{testMarket, otherMarket})
	var newOrderStore = NewOrderStore(&memStore)

	order1 := &Order{
		Order: msg.Order{
			Id:     "d41d8cd98f00b204e9800998ecf8427e",
			Market: testMarket,
		},
	}

	order2 := &Order{
		Order: msg.Order{
			Id:     "ad2dc275947362c45893bbeb30fc3098",
			Market: otherMarket,
		},
	}
	
	order3 := &Order{
		Order: msg.Order{
			Id:     "4e8e41367997cfe705d62ea80592cbcc",
			Market: testMarket,
		},
	}

	err := newOrderStore.Put(ctx, order1)
	assert.Nil(t, err)
	err = newOrderStore.Put(ctx, order2)
	assert.Nil(t, err)
	err = newOrderStore.Put(ctx, order3)
	assert.Nil(t, err)

	orders, err := newOrderStore.All(ctx, testMarket)
	assert.Equal(t, 2, len(orders) )
	orders, err = newOrderStore.All(ctx, otherMarket)
	assert.Equal(t, 1, len(orders) )
}