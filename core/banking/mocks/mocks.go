// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/banking (interfaces: Assets,Notary,Collateral,Witness,TimeService,EpochService,Topology,MarketActivityTracker,ERC20BridgeView,EthereumEventSource)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	assets "code.vegaprotocol.io/vega/core/assets"
	types "code.vegaprotocol.io/vega/core/types"
	validators "code.vegaprotocol.io/vega/core/validators"
	num "code.vegaprotocol.io/vega/libs/num"
	vega "code.vegaprotocol.io/vega/protos/vega"
	v1 "code.vegaprotocol.io/vega/protos/vega/commands/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockAssets is a mock of Assets interface.
type MockAssets struct {
	ctrl     *gomock.Controller
	recorder *MockAssetsMockRecorder
}

// MockAssetsMockRecorder is the mock recorder for MockAssets.
type MockAssetsMockRecorder struct {
	mock *MockAssets
}

// NewMockAssets creates a new mock instance.
func NewMockAssets(ctrl *gomock.Controller) *MockAssets {
	mock := &MockAssets{ctrl: ctrl}
	mock.recorder = &MockAssetsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssets) EXPECT() *MockAssetsMockRecorder {
	return m.recorder
}

// ApplyAssetUpdate mocks base method.
func (m *MockAssets) ApplyAssetUpdate(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyAssetUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyAssetUpdate indicates an expected call of ApplyAssetUpdate.
func (mr *MockAssetsMockRecorder) ApplyAssetUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyAssetUpdate", reflect.TypeOf((*MockAssets)(nil).ApplyAssetUpdate), arg0, arg1)
}

// Enable mocks base method.
func (m *MockAssets) Enable(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockAssetsMockRecorder) Enable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockAssets)(nil).Enable), arg0, arg1)
}

// Get mocks base method.
func (m *MockAssets) Get(arg0 string) (*assets.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*assets.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAssetsMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAssets)(nil).Get), arg0)
}

// MockNotary is a mock of Notary interface.
type MockNotary struct {
	ctrl     *gomock.Controller
	recorder *MockNotaryMockRecorder
}

// MockNotaryMockRecorder is the mock recorder for MockNotary.
type MockNotaryMockRecorder struct {
	mock *MockNotary
}

// NewMockNotary creates a new mock instance.
func NewMockNotary(ctrl *gomock.Controller) *MockNotary {
	mock := &MockNotary{ctrl: ctrl}
	mock.recorder = &MockNotaryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotary) EXPECT() *MockNotaryMockRecorder {
	return m.recorder
}

// IsSigned mocks base method.
func (m *MockNotary) IsSigned(arg0 context.Context, arg1 string, arg2 v1.NodeSignatureKind) ([]v1.NodeSignature, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSigned", arg0, arg1, arg2)
	ret0, _ := ret[0].([]v1.NodeSignature)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// IsSigned indicates an expected call of IsSigned.
func (mr *MockNotaryMockRecorder) IsSigned(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSigned", reflect.TypeOf((*MockNotary)(nil).IsSigned), arg0, arg1, arg2)
}

// OfferSignatures mocks base method.
func (m *MockNotary) OfferSignatures(arg0 v1.NodeSignatureKind, arg1 func(string) []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OfferSignatures", arg0, arg1)
}

// OfferSignatures indicates an expected call of OfferSignatures.
func (mr *MockNotaryMockRecorder) OfferSignatures(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferSignatures", reflect.TypeOf((*MockNotary)(nil).OfferSignatures), arg0, arg1)
}

// StartAggregate mocks base method.
func (m *MockNotary) StartAggregate(arg0 string, arg1 v1.NodeSignatureKind, arg2 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartAggregate", arg0, arg1, arg2)
}

// StartAggregate indicates an expected call of StartAggregate.
func (mr *MockNotaryMockRecorder) StartAggregate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAggregate", reflect.TypeOf((*MockNotary)(nil).StartAggregate), arg0, arg1, arg2)
}

// MockCollateral is a mock of Collateral interface.
type MockCollateral struct {
	ctrl     *gomock.Controller
	recorder *MockCollateralMockRecorder
}

// MockCollateralMockRecorder is the mock recorder for MockCollateral.
type MockCollateralMockRecorder struct {
	mock *MockCollateral
}

// NewMockCollateral creates a new mock instance.
func NewMockCollateral(ctrl *gomock.Controller) *MockCollateral {
	mock := &MockCollateral{ctrl: ctrl}
	mock.recorder = &MockCollateralMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCollateral) EXPECT() *MockCollateralMockRecorder {
	return m.recorder
}

