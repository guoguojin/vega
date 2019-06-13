// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/accounts (interfaces: Accounts)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAccountStore is a mock of Accounts interface
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

// GetAccountsByOwnerAndAsset mocks base method
func (m *MockAccountStore) GetAccountsByOwnerAndAsset(arg0, arg1 string) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountsByOwnerAndAsset", arg0, arg1)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountsByOwnerAndAsset indicates an expected call of GetAccountsByOwnerAndAsset
func (mr *MockAccountStoreMockRecorder) GetAccountsByOwnerAndAsset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountsByOwnerAndAsset", reflect.TypeOf((*MockAccountStore)(nil).GetAccountsByOwnerAndAsset), arg0, arg1)
}

// GetAccountsForOwner mocks base method
func (m *MockAccountStore) GetAccountsForOwner(arg0 string) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountsForOwner", arg0)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountsForOwner indicates an expected call of GetAccountsForOwner
func (mr *MockAccountStoreMockRecorder) GetAccountsForOwner(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountsForOwner", reflect.TypeOf((*MockAccountStore)(nil).GetAccountsForOwner), arg0)
}

// GetAccountsForOwnerByType mocks base method
func (m *MockAccountStore) GetAccountsForOwnerByType(arg0 string, arg1 proto.AccountType) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountsForOwnerByType", arg0, arg1)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountsForOwnerByType indicates an expected call of GetAccountsForOwnerByType
func (mr *MockAccountStoreMockRecorder) GetAccountsForOwnerByType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountsForOwnerByType", reflect.TypeOf((*MockAccountStore)(nil).GetAccountsForOwnerByType), arg0, arg1)
}

// GetMarketAccountsForOwner mocks base method
func (m *MockAccountStore) GetMarketAccountsForOwner(arg0, arg1 string) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketAccountsForOwner", arg0, arg1)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketAccountsForOwner indicates an expected call of GetMarketAccountsForOwner
func (mr *MockAccountStoreMockRecorder) GetMarketAccountsForOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketAccountsForOwner", reflect.TypeOf((*MockAccountStore)(nil).GetMarketAccountsForOwner), arg0, arg1)
}

// GetMarketAssetAccounts mocks base method
func (m *MockAccountStore) GetMarketAssetAccounts(arg0, arg1, arg2 string) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketAssetAccounts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketAssetAccounts indicates an expected call of GetMarketAssetAccounts
func (mr *MockAccountStoreMockRecorder) GetMarketAssetAccounts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketAssetAccounts", reflect.TypeOf((*MockAccountStore)(nil).GetMarketAssetAccounts), arg0, arg1, arg2)
}
