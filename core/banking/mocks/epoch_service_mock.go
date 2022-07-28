// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/banking (interfaces: EpochService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "code.vegaprotocol.io/vega/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockEpochService is a mock of EpochService interface.
type MockEpochService struct {
	ctrl     *gomock.Controller
	recorder *MockEpochServiceMockRecorder
}

// MockEpochServiceMockRecorder is the mock recorder for MockEpochService.
type MockEpochServiceMockRecorder struct {
	mock *MockEpochService
}

// NewMockEpochService creates a new mock instance.
func NewMockEpochService(ctrl *gomock.Controller) *MockEpochService {
	mock := &MockEpochService{ctrl: ctrl}
	mock.recorder = &MockEpochServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEpochService) EXPECT() *MockEpochServiceMockRecorder {
	return m.recorder
}

// NotifyOnEpoch mocks base method.
func (m *MockEpochService) NotifyOnEpoch(arg0, arg1 func(context.Context, types.Epoch)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyOnEpoch", arg0, arg1)
}

// NotifyOnEpoch indicates an expected call of NotifyOnEpoch.
func (mr *MockEpochServiceMockRecorder) NotifyOnEpoch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyOnEpoch", reflect.TypeOf((*MockEpochService)(nil).NotifyOnEpoch), arg0, arg1)
}
