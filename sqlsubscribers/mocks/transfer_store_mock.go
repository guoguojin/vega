// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/sqlsubscribers (interfaces: TransferStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "code.vegaprotocol.io/data-node/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockTransferStore is a mock of TransferStore interface.
type MockTransferStore struct {
	ctrl     *gomock.Controller
	recorder *MockTransferStoreMockRecorder
}

// MockTransferStoreMockRecorder is the mock recorder for MockTransferStore.
type MockTransferStoreMockRecorder struct {
	mock *MockTransferStore
}

// NewMockTransferStore creates a new mock instance.
func NewMockTransferStore(ctrl *gomock.Controller) *MockTransferStore {
	mock := &MockTransferStore{ctrl: ctrl}
	mock.recorder = &MockTransferStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferStore) EXPECT() *MockTransferStoreMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockTransferStore) Upsert(arg0 context.Context, arg1 *entities.Transfer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockTransferStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockTransferStore)(nil).Upsert), arg0, arg1)
}
