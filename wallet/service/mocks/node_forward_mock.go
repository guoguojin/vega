// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/wallet/service (interfaces: NodeForward)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "code.vegaprotocol.io/protos/vega/api/v1"
	v10 "code.vegaprotocol.io/protos/vega/commands/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockNodeForward is a mock of NodeForward interface.
type MockNodeForward struct {
	ctrl     *gomock.Controller
	recorder *MockNodeForwardMockRecorder
}

// MockNodeForwardMockRecorder is the mock recorder for MockNodeForward.
type MockNodeForwardMockRecorder struct {
	mock *MockNodeForward
}

// NewMockNodeForward creates a new mock instance.
func NewMockNodeForward(ctrl *gomock.Controller) *MockNodeForward {
	mock := &MockNodeForward{ctrl: ctrl}
	mock.recorder = &MockNodeForwardMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeForward) EXPECT() *MockNodeForwardMockRecorder {
	return m.recorder
}

// CheckTx mocks base method.
func (m *MockNodeForward) CheckTx(arg0 context.Context, arg1 *v10.Transaction, arg2 int) (*v1.CheckTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CheckTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckTx indicates an expected call of CheckTx.
func (mr *MockNodeForwardMockRecorder) CheckTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTx", reflect.TypeOf((*MockNodeForward)(nil).CheckTx), arg0, arg1, arg2)
}

// HealthCheck mocks base method.
func (m *MockNodeForward) HealthCheck(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockNodeForwardMockRecorder) HealthCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockNodeForward)(nil).HealthCheck), arg0)
}

// LastBlockHeightAndHash mocks base method.
func (m *MockNodeForward) LastBlockHeightAndHash(arg0 context.Context) (*v1.LastBlockHeightResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastBlockHeightAndHash", arg0)
	ret0, _ := ret[0].(*v1.LastBlockHeightResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastBlockHeightAndHash indicates an expected call of LastBlockHeightAndHash.
func (mr *MockNodeForwardMockRecorder) LastBlockHeightAndHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastBlockHeightAndHash", reflect.TypeOf((*MockNodeForward)(nil).LastBlockHeightAndHash), arg0)
}

// SendTx mocks base method.
func (m *MockNodeForward) SendTx(arg0 context.Context, arg1 *v10.Transaction, arg2 v1.SubmitTransactionRequest_Type, arg3 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTx", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTx indicates an expected call of SendTx.
func (mr *MockNodeForwardMockRecorder) SendTx(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTx", reflect.TypeOf((*MockNodeForward)(nil).SendTx), arg0, arg1, arg2, arg3)
}
