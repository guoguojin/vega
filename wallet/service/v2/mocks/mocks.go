// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/wallet/service/v2 (interfaces: ClientAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	jsonrpc "code.vegaprotocol.io/vega/libs/jsonrpc"
	api "code.vegaprotocol.io/vega/wallet/api"
	wallet "code.vegaprotocol.io/vega/wallet/wallet"
	gomock "github.com/golang/mock/gomock"
)

// MockClientAPI is a mock of ClientAPI interface.
type MockClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockClientAPIMockRecorder
}

// MockClientAPIMockRecorder is the mock recorder for MockClientAPI.
type MockClientAPIMockRecorder struct {
	mock *MockClientAPI
}

// NewMockClientAPI creates a new mock instance.
func NewMockClientAPI(ctrl *gomock.Controller) *MockClientAPI {
	mock := &MockClientAPI{ctrl: ctrl}
	mock.recorder = &MockClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientAPI) EXPECT() *MockClientAPIMockRecorder {
	return m.recorder
}

// ConnectWallet mocks base method.
func (m *MockClientAPI) ConnectWallet(arg0 context.Context, arg1 string) (wallet.Wallet, *jsonrpc.ErrorDetails) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectWallet", arg0, arg1)
	ret0, _ := ret[0].(wallet.Wallet)
	ret1, _ := ret[1].(*jsonrpc.ErrorDetails)
	return ret0, ret1
}

// ConnectWallet indicates an expected call of ConnectWallet.
func (mr *MockClientAPIMockRecorder) ConnectWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectWallet", reflect.TypeOf((*MockClientAPI)(nil).ConnectWallet), arg0, arg1)
}

// GetChainID mocks base method.
func (m *MockClientAPI) GetChainID(arg0 context.Context) (jsonrpc.Result, *jsonrpc.ErrorDetails) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainID", arg0)
	ret0, _ := ret[0].(jsonrpc.Result)
	ret1, _ := ret[1].(*jsonrpc.ErrorDetails)
	return ret0, ret1
}

// GetChainID indicates an expected call of GetChainID.
func (mr *MockClientAPIMockRecorder) GetChainID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainID", reflect.TypeOf((*MockClientAPI)(nil).GetChainID), arg0)
}

// ListKeys mocks base method.
func (m *MockClientAPI) ListKeys(arg0 context.Context, arg1 api.ConnectedWallet) (jsonrpc.Result, *jsonrpc.ErrorDetails) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1)
	ret0, _ := ret[0].(jsonrpc.Result)
	ret1, _ := ret[1].(*jsonrpc.ErrorDetails)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockClientAPIMockRecorder) ListKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockClientAPI)(nil).ListKeys), arg0, arg1)
}

// SendTransaction mocks base method.
func (m *MockClientAPI) SendTransaction(arg0 context.Context, arg1 jsonrpc.Params, arg2 api.ConnectedWallet) (jsonrpc.Result, *jsonrpc.ErrorDetails) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", arg0, arg1, arg2)
	ret0, _ := ret[0].(jsonrpc.Result)
	ret1, _ := ret[1].(*jsonrpc.ErrorDetails)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockClientAPIMockRecorder) SendTransaction(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockClientAPI)(nil).SendTransaction), arg0, arg1, arg2)
}

// SignTransaction mocks base method.
func (m *MockClientAPI) SignTransaction(arg0 context.Context, arg1 jsonrpc.Params, arg2 api.ConnectedWallet) (jsonrpc.Result, *jsonrpc.ErrorDetails) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTransaction", arg0, arg1, arg2)
	ret0, _ := ret[0].(jsonrpc.Result)
	ret1, _ := ret[1].(*jsonrpc.ErrorDetails)
	return ret0, ret1
}

// SignTransaction indicates an expected call of SignTransaction.
func (mr *MockClientAPIMockRecorder) SignTransaction(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTransaction", reflect.TypeOf((*MockClientAPI)(nil).SignTransaction), arg0, arg1, arg2)
}
