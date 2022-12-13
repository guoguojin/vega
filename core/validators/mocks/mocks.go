// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/validators (interfaces: NodeWallets,TimeService,Commander,ValidatorTopology,Wallet,ValidatorPerformance,Notary,Signatures,MultiSigTopology)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	txn "code.vegaprotocol.io/vega/core/txn"
	validators "code.vegaprotocol.io/vega/core/validators"
	crypto "code.vegaprotocol.io/vega/libs/crypto"
	v1 "code.vegaprotocol.io/vega/protos/vega/commands/v1"
	v10 "code.vegaprotocol.io/vega/protos/vega/snapshot/v1"
	backoff "github.com/cenkalti/backoff"
	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
)

// MockNodeWallets is a mock of NodeWallets interface.
type MockNodeWallets struct {
	ctrl     *gomock.Controller
	recorder *MockNodeWalletsMockRecorder
}

// MockNodeWalletsMockRecorder is the mock recorder for MockNodeWallets.
type MockNodeWalletsMockRecorder struct {
	mock *MockNodeWallets
}

// NewMockNodeWallets creates a new mock instance.
func NewMockNodeWallets(ctrl *gomock.Controller) *MockNodeWallets {
	mock := &MockNodeWallets{ctrl: ctrl}
	mock.recorder = &MockNodeWalletsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeWallets) EXPECT() *MockNodeWalletsMockRecorder {
	return m.recorder
}

// GetEthereum mocks base method.
func (m *MockNodeWallets) GetEthereum() validators.Signer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEthereum")
	ret0, _ := ret[0].(validators.Signer)
	return ret0
}

// GetEthereum indicates an expected call of GetEthereum.
func (mr *MockNodeWalletsMockRecorder) GetEthereum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEthereum", reflect.TypeOf((*MockNodeWallets)(nil).GetEthereum))
}

// GetEthereumAddress mocks base method.
func (m *MockNodeWallets) GetEthereumAddress() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEthereumAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEthereumAddress indicates an expected call of GetEthereumAddress.
func (mr *MockNodeWalletsMockRecorder) GetEthereumAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEthereumAddress", reflect.TypeOf((*MockNodeWallets)(nil).GetEthereumAddress))
}

// GetTendermintPubkey mocks base method.
func (m *MockNodeWallets) GetTendermintPubkey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTendermintPubkey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTendermintPubkey indicates an expected call of GetTendermintPubkey.
func (mr *MockNodeWalletsMockRecorder) GetTendermintPubkey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTendermintPubkey", reflect.TypeOf((*MockNodeWallets)(nil).GetTendermintPubkey))
}

// GetVega mocks base method.
func (m *MockNodeWallets) GetVega() validators.Wallet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVega")
	ret0, _ := ret[0].(validators.Wallet)
	return ret0
}

// GetVega indicates an expected call of GetVega.
func (mr *MockNodeWalletsMockRecorder) GetVega() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVega", reflect.TypeOf((*MockNodeWallets)(nil).GetVega))
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

// MockCommander is a mock of Commander interface.
type MockCommander struct {
	ctrl     *gomock.Controller
	recorder *MockCommanderMockRecorder
}

// MockCommanderMockRecorder is the mock recorder for MockCommander.
type MockCommanderMockRecorder struct {
	mock *MockCommander
}

// NewMockCommander creates a new mock instance.
func NewMockCommander(ctrl *gomock.Controller) *MockCommander {
	mock := &MockCommander{ctrl: ctrl}
	mock.recorder = &MockCommanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommander) EXPECT() *MockCommanderMockRecorder {
	return m.recorder
}

// Command mocks base method.
func (m *MockCommander) Command(arg0 context.Context, arg1 txn.Command, arg2 protoiface.MessageV1, arg3 func(string, error), arg4 *backoff.ExponentialBackOff) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Command", arg0, arg1, arg2, arg3, arg4)
}

// Command indicates an expected call of Command.
func (mr *MockCommanderMockRecorder) Command(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockCommander)(nil).Command), arg0, arg1, arg2, arg3, arg4)
}

// CommandSync mocks base method.
func (m *MockCommander) CommandSync(arg0 context.Context, arg1 txn.Command, arg2 protoiface.MessageV1, arg3 func(string, error), arg4 *backoff.ExponentialBackOff) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CommandSync", arg0, arg1, arg2, arg3, arg4)
}

// CommandSync indicates an expected call of CommandSync.
func (mr *MockCommanderMockRecorder) CommandSync(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandSync", reflect.TypeOf((*MockCommander)(nil).CommandSync), arg0, arg1, arg2, arg3, arg4)
}

