// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/wallet/network (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	network "code.vegaprotocol.io/vega/wallet/network"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteNetwork mocks base method.
func (m *MockStore) DeleteNetwork(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetwork", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNetwork indicates an expected call of DeleteNetwork.
func (mr *MockStoreMockRecorder) DeleteNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetwork", reflect.TypeOf((*MockStore)(nil).DeleteNetwork), arg0)
}

// GetNetwork mocks base method.
func (m *MockStore) GetNetwork(arg0 string) (*network.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetwork", arg0)
	ret0, _ := ret[0].(*network.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetwork indicates an expected call of GetNetwork.
func (mr *MockStoreMockRecorder) GetNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetwork", reflect.TypeOf((*MockStore)(nil).GetNetwork), arg0)
}

// GetNetworkPath mocks base method.
func (m *MockStore) GetNetworkPath(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkPath", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNetworkPath indicates an expected call of GetNetworkPath.
func (mr *MockStoreMockRecorder) GetNetworkPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkPath", reflect.TypeOf((*MockStore)(nil).GetNetworkPath), arg0)
}

// ListNetworks mocks base method.
func (m *MockStore) ListNetworks() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNetworks")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNetworks indicates an expected call of ListNetworks.
func (mr *MockStoreMockRecorder) ListNetworks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetworks", reflect.TypeOf((*MockStore)(nil).ListNetworks))
}

// NetworkExists mocks base method.
func (m *MockStore) NetworkExists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkExists indicates an expected call of NetworkExists.
func (mr *MockStoreMockRecorder) NetworkExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkExists", reflect.TypeOf((*MockStore)(nil).NetworkExists), arg0)
}

// SaveNetwork mocks base method.
func (m *MockStore) SaveNetwork(arg0 *network.Network) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNetwork", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNetwork indicates an expected call of SaveNetwork.
func (mr *MockStoreMockRecorder) SaveNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNetwork", reflect.TypeOf((*MockStore)(nil).SaveNetwork), arg0)
}
