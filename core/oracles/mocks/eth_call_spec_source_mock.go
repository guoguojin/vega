// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/oracles (interfaces: EthCallSpecSource)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	oracles "code.vegaprotocol.io/vega/core/oracles"
	gomock "github.com/golang/mock/gomock"
)

// MockEthCallSpecSource is a mock of EthCallSpecSource interface.
type MockEthCallSpecSource struct {
	ctrl     *gomock.Controller
	recorder *MockEthCallSpecSourceMockRecorder
}

// MockEthCallSpecSourceMockRecorder is the mock recorder for MockEthCallSpecSource.
type MockEthCallSpecSourceMockRecorder struct {
	mock *MockEthCallSpecSource
}

// NewMockEthCallSpecSource creates a new mock instance.
func NewMockEthCallSpecSource(ctrl *gomock.Controller) *MockEthCallSpecSource {
	mock := &MockEthCallSpecSource{ctrl: ctrl}
	mock.recorder = &MockEthCallSpecSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthCallSpecSource) EXPECT() *MockEthCallSpecSourceMockRecorder {
	return m.recorder
}

// GetCall mocks base method.
func (m *MockEthCallSpecSource) GetCall(arg0 string) (oracles.EthCall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCall", arg0)
	ret0, _ := ret[0].(oracles.EthCall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCall indicates an expected call of GetCall.
func (mr *MockEthCallSpecSourceMockRecorder) GetCall(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCall", reflect.TypeOf((*MockEthCallSpecSource)(nil).GetCall), arg0)
}
