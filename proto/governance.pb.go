// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/governance.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Proposal state transition:
// Open ->
//   - Passed -> Enacted.
//   - Passed -> Failed.
//   - Declined
// Rejected
// Proposal can enter Failed state from any other state.
type Proposal_State int32

const (
	// Default value, always invalid.
	Proposal_STATE_UNSPECIFIED Proposal_State = 0
	// Proposal enactment has failed - even though proposal has passed, its execusion could not be performed.
	Proposal_STATE_FAILED Proposal_State = 1
	// Proposal is open for voting.
	Proposal_STATE_OPEN Proposal_State = 2
	// Proposal has gained enough support to be executed.
	Proposal_STATE_PASSED Proposal_State = 3
	// Proposal wasn't accepted (proposal terms failed validation due to wrong configuration or failing to meet network requirements).
	Proposal_STATE_REJECTED Proposal_State = 4
	// Proposal didn't get enough votes (either failing to gain required participation or majority level).
	Proposal_STATE_DECLINED Proposal_State = 5
	// Proposal has been executed and the changes under this proposal have now been applied.
	Proposal_STATE_ENACTED Proposal_State = 6
)

var Proposal_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "STATE_FAILED",
	2: "STATE_OPEN",
	3: "STATE_PASSED",
	4: "STATE_REJECTED",
	5: "STATE_DECLINED",
	6: "STATE_ENACTED",
}

var Proposal_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"STATE_FAILED":      1,
	"STATE_OPEN":        2,
	"STATE_PASSED":      3,
	"STATE_REJECTED":    4,
	"STATE_DECLINED":    5,
	"STATE_ENACTED":     6,
}

func (x Proposal_State) String() string {
	return proto.EnumName(Proposal_State_name, int32(x))
}

func (Proposal_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{7, 0}
}

type Vote_Value int32

const (
	// Default value, always invalid.
	Vote_VALUE_UNSPECIFIED Vote_Value = 0
	// A vote against the proposal.
	Vote_VALUE_NO Vote_Value = 1
	// A vote in favour of the proposal.
	Vote_VALUE_YES Vote_Value = 2
)

var Vote_Value_name = map[int32]string{
	0: "VALUE_UNSPECIFIED",
	1: "VALUE_NO",
	2: "VALUE_YES",
}

var Vote_Value_value = map[string]int32{
	"VALUE_UNSPECIFIED": 0,
	"VALUE_NO":          1,
	"VALUE_YES":         2,
}

func (x Vote_Value) String() string {
	return proto.EnumName(Vote_Value_name, int32(x))
}

func (Vote_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{8, 0}
}

