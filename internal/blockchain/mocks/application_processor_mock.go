// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/blockchain (interfaces: ApplicationProcessor)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockApplicationProcessor is a mock of ApplicationProcessor interface
type MockApplicationProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationProcessorMockRecorder
}

// MockApplicationProcessorMockRecorder is the mock recorder for MockApplicationProcessor
type MockApplicationProcessorMockRecorder struct {
	mock *MockApplicationProcessor
}

// NewMockApplicationProcessor creates a new mock instance
func NewMockApplicationProcessor(ctrl *gomock.Controller) *MockApplicationProcessor {
	mock := &MockApplicationProcessor{ctrl: ctrl}
	mock.recorder = &MockApplicationProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationProcessor) EXPECT() *MockApplicationProcessorMockRecorder {
	return m.recorder
}

// Process mocks base method
func (m *MockApplicationProcessor) Process(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Process indicates an expected call of Process
func (mr *MockApplicationProcessorMockRecorder) Process(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockApplicationProcessor)(nil).Process), arg0)
}

// Validate mocks base method
func (m *MockApplicationProcessor) Validate(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockApplicationProcessorMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockApplicationProcessor)(nil).Validate), arg0)
}