// Deposit mocks base method.
func (m *MockCollateral) Deposit(arg0 context.Context, arg1, arg2 string, arg3 *num.Uint) (*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deposit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deposit indicates an expected call of Deposit.
func (mr *MockCollateralMockRecorder) Deposit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deposit", reflect.TypeOf((*MockCollateral)(nil).Deposit), arg0, arg1, arg2, arg3)
}

// EnableAsset mocks base method.
func (m *MockCollateral) EnableAsset(arg0 context.Context, arg1 types.Asset) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableAsset", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableAsset indicates an expected call of EnableAsset.
func (mr *MockCollateralMockRecorder) EnableAsset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAsset", reflect.TypeOf((*MockCollateral)(nil).EnableAsset), arg0, arg1)
}

// GetPartyGeneralAccount mocks base method.
func (m *MockCollateral) GetPartyGeneralAccount(arg0, arg1 string) (*types.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyGeneralAccount", arg0, arg1)
	ret0, _ := ret[0].(*types.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyGeneralAccount indicates an expected call of GetPartyGeneralAccount.
func (mr *MockCollateralMockRecorder) GetPartyGeneralAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyGeneralAccount", reflect.TypeOf((*MockCollateral)(nil).GetPartyGeneralAccount), arg0, arg1)
}

// GetSystemAccountBalance mocks base method.
func (m *MockCollateral) GetSystemAccountBalance(arg0, arg1 string, arg2 vega.AccountType) (*num.Uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSystemAccountBalance", arg0, arg1, arg2)
	ret0, _ := ret[0].(*num.Uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSystemAccountBalance indicates an expected call of GetSystemAccountBalance.
func (mr *MockCollateralMockRecorder) GetSystemAccountBalance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemAccountBalance", reflect.TypeOf((*MockCollateral)(nil).GetSystemAccountBalance), arg0, arg1, arg2)
}

// GovernanceTransferFunds mocks base method.
func (m *MockCollateral) GovernanceTransferFunds(arg0 context.Context, arg1 []*types.Transfer, arg2 []vega.AccountType, arg3 []string) ([]*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GovernanceTransferFunds", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GovernanceTransferFunds indicates an expected call of GovernanceTransferFunds.
func (mr *MockCollateralMockRecorder) GovernanceTransferFunds(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GovernanceTransferFunds", reflect.TypeOf((*MockCollateral)(nil).GovernanceTransferFunds), arg0, arg1, arg2, arg3)
}

// PropagateAssetUpdate mocks base method.
func (m *MockCollateral) PropagateAssetUpdate(arg0 context.Context, arg1 types.Asset) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PropagateAssetUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PropagateAssetUpdate indicates an expected call of PropagateAssetUpdate.
func (mr *MockCollateralMockRecorder) PropagateAssetUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PropagateAssetUpdate", reflect.TypeOf((*MockCollateral)(nil).PropagateAssetUpdate), arg0, arg1)
}

// TransferFunds mocks base method.
func (m *MockCollateral) TransferFunds(arg0 context.Context, arg1 []*types.Transfer, arg2 []vega.AccountType, arg3 []string, arg4 []*types.Transfer, arg5 []vega.AccountType) ([]*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferFunds", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferFunds indicates an expected call of TransferFunds.
func (mr *MockCollateralMockRecorder) TransferFunds(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferFunds", reflect.TypeOf((*MockCollateral)(nil).TransferFunds), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Withdraw mocks base method.
func (m *MockCollateral) Withdraw(arg0 context.Context, arg1, arg2 string, arg3 *num.Uint) (*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockCollateralMockRecorder) Withdraw(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockCollateral)(nil).Withdraw), arg0, arg1, arg2, arg3)
}

// MockWitness is a mock of Witness interface.
type MockWitness struct {
	ctrl     *gomock.Controller
	recorder *MockWitnessMockRecorder
}

// MockWitnessMockRecorder is the mock recorder for MockWitness.
type MockWitnessMockRecorder struct {
	mock *MockWitness
}

// NewMockWitness creates a new mock instance.
func NewMockWitness(ctrl *gomock.Controller) *MockWitness {
	mock := &MockWitness{ctrl: ctrl}
	mock.recorder = &MockWitnessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWitness) EXPECT() *MockWitnessMockRecorder {
	return m.recorder
}

// RestoreResource mocks base method.
func (m *MockWitness) RestoreResource(arg0 validators.Resource, arg1 func(interface{}, bool)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreResource", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreResource indicates an expected call of RestoreResource.
func (mr *MockWitnessMockRecorder) RestoreResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreResource", reflect.TypeOf((*MockWitness)(nil).RestoreResource), arg0, arg1)
}

// StartCheck mocks base method.
func (m *MockWitness) StartCheck(arg0 validators.Resource, arg1 func(interface{}, bool), arg2 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCheck", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartCheck indicates an expected call of StartCheck.
func (mr *MockWitnessMockRecorder) StartCheck(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCheck", reflect.TypeOf((*MockWitness)(nil).StartCheck), arg0, arg1, arg2)
}

// MockTimeService is a mock of TimeService interface.
type MockTimeService struct {
	ctrl     *gomock.Controller
	recorder *MockTimeServiceMockRecorder
}

// MockTimeServiceMockRecorder is the mock recorder for MockTimeService.
type MockTimeServiceMockRecorder struct {
	mock *MockTimeService
}

// NewMockTimeService creates a new mock instance.
func NewMockTimeService(ctrl *gomock.Controller) *MockTimeService {
	mock := &MockTimeService{ctrl: ctrl}
	mock.recorder = &MockTimeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimeService) EXPECT() *MockTimeServiceMockRecorder {
	return m.recorder
}

// GetTimeNow mocks base method.
func (m *MockTimeService) GetTimeNow() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeNow")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimeNow indicates an expected call of GetTimeNow.
func (mr *MockTimeServiceMockRecorder) GetTimeNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeNow", reflect.TypeOf((*MockTimeService)(nil).GetTimeNow))
}

// MockEpochService is a mock of EpochService interface.
type MockEpochService struct {
	ctrl     *gomock.Controller
	recorder *MockEpochServiceMockRecorder
}

// MockEpochServiceMockRecorder is the mock recorder for MockEpochService.
type MockEpochServiceMockRecorder struct {
	mock *MockEpochService
}

// NewMockEpochService creates a new mock instance.
func NewMockEpochService(ctrl *gomock.Controller) *MockEpochService {
	mock := &MockEpochService{ctrl: ctrl}
	mock.recorder = &MockEpochServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEpochService) EXPECT() *MockEpochServiceMockRecorder {
	return m.recorder
}

// NotifyOnEpoch mocks base method.
func (m *MockEpochService) NotifyOnEpoch(arg0, arg1 func(context.Context, types.Epoch)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyOnEpoch", arg0, arg1)
}

// NotifyOnEpoch indicates an expected call of NotifyOnEpoch.
func (mr *MockEpochServiceMockRecorder) NotifyOnEpoch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyOnEpoch", reflect.TypeOf((*MockEpochService)(nil).NotifyOnEpoch), arg0, arg1)
}

