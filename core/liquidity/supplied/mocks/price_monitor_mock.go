// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/liquidity/supplied (interfaces: PriceMonitor)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	num "code.vegaprotocol.io/vega/core/types/num"
	gomock "github.com/golang/mock/gomock"
)

// MockPriceMonitor is a mock of PriceMonitor interface.
type MockPriceMonitor struct {
	ctrl     *gomock.Controller
	recorder *MockPriceMonitorMockRecorder
}

// MockPriceMonitorMockRecorder is the mock recorder for MockPriceMonitor.
type MockPriceMonitorMockRecorder struct {
	mock *MockPriceMonitor
}

// NewMockPriceMonitor creates a new mock instance.
func NewMockPriceMonitor(ctrl *gomock.Controller) *MockPriceMonitor {
	mock := &MockPriceMonitor{ctrl: ctrl}
	mock.recorder = &MockPriceMonitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPriceMonitor) EXPECT() *MockPriceMonitorMockRecorder {
	return m.recorder
}

// GetValidPriceRange mocks base method.
func (m *MockPriceMonitor) GetValidPriceRange() (num.WrappedDecimal, num.WrappedDecimal) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidPriceRange")
	ret0, _ := ret[0].(num.WrappedDecimal)
	ret1, _ := ret[1].(num.WrappedDecimal)
	return ret0, ret1
}

// GetValidPriceRange indicates an expected call of GetValidPriceRange.
func (mr *MockPriceMonitorMockRecorder) GetValidPriceRange() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidPriceRange", reflect.TypeOf((*MockPriceMonitor)(nil).GetValidPriceRange))
}