// MockValidatorTopology is a mock of ValidatorTopology interface.
type MockValidatorTopology struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorTopologyMockRecorder
}

// MockValidatorTopologyMockRecorder is the mock recorder for MockValidatorTopology.
type MockValidatorTopologyMockRecorder struct {
	mock *MockValidatorTopology
}

// NewMockValidatorTopology creates a new mock instance.
func NewMockValidatorTopology(ctrl *gomock.Controller) *MockValidatorTopology {
	mock := &MockValidatorTopology{ctrl: ctrl}
	mock.recorder = &MockValidatorTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorTopology) EXPECT() *MockValidatorTopologyMockRecorder {
	return m.recorder
}

// AllNodeIDs mocks base method.
func (m *MockValidatorTopology) AllNodeIDs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllNodeIDs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// AllNodeIDs indicates an expected call of AllNodeIDs.
func (mr *MockValidatorTopologyMockRecorder) AllNodeIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllNodeIDs", reflect.TypeOf((*MockValidatorTopology)(nil).AllNodeIDs))
}

// GetTotalVotingPower mocks base method.
func (m *MockValidatorTopology) GetTotalVotingPower() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalVotingPower")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetTotalVotingPower indicates an expected call of GetTotalVotingPower.
func (mr *MockValidatorTopologyMockRecorder) GetTotalVotingPower() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalVotingPower", reflect.TypeOf((*MockValidatorTopology)(nil).GetTotalVotingPower))
}

// GetVotingPower mocks base method.
func (m *MockValidatorTopology) GetVotingPower(arg0 string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVotingPower", arg0)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetVotingPower indicates an expected call of GetVotingPower.
func (mr *MockValidatorTopologyMockRecorder) GetVotingPower(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotingPower", reflect.TypeOf((*MockValidatorTopology)(nil).GetVotingPower), arg0)
}

// IsTendermintValidator mocks base method.
func (m *MockValidatorTopology) IsTendermintValidator(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTendermintValidator", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTendermintValidator indicates an expected call of IsTendermintValidator.
func (mr *MockValidatorTopologyMockRecorder) IsTendermintValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTendermintValidator", reflect.TypeOf((*MockValidatorTopology)(nil).IsTendermintValidator), arg0)
}

// IsValidator mocks base method.
func (m *MockValidatorTopology) IsValidator() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidator")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidator indicates an expected call of IsValidator.
func (mr *MockValidatorTopologyMockRecorder) IsValidator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidator", reflect.TypeOf((*MockValidatorTopology)(nil).IsValidator))
}

// IsValidatorVegaPubKey mocks base method.
func (m *MockValidatorTopology) IsValidatorVegaPubKey(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidatorVegaPubKey", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidatorVegaPubKey indicates an expected call of IsValidatorVegaPubKey.
func (mr *MockValidatorTopologyMockRecorder) IsValidatorVegaPubKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidatorVegaPubKey", reflect.TypeOf((*MockValidatorTopology)(nil).IsValidatorVegaPubKey), arg0)
}

// SelfVegaPubKey mocks base method.
func (m *MockValidatorTopology) SelfVegaPubKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelfVegaPubKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// SelfVegaPubKey indicates an expected call of SelfVegaPubKey.
func (mr *MockValidatorTopologyMockRecorder) SelfVegaPubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelfVegaPubKey", reflect.TypeOf((*MockValidatorTopology)(nil).SelfVegaPubKey))
}

// MockWallet is a mock of Wallet interface.
type MockWallet struct {
	ctrl     *gomock.Controller
	recorder *MockWalletMockRecorder
}

// MockWalletMockRecorder is the mock recorder for MockWallet.
type MockWalletMockRecorder struct {
	mock *MockWallet
}

// NewMockWallet creates a new mock instance.
func NewMockWallet(ctrl *gomock.Controller) *MockWallet {
	mock := &MockWallet{ctrl: ctrl}
	mock.recorder = &MockWalletMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWallet) EXPECT() *MockWalletMockRecorder {
	return m.recorder
}

// Algo mocks base method.
func (m *MockWallet) Algo() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Algo")
	ret0, _ := ret[0].(string)
	return ret0
}

// Algo indicates an expected call of Algo.
func (mr *MockWalletMockRecorder) Algo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Algo", reflect.TypeOf((*MockWallet)(nil).Algo))
}

// ID mocks base method.
func (m *MockWallet) ID() crypto.PublicKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(crypto.PublicKey)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockWalletMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockWallet)(nil).ID))
}

// PubKey mocks base method.
func (m *MockWallet) PubKey() crypto.PublicKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubKey")
	ret0, _ := ret[0].(crypto.PublicKey)
	return ret0
}

