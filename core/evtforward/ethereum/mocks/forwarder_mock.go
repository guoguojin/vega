// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/evtforward/ethereum (interfaces: Forwarder)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	v1 "code.vegaprotocol.io/vega/protos/vega/commands/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockForwarder is a mock of Forwarder interface.
type MockForwarder struct {
	ctrl     *gomock.Controller
	recorder *MockForwarderMockRecorder
}

// MockForwarderMockRecorder is the mock recorder for MockForwarder.
type MockForwarderMockRecorder struct {
	mock *MockForwarder
}

// NewMockForwarder creates a new mock instance.
func NewMockForwarder(ctrl *gomock.Controller) *MockForwarder {
	mock := &MockForwarder{ctrl: ctrl}
	mock.recorder = &MockForwarderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockForwarder) EXPECT() *MockForwarderMockRecorder {
	return m.recorder
}

// ForwardFromSelf mocks base method.
func (m *MockForwarder) ForwardFromSelf(arg0 *v1.ChainEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ForwardFromSelf", arg0)
}

// ForwardFromSelf indicates an expected call of ForwardFromSelf.
func (mr *MockForwarderMockRecorder) ForwardFromSelf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForwardFromSelf", reflect.TypeOf((*MockForwarder)(nil).ForwardFromSelf), arg0)
}
