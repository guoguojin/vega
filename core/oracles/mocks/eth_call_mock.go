// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/oracles (interfaces: EthCall)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	ethereum "github.com/ethereum/go-ethereum"
	gomock "github.com/golang/mock/gomock"
)

// MockEthCall is a mock of EthCall interface.
type MockEthCall struct {
	ctrl     *gomock.Controller
	recorder *MockEthCallMockRecorder
}

// MockEthCallMockRecorder is the mock recorder for MockEthCall.
type MockEthCallMockRecorder struct {
	mock *MockEthCall
}

// NewMockEthCall creates a new mock instance.
func NewMockEthCall(ctrl *gomock.Controller) *MockEthCall {
	mock := &MockEthCall{ctrl: ctrl}
	mock.recorder = &MockEthCallMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthCall) EXPECT() *MockEthCallMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockEthCall) Call(arg0 context.Context, arg1 ethereum.ContractCaller, arg2 *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockEthCallMockRecorder) Call(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockEthCall)(nil).Call), arg0, arg1, arg2)
}

// Normalise mocks base method.
func (m *MockEthCall) Normalise(arg0 []byte) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Normalise", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Normalise indicates an expected call of Normalise.
func (mr *MockEthCallMockRecorder) Normalise(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Normalise", reflect.TypeOf((*MockEthCall)(nil).Normalise), arg0)
}

// PassesFilters mocks base method.
func (m *MockEthCall) PassesFilters(arg0 []byte, arg1, arg2 uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PassesFilters", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// PassesFilters indicates an expected call of PassesFilters.
func (mr *MockEthCallMockRecorder) PassesFilters(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PassesFilters", reflect.TypeOf((*MockEthCall)(nil).PassesFilters), arg0, arg1, arg2)
}

// RequiredConfirmations mocks base method.
func (m *MockEthCall) RequiredConfirmations() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequiredConfirmations")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// RequiredConfirmations indicates an expected call of RequiredConfirmations.
func (mr *MockEthCallMockRecorder) RequiredConfirmations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequiredConfirmations", reflect.TypeOf((*MockEthCall)(nil).RequiredConfirmations))
}