// PubKey indicates an expected call of PubKey.
func (mr *MockWalletMockRecorder) PubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubKey", reflect.TypeOf((*MockWallet)(nil).PubKey))
}

// Sign mocks base method.
func (m *MockWallet) Sign(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockWalletMockRecorder) Sign(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockWallet)(nil).Sign), arg0)
}

// MockValidatorPerformance is a mock of ValidatorPerformance interface.
type MockValidatorPerformance struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorPerformanceMockRecorder
}

// MockValidatorPerformanceMockRecorder is the mock recorder for MockValidatorPerformance.
type MockValidatorPerformanceMockRecorder struct {
	mock *MockValidatorPerformance
}

// NewMockValidatorPerformance creates a new mock instance.
func NewMockValidatorPerformance(ctrl *gomock.Controller) *MockValidatorPerformance {
	mock := &MockValidatorPerformance{ctrl: ctrl}
	mock.recorder = &MockValidatorPerformanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorPerformance) EXPECT() *MockValidatorPerformanceMockRecorder {
	return m.recorder
}

// BeginBlock mocks base method.
func (m *MockValidatorPerformance) BeginBlock(arg0 context.Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeginBlock", arg0, arg1)
}

// BeginBlock indicates an expected call of BeginBlock.
func (mr *MockValidatorPerformanceMockRecorder) BeginBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginBlock", reflect.TypeOf((*MockValidatorPerformance)(nil).BeginBlock), arg0, arg1)
}

// Deserialize mocks base method.
func (m *MockValidatorPerformance) Deserialize(arg0 *v10.ValidatorPerformance) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Deserialize", arg0)
}

// Deserialize indicates an expected call of Deserialize.
func (mr *MockValidatorPerformanceMockRecorder) Deserialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deserialize", reflect.TypeOf((*MockValidatorPerformance)(nil).Deserialize), arg0)
}

// Reset mocks base method.
func (m *MockValidatorPerformance) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockValidatorPerformanceMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockValidatorPerformance)(nil).Reset))
}

// Serialize mocks base method.
func (m *MockValidatorPerformance) Serialize() *v10.ValidatorPerformance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serialize")
	ret0, _ := ret[0].(*v10.ValidatorPerformance)
	return ret0
}

// Serialize indicates an expected call of Serialize.
func (mr *MockValidatorPerformanceMockRecorder) Serialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serialize", reflect.TypeOf((*MockValidatorPerformance)(nil).Serialize))
}

// ValidatorPerformanceScore mocks base method.
func (m *MockValidatorPerformance) ValidatorPerformanceScore(arg0 string, arg1, arg2 int64, arg3 decimal.Decimal) decimal.Decimal {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorPerformanceScore", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(decimal.Decimal)
	return ret0
}

// ValidatorPerformanceScore indicates an expected call of ValidatorPerformanceScore.
func (mr *MockValidatorPerformanceMockRecorder) ValidatorPerformanceScore(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorPerformanceScore", reflect.TypeOf((*MockValidatorPerformance)(nil).ValidatorPerformanceScore), arg0, arg1, arg2, arg3)
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

// MockSignatures is a mock of Signatures interface.
type MockSignatures struct {
	ctrl     *gomock.Controller
	recorder *MockSignaturesMockRecorder
}

// MockSignaturesMockRecorder is the mock recorder for MockSignatures.
type MockSignaturesMockRecorder struct {
	mock *MockSignatures
}

// NewMockSignatures creates a new mock instance.
func NewMockSignatures(ctrl *gomock.Controller) *MockSignatures {
	mock := &MockSignatures{ctrl: ctrl}
	mock.recorder = &MockSignaturesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSignatures) EXPECT() *MockSignaturesMockRecorder {
	return m.recorder
}

// ClearStaleSignatures mocks base method.
func (m *MockSignatures) ClearStaleSignatures() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearStaleSignatures")
}

// ClearStaleSignatures indicates an expected call of ClearStaleSignatures.
func (mr *MockSignaturesMockRecorder) ClearStaleSignatures() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearStaleSignatures", reflect.TypeOf((*MockSignatures)(nil).ClearStaleSignatures))
}

