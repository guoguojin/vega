// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/execution (interfaces: TimeService,Assets,StateVarEngine,Collateral,OracleEngine,EpochEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	assets "code.vegaprotocol.io/vega/core/assets"
	events "code.vegaprotocol.io/vega/core/events"
	oracles "code.vegaprotocol.io/vega/core/oracles"
	types "code.vegaprotocol.io/vega/core/types"
	statevar "code.vegaprotocol.io/vega/core/types/statevar"
	num "code.vegaprotocol.io/vega/libs/num"
	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
)

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

// MockStateVarEngine is a mock of StateVarEngine interface.
type MockStateVarEngine struct {
	ctrl     *gomock.Controller
	recorder *MockStateVarEngineMockRecorder
}

// MockStateVarEngineMockRecorder is the mock recorder for MockStateVarEngine.
type MockStateVarEngineMockRecorder struct {
	mock *MockStateVarEngine
}

// NewMockStateVarEngine creates a new mock instance.
func NewMockStateVarEngine(ctrl *gomock.Controller) *MockStateVarEngine {
	mock := &MockStateVarEngine{ctrl: ctrl}
	mock.recorder = &MockStateVarEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateVarEngine) EXPECT() *MockStateVarEngineMockRecorder {
	return m.recorder
}

// NewEvent mocks base method.
func (m *MockStateVarEngine) NewEvent(arg0, arg1 string, arg2 statevar.EventType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewEvent", arg0, arg1, arg2)
}

// NewEvent indicates an expected call of NewEvent.
func (mr *MockStateVarEngineMockRecorder) NewEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEvent", reflect.TypeOf((*MockStateVarEngine)(nil).NewEvent), arg0, arg1, arg2)
}

// ReadyForTimeTrigger mocks base method.
func (m *MockStateVarEngine) ReadyForTimeTrigger(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReadyForTimeTrigger", arg0, arg1)
}

// ReadyForTimeTrigger indicates an expected call of ReadyForTimeTrigger.
func (mr *MockStateVarEngineMockRecorder) ReadyForTimeTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadyForTimeTrigger", reflect.TypeOf((*MockStateVarEngine)(nil).ReadyForTimeTrigger), arg0, arg1)
}

