// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/api (interfaces: TransferResponseService)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/data-node/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTransferResponseService is a mock of TransferResponseService interface
type MockTransferResponseService struct {
	ctrl     *gomock.Controller
	recorder *MockTransferResponseServiceMockRecorder
}

// MockTransferResponseServiceMockRecorder is the mock recorder for MockTransferResponseService
type MockTransferResponseServiceMockRecorder struct {
	mock *MockTransferResponseService
}

// NewMockTransferResponseService creates a new mock instance
func NewMockTransferResponseService(ctrl *gomock.Controller) *MockTransferResponseService {
	mock := &MockTransferResponseService{ctrl: ctrl}
	mock.recorder = &MockTransferResponseServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransferResponseService) EXPECT() *MockTransferResponseServiceMockRecorder {
	return m.recorder
}

// ObserveTransferResponses mocks base method
func (m *MockTransferResponseService) ObserveTransferResponses(arg0 context.Context, arg1 int) (<-chan []*proto.TransferResponse, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObserveTransferResponses", arg0, arg1)
	ret0, _ := ret[0].(<-chan []*proto.TransferResponse)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// ObserveTransferResponses indicates an expected call of ObserveTransferResponses
func (mr *MockTransferResponseServiceMockRecorder) ObserveTransferResponses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveTransferResponses", reflect.TypeOf((*MockTransferResponseService)(nil).ObserveTransferResponses), arg0, arg1)
}
