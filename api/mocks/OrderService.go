// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "context"
import core "vega/core"
import datastore "vega/datastore"
import filters "vega/filters"
import mock "github.com/stretchr/testify/mock"
import msg "vega/msg"

// OrderService is an autogenerated mock type for the OrderService type
type OrderService struct {
	mock.Mock
}

// AmendOrder provides a mock function with given fields: ctx, amendment
func (_m *OrderService) AmendOrder(ctx context.Context, amendment *msg.Amendment) (bool, error) {
	ret := _m.Called(ctx, amendment)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *msg.Amendment) bool); ok {
		r0 = rf(ctx, amendment)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *msg.Amendment) error); ok {
		r1 = rf(ctx, amendment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelOrder provides a mock function with given fields: ctx, order
func (_m *OrderService) CancelOrder(ctx context.Context, order *msg.Order) (bool, error) {
	ret := _m.Called(ctx, order)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *msg.Order) bool); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *msg.Order) error); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrder provides a mock function with given fields: ctx, order
func (_m *OrderService) CreateOrder(ctx context.Context, order *msg.Order) (bool, string, error) {
	ret := _m.Called(ctx, order)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *msg.Order) bool); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, *msg.Order) string); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *msg.Order) error); ok {
		r2 = rf(ctx, order)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByMarket provides a mock function with given fields: ctx, market, _a2
func (_m *OrderService) GetByMarket(ctx context.Context, market string, _a2 *filters.OrderQueryFilters) ([]*msg.Order, error) {
	ret := _m.Called(ctx, market, _a2)

	var r0 []*msg.Order
	if rf, ok := ret.Get(0).(func(context.Context, string, *filters.OrderQueryFilters) []*msg.Order); ok {
		r0 = rf(ctx, market, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*msg.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *filters.OrderQueryFilters) error); ok {
		r1 = rf(ctx, market, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByMarketAndId provides a mock function with given fields: ctx, market, id
func (_m *OrderService) GetByMarketAndId(ctx context.Context, market string, id string) (*msg.Order, error) {
	ret := _m.Called(ctx, market, id)

	var r0 *msg.Order
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *msg.Order); ok {
		r0 = rf(ctx, market, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msg.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, market, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByParty provides a mock function with given fields: ctx, party, _a2
func (_m *OrderService) GetByParty(ctx context.Context, party string, _a2 *filters.OrderQueryFilters) ([]*msg.Order, error) {
	ret := _m.Called(ctx, party, _a2)

	var r0 []*msg.Order
	if rf, ok := ret.Get(0).(func(context.Context, string, *filters.OrderQueryFilters) []*msg.Order); ok {
		r0 = rf(ctx, party, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*msg.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *filters.OrderQueryFilters) error); ok {
		r1 = rf(ctx, party, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPartyAndId provides a mock function with given fields: ctx, market, id
func (_m *OrderService) GetByPartyAndId(ctx context.Context, market string, id string) (*msg.Order, error) {
	ret := _m.Called(ctx, market, id)

	var r0 *msg.Order
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *msg.Order); ok {
		r0 = rf(ctx, market, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msg.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, market, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMarketDepth provides a mock function with given fields: ctx, market
func (_m *OrderService) GetMarketDepth(ctx context.Context, market string) (*msg.MarketDepth, error) {
	ret := _m.Called(ctx, market)

	var r0 *msg.MarketDepth
	if rf, ok := ret.Get(0).(func(context.Context, string) *msg.MarketDepth); ok {
		r0 = rf(ctx, market)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msg.MarketDepth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, market)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMarkets provides a mock function with given fields: ctx
func (_m *OrderService) GetMarkets(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStatistics provides a mock function with given fields: ctx
func (_m *OrderService) GetStatistics(ctx context.Context) (*msg.Statistics, error) {
	ret := _m.Called(ctx)

	var r0 *msg.Statistics
	if rf, ok := ret.Get(0).(func(context.Context) *msg.Statistics); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msg.Statistics)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: vega, orderStore
func (_m *OrderService) Init(vega *core.Vega, orderStore datastore.OrderStore) {
	_m.Called(vega, orderStore)
}

// ObserveMarketDepth provides a mock function with given fields: ctx, market
func (_m *OrderService) ObserveMarketDepth(ctx context.Context, market string) (<-chan msg.MarketDepth, uint64) {
	ret := _m.Called(ctx, market)

	var r0 <-chan msg.MarketDepth
	if rf, ok := ret.Get(0).(func(context.Context, string) <-chan msg.MarketDepth); ok {
		r0 = rf(ctx, market)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan msg.MarketDepth)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, string) uint64); ok {
		r1 = rf(ctx, market)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}

// ObserveOrders provides a mock function with given fields: ctx, market, party
func (_m *OrderService) ObserveOrders(ctx context.Context, market *string, party *string) (<-chan msg.Order, uint64) {
	ret := _m.Called(ctx, market, party)

	var r0 <-chan msg.Order
	if rf, ok := ret.Get(0).(func(context.Context, *string, *string) <-chan msg.Order); ok {
		r0 = rf(ctx, market, party)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan msg.Order)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, *string, *string) uint64); ok {
		r1 = rf(ctx, market, party)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}
