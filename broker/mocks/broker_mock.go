// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/broker (interfaces: BrokerI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	broker "code.vegaprotocol.io/vega/broker"
	events "code.vegaprotocol.io/vega/events"
	gomock "github.com/golang/mock/gomock"
)

// MockBrokerI is a mock of BrokerI interface.
type MockBrokerI struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerIMockRecorder
}

// MockBrokerIMockRecorder is the mock recorder for MockBrokerI.
type MockBrokerIMockRecorder struct {
	mock *MockBrokerI
}

// NewMockBrokerI creates a new mock instance.
func NewMockBrokerI(ctrl *gomock.Controller) *MockBrokerI {
	mock := &MockBrokerI{ctrl: ctrl}
	mock.recorder = &MockBrokerIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBrokerI) EXPECT() *MockBrokerIMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockBrokerI) Send(arg0 events.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0)
}

// Send indicates an expected call of Send.
func (mr *MockBrokerIMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBrokerI)(nil).Send), arg0)
}

// SendBatch mocks base method.
func (m *MockBrokerI) SendBatch(arg0 []events.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendBatch", arg0)
}

// SendBatch indicates an expected call of SendBatch.
func (mr *MockBrokerIMockRecorder) SendBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBatch", reflect.TypeOf((*MockBrokerI)(nil).SendBatch), arg0)
}

// SetStreaming mocks base method.
func (m *MockBrokerI) SetStreaming(arg0 bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStreaming", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetStreaming indicates an expected call of SetStreaming.
func (mr *MockBrokerIMockRecorder) SetStreaming(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStreaming", reflect.TypeOf((*MockBrokerI)(nil).SetStreaming), arg0)
}

// Subscribe mocks base method.
func (m *MockBrokerI) Subscribe(arg0 broker.Subscriber) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockBrokerIMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockBrokerI)(nil).Subscribe), arg0)
}

// SubscribeBatch mocks base method.
func (m *MockBrokerI) SubscribeBatch(arg0 ...broker.Subscriber) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SubscribeBatch", varargs...)
}

// SubscribeBatch indicates an expected call of SubscribeBatch.
func (mr *MockBrokerIMockRecorder) SubscribeBatch(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeBatch", reflect.TypeOf((*MockBrokerI)(nil).SubscribeBatch), arg0...)
}

// Unsubscribe mocks base method.
func (m *MockBrokerI) Unsubscribe(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unsubscribe", arg0)
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockBrokerIMockRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockBrokerI)(nil).Unsubscribe), arg0)
}
