// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/assets (interfaces: Plugin)

// Package mocks is a generated GoMock package.
package mocks

import (
	vega "code.vegaprotocol.io/protos/vega"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPlugin is a mock of Plugin interface
type MockPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockPluginMockRecorder
}

// MockPluginMockRecorder is the mock recorder for MockPlugin
type MockPluginMockRecorder struct {
	mock *MockPlugin
}

// NewMockPlugin creates a new mock instance
func NewMockPlugin(ctrl *gomock.Controller) *MockPlugin {
	mock := &MockPlugin{ctrl: ctrl}
	mock.recorder = &MockPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlugin) EXPECT() *MockPluginMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockPlugin) GetAll() []vega.Asset {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]vega.Asset)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (mr *MockPluginMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPlugin)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockPlugin) GetByID(arg0 string) (*vega.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*vega.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockPluginMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockPlugin)(nil).GetByID), arg0)
}