// MockTopology is a mock of Topology interface.
type MockTopology struct {
	ctrl     *gomock.Controller
	recorder *MockTopologyMockRecorder
}

// MockTopologyMockRecorder is the mock recorder for MockTopology.
type MockTopologyMockRecorder struct {
	mock *MockTopology
}

// NewMockTopology creates a new mock instance.
func NewMockTopology(ctrl *gomock.Controller) *MockTopology {
	mock := &MockTopology{ctrl: ctrl}
	mock.recorder = &MockTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTopology) EXPECT() *MockTopologyMockRecorder {
	return m.recorder
}

// IsValidator mocks base method.
func (m *MockTopology) IsValidator() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidator")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidator indicates an expected call of IsValidator.
func (mr *MockTopologyMockRecorder) IsValidator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidator", reflect.TypeOf((*MockTopology)(nil).IsValidator))
}

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
func (m *MockMarketActivityTracker) GetMarketsWithEligibleProposer(arg0 string, arg1 []string, arg2, arg3 string) []*types.MarketContributionScore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketsWithEligibleProposer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*types.MarketContributionScore)
	return ret0
}

// GetMarketsWithEligibleProposer indicates an expected call of GetMarketsWithEligibleProposer.
func (mr *MockMarketActivityTrackerMockRecorder) GetMarketsWithEligibleProposer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketsWithEligibleProposer", reflect.TypeOf((*MockMarketActivityTracker)(nil).GetMarketsWithEligibleProposer), arg0, arg1, arg2, arg3)
}

// MarkPaidProposer mocks base method.
func (m *MockMarketActivityTracker) MarkPaidProposer(arg0, arg1 string, arg2 []string, arg3 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkPaidProposer", arg0, arg1, arg2, arg3)
}

// MarkPaidProposer indicates an expected call of MarkPaidProposer.
func (mr *MockMarketActivityTrackerMockRecorder) MarkPaidProposer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkPaidProposer", reflect.TypeOf((*MockMarketActivityTracker)(nil).MarkPaidProposer), arg0, arg1, arg2, arg3)
}

