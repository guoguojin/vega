// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/staking (interfaces: EvtForwarder)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	v1 "code.vegaprotocol.io/protos/vega/commands/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockEvtForwarder is a mock of EvtForwarder interface.
type MockEvtForwarder struct {
	ctrl     *gomock.Controller
	recorder *MockEvtForwarderMockRecorder
}

// MockEvtForwarderMockRecorder is the mock recorder for MockEvtForwarder.
type MockEvtForwarderMockRecorder struct {
	mock *MockEvtForwarder
}

// NewMockEvtForwarder creates a new mock instance.
func NewMockEvtForwarder(ctrl *gomock.Controller) *MockEvtForwarder {
	mock := &MockEvtForwarder{ctrl: ctrl}
	mock.recorder = &MockEvtForwarderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvtForwarder) EXPECT() *MockEvtForwarderMockRecorder {
	return m.recorder
}

// ForwardFromSelf mocks base method.
func (m *MockEvtForwarder) ForwardFromSelf(arg0 *v1.ChainEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ForwardFromSelf", arg0)
}

// ForwardFromSelf indicates an expected call of ForwardFromSelf.
func (mr *MockEvtForwarderMockRecorder) ForwardFromSelf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForwardFromSelf", reflect.TypeOf((*MockEvtForwarder)(nil).ForwardFromSelf), arg0)
}
