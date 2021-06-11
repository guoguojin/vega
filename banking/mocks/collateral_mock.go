// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/banking (interfaces: Collateral)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	types "code.vegaprotocol.io/vega/types"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCollateral is a mock of Collateral interface
type MockCollateral struct {
	ctrl     *gomock.Controller
	recorder *MockCollateralMockRecorder
}

// MockCollateralMockRecorder is the mock recorder for MockCollateral
type MockCollateralMockRecorder struct {
	mock *MockCollateral
}

// NewMockCollateral creates a new mock instance
func NewMockCollateral(ctrl *gomock.Controller) *MockCollateral {
	mock := &MockCollateral{ctrl: ctrl}
	mock.recorder = &MockCollateralMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollateral) EXPECT() *MockCollateralMockRecorder {
	return m.recorder
}

// Deposit mocks base method
func (m *MockCollateral) Deposit(arg0 context.Context, arg1, arg2 string, arg3 uint64) (*types.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deposit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deposit indicates an expected call of Deposit
func (mr *MockCollateralMockRecorder) Deposit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deposit", reflect.TypeOf((*MockCollateral)(nil).Deposit), arg0, arg1, arg2, arg3)
}

// EnableAsset mocks base method
func (m *MockCollateral) EnableAsset(arg0 context.Context, arg1 proto.Asset) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableAsset", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableAsset indicates an expected call of EnableAsset
func (mr *MockCollateralMockRecorder) EnableAsset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAsset", reflect.TypeOf((*MockCollateral)(nil).EnableAsset), arg0, arg1)
}

// HasBalance mocks base method
func (m *MockCollateral) HasBalance(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasBalance", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasBalance indicates an expected call of HasBalance
func (mr *MockCollateralMockRecorder) HasBalance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasBalance", reflect.TypeOf((*MockCollateral)(nil).HasBalance), arg0)
}

// LockFundsForWithdraw mocks base method
func (m *MockCollateral) LockFundsForWithdraw(arg0 context.Context, arg1, arg2 string, arg3 uint64) (*types.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockFundsForWithdraw", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LockFundsForWithdraw indicates an expected call of LockFundsForWithdraw
func (mr *MockCollateralMockRecorder) LockFundsForWithdraw(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockFundsForWithdraw", reflect.TypeOf((*MockCollateral)(nil).LockFundsForWithdraw), arg0, arg1, arg2, arg3)
}

// Withdraw mocks base method
func (m *MockCollateral) Withdraw(arg0 context.Context, arg1, arg2 string, arg3 uint64) (*types.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Withdraw indicates an expected call of Withdraw
func (mr *MockCollateralMockRecorder) Withdraw(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockCollateral)(nil).Withdraw), arg0, arg1, arg2, arg3)
}
