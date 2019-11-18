// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/execution (interfaces: CandleBuf)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockCandleBuf is a mock of CandleBuf interface
type MockCandleBuf struct {
	ctrl     *gomock.Controller
	recorder *MockCandleBufMockRecorder
}

// MockCandleBufMockRecorder is the mock recorder for MockCandleBuf
type MockCandleBufMockRecorder struct {
	mock *MockCandleBuf
}

// NewMockCandleBuf creates a new mock instance
func NewMockCandleBuf(ctrl *gomock.Controller) *MockCandleBuf {
	mock := &MockCandleBuf{ctrl: ctrl}
	mock.recorder = &MockCandleBufMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCandleBuf) EXPECT() *MockCandleBufMockRecorder {
	return m.recorder
}

// AddTrade mocks base method
func (m *MockCandleBuf) AddTrade(arg0 proto.Trade) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTrade", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTrade indicates an expected call of AddTrade
func (mr *MockCandleBufMockRecorder) AddTrade(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrade", reflect.TypeOf((*MockCandleBuf)(nil).AddTrade), arg0)
}

// Flush mocks base method
func (m *MockCandleBuf) Flush(arg0 string, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockCandleBufMockRecorder) Flush(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockCandleBuf)(nil).Flush), arg0, arg1)
}

// Start mocks base method
func (m *MockCandleBuf) Start(arg0 string, arg1 time.Time) (map[string]proto.Candle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(map[string]proto.Candle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start
func (mr *MockCandleBufMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockCandleBuf)(nil).Start), arg0, arg1)
}