type NetworkConfiguration struct {
	// Constrains minimum duration since submission (in seconds) when vote closing time is allowed to be set for a proposal.
	MinCloseInSeconds int64 `protobuf:"varint,1,opt,name=minCloseInSeconds,proto3" json:"minCloseInSeconds,omitempty"`
	// Constrains maximum duration since submission (in seconds) when vote closing time is allowed to be set for a proposal.
	MaxCloseInSeconds int64 `protobuf:"varint,2,opt,name=maxCloseInSeconds,proto3" json:"maxCloseInSeconds,omitempty"`
	// Constrains minimum duration since submission (in seconds) when enactment is allowed to be set for a proposal.
	MinEnactInSeconds int64 `protobuf:"varint,3,opt,name=minEnactInSeconds,proto3" json:"minEnactInSeconds,omitempty"`
	// Constrains maximum duration since submission (in seconds) when enactment is allowed to be set for a proposal.
	MaxEnactInSeconds int64 `protobuf:"varint,4,opt,name=maxEnactInSeconds,proto3" json:"maxEnactInSeconds,omitempty"`
	// Participation level required for any proposal to pass. Value from `0` to `1`.
	RequiredParticipation float32 `protobuf:"fixed32,5,opt,name=requiredParticipation,proto3" json:"requiredParticipation,omitempty"`
	// Majority level required for any proposal to pass. Value from `0.5` to `1`.
	RequiredMajority float32 `protobuf:"fixed32,6,opt,name=requiredMajority,proto3" json:"requiredMajority,omitempty"`
	// Minimum balance required for a party to be able to submit a new proposal. Value greater than `0` to `1`.
	MinProposingBalance float32 `protobuf:"fixed32,7,opt,name=minProposingBalance,proto3" json:"minProposingBalance,omitempty"`
	// Minimum balance required for a party to be able to cast a vote. Value greater than `0` to `1`.
	MinVotingBalance     float32  `protobuf:"fixed32,8,opt,name=minVotingBalance,proto3" json:"minVotingBalance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkConfiguration) Reset()         { *m = NetworkConfiguration{} }
func (m *NetworkConfiguration) String() string { return proto.CompactTextString(m) }
func (*NetworkConfiguration) ProtoMessage()    {}
func (*NetworkConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{0}
}

func (m *NetworkConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkConfiguration.Unmarshal(m, b)
}
func (m *NetworkConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkConfiguration.Marshal(b, m, deterministic)
}
func (m *NetworkConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkConfiguration.Merge(m, src)
}
func (m *NetworkConfiguration) XXX_Size() int {
	return xxx_messageInfo_NetworkConfiguration.Size(m)
}
func (m *NetworkConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkConfiguration proto.InternalMessageInfo

func (m *NetworkConfiguration) GetMinCloseInSeconds() int64 {
	if m != nil {
		return m.MinCloseInSeconds
	}
	return 0
}

func (m *NetworkConfiguration) GetMaxCloseInSeconds() int64 {
	if m != nil {
		return m.MaxCloseInSeconds
	}
	return 0
}

func (m *NetworkConfiguration) GetMinEnactInSeconds() int64 {
	if m != nil {
		return m.MinEnactInSeconds
	}
	return 0
}

func (m *NetworkConfiguration) GetMaxEnactInSeconds() int64 {
	if m != nil {
		return m.MaxEnactInSeconds
	}
	return 0
}

func (m *NetworkConfiguration) GetRequiredParticipation() float32 {
	if m != nil {
		return m.RequiredParticipation
	}
	return 0
}

func (m *NetworkConfiguration) GetRequiredMajority() float32 {
	if m != nil {
		return m.RequiredMajority
	}
	return 0
}

func (m *NetworkConfiguration) GetMinProposingBalance() float32 {
	if m != nil {
		return m.MinProposingBalance
	}
	return 0
}

func (m *NetworkConfiguration) GetMinVotingBalance() float32 {
	if m != nil {
		return m.MinVotingBalance
	}
	return 0
}

type UpdateMarket struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMarket) Reset()         { *m = UpdateMarket{} }
func (m *UpdateMarket) String() string { return proto.CompactTextString(m) }
func (*UpdateMarket) ProtoMessage()    {}
func (*UpdateMarket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{1}
}

func (m *UpdateMarket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMarket.Unmarshal(m, b)
}
func (m *UpdateMarket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMarket.Marshal(b, m, deterministic)
}
func (m *UpdateMarket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMarket.Merge(m, src)
}
func (m *UpdateMarket) XXX_Size() int {
	return xxx_messageInfo_UpdateMarket.Size(m)
}
func (m *UpdateMarket) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMarket.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMarket proto.InternalMessageInfo

type NewMarket struct {
	Changes              *Market  `protobuf:"bytes,1,opt,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewMarket) Reset()         { *m = NewMarket{} }
func (m *NewMarket) String() string { return proto.CompactTextString(m) }
func (*NewMarket) ProtoMessage()    {}
func (*NewMarket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{2}
}

func (m *NewMarket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewMarket.Unmarshal(m, b)
}
func (m *NewMarket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewMarket.Marshal(b, m, deterministic)
}
func (m *NewMarket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewMarket.Merge(m, src)
}
func (m *NewMarket) XXX_Size() int {
	return xxx_messageInfo_NewMarket.Size(m)
}
func (m *NewMarket) XXX_DiscardUnknown() {
	xxx_messageInfo_NewMarket.DiscardUnknown(m)
}

var xxx_messageInfo_NewMarket proto.InternalMessageInfo

func (m *NewMarket) GetChanges() *Market {
	if m != nil {
		return m.Changes
	}
	return nil
}

type UpdateNetwork struct {
	Changes              *NetworkConfiguration `protobuf:"bytes,1,opt,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateNetwork) Reset()         { *m = UpdateNetwork{} }
func (m *UpdateNetwork) String() string { return proto.CompactTextString(m) }
func (*UpdateNetwork) ProtoMessage()    {}
func (*UpdateNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{3}
}

func (m *UpdateNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNetwork.Unmarshal(m, b)
}
func (m *UpdateNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNetwork.Marshal(b, m, deterministic)
}
func (m *UpdateNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNetwork.Merge(m, src)
}
func (m *UpdateNetwork) XXX_Size() int {
	return xxx_messageInfo_UpdateNetwork.Size(m)
}
func (m *UpdateNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNetwork proto.InternalMessageInfo

func (m *UpdateNetwork) GetChanges() *NetworkConfiguration {
	if m != nil {
		return m.Changes
	}
	return nil
}

// To be implemented
type NewAsset struct {
	Changes              *AssetSource `protobuf:"bytes,1,opt,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NewAsset) Reset()         { *m = NewAsset{} }
func (m *NewAsset) String() string { return proto.CompactTextString(m) }
func (*NewAsset) ProtoMessage()    {}
func (*NewAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{4}
}

func (m *NewAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewAsset.Unmarshal(m, b)
}
func (m *NewAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewAsset.Marshal(b, m, deterministic)
}
func (m *NewAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewAsset.Merge(m, src)
}
func (m *NewAsset) XXX_Size() int {
	return xxx_messageInfo_NewAsset.Size(m)
}
func (m *NewAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_NewAsset.DiscardUnknown(m)
}

var xxx_messageInfo_NewAsset proto.InternalMessageInfo

func (m *NewAsset) GetChanges() *AssetSource {
	if m != nil {
		return m.Changes
	}
	return nil
}

type ProposalTerms struct {
	// Timestamp (Unix time in seconds) when voting closes for this proposal.
	// Constrained by `minCloseInSeconds` and `maxCloseInSeconds` network parameters.
	ClosingTimestamp int64 `protobuf:"varint,1,opt,name=closingTimestamp,proto3" json:"closingTimestamp,omitempty"`
	// Timestamp (Unix time in seconds) when proposal gets enacted (if passed).
	// Constrained by `minEnactInSeconds` and `maxEnactInSeconds` network parameters.
	EnactmentTimestamp  int64 `protobuf:"varint,2,opt,name=enactmentTimestamp,proto3" json:"enactmentTimestamp,omitempty"`
	ValidationTimestamp int64 `protobuf:"varint,5,opt,name=validationTimestamp,proto3" json:"validationTimestamp,omitempty"`
	// Actual changes being proposed
	//
	// Types that are valid to be assigned to Change:
	//	*ProposalTerms_UpdateMarket
	//	*ProposalTerms_NewMarket
	//	*ProposalTerms_UpdateNetwork
	//	*ProposalTerms_NewAsset
	Change               isProposalTerms_Change `protobuf_oneof:"change"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ProposalTerms) Reset()         { *m = ProposalTerms{} }
func (m *ProposalTerms) String() string { return proto.CompactTextString(m) }
func (*ProposalTerms) ProtoMessage()    {}
func (*ProposalTerms) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{5}
}

func (m *ProposalTerms) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalTerms.Unmarshal(m, b)
}
func (m *ProposalTerms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalTerms.Marshal(b, m, deterministic)
}
func (m *ProposalTerms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalTerms.Merge(m, src)
}
func (m *ProposalTerms) XXX_Size() int {
	return xxx_messageInfo_ProposalTerms.Size(m)
}
func (m *ProposalTerms) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalTerms.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalTerms proto.InternalMessageInfo

func (m *ProposalTerms) GetClosingTimestamp() int64 {
	if m != nil {
		return m.ClosingTimestamp
	}
	return 0
}

func (m *ProposalTerms) GetEnactmentTimestamp() int64 {
	if m != nil {
		return m.EnactmentTimestamp
	}
	return 0
}

func (m *ProposalTerms) GetValidationTimestamp() int64 {
	if m != nil {
		return m.ValidationTimestamp
	}
	return 0
}

type isProposalTerms_Change interface {
	isProposalTerms_Change()
}

type ProposalTerms_UpdateMarket struct {
	UpdateMarket *UpdateMarket `protobuf:"bytes,101,opt,name=updateMarket,proto3,oneof"`
}

type ProposalTerms_NewMarket struct {
	NewMarket *NewMarket `protobuf:"bytes,102,opt,name=newMarket,proto3,oneof"`
}

type ProposalTerms_UpdateNetwork struct {
	UpdateNetwork *UpdateNetwork `protobuf:"bytes,103,opt,name=updateNetwork,proto3,oneof"`
}

type ProposalTerms_NewAsset struct {
	NewAsset *NewAsset `protobuf:"bytes,104,opt,name=newAsset,proto3,oneof"`
}

func (*ProposalTerms_UpdateMarket) isProposalTerms_Change() {}

func (*ProposalTerms_NewMarket) isProposalTerms_Change() {}

func (*ProposalTerms_UpdateNetwork) isProposalTerms_Change() {}

func (*ProposalTerms_NewAsset) isProposalTerms_Change() {}

func (m *ProposalTerms) GetChange() isProposalTerms_Change {
	if m != nil {
		return m.Change
	}
	return nil
}

func (m *ProposalTerms) GetUpdateMarket() *UpdateMarket {
	if x, ok := m.GetChange().(*ProposalTerms_UpdateMarket); ok {
		return x.UpdateMarket
	}
	return nil
}

func (m *ProposalTerms) GetNewMarket() *NewMarket {
	if x, ok := m.GetChange().(*ProposalTerms_NewMarket); ok {
		return x.NewMarket
	}
	return nil
}

func (m *ProposalTerms) GetUpdateNetwork() *UpdateNetwork {
	if x, ok := m.GetChange().(*ProposalTerms_UpdateNetwork); ok {
		return x.UpdateNetwork
	}
	return nil
}

func (m *ProposalTerms) GetNewAsset() *NewAsset {
	if x, ok := m.GetChange().(*ProposalTerms_NewAsset); ok {
		return x.NewAsset
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ProposalTerms) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ProposalTerms_UpdateMarket)(nil),
		(*ProposalTerms_NewMarket)(nil),
		(*ProposalTerms_UpdateNetwork)(nil),
		(*ProposalTerms_NewAsset)(nil),
	}
}

type GovernanceData struct {
	// Proposal
	Proposal *Proposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
	// All "yes" votes in favour of the proposal above.
	Yes []*Vote `protobuf:"bytes,2,rep,name=yes,proto3" json:"yes,omitempty"`
	// All "no" votes against the proposal above.
	No                   []*Vote  `protobuf:"bytes,3,rep,name=no,proto3" json:"no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GovernanceData) Reset()         { *m = GovernanceData{} }
func (m *GovernanceData) String() string { return proto.CompactTextString(m) }
func (*GovernanceData) ProtoMessage()    {}
func (*GovernanceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{6}
}

func (m *GovernanceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GovernanceData.Unmarshal(m, b)
}
func (m *GovernanceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GovernanceData.Marshal(b, m, deterministic)
}
func (m *GovernanceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GovernanceData.Merge(m, src)
}
func (m *GovernanceData) XXX_Size() int {
	return xxx_messageInfo_GovernanceData.Size(m)
}
func (m *GovernanceData) XXX_DiscardUnknown() {
	xxx_messageInfo_GovernanceData.DiscardUnknown(m)
}

var xxx_messageInfo_GovernanceData proto.InternalMessageInfo

func (m *GovernanceData) GetProposal() *Proposal {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func (m *GovernanceData) GetYes() []*Vote {
	if m != nil {
		return m.Yes
	}
	return nil
}

func (m *GovernanceData) GetNo() []*Vote {
	if m != nil {
		return m.No
	}
	return nil
}

type Proposal struct {
	// Proposal unique identifier.
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Proposal reference.
	Reference string `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
	// Proposal author, identifier of the party submitting the proposal.
	PartyID string `protobuf:"bytes,3,opt,name=partyID,proto3" json:"partyID,omitempty"`
	// Proposal state (see Proposal.State definition)
	State Proposal_State `protobuf:"varint,4,opt,name=state,proto3,enum=vega.Proposal_State" json:"state,omitempty"`
	// Proposal timestamp for date and time (in nanoseconds) when proposal was submitted to the network.
	Timestamp int64 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Proposal configuration and the actual change that is meant to be executed when proposal is enacted.
	Terms                *ProposalTerms `protobuf:"bytes,6,opt,name=terms,proto3" json:"terms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{7}
}

func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Proposal) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *Proposal) GetPartyID() string {
	if m != nil {
		return m.PartyID
	}
	return ""
}

func (m *Proposal) GetState() Proposal_State {
	if m != nil {
		return m.State
	}
	return Proposal_STATE_UNSPECIFIED
}

func (m *Proposal) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Proposal) GetTerms() *ProposalTerms {
	if m != nil {
		return m.Terms
	}
	return nil
}

type Vote struct {
	// Voter's party identifier.
	PartyID string `protobuf:"bytes,1,opt,name=partyID,proto3" json:"partyID,omitempty"`
	// Actual vote.
	Value Vote_Value `protobuf:"varint,2,opt,name=value,proto3,enum=vega.Vote_Value" json:"value,omitempty"`
	// Identifier of the proposal being voted on.
	ProposalID string `protobuf:"bytes,3,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	// Vote timestamp for date and time (in nanoseconds) when vote was submitted to the network.
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{8}
}

func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetPartyID() string {
	if m != nil {
		return m.PartyID
	}
	return ""
}

func (m *Vote) GetValue() Vote_Value {
	if m != nil {
		return m.Value
	}
	return Vote_VALUE_UNSPECIFIED
}

func (m *Vote) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

func (m *Vote) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("vega.Proposal_State", Proposal_State_name, Proposal_State_value)
	proto.RegisterEnum("vega.Vote_Value", Vote_Value_name, Vote_Value_value)
	proto.RegisterType((*NetworkConfiguration)(nil), "vega.NetworkConfiguration")
	proto.RegisterType((*UpdateMarket)(nil), "vega.UpdateMarket")
	proto.RegisterType((*NewMarket)(nil), "vega.NewMarket")
	proto.RegisterType((*UpdateNetwork)(nil), "vega.UpdateNetwork")
	proto.RegisterType((*NewAsset)(nil), "vega.NewAsset")
	proto.RegisterType((*ProposalTerms)(nil), "vega.ProposalTerms")
	proto.RegisterType((*GovernanceData)(nil), "vega.GovernanceData")
	proto.RegisterType((*Proposal)(nil), "vega.Proposal")
	proto.RegisterType((*Vote)(nil), "vega.Vote")
}

func init() { proto.RegisterFile("proto/governance.proto", fileDescriptor_c891e73c7d2524a3) }

var fileDescriptor_c891e73c7d2524a3 = []byte{
	// 898 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x6e, 0xe2, 0x46,
	0x18, 0xc5, 0xe6, 0x27, 0xf0, 0x2d, 0x50, 0x67, 0xb2, 0xbb, 0xb2, 0xa2, 0x95, 0x92, 0xfa, 0xa2,
	0x8a, 0xaa, 0x5d, 0xe8, 0x12, 0x69, 0x55, 0x29, 0x95, 0x56, 0x80, 0xbd, 0x8d, 0xab, 0x84, 0xa5,
	0x86, 0x44, 0x6d, 0x6f, 0xaa, 0x59, 0x33, 0x21, 0x6e, 0xf0, 0x0c, 0xb5, 0x07, 0x68, 0xde, 0xa0,
	0x37, 0xfb, 0x86, 0x95, 0x22, 0xe5, 0x01, 0xaa, 0x7d, 0x84, 0x95, 0x67, 0x8c, 0x0d, 0x36, 0x4a,
	0x6e, 0x1c, 0x9f, 0xef, 0x7c, 0x67, 0x66, 0xbe, 0x73, 0x3c, 0xc0, 0xcb, 0x79, 0xc0, 0x38, 0x6b,
	0x4f, 0xd9, 0x92, 0x04, 0x14, 0x53, 0x97, 0xb4, 0x04, 0x80, 0x4a, 0x4b, 0x32, 0xc5, 0x87, 0xef,
	0xa6, 0x1e, 0xbf, 0x5d, 0x7c, 0x6a, 0xb9, 0xcc, 0x6f, 0xfb, 0x2b, 0x8f, 0xdf, 0xb1, 0x55, 0x7b,
	0xca, 0xde, 0x08, 0xca, 0x9b, 0x25, 0x9e, 0x79, 0x13, 0xcc, 0x59, 0x10, 0xb6, 0x93, 0x7f, 0x65,
	0xf7, 0xe1, 0x81, 0x54, 0xf5, 0x71, 0x70, 0x47, 0x78, 0x18, 0x83, 0x48, 0x82, 0x38, 0x0c, 0x13,
	0xcc, 0xf8, 0xbf, 0x08, 0xcf, 0x07, 0x84, 0xaf, 0x58, 0x70, 0xd7, 0x67, 0xf4, 0xc6, 0x9b, 0x2e,
	0x02, 0xcc, 0x3d, 0x46, 0xd1, 0x6b, 0xd8, 0xf7, 0x3d, 0xda, 0x9f, 0xb1, 0x90, 0xd8, 0x74, 0x44,
	0x5c, 0x46, 0x27, 0xa1, 0xae, 0x1c, 0x2b, 0x27, 0x45, 0x27, 0x5f, 0x10, 0x6c, 0xfc, 0x4f, 0x86,
	0xad, 0xc6, 0xec, 0x6c, 0x21, 0xd6, 0xb6, 0x28, 0x76, 0x79, 0xca, 0x2e, 0x26, 0xda, 0xdb, 0x85,
	0x58, 0x3b, 0xc3, 0x2e, 0x25, 0xda, 0x19, 0xf6, 0x05, 0xbc, 0x08, 0xc8, 0xdf, 0x0b, 0x2f, 0x20,
	0x93, 0x21, 0x0e, 0xb8, 0xe7, 0x7a, 0x73, 0x71, 0x20, 0xbd, 0x7c, 0xac, 0x9c, 0xa8, 0xbd, 0x97,
	0x8f, 0x0f, 0x47, 0xc8, 0x2e, 0xc4, 0x7f, 0xbf, 0xca, 0xc7, 0x97, 0xf7, 0xce, 0xee, 0x26, 0xd4,
	0x03, 0x6d, 0x5d, 0xb8, 0xc4, 0x7f, 0xb1, 0xc0, 0xe3, 0xf7, 0x7a, 0xe5, 0x49, 0xa1, 0x1c, 0x1f,
	0x9d, 0xc3, 0x81, 0xef, 0xd1, 0x61, 0xc0, 0xe6, 0x2c, 0xf4, 0xe8, 0xb4, 0x87, 0x67, 0x91, 0xcd,
	0xfa, 0x5e, 0x2a, 0xf3, 0x36, 0x2f, 0xb3, 0xab, 0x25, 0xda, 0x8d, 0xef, 0xd1, 0x6b, 0xc6, 0x37,
	0x64, 0xaa, 0x4f, 0xca, 0xe4, 0xf8, 0x46, 0x13, 0xea, 0x57, 0xf3, 0x09, 0xe6, 0xe4, 0x52, 0x64,
	0xc3, 0x38, 0x83, 0xda, 0x80, 0xac, 0xe4, 0x0b, 0x6a, 0xc1, 0x9e, 0x7b, 0x8b, 0xe9, 0x94, 0x48,
	0xab, 0x9f, 0x75, 0xea, 0xad, 0x28, 0x86, 0x2d, 0x59, 0xee, 0x55, 0x1e, 0x1f, 0x8e, 0xd4, 0x63,
	0xc5, 0x59, 0x93, 0x8c, 0x4b, 0x68, 0x48, 0xb1, 0x38, 0x42, 0xe8, 0xa7, 0xac, 0xc0, 0xa1, 0x14,
	0xd8, 0x15, 0xb1, 0xbc, 0xdc, 0x7b, 0xa8, 0x0e, 0xc8, 0xaa, 0x1b, 0xe5, 0x13, 0x9d, 0x66, 0x95,
	0xf6, 0xa5, 0x92, 0xa8, 0x8e, 0xd8, 0x22, 0x70, 0x49, 0x5e, 0xe0, 0x73, 0x11, 0x1a, 0x72, 0x6a,
	0x78, 0x36, 0x26, 0x81, 0x1f, 0xa2, 0x0e, 0x68, 0xee, 0x4c, 0x0c, 0x71, 0xec, 0xf9, 0x24, 0xe4,
	0xd8, 0x9f, 0xcb, 0x14, 0xcb, 0x66, 0xad, 0xe0, 0xe4, 0xea, 0xe8, 0x1d, 0x20, 0x12, 0x85, 0xca,
	0x27, 0x94, 0xa7, 0x5d, 0xea, 0x56, 0xd7, 0x0e, 0x06, 0xfa, 0x01, 0x0e, 0xe2, 0xef, 0xd0, 0x63,
	0x34, 0x6d, 0x2c, 0x8b, 0xa8, 0xee, 0x2a, 0xa1, 0x1f, 0xa1, 0xbe, 0xd8, 0x30, 0x43, 0x27, 0xe2,
	0xa4, 0x48, 0x9e, 0x74, 0xd3, 0xa6, 0xf3, 0x82, 0xb3, 0xc5, 0x44, 0x6d, 0xa8, 0xd1, 0xb5, 0x6d,
	0xfa, 0x8d, 0x68, 0xfb, 0x66, 0x3d, 0xea, 0x55, 0xd2, 0x93, 0x72, 0xd0, 0x19, 0x34, 0x16, 0x9b,
	0x56, 0xe9, 0x53, 0xd1, 0x74, 0xb0, 0xb9, 0x56, 0x5c, 0x3a, 0x2f, 0x38, 0xdb, 0x5c, 0xf4, 0x1a,
	0xaa, 0x34, 0x36, 0x46, 0xbf, 0x15, 0x7d, 0xcd, 0x64, 0x31, 0x81, 0x9e, 0x17, 0x9c, 0x84, 0xd1,
	0xab, 0x42, 0x45, 0x1a, 0x62, 0x2c, 0xa1, 0xf9, 0x73, 0x72, 0xb1, 0x99, 0x98, 0x63, 0xf4, 0x3d,
	0x54, 0xe7, 0xb1, 0x41, 0xb1, 0xaf, 0xb1, 0xd2, 0xda, 0x36, 0x27, 0xa9, 0xa3, 0x57, 0x50, 0xbc,
	0x27, 0xd1, 0x35, 0x52, 0x3c, 0x79, 0xd6, 0x01, 0x49, 0xbb, 0x66, 0x9c, 0x38, 0x11, 0x8c, 0x0e,
	0x41, 0xa5, 0x4c, 0x2f, 0xe6, 0x8a, 0x2a, 0x65, 0xc6, 0x17, 0x15, 0xaa, 0x6b, 0x41, 0xd4, 0x04,
	0xd5, 0x36, 0xc5, 0x62, 0x35, 0x47, 0xb5, 0x4d, 0xf4, 0x0a, 0x6a, 0x01, 0xb9, 0x21, 0x01, 0x89,
	0x3e, 0x1f, 0x55, 0xc0, 0x29, 0x80, 0x8e, 0x61, 0x6f, 0x8e, 0x03, 0x7e, 0x6f, 0x9b, 0xe2, 0x46,
	0xaa, 0x49, 0xc7, 0x7f, 0x53, 0x9c, 0x35, 0x8c, 0x4e, 0xa1, 0x1c, 0x72, 0xcc, 0x89, 0xb8, 0x83,
	0x9a, 0x9d, 0xe7, 0xdb, 0xfb, 0x6f, 0x8d, 0xa2, 0x5a, 0x6f, 0xef, 0xf1, 0xe1, 0xa8, 0xf8, 0xaf,
	0xa2, 0x38, 0x92, 0x1b, 0x2d, 0xca, 0x33, 0x89, 0x48, 0x01, 0xf4, 0x16, 0xca, 0x3c, 0x8a, 0xab,
	0xb8, 0x5b, 0x12, 0x53, 0xb6, 0x92, 0x9c, 0x84, 0x5d, 0x32, 0x8d, 0xcf, 0x0a, 0x94, 0xc5, 0x52,
	0xe8, 0x05, 0xec, 0x8f, 0xc6, 0xdd, 0xb1, 0xf5, 0xe7, 0xd5, 0x60, 0x34, 0xb4, 0xfa, 0xf6, 0x07,
	0xdb, 0x32, 0xb5, 0x02, 0xd2, 0xa0, 0x2e, 0xe1, 0x0f, 0x5d, 0xfb, 0xc2, 0x32, 0x35, 0x05, 0x35,
	0x01, 0x24, 0xf2, 0x71, 0x68, 0x0d, 0x34, 0x35, 0x65, 0x0c, 0xbb, 0xa3, 0x91, 0x65, 0x6a, 0x45,
	0x84, 0xa0, 0x29, 0x11, 0xc7, 0xfa, 0xc5, 0xea, 0x8f, 0x2d, 0x53, 0x2b, 0xa5, 0x98, 0x69, 0xf5,
	0x2f, 0xec, 0x81, 0x65, 0x6a, 0x65, 0xb4, 0x0f, 0x0d, 0x89, 0x59, 0x83, 0xae, 0xa0, 0x55, 0x8c,
	0xff, 0x14, 0x28, 0x45, 0xf3, 0xdf, 0x1c, 0xa0, 0xb2, 0x7b, 0x80, 0x6d, 0x28, 0x2f, 0xf1, 0x6c,
	0x21, 0x87, 0xdf, 0xec, 0x68, 0xa9, 0x79, 0xad, 0xeb, 0x08, 0xdf, 0x18, 0x9e, 0xe0, 0xa1, 0xef,
	0x00, 0xd6, 0xa1, 0xc8, 0xd9, 0xb2, 0x51, 0xd9, 0x1e, 0x72, 0x29, 0x33, 0x64, 0xe3, 0x0c, 0xca,
	0x42, 0x3e, 0x1a, 0xd8, 0x75, 0xf7, 0xe2, 0x2a, 0x3b, 0xb0, 0x3a, 0x54, 0x25, 0x3c, 0xf8, 0xa8,
	0x29, 0xa8, 0x01, 0x35, 0xf9, 0xf6, 0xbb, 0x35, 0xd2, 0xd4, 0xde, 0xb7, 0x7f, 0x1c, 0xb9, 0x6c,
	0x42, 0xc4, 0x56, 0xc5, 0x2f, 0xa7, 0xcb, 0x66, 0x2d, 0x8f, 0xb5, 0xa3, 0xf7, 0xb6, 0x00, 0x3e,
	0x55, 0xc4, 0xe3, 0xf4, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xb1, 0x37, 0x8d, 0xd2, 0x07,
	0x00, 0x00,
}
