// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/accounts (interfaces: AccountStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/data-node/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAccountStore is a mock of AccountStore interface
type MockAccountStore struct {
	ctrl     *gomock.Controller
	recorder *MockAccountStoreMockRecorder
}

// MockAccountStoreMockRecorder is the mock recorder for MockAccountStore
type MockAccountStoreMockRecorder struct {
	mock *MockAccountStore
}

// NewMockAccountStore creates a new mock instance
func NewMockAccountStore(ctrl *gomock.Controller) *MockAccountStore {
	mock := &MockAccountStore{ctrl: ctrl}
	mock.recorder = &MockAccountStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountStore) EXPECT() *MockAccountStoreMockRecorder {
	return m.recorder
}

// GetFeeInfrastructureAccounts mocks base method
func (m *MockAccountStore) GetFeeInfrastructureAccounts(arg0 string) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeeInfrastructureAccounts", arg0)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeeInfrastructureAccounts indicates an expected call of GetFeeInfrastructureAccounts
func (mr *MockAccountStoreMockRecorder) GetFeeInfrastructureAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeeInfrastructureAccounts", reflect.TypeOf((*MockAccountStore)(nil).GetFeeInfrastructureAccounts), arg0)
}

// GetMarketAccounts mocks base method
func (m *MockAccountStore) GetMarketAccounts(arg0, arg1 string) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketAccounts", arg0, arg1)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketAccounts indicates an expected call of GetMarketAccounts
func (mr *MockAccountStoreMockRecorder) GetMarketAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketAccounts", reflect.TypeOf((*MockAccountStore)(nil).GetMarketAccounts), arg0, arg1)
}

// GetPartyAccounts mocks base method
func (m *MockAccountStore) GetPartyAccounts(arg0, arg1, arg2 string, arg3 proto.AccountType) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyAccounts", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyAccounts indicates an expected call of GetPartyAccounts
func (mr *MockAccountStoreMockRecorder) GetPartyAccounts(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyAccounts", reflect.TypeOf((*MockAccountStore)(nil).GetPartyAccounts), arg0, arg1, arg2, arg3)
}

// Subscribe mocks base method
func (m *MockAccountStore) Subscribe(arg0 chan []*proto.Account) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockAccountStoreMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockAccountStore)(nil).Subscribe), arg0)
}

// Unsubscribe mocks base method
func (m *MockAccountStore) Unsubscribe(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockAccountStoreMockRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockAccountStore)(nil).Unsubscribe), arg0)
}
