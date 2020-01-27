// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/markets (interfaces: OrderStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrderStore is a mock of OrderStore interface
type MockOrderStore struct {
	ctrl     *gomock.Controller
	recorder *MockOrderStoreMockRecorder
}

// MockOrderStoreMockRecorder is the mock recorder for MockOrderStore
type MockOrderStoreMockRecorder struct {
	mock *MockOrderStore
}

// NewMockOrderStore creates a new mock instance
func NewMockOrderStore(ctrl *gomock.Controller) *MockOrderStore {
	mock := &MockOrderStore{ctrl: ctrl}
	mock.recorder = &MockOrderStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderStore) EXPECT() *MockOrderStoreMockRecorder {
	return m.recorder
}

// GetMarketDepth mocks base method
func (m *MockOrderStore) GetMarketDepth(arg0 context.Context, arg1 string, arg2 uint64) (*proto.MarketDepth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketDepth", arg0, arg1, arg2)
	ret0, _ := ret[0].(*proto.MarketDepth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketDepth indicates an expected call of GetMarketDepth
func (mr *MockOrderStoreMockRecorder) GetMarketDepth(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketDepth", reflect.TypeOf((*MockOrderStore)(nil).GetMarketDepth), arg0, arg1, arg2)
}

// Subscribe mocks base method
func (m *MockOrderStore) Subscribe(arg0 chan<- []proto.Order) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockOrderStoreMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockOrderStore)(nil).Subscribe), arg0)
}

// Unsubscribe mocks base method
func (m *MockOrderStore) Unsubscribe(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockOrderStoreMockRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockOrderStore)(nil).Unsubscribe), arg0)
}
