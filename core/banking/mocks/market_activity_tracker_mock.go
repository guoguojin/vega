// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/banking (interfaces: MarketActivityTracker)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	vega "code.vegaprotocol.io/protos/vega"
	types "code.vegaprotocol.io/vega/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockMarketActivityTracker is a mock of MarketActivityTracker interface.
type MockMarketActivityTracker struct {
	ctrl     *gomock.Controller
	recorder *MockMarketActivityTrackerMockRecorder
}

// MockMarketActivityTrackerMockRecorder is the mock recorder for MockMarketActivityTracker.
type MockMarketActivityTrackerMockRecorder struct {
	mock *MockMarketActivityTracker
}

// NewMockMarketActivityTracker creates a new mock instance.
func NewMockMarketActivityTracker(ctrl *gomock.Controller) *MockMarketActivityTracker {
	mock := &MockMarketActivityTracker{ctrl: ctrl}
	mock.recorder = &MockMarketActivityTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarketActivityTracker) EXPECT() *MockMarketActivityTrackerMockRecorder {
	return m.recorder
}

// GetMarketScores mocks base method.
func (m *MockMarketActivityTracker) GetMarketScores(arg0 string, arg1 []string, arg2 vega.DispatchMetric) []*types.MarketContributionScore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketScores", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.MarketContributionScore)
	return ret0
}

// GetMarketScores indicates an expected call of GetMarketScores.
func (mr *MockMarketActivityTrackerMockRecorder) GetMarketScores(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketScores", reflect.TypeOf((*MockMarketActivityTracker)(nil).GetMarketScores), arg0, arg1, arg2)
}

// GetMarketsWithEligibleProposer mocks base method.
func (m *MockMarketActivityTracker) GetMarketsWithEligibleProposer(arg0 string, arg1 []string) []*types.MarketContributionScore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketsWithEligibleProposer", arg0, arg1)
	ret0, _ := ret[0].([]*types.MarketContributionScore)
	return ret0
}

// GetMarketsWithEligibleProposer indicates an expected call of GetMarketsWithEligibleProposer.
func (mr *MockMarketActivityTrackerMockRecorder) GetMarketsWithEligibleProposer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketsWithEligibleProposer", reflect.TypeOf((*MockMarketActivityTracker)(nil).GetMarketsWithEligibleProposer), arg0, arg1)
}
