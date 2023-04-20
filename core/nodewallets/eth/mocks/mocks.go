// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/nodewallets/eth (interfaces: EthereumWallet)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	registry "code.vegaprotocol.io/vega/core/nodewallets/registry"
	crypto "code.vegaprotocol.io/vega/libs/crypto"
	gomock "github.com/golang/mock/gomock"
)

// MockEthereumWallet is a mock of EthereumWallet interface.
type MockEthereumWallet struct {
	ctrl     *gomock.Controller
	recorder *MockEthereumWalletMockRecorder
}

// MockEthereumWalletMockRecorder is the mock recorder for MockEthereumWallet.
type MockEthereumWalletMockRecorder struct {
	mock *MockEthereumWallet
}

// NewMockEthereumWallet creates a new mock instance.
func NewMockEthereumWallet(ctrl *gomock.Controller) *MockEthereumWallet {
	mock := &MockEthereumWallet{ctrl: ctrl}
	mock.recorder = &MockEthereumWalletMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthereumWallet) EXPECT() *MockEthereumWalletMockRecorder {
	return m.recorder
}

// Algo mocks base method.
func (m *MockEthereumWallet) Algo() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Algo")
	ret0, _ := ret[0].(string)
	return ret0
}

// Algo indicates an expected call of Algo.
func (mr *MockEthereumWalletMockRecorder) Algo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Algo", reflect.TypeOf((*MockEthereumWallet)(nil).Algo))
}

// Chain mocks base method.
func (m *MockEthereumWallet) Chain() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chain")
	ret0, _ := ret[0].(string)
	return ret0
}

// Chain indicates an expected call of Chain.
func (mr *MockEthereumWalletMockRecorder) Chain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chain", reflect.TypeOf((*MockEthereumWallet)(nil).Chain))
}

// Cleanup mocks base method.
func (m *MockEthereumWallet) Cleanup() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cleanup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Cleanup indicates an expected call of Cleanup.
func (mr *MockEthereumWalletMockRecorder) Cleanup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockEthereumWallet)(nil).Cleanup))
}

// Name mocks base method.
func (m *MockEthereumWallet) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockEthereumWalletMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockEthereumWallet)(nil).Name))
}

// PubKey mocks base method.
func (m *MockEthereumWallet) PubKey() crypto.PublicKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubKey")
	ret0, _ := ret[0].(crypto.PublicKey)
	return ret0
}

// PubKey indicates an expected call of PubKey.
func (mr *MockEthereumWalletMockRecorder) PubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubKey", reflect.TypeOf((*MockEthereumWallet)(nil).PubKey))
}

// Reload mocks base method.
func (m *MockEthereumWallet) Reload(arg0 registry.EthereumWalletDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockEthereumWalletMockRecorder) Reload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockEthereumWallet)(nil).Reload), arg0)
}

// Sign mocks base method.
func (m *MockEthereumWallet) Sign(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockEthereumWalletMockRecorder) Sign(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockEthereumWallet)(nil).Sign), arg0)
}

// Version mocks base method.
func (m *MockEthereumWallet) Version() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockEthereumWalletMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockEthereumWallet)(nil).Version))
}
