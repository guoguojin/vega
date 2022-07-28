// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/liquidity (interfaces: OrderBook)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "code.vegaprotocol.io/vega/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockOrderBook is a mock of OrderBook interface.
type MockOrderBook struct {
	ctrl     *gomock.Controller
	recorder *MockOrderBookMockRecorder
}

// MockOrderBookMockRecorder is the mock recorder for MockOrderBook.
type MockOrderBookMockRecorder struct {
	mock *MockOrderBook
}

// NewMockOrderBook creates a new mock instance.
func NewMockOrderBook(ctrl *gomock.Controller) *MockOrderBook {
	mock := &MockOrderBook{ctrl: ctrl}
	mock.recorder = &MockOrderBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderBook) EXPECT() *MockOrderBookMockRecorder {
	return m.recorder
}

// GetOrderByID mocks base method.
func (m *MockOrderBook) GetOrderByID(arg0 string) (*types.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByID", arg0)
	ret0, _ := ret[0].(*types.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByID indicates an expected call of GetOrderByID.
func (mr *MockOrderBookMockRecorder) GetOrderByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByID", reflect.TypeOf((*MockOrderBook)(nil).GetOrderByID), arg0)
}
