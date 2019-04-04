// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/blockchain (interfaces: ServiceTime)

// Package mocks is a generated GoMock package.
package mocks

import (
	vegatime "code.vegaprotocol.io/vega/internal/vegatime"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockServiceTime is a mock of ServiceTime interface
type MockServiceTime struct {
	ctrl     *gomock.Controller
	recorder *MockServiceTimeMockRecorder
}

// MockServiceTimeMockRecorder is the mock recorder for MockServiceTime
type MockServiceTimeMockRecorder struct {
	mock *MockServiceTime
}

// NewMockServiceTime creates a new mock instance
func NewMockServiceTime(ctrl *gomock.Controller) *MockServiceTime {
	mock := &MockServiceTime{ctrl: ctrl}
	mock.recorder = &MockServiceTimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceTime) EXPECT() *MockServiceTimeMockRecorder {
	return m.recorder
}

// GetTimeNow mocks base method
func (m *MockServiceTime) GetTimeNow() (vegatime.Stamp, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeNow")
	ret0, _ := ret[0].(vegatime.Stamp)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTimeNow indicates an expected call of GetTimeNow
func (mr *MockServiceTimeMockRecorder) GetTimeNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeNow", reflect.TypeOf((*MockServiceTime)(nil).GetTimeNow))
}
