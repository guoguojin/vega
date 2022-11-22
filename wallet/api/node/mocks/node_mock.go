// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/wallet/api/node (interfaces: GRPCAdapter)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "code.vegaprotocol.io/vega/protos/vega/api/v1"
	types "code.vegaprotocol.io/vega/wallet/api/node/types"
	gomock "github.com/golang/mock/gomock"
)

// MockGRPCAdapter is a mock of GRPCAdapter interface.
type MockGRPCAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCAdapterMockRecorder
}

// MockGRPCAdapterMockRecorder is the mock recorder for MockGRPCAdapter.
type MockGRPCAdapterMockRecorder struct {
	mock *MockGRPCAdapter
}

// NewMockGRPCAdapter creates a new mock instance.
func NewMockGRPCAdapter(ctrl *gomock.Controller) *MockGRPCAdapter {
	mock := &MockGRPCAdapter{ctrl: ctrl}
	mock.recorder = &MockGRPCAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGRPCAdapter) EXPECT() *MockGRPCAdapterMockRecorder {
	return m.recorder
}

// Host mocks base method.
func (m *MockGRPCAdapter) Host() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(string)
	return ret0
}

// Host indicates an expected call of Host.
func (mr *MockGRPCAdapterMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockGRPCAdapter)(nil).Host))
}

// LastBlock mocks base method.
func (m *MockGRPCAdapter) LastBlock(arg0 context.Context) (types.LastBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastBlock", arg0)
	ret0, _ := ret[0].(types.LastBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastBlock indicates an expected call of LastBlock.
func (mr *MockGRPCAdapterMockRecorder) LastBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastBlock", reflect.TypeOf((*MockGRPCAdapter)(nil).LastBlock), arg0)
}

// Statistics mocks base method.
func (m *MockGRPCAdapter) Statistics(arg0 context.Context) (types.Statistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Statistics", arg0)
	ret0, _ := ret[0].(types.Statistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Statistics indicates an expected call of Statistics.
func (mr *MockGRPCAdapterMockRecorder) Statistics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Statistics", reflect.TypeOf((*MockGRPCAdapter)(nil).Statistics), arg0)
}

// Stop mocks base method.
func (m *MockGRPCAdapter) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockGRPCAdapterMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockGRPCAdapter)(nil).Stop))
}

// SubmitTransaction mocks base method.
func (m *MockGRPCAdapter) SubmitTransaction(arg0 context.Context, arg1 *v1.SubmitTransactionRequest) (*v1.SubmitTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTransaction", arg0, arg1)
	ret0, _ := ret[0].(*v1.SubmitTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitTransaction indicates an expected call of SubmitTransaction.
func (mr *MockGRPCAdapterMockRecorder) SubmitTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransaction", reflect.TypeOf((*MockGRPCAdapter)(nil).SubmitTransaction), arg0, arg1)
}
