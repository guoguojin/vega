// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/monitoring (interfaces: BlockchainClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	coretypes "github.com/tendermint/tendermint/rpc/coretypes"
)

// MockBlockchainClient is a mock of BlockchainClient interface.
type MockBlockchainClient struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainClientMockRecorder
}

// MockBlockchainClientMockRecorder is the mock recorder for MockBlockchainClient.
type MockBlockchainClientMockRecorder struct {
	mock *MockBlockchainClient
}

// NewMockBlockchainClient creates a new mock instance.
func NewMockBlockchainClient(ctrl *gomock.Controller) *MockBlockchainClient {
	mock := &MockBlockchainClient{ctrl: ctrl}
	mock.recorder = &MockBlockchainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockchainClient) EXPECT() *MockBlockchainClientMockRecorder {
	return m.recorder
}

// GetStatus mocks base method.
func (m *MockBlockchainClient) GetStatus(arg0 context.Context) (*coretypes.ResultStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus", arg0)
	ret0, _ := ret[0].(*coretypes.ResultStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockBlockchainClientMockRecorder) GetStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockBlockchainClient)(nil).GetStatus), arg0)
}

// GetUnconfirmedTxCount mocks base method.
func (m *MockBlockchainClient) GetUnconfirmedTxCount(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnconfirmedTxCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnconfirmedTxCount indicates an expected call of GetUnconfirmedTxCount.
func (mr *MockBlockchainClientMockRecorder) GetUnconfirmedTxCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnconfirmedTxCount", reflect.TypeOf((*MockBlockchainClient)(nil).GetUnconfirmedTxCount), arg0)
}

// Health mocks base method.
func (m *MockBlockchainClient) Health() (*coretypes.ResultHealth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health")
	ret0, _ := ret[0].(*coretypes.ResultHealth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health.
func (mr *MockBlockchainClientMockRecorder) Health() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockBlockchainClient)(nil).Health))
}
