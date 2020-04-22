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
	// Proposal could not be enacted after being accepted by the network
	Proposal_FAILED Proposal_State = 0
	// Proposal is open for voting.
	Proposal_OPEN Proposal_State = 1
	// Proposal has gained enough support to be executed.
	Proposal_PASSED Proposal_State = 2
	// Proposal wasn't accepted (validation failed, author not allowed to submit proposals)
	Proposal_REJECTED Proposal_State = 3
	// Proposal didn't get enough votes
	Proposal_DECLINED Proposal_State = 4
	// Proposal has been executed and the changes under this proposal have now been applied.
	Proposal_ENACTED Proposal_State = 5
)

var Proposal_State_name = map[int32]string{
	0: "FAILED",
	1: "OPEN",
	2: "PASSED",
	3: "REJECTED",
	4: "DECLINED",
	5: "ENACTED",
}

var Proposal_State_value = map[string]int32{
	"FAILED":   0,
	"OPEN":     1,
	"PASSED":   2,
	"REJECTED": 3,
	"DECLINED": 4,
	"ENACTED":  5,
}

func (x Proposal_State) String() string {
	return proto.EnumName(Proposal_State_name, int32(x))
}

func (Proposal_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{7, 0}
}

type Vote_Value int32

const (
	Vote_NO  Vote_Value = 0
	Vote_YES Vote_Value = 1
)

var Vote_Value_name = map[int32]string{
	0: "NO",
	1: "YES",
}

var Vote_Value_value = map[string]int32{
	"NO":  0,
	"YES": 1,
}

func (x Vote_Value) String() string {
	return proto.EnumName(Vote_Value_name, int32(x))
}

func (Vote_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c891e73c7d2524a3, []int{8, 0}
}

type NetworkConfiguration struct {
	MinCloseInSeconds     int64    `protobuf:"varint,1,opt,name=minCloseInSeconds,proto3" json:"minCloseInSeconds,omitempty"`
	MaxCloseInSeconds     int64    `protobuf:"varint,2,opt,name=maxCloseInSeconds,proto3" json:"maxCloseInSeconds,omitempty"`
	MinEnactInSeconds     int64    `protobuf:"varint,3,opt,name=minEnactInSeconds,proto3" json:"minEnactInSeconds,omitempty"`
	MaxEnactInSeconds     int64    `protobuf:"varint,4,opt,name=maxEnactInSeconds,proto3" json:"maxEnactInSeconds,omitempty"`
	MinParticipationStake uint64   `protobuf:"varint,5,opt,name=minParticipationStake,proto3" json:"minParticipationStake,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
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

func (m *NetworkConfiguration) GetMinParticipationStake() uint64 {
	if m != nil {
		return m.MinParticipationStake
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
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *NewAsset) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type ProposalTerms struct {
	ClosingTimestamp      int64  `protobuf:"varint,1,opt,name=closingTimestamp,proto3" json:"closingTimestamp,omitempty"`
	EnactmentTimestamp    int64  `protobuf:"varint,2,opt,name=enactmentTimestamp,proto3" json:"enactmentTimestamp,omitempty"`
	MinParticipationStake uint64 `protobuf:"varint,3,opt,name=minParticipationStake,proto3" json:"minParticipationStake,omitempty"`
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

func (m *ProposalTerms) GetMinParticipationStake() uint64 {
	if m != nil {
		return m.MinParticipationStake
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
	Proposal             *Proposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
	Yes                  []*Vote   `protobuf:"bytes,2,rep,name=yes,proto3" json:"yes,omitempty"`
	No                   []*Vote   `protobuf:"bytes,3,rep,name=no,proto3" json:"no,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
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
	ID                   string         `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Reference            string         `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
	PartyID              string         `protobuf:"bytes,3,opt,name=partyID,proto3" json:"partyID,omitempty"`
	State                Proposal_State `protobuf:"varint,4,opt,name=state,proto3,enum=vega.Proposal_State" json:"state,omitempty"`
	Timestamp            int64          `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
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
	return Proposal_FAILED
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
	PartyID              string     `protobuf:"bytes,1,opt,name=partyID,proto3" json:"partyID,omitempty"`
	Value                Vote_Value `protobuf:"varint,2,opt,name=value,proto3,enum=vega.Vote_Value" json:"value,omitempty"`
	ProposalID           string     `protobuf:"bytes,3,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
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
	return Vote_NO
}