// MockERC20BridgeView is a mock of ERC20BridgeView interface.
type MockERC20BridgeView struct {
	ctrl     *gomock.Controller
	recorder *MockERC20BridgeViewMockRecorder
}

// MockERC20BridgeViewMockRecorder is the mock recorder for MockERC20BridgeView.
type MockERC20BridgeViewMockRecorder struct {
	mock *MockERC20BridgeView
}

// NewMockERC20BridgeView creates a new mock instance.
func NewMockERC20BridgeView(ctrl *gomock.Controller) *MockERC20BridgeView {
	mock := &MockERC20BridgeView{ctrl: ctrl}
	mock.recorder = &MockERC20BridgeViewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockERC20BridgeView) EXPECT() *MockERC20BridgeViewMockRecorder {
	return m.recorder
}

// FindAssetLimitsUpdated mocks base method.
func (m *MockERC20BridgeView) FindAssetLimitsUpdated(arg0 *types.ERC20AssetLimitsUpdated, arg1, arg2 uint64, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAssetLimitsUpdated", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindAssetLimitsUpdated indicates an expected call of FindAssetLimitsUpdated.
func (mr *MockERC20BridgeViewMockRecorder) FindAssetLimitsUpdated(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAssetLimitsUpdated", reflect.TypeOf((*MockERC20BridgeView)(nil).FindAssetLimitsUpdated), arg0, arg1, arg2, arg3)
}

// FindAssetList mocks base method.
func (m *MockERC20BridgeView) FindAssetList(arg0 *types.ERC20AssetList, arg1, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAssetList", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindAssetList indicates an expected call of FindAssetList.
func (mr *MockERC20BridgeViewMockRecorder) FindAssetList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAssetList", reflect.TypeOf((*MockERC20BridgeView)(nil).FindAssetList), arg0, arg1, arg2)
}

// FindBridgeResumed mocks base method.
func (m *MockERC20BridgeView) FindBridgeResumed(arg0 *types.ERC20EventBridgeResumed, arg1, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBridgeResumed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindBridgeResumed indicates an expected call of FindBridgeResumed.
func (mr *MockERC20BridgeViewMockRecorder) FindBridgeResumed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBridgeResumed", reflect.TypeOf((*MockERC20BridgeView)(nil).FindBridgeResumed), arg0, arg1, arg2)
}

// FindBridgeStopped mocks base method.
func (m *MockERC20BridgeView) FindBridgeStopped(arg0 *types.ERC20EventBridgeStopped, arg1, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBridgeStopped", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindBridgeStopped indicates an expected call of FindBridgeStopped.
func (mr *MockERC20BridgeViewMockRecorder) FindBridgeStopped(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBridgeStopped", reflect.TypeOf((*MockERC20BridgeView)(nil).FindBridgeStopped), arg0, arg1, arg2)
}

// FindDeposit mocks base method.
func (m *MockERC20BridgeView) FindDeposit(arg0 *types.ERC20Deposit, arg1, arg2 uint64, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeposit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindDeposit indicates an expected call of FindDeposit.
func (mr *MockERC20BridgeViewMockRecorder) FindDeposit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeposit", reflect.TypeOf((*MockERC20BridgeView)(nil).FindDeposit), arg0, arg1, arg2, arg3)
}

// MockEthereumEventSource is a mock of EthereumEventSource interface.
type MockEthereumEventSource struct {
	ctrl     *gomock.Controller
	recorder *MockEthereumEventSourceMockRecorder
}

// MockEthereumEventSourceMockRecorder is the mock recorder for MockEthereumEventSource.
type MockEthereumEventSourceMockRecorder struct {
	mock *MockEthereumEventSource
}

// NewMockEthereumEventSource creates a new mock instance.
func NewMockEthereumEventSource(ctrl *gomock.Controller) *MockEthereumEventSource {
	mock := &MockEthereumEventSource{ctrl: ctrl}
	mock.recorder = &MockEthereumEventSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthereumEventSource) EXPECT() *MockEthereumEventSourceMockRecorder {
	return m.recorder
}

// UpdateCollateralStartingBlock mocks base method.
func (m *MockEthereumEventSource) UpdateCollateralStartingBlock(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateCollateralStartingBlock", arg0)
}

// UpdateCollateralStartingBlock indicates an expected call of UpdateCollateralStartingBlock.
func (mr *MockEthereumEventSourceMockRecorder) UpdateCollateralStartingBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCollateralStartingBlock", reflect.TypeOf((*MockEthereumEventSource)(nil).UpdateCollateralStartingBlock), arg0)
}
