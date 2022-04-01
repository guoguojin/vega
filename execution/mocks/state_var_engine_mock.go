// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/execution (interfaces: StateVarEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	statevar "code.vegaprotocol.io/vega/types/statevar"
	gomock "github.com/golang/mock/gomock"
)

// MockStateVarEngine is a mock of StateVarEngine interface.
type MockStateVarEngine struct {
	ctrl     *gomock.Controller
	recorder *MockStateVarEngineMockRecorder
}

// MockStateVarEngineMockRecorder is the mock recorder for MockStateVarEngine.
type MockStateVarEngineMockRecorder struct {
	mock *MockStateVarEngine
}

// NewMockStateVarEngine creates a new mock instance.
func NewMockStateVarEngine(ctrl *gomock.Controller) *MockStateVarEngine {
	mock := &MockStateVarEngine{ctrl: ctrl}
	mock.recorder = &MockStateVarEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateVarEngine) EXPECT() *MockStateVarEngineMockRecorder {
	return m.recorder
}

// NewEvent mocks base method.
func (m *MockStateVarEngine) NewEvent(arg0, arg1 string, arg2 statevar.StateVarEventType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewEvent", arg0, arg1, arg2)
}

// NewEvent indicates an expected call of NewEvent.
func (mr *MockStateVarEngineMockRecorder) NewEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEvent", reflect.TypeOf((*MockStateVarEngine)(nil).NewEvent), arg0, arg1, arg2)
}

// ReadyForTimeTrigger mocks base method.
func (m *MockStateVarEngine) ReadyForTimeTrigger(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReadyForTimeTrigger", arg0, arg1)
}

// ReadyForTimeTrigger indicates an expected call of ReadyForTimeTrigger.
func (mr *MockStateVarEngineMockRecorder) ReadyForTimeTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadyForTimeTrigger", reflect.TypeOf((*MockStateVarEngine)(nil).ReadyForTimeTrigger), arg0, arg1)
}

// RegisterStateVariable mocks base method.
func (m *MockStateVarEngine) RegisterStateVariable(arg0, arg1, arg2 string, arg3 statevar.Converter, arg4 func(string, statevar.FinaliseCalculation), arg5 []statevar.StateVarEventType, arg6 func(context.Context, statevar.StateVariableResult) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterStateVariable", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterStateVariable indicates an expected call of RegisterStateVariable.
func (mr *MockStateVarEngineMockRecorder) RegisterStateVariable(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterStateVariable", reflect.TypeOf((*MockStateVarEngine)(nil).RegisterStateVariable), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// RemoveTimeTriggers mocks base method.
func (m *MockStateVarEngine) RemoveTimeTriggers(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveTimeTriggers", arg0, arg1)
}

// RemoveTimeTriggers indicates an expected call of RemoveTimeTriggers.
func (mr *MockStateVarEngineMockRecorder) RemoveTimeTriggers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTimeTriggers", reflect.TypeOf((*MockStateVarEngine)(nil).RemoveTimeTriggers), arg0, arg1)
}