func (m *Vote) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
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
	// 765 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x6e, 0xe2, 0x46,
	0x18, 0xc5, 0x36, 0xe6, 0xe7, 0x0b, 0xa1, 0xee, 0xec, 0xb6, 0xb2, 0xd0, 0x4a, 0xa1, 0xbe, 0xa8,
	0xa2, 0x6a, 0xd7, 0x56, 0xd9, 0x6a, 0x55, 0x69, 0xf7, 0x06, 0xb0, 0xdb, 0x50, 0xed, 0x12, 0x64,
	0xd2, 0xa8, 0xed, 0xdd, 0xc4, 0x4c, 0x1c, 0x0b, 0x3c, 0x83, 0xec, 0x01, 0x9a, 0xfb, 0x5e, 0xf4,
	0x15, 0xaa, 0xbe, 0x46, 0xdf, 0x27, 0x52, 0x9e, 0xa4, 0xf2, 0x8c, 0x31, 0xe6, 0x67, 0xaf, 0x92,
	0x39, 0xe7, 0x7c, 0x67, 0xe0, 0x9c, 0xf9, 0x80, 0xaf, 0x97, 0x09, 0xe3, 0xcc, 0x09, 0xd9, 0x9a,
	0x24, 0x14, 0xd3, 0x80, 0xd8, 0x02, 0x40, 0xd5, 0x35, 0x09, 0x71, 0xe7, 0x5d, 0x18, 0xf1, 0x87,
	0xd5, 0x9d, 0x1d, 0xb0, 0xd8, 0x89, 0x37, 0x11, 0x9f, 0xb3, 0x8d, 0x13, 0xb2, 0x37, 0x42, 0xf2,
	0x66, 0x8d, 0x17, 0xd1, 0x0c, 0x73, 0x96, 0xa4, 0x4e, 0xf1, 0xaf, 0x9c, 0xee, 0xbc, 0x90, 0xae,
	0x31, 0x4e, 0xe6, 0x84, 0xa7, 0x12, 0xb4, 0xfe, 0x52, 0xe1, 0xe5, 0x98, 0xf0, 0x0d, 0x4b, 0xe6,
	0x43, 0x46, 0xef, 0xa3, 0x70, 0x95, 0x60, 0x1e, 0x31, 0x8a, 0x5e, 0xc3, 0x97, 0x71, 0x44, 0x87,
	0x0b, 0x96, 0x92, 0x11, 0x9d, 0x92, 0x80, 0xd1, 0x59, 0x6a, 0x2a, 0x5d, 0xe5, 0x52, 0xf3, 0x8f,
	0x09, 0xa1, 0xc6, 0x7f, 0x1e, 0xa8, 0xd5, 0x5c, 0x7d, 0x48, 0xe4, 0xde, 0x1e, 0xc5, 0x01, 0xdf,
	0xa9, 0xb5, 0xc2, 0x7b, 0x9f, 0xc8, 0xbd, 0x0f, 0xd4, 0xd5, 0xc2, 0xfb, 0x40, 0xfd, 0x03, 0x7c,
	0x15, 0x47, 0x74, 0x82, 0x13, 0x1e, 0x05, 0xd1, 0x52, 0x7c, 0x97, 0x29, 0xc7, 0x73, 0x62, 0xea,
	0x5d, 0xe5, 0xb2, 0xea, 0x9f, 0x26, 0xad, 0x36, 0xb4, 0x7e, 0x5d, 0xce, 0x30, 0x27, 0x9f, 0x44,
	0x3a, 0xd6, 0x7b, 0x68, 0x8e, 0xc9, 0x46, 0x1e, 0x90, 0x0d, 0xf5, 0xe0, 0x01, 0xd3, 0x90, 0xc8,
	0x00, 0xce, 0x7a, 0x2d, 0x3b, 0x2b, 0xc2, 0x96, 0xf4, 0xa0, 0xf6, 0xfc, 0x74, 0xa1, 0x76, 0x15,
	0x7f, 0x2b, 0xb2, 0x3e, 0xc1, 0xb9, 0x34, 0xcb, 0x83, 0x45, 0x1f, 0x0e, 0x0d, 0x3a, 0xd2, 0xe0,
	0x54, 0xf0, 0xc7, 0x76, 0x1d, 0x68, 0x8c, 0xc9, 0xa6, 0x9f, 0xa6, 0x84, 0xa3, 0x36, 0xa8, 0x23,
	0x57, 0x98, 0x34, 0x7d, 0x75, 0xe4, 0x5a, 0xff, 0x68, 0x70, 0x3e, 0x49, 0xd8, 0x92, 0xa5, 0x78,
	0x71, 0x43, 0x92, 0x38, 0x45, 0x3d, 0x30, 0x82, 0x05, 0x4b, 0x23, 0x1a, 0xde, 0x44, 0x31, 0x49,
	0x39, 0x8e, 0x97, 0xb2, 0x36, 0x69, 0x6c, 0x54, 0xfc, 0x23, 0x1e, 0xbd, 0x03, 0x44, 0xb2, 0x14,
	0x63, 0x42, 0xf9, 0x6e, 0x4a, 0xdd, 0x9b, 0x3a, 0xa1, 0x40, 0x1f, 0x3e, 0x97, 0x75, 0xd6, 0x65,
	0xb5, 0x18, 0x3d, 0x2d, 0x42, 0x3f, 0x42, 0x6b, 0x55, 0xca, 0xdc, 0x24, 0x22, 0x1a, 0x24, 0xa3,
	0x29, 0xb7, 0x71, 0x55, 0xf1, 0xf7, 0x94, 0xc8, 0x81, 0x26, 0xdd, 0xb6, 0x63, 0xde, 0x8b, 0xb1,
	0x2f, 0xb6, 0x89, 0x6e, 0x8a, 0x99, 0x9d, 0x06, 0xbd, 0x87, 0xf3, 0x55, 0xb9, 0x11, 0x33, 0x14,
	0x43, 0x2f, 0xca, 0x77, 0xe5, 0xd4, 0x55, 0xc5, 0xdf, 0xd7, 0xa2, 0xd7, 0xd0, 0xa0, 0x79, 0xfe,
	0xe6, 0x83, 0x98, 0x6b, 0x17, 0x97, 0x09, 0xf4, 0xaa, 0xe2, 0x17, 0x8a, 0x41, 0x03, 0x6a, 0xb2,
	0x38, 0x6b, 0x0d, 0xed, 0x9f, 0x8b, 0x0d, 0x76, 0x31, 0xc7, 0xe8, 0x3b, 0x68, 0x2c, 0xf3, 0xb2,
	0xf2, 0x87, 0x90, 0x3b, 0x6d, 0x2b, 0xf4, 0x0b, 0x1e, 0xbd, 0x02, 0xed, 0x91, 0x64, 0x3b, 0xa4,
	0x5d, 0x9e, 0xf5, 0x40, 0xca, 0x6e, 0x19, 0x27, 0x7e, 0x06, 0xa3, 0x0e, 0xa8, 0x94, 0x99, 0xda,
	0x11, 0xa9, 0x52, 0x66, 0xfd, 0xa7, 0x42, 0x63, 0x6b, 0x78, 0xf8, 0x60, 0xd0, 0x2b, 0x68, 0x26,
	0xe4, 0x9e, 0x24, 0x84, 0x06, 0x44, 0x34, 0xdc, 0xf4, 0x77, 0x00, 0xea, 0x42, 0x7d, 0x89, 0x13,
	0xfe, 0x38, 0x72, 0x45, 0x85, 0x4d, 0x59, 0xe1, 0x6f, 0x8a, 0xbf, 0x85, 0xd1, 0x5b, 0xd0, 0x53,
	0x8e, 0x39, 0x11, 0x0b, 0xd8, 0xee, 0xbd, 0xdc, 0xff, 0xfc, 0xf6, 0x34, 0xe3, 0x06, 0xf5, 0xe7,
	0xa7, 0x0b, 0xed, 0x6f, 0x45, 0xf1, 0xa5, 0x36, 0xbb, 0x94, 0x17, 0xcf, 0x4a, 0x17, 0x9b, 0xbb,
	0x03, 0xd0, 0xf7, 0xa0, 0xf3, 0xec, 0xe9, 0x9a, 0xb5, 0x72, 0x29, 0x7b, 0xaf, 0xba, 0x58, 0x0a,
	0xa9, 0xb4, 0x7c, 0xd0, 0xc5, 0x4d, 0x08, 0xa0, 0xf6, 0x53, 0x7f, 0xf4, 0xd1, 0x73, 0x8d, 0x0a,
	0x6a, 0x40, 0xf5, 0x7a, 0xe2, 0x8d, 0x0d, 0x25, 0x43, 0x27, 0xfd, 0xe9, 0xd4, 0x73, 0x0d, 0x15,
	0xb5, 0xa0, 0xe1, 0x7b, 0xbf, 0x78, 0xc3, 0x1b, 0xcf, 0x35, 0xb4, 0xec, 0xe4, 0x7a, 0xc3, 0x8f,
	0xa3, 0xb1, 0xe7, 0x1a, 0x55, 0x74, 0x06, 0x75, 0x6f, 0xdc, 0x17, 0x94, 0x6e, 0xfd, 0xab, 0x40,
	0x35, 0xcb, 0xb0, 0x1c, 0x82, 0x72, 0x3a, 0x04, 0x07, 0xf4, 0x35, 0x5e, 0xac, 0x64, 0x80, 0xed,
	0x9e, 0xb1, 0x2b, 0xc0, 0xbe, 0xcd, 0xf0, 0x52, 0x00, 0x42, 0x87, 0xbe, 0x05, 0xd8, 0x16, 0x7b,
	0x14, 0x6d, 0x89, 0xb1, 0x4c, 0xd0, 0x85, 0x01, 0xaa, 0x81, 0x3a, 0xbe, 0x36, 0x2a, 0xa8, 0x0e,
	0xda, 0xef, 0xde, 0xd4, 0x50, 0x06, 0xdf, 0xfc, 0x71, 0x11, 0xb0, 0x19, 0x11, 0x37, 0x89, 0x5f,
	0xee, 0x80, 0x2d, 0xec, 0x88, 0x39, 0xd9, 0xd9, 0x11, 0xc0, 0x5d, 0x4d, 0xfc, 0x79, 0xfb, 0x7f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x20, 0x2a, 0x37, 0x3e, 0x06, 0x00, 0x00,
}