// EmitValidatorAddedSignatures mocks base method.
func (m *MockSignatures) EmitValidatorAddedSignatures(arg0 context.Context, arg1, arg2 string, arg3 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmitValidatorAddedSignatures", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EmitValidatorAddedSignatures indicates an expected call of EmitValidatorAddedSignatures.
func (mr *MockSignaturesMockRecorder) EmitValidatorAddedSignatures(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitValidatorAddedSignatures", reflect.TypeOf((*MockSignatures)(nil).EmitValidatorAddedSignatures), arg0, arg1, arg2, arg3)
}

// EmitValidatorRemovedSignatures mocks base method.
func (m *MockSignatures) EmitValidatorRemovedSignatures(arg0 context.Context, arg1, arg2 string, arg3 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmitValidatorRemovedSignatures", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EmitValidatorRemovedSignatures indicates an expected call of EmitValidatorRemovedSignatures.
func (mr *MockSignaturesMockRecorder) EmitValidatorRemovedSignatures(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitValidatorRemovedSignatures", reflect.TypeOf((*MockSignatures)(nil).EmitValidatorRemovedSignatures), arg0, arg1, arg2, arg3)
}

// PreparePromotionsSignatures mocks base method.
func (m *MockSignatures) PreparePromotionsSignatures(arg0 context.Context, arg1 time.Time, arg2 uint64, arg3, arg4 map[string]validators.StatusAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PreparePromotionsSignatures", arg0, arg1, arg2, arg3, arg4)
}

// PreparePromotionsSignatures indicates an expected call of PreparePromotionsSignatures.
func (mr *MockSignaturesMockRecorder) PreparePromotionsSignatures(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreparePromotionsSignatures", reflect.TypeOf((*MockSignatures)(nil).PreparePromotionsSignatures), arg0, arg1, arg2, arg3, arg4)
}

// PrepareValidatorSignatures mocks base method.
func (m *MockSignatures) PrepareValidatorSignatures(arg0 context.Context, arg1 []validators.NodeIDAddress, arg2 uint64, arg3 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PrepareValidatorSignatures", arg0, arg1, arg2, arg3)
}

// PrepareValidatorSignatures indicates an expected call of PrepareValidatorSignatures.
func (mr *MockSignaturesMockRecorder) PrepareValidatorSignatures(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareValidatorSignatures", reflect.TypeOf((*MockSignatures)(nil).PrepareValidatorSignatures), arg0, arg1, arg2, arg3)
}

// RestorePendingSignatures mocks base method.
func (m *MockSignatures) RestorePendingSignatures(arg0 *v10.ToplogySignatures) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RestorePendingSignatures", arg0)
}

// RestorePendingSignatures indicates an expected call of RestorePendingSignatures.
func (mr *MockSignaturesMockRecorder) RestorePendingSignatures(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestorePendingSignatures", reflect.TypeOf((*MockSignatures)(nil).RestorePendingSignatures), arg0)
}

// SerialisePendingSignatures mocks base method.
func (m *MockSignatures) SerialisePendingSignatures() *v10.ToplogySignatures {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SerialisePendingSignatures")
	ret0, _ := ret[0].(*v10.ToplogySignatures)
	return ret0
}

// SerialisePendingSignatures indicates an expected call of SerialisePendingSignatures.
func (mr *MockSignaturesMockRecorder) SerialisePendingSignatures() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SerialisePendingSignatures", reflect.TypeOf((*MockSignatures)(nil).SerialisePendingSignatures))
}

// MockMultiSigTopology is a mock of MultiSigTopology interface.
type MockMultiSigTopology struct {
	ctrl     *gomock.Controller
	recorder *MockMultiSigTopologyMockRecorder
}

// MockMultiSigTopologyMockRecorder is the mock recorder for MockMultiSigTopology.
type MockMultiSigTopologyMockRecorder struct {
	mock *MockMultiSigTopology
}

// NewMockMultiSigTopology creates a new mock instance.
func NewMockMultiSigTopology(ctrl *gomock.Controller) *MockMultiSigTopology {
	mock := &MockMultiSigTopology{ctrl: ctrl}
	mock.recorder = &MockMultiSigTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMultiSigTopology) EXPECT() *MockMultiSigTopologyMockRecorder {
	return m.recorder
}

// ExcessSigners mocks base method.
func (m *MockMultiSigTopology) ExcessSigners(arg0 []string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExcessSigners", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExcessSigners indicates an expected call of ExcessSigners.
func (mr *MockMultiSigTopologyMockRecorder) ExcessSigners(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExcessSigners", reflect.TypeOf((*MockMultiSigTopology)(nil).ExcessSigners), arg0)
}

// IsSigner mocks base method.
func (m *MockMultiSigTopology) IsSigner(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSigner", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSigner indicates an expected call of IsSigner.
func (mr *MockMultiSigTopologyMockRecorder) IsSigner(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSigner", reflect.TypeOf((*MockMultiSigTopology)(nil).IsSigner), arg0)
}
