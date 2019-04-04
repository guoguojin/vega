// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/blockchain (interfaces: ProcessorService)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProcessorService is a mock of ProcessorService interface
type MockProcessorService struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorServiceMockRecorder
}

// MockProcessorServiceMockRecorder is the mock recorder for MockProcessorService
type MockProcessorServiceMockRecorder struct {
	mock *MockProcessorService
}

// NewMockProcessorService creates a new mock instance
func NewMockProcessorService(ctrl *gomock.Controller) *MockProcessorService {
	mock := &MockProcessorService{ctrl: ctrl}
	mock.recorder = &MockProcessorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcessorService) EXPECT() *MockProcessorServiceMockRecorder {
	return m.recorder
}

// AmendOrder mocks base method
func (m *MockProcessorService) AmendOrder(arg0 *proto.OrderAmendment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AmendOrder", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AmendOrder indicates an expected call of AmendOrder
func (mr *MockProcessorServiceMockRecorder) AmendOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AmendOrder", reflect.TypeOf((*MockProcessorService)(nil).AmendOrder), arg0)
}

// CancelOrder mocks base method
func (m *MockProcessorService) CancelOrder(arg0 *proto.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelOrder indicates an expected call of CancelOrder
func (mr *MockProcessorServiceMockRecorder) CancelOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockProcessorService)(nil).CancelOrder), arg0)
}

// SubmitOrder mocks base method
func (m *MockProcessorService) SubmitOrder(arg0 *proto.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitOrder", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitOrder indicates an expected call of SubmitOrder
func (mr *MockProcessorServiceMockRecorder) SubmitOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOrder", reflect.TypeOf((*MockProcessorService)(nil).SubmitOrder), arg0)
}
