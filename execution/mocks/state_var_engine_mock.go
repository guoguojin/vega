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

// AddStateVariable mocks base method.
func (m *MockStateVarEngine) AddStateVariable(arg0, arg1 string, arg2 statevar.Converter, arg3 func(string, statevar.FinaliseCalculation), arg4 []statevar.StateVarEventType, arg5 func(context.Context, statevar.StateVariableResult) error, arg6 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStateVariable", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddStateVariable indicates an expected call of AddStateVariable.
func (mr *MockStateVarEngineMockRecorder) AddStateVariable(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStateVariable", reflect.TypeOf((*MockStateVarEngine)(nil).AddStateVariable), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
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
