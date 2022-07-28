// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/evtforward/ethereum (interfaces: Filterer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ethereum "code.vegaprotocol.io/vega/core/evtforward/ethereum"
	gomock "github.com/golang/mock/gomock"
)

// MockFilterer is a mock of Filterer interface.
type MockFilterer struct {
	ctrl     *gomock.Controller
	recorder *MockFiltererMockRecorder
}

// MockFiltererMockRecorder is the mock recorder for MockFilterer.
type MockFiltererMockRecorder struct {
	mock *MockFilterer
}

// NewMockFilterer creates a new mock instance.
func NewMockFilterer(ctrl *gomock.Controller) *MockFilterer {
	mock := &MockFilterer{ctrl: ctrl}
	mock.recorder = &MockFiltererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilterer) EXPECT() *MockFiltererMockRecorder {
	return m.recorder
}

// CurrentHeight mocks base method.
func (m *MockFilterer) CurrentHeight(arg0 context.Context) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentHeight", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// CurrentHeight indicates an expected call of CurrentHeight.
func (mr *MockFiltererMockRecorder) CurrentHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentHeight", reflect.TypeOf((*MockFilterer)(nil).CurrentHeight), arg0)
}

// FilterCollateralEvents mocks base method.
func (m *MockFilterer) FilterCollateralEvents(arg0 context.Context, arg1, arg2 uint64, arg3 ethereum.OnEventFound) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FilterCollateralEvents", arg0, arg1, arg2, arg3)
}

// FilterCollateralEvents indicates an expected call of FilterCollateralEvents.
func (mr *MockFiltererMockRecorder) FilterCollateralEvents(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterCollateralEvents", reflect.TypeOf((*MockFilterer)(nil).FilterCollateralEvents), arg0, arg1, arg2, arg3)
}

// FilterMultisigControlEvents mocks base method.
func (m *MockFilterer) FilterMultisigControlEvents(arg0 context.Context, arg1, arg2 uint64, arg3 ethereum.OnEventFound) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FilterMultisigControlEvents", arg0, arg1, arg2, arg3)
}

// FilterMultisigControlEvents indicates an expected call of FilterMultisigControlEvents.
func (mr *MockFiltererMockRecorder) FilterMultisigControlEvents(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterMultisigControlEvents", reflect.TypeOf((*MockFilterer)(nil).FilterMultisigControlEvents), arg0, arg1, arg2, arg3)
}

// FilterStakingEvents mocks base method.
func (m *MockFilterer) FilterStakingEvents(arg0 context.Context, arg1, arg2 uint64, arg3 ethereum.OnEventFound) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FilterStakingEvents", arg0, arg1, arg2, arg3)
}

// FilterStakingEvents indicates an expected call of FilterStakingEvents.
func (mr *MockFiltererMockRecorder) FilterStakingEvents(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterStakingEvents", reflect.TypeOf((*MockFilterer)(nil).FilterStakingEvents), arg0, arg1, arg2, arg3)
}

// FilterVestingEvents mocks base method.
func (m *MockFilterer) FilterVestingEvents(arg0 context.Context, arg1, arg2 uint64, arg3 ethereum.OnEventFound) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FilterVestingEvents", arg0, arg1, arg2, arg3)
}

// FilterVestingEvents indicates an expected call of FilterVestingEvents.
func (mr *MockFiltererMockRecorder) FilterVestingEvents(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterVestingEvents", reflect.TypeOf((*MockFilterer)(nil).FilterVestingEvents), arg0, arg1, arg2, arg3)
}
