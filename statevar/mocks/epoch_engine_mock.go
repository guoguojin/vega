// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/statevar (interfaces: EpochEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "code.vegaprotocol.io/vega/types"
	gomock "github.com/golang/mock/gomock"
)

// MockEpochEngine is a mock of EpochEngine interface.
type MockEpochEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEpochEngineMockRecorder
}

// MockEpochEngineMockRecorder is the mock recorder for MockEpochEngine.
type MockEpochEngineMockRecorder struct {
	mock *MockEpochEngine
}

// NewMockEpochEngine creates a new mock instance.
func NewMockEpochEngine(ctrl *gomock.Controller) *MockEpochEngine {
	mock := &MockEpochEngine{ctrl: ctrl}
	mock.recorder = &MockEpochEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEpochEngine) EXPECT() *MockEpochEngineMockRecorder {
	return m.recorder
}

// NotifyOnEpoch mocks base method.
func (m *MockEpochEngine) NotifyOnEpoch(arg0 func(context.Context, types.Epoch)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyOnEpoch", arg0)
}

// NotifyOnEpoch indicates an expected call of NotifyOnEpoch.
func (mr *MockEpochEngineMockRecorder) NotifyOnEpoch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyOnEpoch", reflect.TypeOf((*MockEpochEngine)(nil).NotifyOnEpoch), arg0)
}