// RegisterStateVariable mocks base method.
func (m *MockStateVarEngine) RegisterStateVariable(arg0, arg1, arg2 string, arg3 statevar.Converter, arg4 func(string, statevar.FinaliseCalculation), arg5 []statevar.EventType, arg6 func(context.Context, statevar.StateVariableResult) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterStateVariable", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterStateVariable indicates an expected call of RegisterStateVariable.
func (mr *MockStateVarEngineMockRecorder) RegisterStateVariable(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterStateVariable", reflect.TypeOf((*MockStateVarEngine)(nil).RegisterStateVariable), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// UnregisterStateVariable mocks base method.
func (m *MockStateVarEngine) UnregisterStateVariable(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnregisterStateVariable", arg0, arg1)
}

// UnregisterStateVariable indicates an expected call of UnregisterStateVariable.
func (mr *MockStateVarEngineMockRecorder) UnregisterStateVariable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterStateVariable", reflect.TypeOf((*MockStateVarEngine)(nil).UnregisterStateVariable), arg0, arg1)
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

// AssetExists mocks base method.
func (m *MockCollateral) AssetExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssetExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AssetExists indicates an expected call of AssetExists.
func (mr *MockCollateralMockRecorder) AssetExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssetExists", reflect.TypeOf((*MockCollateral)(nil).AssetExists), arg0)
}

// BondUpdate mocks base method.
func (m *MockCollateral) BondUpdate(arg0 context.Context, arg1 string, arg2 *types.Transfer) (*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BondUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BondUpdate indicates an expected call of BondUpdate.
func (mr *MockCollateralMockRecorder) BondUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BondUpdate", reflect.TypeOf((*MockCollateral)(nil).BondUpdate), arg0, arg1, arg2)
}

// CanCoverBond mocks base method.
func (m *MockCollateral) CanCoverBond(arg0, arg1, arg2 string, arg3 *num.Uint) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanCoverBond", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanCoverBond indicates an expected call of CanCoverBond.
func (mr *MockCollateralMockRecorder) CanCoverBond(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanCoverBond", reflect.TypeOf((*MockCollateral)(nil).CanCoverBond), arg0, arg1, arg2, arg3)
}

// ClearMarket mocks base method.
func (m *MockCollateral) ClearMarket(arg0 context.Context, arg1, arg2 string, arg3 []string) ([]*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearMarket", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClearMarket indicates an expected call of ClearMarket.
func (mr *MockCollateralMockRecorder) ClearMarket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearMarket", reflect.TypeOf((*MockCollateral)(nil).ClearMarket), arg0, arg1, arg2, arg3)
}

// ClearPartyMarginAccount mocks base method.
func (m *MockCollateral) ClearPartyMarginAccount(arg0 context.Context, arg1, arg2, arg3 string) (*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearPartyMarginAccount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClearPartyMarginAccount indicates an expected call of ClearPartyMarginAccount.
func (mr *MockCollateralMockRecorder) ClearPartyMarginAccount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearPartyMarginAccount", reflect.TypeOf((*MockCollateral)(nil).ClearPartyMarginAccount), arg0, arg1, arg2, arg3)
}

// CreateMarketAccounts mocks base method.
func (m *MockCollateral) CreateMarketAccounts(arg0 context.Context, arg1, arg2 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMarketAccounts", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateMarketAccounts indicates an expected call of CreateMarketAccounts.
func (mr *MockCollateralMockRecorder) CreateMarketAccounts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMarketAccounts", reflect.TypeOf((*MockCollateral)(nil).CreateMarketAccounts), arg0, arg1, arg2)
}

// CreatePartyMarginAccount mocks base method.
func (m *MockCollateral) CreatePartyMarginAccount(arg0 context.Context, arg1, arg2, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePartyMarginAccount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePartyMarginAccount indicates an expected call of CreatePartyMarginAccount.
func (mr *MockCollateralMockRecorder) CreatePartyMarginAccount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePartyMarginAccount", reflect.TypeOf((*MockCollateral)(nil).CreatePartyMarginAccount), arg0, arg1, arg2, arg3)
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

// FinalSettlement mocks base method.
func (m *MockCollateral) FinalSettlement(arg0 context.Context, arg1 string, arg2 []*types.Transfer) ([]*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalSettlement", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FinalSettlement indicates an expected call of FinalSettlement.
func (mr *MockCollateralMockRecorder) FinalSettlement(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalSettlement", reflect.TypeOf((*MockCollateral)(nil).FinalSettlement), arg0, arg1, arg2)
}

// GetAssetQuantum mocks base method.
func (m *MockCollateral) GetAssetQuantum(arg0 string) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetQuantum", arg0)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetQuantum indicates an expected call of GetAssetQuantum.
func (mr *MockCollateralMockRecorder) GetAssetQuantum(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetQuantum", reflect.TypeOf((*MockCollateral)(nil).GetAssetQuantum), arg0)
}

// GetInsurancePoolBalance mocks base method.
func (m *MockCollateral) GetInsurancePoolBalance(arg0, arg1 string) (*num.Uint, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInsurancePoolBalance", arg0, arg1)
	ret0, _ := ret[0].(*num.Uint)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetInsurancePoolBalance indicates an expected call of GetInsurancePoolBalance.
func (mr *MockCollateralMockRecorder) GetInsurancePoolBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsurancePoolBalance", reflect.TypeOf((*MockCollateral)(nil).GetInsurancePoolBalance), arg0, arg1)
}

// GetMarketLiquidityFeeAccount mocks base method.
func (m *MockCollateral) GetMarketLiquidityFeeAccount(arg0, arg1 string) (*types.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketLiquidityFeeAccount", arg0, arg1)
	ret0, _ := ret[0].(*types.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketLiquidityFeeAccount indicates an expected call of GetMarketLiquidityFeeAccount.
func (mr *MockCollateralMockRecorder) GetMarketLiquidityFeeAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketLiquidityFeeAccount", reflect.TypeOf((*MockCollateral)(nil).GetMarketLiquidityFeeAccount), arg0, arg1)
}

// GetOrCreatePartyBondAccount mocks base method.
func (m *MockCollateral) GetOrCreatePartyBondAccount(arg0 context.Context, arg1, arg2, arg3 string) (*types.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreatePartyBondAccount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreatePartyBondAccount indicates an expected call of GetOrCreatePartyBondAccount.
func (mr *MockCollateralMockRecorder) GetOrCreatePartyBondAccount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreatePartyBondAccount", reflect.TypeOf((*MockCollateral)(nil).GetOrCreatePartyBondAccount), arg0, arg1, arg2, arg3)
}

// GetPartyBondAccount mocks base method.
func (m *MockCollateral) GetPartyBondAccount(arg0, arg1, arg2 string) (*types.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyBondAccount", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyBondAccount indicates an expected call of GetPartyBondAccount.
func (mr *MockCollateralMockRecorder) GetPartyBondAccount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyBondAccount", reflect.TypeOf((*MockCollateral)(nil).GetPartyBondAccount), arg0, arg1, arg2)
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

// GetPartyMargin mocks base method.
func (m *MockCollateral) GetPartyMargin(arg0 events.MarketPosition, arg1, arg2 string) (events.Margin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyMargin", arg0, arg1, arg2)
	ret0, _ := ret[0].(events.Margin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyMargin indicates an expected call of GetPartyMargin.
func (mr *MockCollateralMockRecorder) GetPartyMargin(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyMargin", reflect.TypeOf((*MockCollateral)(nil).GetPartyMargin), arg0, arg1, arg2)
}

// GetPartyMarginAccount mocks base method.
func (m *MockCollateral) GetPartyMarginAccount(arg0, arg1, arg2 string) (*types.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyMarginAccount", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyMarginAccount indicates an expected call of GetPartyMarginAccount.
func (mr *MockCollateralMockRecorder) GetPartyMarginAccount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyMarginAccount", reflect.TypeOf((*MockCollateral)(nil).GetPartyMarginAccount), arg0, arg1, arg2)
}

// HasGeneralAccount mocks base method.
func (m *MockCollateral) HasGeneralAccount(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasGeneralAccount", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasGeneralAccount indicates an expected call of HasGeneralAccount.
func (mr *MockCollateralMockRecorder) HasGeneralAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasGeneralAccount", reflect.TypeOf((*MockCollateral)(nil).HasGeneralAccount), arg0, arg1)
}

// Hash mocks base method.
func (m *MockCollateral) Hash() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Hash indicates an expected call of Hash.
func (mr *MockCollateralMockRecorder) Hash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockCollateral)(nil).Hash))
}

// MarginUpdate mocks base method.
func (m *MockCollateral) MarginUpdate(arg0 context.Context, arg1 string, arg2 []events.Risk) ([]*types.LedgerMovement, []events.Margin, []events.Margin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarginUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.LedgerMovement)
	ret1, _ := ret[1].([]events.Margin)
	ret2, _ := ret[2].([]events.Margin)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// MarginUpdate indicates an expected call of MarginUpdate.
func (mr *MockCollateralMockRecorder) MarginUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarginUpdate", reflect.TypeOf((*MockCollateral)(nil).MarginUpdate), arg0, arg1, arg2)
}

// MarginUpdateOnOrder mocks base method.
func (m *MockCollateral) MarginUpdateOnOrder(arg0 context.Context, arg1 string, arg2 events.Risk) (*types.LedgerMovement, events.Margin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarginUpdateOnOrder", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.LedgerMovement)
	ret1, _ := ret[1].(events.Margin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MarginUpdateOnOrder indicates an expected call of MarginUpdateOnOrder.
func (mr *MockCollateralMockRecorder) MarginUpdateOnOrder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarginUpdateOnOrder", reflect.TypeOf((*MockCollateral)(nil).MarginUpdateOnOrder), arg0, arg1, arg2)
}

// MarkToMarket mocks base method.
func (m *MockCollateral) MarkToMarket(arg0 context.Context, arg1 string, arg2 []events.Transfer, arg3 string) ([]events.Margin, []*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkToMarket", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]events.Margin)
	ret1, _ := ret[1].([]*types.LedgerMovement)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MarkToMarket indicates an expected call of MarkToMarket.
func (mr *MockCollateralMockRecorder) MarkToMarket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkToMarket", reflect.TypeOf((*MockCollateral)(nil).MarkToMarket), arg0, arg1, arg2, arg3)
}

// RemoveBondAccount mocks base method.
func (m *MockCollateral) RemoveBondAccount(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBondAccount", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveBondAccount indicates an expected call of RemoveBondAccount.
func (mr *MockCollateralMockRecorder) RemoveBondAccount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBondAccount", reflect.TypeOf((*MockCollateral)(nil).RemoveBondAccount), arg0, arg1, arg2)
}

// RemoveDistressed mocks base method.
func (m *MockCollateral) RemoveDistressed(arg0 context.Context, arg1 []events.MarketPosition, arg2, arg3 string) (*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDistressed", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveDistressed indicates an expected call of RemoveDistressed.
func (mr *MockCollateralMockRecorder) RemoveDistressed(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDistressed", reflect.TypeOf((*MockCollateral)(nil).RemoveDistressed), arg0, arg1, arg2, arg3)
}

// RollbackMarginUpdateOnOrder mocks base method.
func (m *MockCollateral) RollbackMarginUpdateOnOrder(arg0 context.Context, arg1, arg2 string, arg3 *types.Transfer) (*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackMarginUpdateOnOrder", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RollbackMarginUpdateOnOrder indicates an expected call of RollbackMarginUpdateOnOrder.
func (mr *MockCollateralMockRecorder) RollbackMarginUpdateOnOrder(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackMarginUpdateOnOrder", reflect.TypeOf((*MockCollateral)(nil).RollbackMarginUpdateOnOrder), arg0, arg1, arg2, arg3)
}

// SuccessorInsuranceFraction mocks base method.
func (m *MockCollateral) SuccessorInsuranceFraction(arg0 context.Context, arg1, arg2, arg3 string, arg4 decimal.Decimal) *types.LedgerMovement {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuccessorInsuranceFraction", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types.LedgerMovement)
	return ret0
}

// SuccessorInsuranceFraction indicates an expected call of SuccessorInsuranceFraction.
func (mr *MockCollateralMockRecorder) SuccessorInsuranceFraction(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuccessorInsuranceFraction", reflect.TypeOf((*MockCollateral)(nil).SuccessorInsuranceFraction), arg0, arg1, arg2, arg3, arg4)
}

// TransferFees mocks base method.
func (m *MockCollateral) TransferFees(arg0 context.Context, arg1, arg2 string, arg3 events.FeesTransfer) ([]*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferFees", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferFees indicates an expected call of TransferFees.
func (mr *MockCollateralMockRecorder) TransferFees(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferFees", reflect.TypeOf((*MockCollateral)(nil).TransferFees), arg0, arg1, arg2, arg3)
}

// TransferFeesContinuousTrading mocks base method.
func (m *MockCollateral) TransferFeesContinuousTrading(arg0 context.Context, arg1, arg2 string, arg3 events.FeesTransfer) ([]*types.LedgerMovement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferFeesContinuousTrading", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*types.LedgerMovement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferFeesContinuousTrading indicates an expected call of TransferFeesContinuousTrading.
func (mr *MockCollateralMockRecorder) TransferFeesContinuousTrading(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferFeesContinuousTrading", reflect.TypeOf((*MockCollateral)(nil).TransferFeesContinuousTrading), arg0, arg1, arg2, arg3)
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

// MockOracleEngine is a mock of OracleEngine interface.
type MockOracleEngine struct {
	ctrl     *gomock.Controller
	recorder *MockOracleEngineMockRecorder
}

// MockOracleEngineMockRecorder is the mock recorder for MockOracleEngine.
type MockOracleEngineMockRecorder struct {
	mock *MockOracleEngine
}

// NewMockOracleEngine creates a new mock instance.
func NewMockOracleEngine(ctrl *gomock.Controller) *MockOracleEngine {
	mock := &MockOracleEngine{ctrl: ctrl}
	mock.recorder = &MockOracleEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOracleEngine) EXPECT() *MockOracleEngineMockRecorder {
	return m.recorder
}

// ListensToSigners mocks base method.
func (m *MockOracleEngine) ListensToSigners(arg0 oracles.OracleData) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListensToSigners", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ListensToSigners indicates an expected call of ListensToSigners.
func (mr *MockOracleEngineMockRecorder) ListensToSigners(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListensToSigners", reflect.TypeOf((*MockOracleEngine)(nil).ListensToSigners), arg0)
}

// Subscribe mocks base method.
func (m *MockOracleEngine) Subscribe(arg0 context.Context, arg1 oracles.OracleSpec, arg2 oracles.OnMatchedOracleData) (oracles.SubscriptionID, oracles.Unsubscriber) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1, arg2)
	ret0, _ := ret[0].(oracles.SubscriptionID)
	ret1, _ := ret[1].(oracles.Unsubscriber)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockOracleEngineMockRecorder) Subscribe(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockOracleEngine)(nil).Subscribe), arg0, arg1, arg2)
}

// Unsubscribe mocks base method.
func (m *MockOracleEngine) Unsubscribe(arg0 context.Context, arg1 oracles.SubscriptionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unsubscribe", arg0, arg1)
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockOracleEngineMockRecorder) Unsubscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockOracleEngine)(nil).Unsubscribe), arg0, arg1)
}

// MockEpochEngine is a mock of EpochEngine interface.
type MockEpochEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEpochEngineMockRecorder
}

// MockEpochEngineMockRecorder is the mock recorder for MockEpochEngine.
type MockEpochEngineMockRecorder struct {
	mock *MockEpochEngine
}

// NewMockEpochEngine creates a new mock instance.
func NewMockEpochEngine(ctrl *gomock.Controller) *MockEpochEngine {
	mock := &MockEpochEngine{ctrl: ctrl}
	mock.recorder = &MockEpochEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEpochEngine) EXPECT() *MockEpochEngineMockRecorder {
	return m.recorder
}

// NotifyOnEpoch mocks base method.
func (m *MockEpochEngine) NotifyOnEpoch(arg0, arg1 func(context.Context, types.Epoch)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyOnEpoch", arg0, arg1)
}

// NotifyOnEpoch indicates an expected call of NotifyOnEpoch.
func (mr *MockEpochEngineMockRecorder) NotifyOnEpoch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyOnEpoch", reflect.TypeOf((*MockEpochEngine)(nil).NotifyOnEpoch), arg0, arg1)
}
