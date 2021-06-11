// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/api (interfaces: BlockchainClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	api "code.vegaprotocol.io/vega/proto/api"
	v1 "code.vegaprotocol.io/vega/proto/commands/v1"
	context "context"
	gomock "github.com/golang/mock/gomock"
	types "github.com/tendermint/tendermint/rpc/core/types"
	reflect "reflect"
	time "time"
)

// MockBlockchainClient is a mock of BlockchainClient interface
type MockBlockchainClient struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainClientMockRecorder
}

// MockBlockchainClientMockRecorder is the mock recorder for MockBlockchainClient
type MockBlockchainClientMockRecorder struct {
	mock *MockBlockchainClient
}

// NewMockBlockchainClient creates a new mock instance
func NewMockBlockchainClient(ctrl *gomock.Controller) *MockBlockchainClient {
	mock := &MockBlockchainClient{ctrl: ctrl}
	mock.recorder = &MockBlockchainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockchainClient) EXPECT() *MockBlockchainClientMockRecorder {
	return m.recorder
}

// GetChainID mocks base method
func (m *MockBlockchainClient) GetChainID(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainID indicates an expected call of GetChainID
func (mr *MockBlockchainClientMockRecorder) GetChainID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainID", reflect.TypeOf((*MockBlockchainClient)(nil).GetChainID), arg0)
}

// GetGenesisTime mocks base method
func (m *MockBlockchainClient) GetGenesisTime(arg0 context.Context) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenesisTime", arg0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenesisTime indicates an expected call of GetGenesisTime
func (mr *MockBlockchainClientMockRecorder) GetGenesisTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenesisTime", reflect.TypeOf((*MockBlockchainClient)(nil).GetGenesisTime), arg0)
}

// GetNetworkInfo mocks base method
func (m *MockBlockchainClient) GetNetworkInfo(arg0 context.Context) (*types.ResultNetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkInfo", arg0)
	ret0, _ := ret[0].(*types.ResultNetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkInfo indicates an expected call of GetNetworkInfo
func (mr *MockBlockchainClientMockRecorder) GetNetworkInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkInfo", reflect.TypeOf((*MockBlockchainClient)(nil).GetNetworkInfo), arg0)
}

// GetStatus mocks base method
func (m *MockBlockchainClient) GetStatus(arg0 context.Context) (*types.ResultStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus", arg0)
	ret0, _ := ret[0].(*types.ResultStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus
func (mr *MockBlockchainClientMockRecorder) GetStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockBlockchainClient)(nil).GetStatus), arg0)
}

// GetUnconfirmedTxCount mocks base method
func (m *MockBlockchainClient) GetUnconfirmedTxCount(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnconfirmedTxCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnconfirmedTxCount indicates an expected call of GetUnconfirmedTxCount
func (mr *MockBlockchainClientMockRecorder) GetUnconfirmedTxCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnconfirmedTxCount", reflect.TypeOf((*MockBlockchainClient)(nil).GetUnconfirmedTxCount), arg0)
}

// Health mocks base method
func (m *MockBlockchainClient) Health() (*types.ResultHealth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health")
	ret0, _ := ret[0].(*types.ResultHealth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health
func (mr *MockBlockchainClientMockRecorder) Health() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockBlockchainClient)(nil).Health))
}

// SubmitTransaction mocks base method
func (m *MockBlockchainClient) SubmitTransaction(arg0 context.Context, arg1 *proto.SignedBundle, arg2 api.SubmitTransactionRequest_Type) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTransaction", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitTransaction indicates an expected call of SubmitTransaction
func (mr *MockBlockchainClientMockRecorder) SubmitTransaction(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransaction", reflect.TypeOf((*MockBlockchainClient)(nil).SubmitTransaction), arg0, arg1, arg2)
}

// SubmitTransactionV2 mocks base method
func (m *MockBlockchainClient) SubmitTransactionV2(arg0 context.Context, arg1 *v1.Transaction, arg2 api.SubmitTransactionV2Request_Type) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTransactionV2", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitTransactionV2 indicates an expected call of SubmitTransactionV2
func (mr *MockBlockchainClientMockRecorder) SubmitTransactionV2(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransactionV2", reflect.TypeOf((*MockBlockchainClient)(nil).SubmitTransactionV2), arg0, arg1, arg2)
}
