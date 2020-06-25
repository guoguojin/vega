// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: Notary)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
)

// MockNotary is a mock of Notary interface
type MockNotary struct {
	ctrl     *gomock.Controller
	recorder *MockNotaryMockRecorder
}

// MockNotaryMockRecorder is the mock recorder for MockNotary
type MockNotaryMockRecorder struct {
	mock *MockNotary
}

// NewMockNotary creates a new mock instance
func NewMockNotary(ctrl *gomock.Controller) *MockNotary {
	mock := &MockNotary{ctrl: ctrl}
	mock.recorder = &MockNotaryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNotary) EXPECT() *MockNotaryMockRecorder {
	return m.recorder
}

// AddSig mocks base method
func (m *MockNotary) AddSig(arg0 context.Context, arg1 []byte, arg2 proto.NodeSignature) ([]proto.NodeSignature, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSig", arg0, arg1, arg2)
	ret0, _ := ret[0].([]proto.NodeSignature)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddSig indicates an expected call of AddSig
func (mr *MockNotaryMockRecorder) AddSig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSig", reflect.TypeOf((*MockNotary)(nil).AddSig), arg0, arg1, arg2)
}

// StartAggregate mocks base method
func (m *MockNotary) StartAggregate(arg0 string, arg1 proto.NodeSignatureKind) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartAggregate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartAggregate indicates an expected call of StartAggregate
func (mr *MockNotaryMockRecorder) StartAggregate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAggregate", reflect.TypeOf((*MockNotary)(nil).StartAggregate), arg0, arg1)
}
