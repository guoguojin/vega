// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet/v1/wallet.proto

package v1

import (
	v1 "code.vegaprotocol.io/vega/proto/commands/v1"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SubmitTransactionRequest struct {
	PubKey    string `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Propagate bool   `protobuf:"varint,2,opt,name=propagate,proto3" json:"propagate,omitempty"`
	// Types that are valid to be assigned to Command:
	//	*SubmitTransactionRequest_OrderSubmission
	//	*SubmitTransactionRequest_OrderCancellation
	//	*SubmitTransactionRequest_OrderAmendment
	//	*SubmitTransactionRequest_WithdrawSubmission
	//	*SubmitTransactionRequest_ProposalSubmission
	//	*SubmitTransactionRequest_VoteSubmission
	//	*SubmitTransactionRequest_LiquidityProvisionSubmission
	//	*SubmitTransactionRequest_NodeRegistration
	//	*SubmitTransactionRequest_NodeVote
	//	*SubmitTransactionRequest_NodeSignature
	//	*SubmitTransactionRequest_ChainEvent
	//	*SubmitTransactionRequest_OracleDataSubmission
	//	*SubmitTransactionRequest_DelegateSubmission
	//	*SubmitTransactionRequest_UndelegateAtEpochEndSubmission
	Command              isSubmitTransactionRequest_Command `protobuf_oneof:"command"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *SubmitTransactionRequest) Reset()         { *m = SubmitTransactionRequest{} }
func (m *SubmitTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionRequest) ProtoMessage()    {}
func (*SubmitTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a45fb539ff6b1e, []int{0}
}

func (m *SubmitTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionRequest.Unmarshal(m, b)
}
func (m *SubmitTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionRequest.Marshal(b, m, deterministic)
}
func (m *SubmitTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionRequest.Merge(m, src)
}
func (m *SubmitTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionRequest.Size(m)
}
func (m *SubmitTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionRequest proto.InternalMessageInfo

func (m *SubmitTransactionRequest) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *SubmitTransactionRequest) GetPropagate() bool {
	if m != nil {
		return m.Propagate
	}
	return false
}

type isSubmitTransactionRequest_Command interface {
	isSubmitTransactionRequest_Command()
}

type SubmitTransactionRequest_OrderSubmission struct {
	OrderSubmission *v1.OrderSubmission `protobuf:"bytes,1001,opt,name=order_submission,json=orderSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_OrderCancellation struct {
	OrderCancellation *v1.OrderCancellation `protobuf:"bytes,1002,opt,name=order_cancellation,json=orderCancellation,proto3,oneof"`
}

type SubmitTransactionRequest_OrderAmendment struct {
	OrderAmendment *v1.OrderAmendment `protobuf:"bytes,1003,opt,name=order_amendment,json=orderAmendment,proto3,oneof"`
}

type SubmitTransactionRequest_WithdrawSubmission struct {
	WithdrawSubmission *v1.WithdrawSubmission `protobuf:"bytes,1004,opt,name=withdraw_submission,json=withdrawSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_ProposalSubmission struct {
	ProposalSubmission *v1.ProposalSubmission `protobuf:"bytes,1005,opt,name=proposal_submission,json=proposalSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_VoteSubmission struct {
	VoteSubmission *v1.VoteSubmission `protobuf:"bytes,1006,opt,name=vote_submission,json=voteSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_LiquidityProvisionSubmission struct {
	LiquidityProvisionSubmission *v1.LiquidityProvisionSubmission `protobuf:"bytes,1007,opt,name=liquidity_provision_submission,json=liquidityProvisionSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_NodeRegistration struct {
	NodeRegistration *v1.NodeRegistration `protobuf:"bytes,2001,opt,name=node_registration,json=nodeRegistration,proto3,oneof"`
}

type SubmitTransactionRequest_NodeVote struct {
	NodeVote *v1.NodeVote `protobuf:"bytes,2002,opt,name=node_vote,json=nodeVote,proto3,oneof"`
}

type SubmitTransactionRequest_NodeSignature struct {
	NodeSignature *v1.NodeSignature `protobuf:"bytes,2003,opt,name=node_signature,json=nodeSignature,proto3,oneof"`
}

type SubmitTransactionRequest_ChainEvent struct {
	ChainEvent *v1.ChainEvent `protobuf:"bytes,2004,opt,name=chain_event,json=chainEvent,proto3,oneof"`
}

type SubmitTransactionRequest_OracleDataSubmission struct {
	OracleDataSubmission *v1.OracleDataSubmission `protobuf:"bytes,3001,opt,name=oracle_data_submission,json=oracleDataSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_DelegateSubmission struct {
	DelegateSubmission *v1.DelegateSubmission `protobuf:"bytes,4001,opt,name=delegate_submission,json=delegateSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_UndelegateAtEpochEndSubmission struct {
	UndelegateAtEpochEndSubmission *v1.UndelegateAtEpochEndSubmission `protobuf:"bytes,4002,opt,name=undelegate_at_epoch_end_submission,json=undelegateAtEpochEndSubmission,proto3,oneof"`
}

func (*SubmitTransactionRequest_OrderSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_OrderCancellation) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_OrderAmendment) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_WithdrawSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_ProposalSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_VoteSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_LiquidityProvisionSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_NodeRegistration) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_NodeVote) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_NodeSignature) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_ChainEvent) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_OracleDataSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_DelegateSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_UndelegateAtEpochEndSubmission) isSubmitTransactionRequest_Command() {
}

func (m *SubmitTransactionRequest) GetCommand() isSubmitTransactionRequest_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *SubmitTransactionRequest) GetOrderSubmission() *v1.OrderSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_OrderSubmission); ok {
		return x.OrderSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetOrderCancellation() *v1.OrderCancellation {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_OrderCancellation); ok {
		return x.OrderCancellation
	}
	return nil
}

func (m *SubmitTransactionRequest) GetOrderAmendment() *v1.OrderAmendment {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_OrderAmendment); ok {
		return x.OrderAmendment
	}
	return nil
}

func (m *SubmitTransactionRequest) GetWithdrawSubmission() *v1.WithdrawSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_WithdrawSubmission); ok {
		return x.WithdrawSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetProposalSubmission() *v1.ProposalSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_ProposalSubmission); ok {
		return x.ProposalSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetVoteSubmission() *v1.VoteSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_VoteSubmission); ok {
		return x.VoteSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetLiquidityProvisionSubmission() *v1.LiquidityProvisionSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_LiquidityProvisionSubmission); ok {
		return x.LiquidityProvisionSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetNodeRegistration() *v1.NodeRegistration {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_NodeRegistration); ok {
		return x.NodeRegistration
	}
	return nil
}

func (m *SubmitTransactionRequest) GetNodeVote() *v1.NodeVote {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_NodeVote); ok {
		return x.NodeVote
	}
	return nil
}

func (m *SubmitTransactionRequest) GetNodeSignature() *v1.NodeSignature {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_NodeSignature); ok {
		return x.NodeSignature
	}
	return nil
}

func (m *SubmitTransactionRequest) GetChainEvent() *v1.ChainEvent {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_ChainEvent); ok {
		return x.ChainEvent
	}
	return nil
}

func (m *SubmitTransactionRequest) GetOracleDataSubmission() *v1.OracleDataSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_OracleDataSubmission); ok {
		return x.OracleDataSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetDelegateSubmission() *v1.DelegateSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_DelegateSubmission); ok {
		return x.DelegateSubmission
	}
	return nil
}

func (m *SubmitTransactionRequest) GetUndelegateAtEpochEndSubmission() *v1.UndelegateAtEpochEndSubmission {
	if x, ok := m.GetCommand().(*SubmitTransactionRequest_UndelegateAtEpochEndSubmission); ok {
		return x.UndelegateAtEpochEndSubmission
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubmitTransactionRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubmitTransactionRequest_OrderSubmission)(nil),
		(*SubmitTransactionRequest_OrderCancellation)(nil),
		(*SubmitTransactionRequest_OrderAmendment)(nil),
		(*SubmitTransactionRequest_WithdrawSubmission)(nil),
		(*SubmitTransactionRequest_ProposalSubmission)(nil),
		(*SubmitTransactionRequest_VoteSubmission)(nil),
		(*SubmitTransactionRequest_LiquidityProvisionSubmission)(nil),
		(*SubmitTransactionRequest_NodeRegistration)(nil),
		(*SubmitTransactionRequest_NodeVote)(nil),
		(*SubmitTransactionRequest_NodeSignature)(nil),
		(*SubmitTransactionRequest_ChainEvent)(nil),
		(*SubmitTransactionRequest_OracleDataSubmission)(nil),
		(*SubmitTransactionRequest_DelegateSubmission)(nil),
		(*SubmitTransactionRequest_UndelegateAtEpochEndSubmission)(nil),
	}
}

func init() {
	proto.RegisterType((*SubmitTransactionRequest)(nil), "vega.wallet.v1.SubmitTransactionRequest")
}

func init() { proto.RegisterFile("wallet/v1/wallet.proto", fileDescriptor_73a45fb539ff6b1e) }

var fileDescriptor_73a45fb539ff6b1e = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x5b, 0x2e, 0xb6, 0xd5, 0x13, 0xdd, 0x66, 0xd0, 0x16, 0xaa, 0x69, 0x2b, 0x63, 0x42,
	0x43, 0x48, 0x29, 0x83, 0x3b, 0xae, 0xd8, 0x97, 0x14, 0xc4, 0xc4, 0xa6, 0x8c, 0x01, 0xe2, 0x26,
	0x72, 0xe3, 0xa3, 0xce, 0x22, 0xb5, 0x33, 0xc7, 0x49, 0xb5, 0x2b, 0x9e, 0x05, 0x2e, 0x78, 0x07,
	0x1e, 0x83, 0x8f, 0x97, 0xe0, 0xfb, 0x15, 0x90, 0xe3, 0x34, 0x38, 0x4b, 0xbb, 0x3b, 0xe7, 0xef,
	0xff, 0xf9, 0xfd, 0xe3, 0x93, 0x13, 0xa3, 0xe5, 0x11, 0x89, 0x22, 0x50, 0xbd, 0x6c, 0xbb, 0x67,
	0x56, 0x6e, 0x2c, 0x85, 0x12, 0xb8, 0x9d, 0xc1, 0x80, 0xb8, 0x85, 0x94, 0x6d, 0x77, 0x3a, 0xa1,
	0x18, 0x0e, 0x09, 0xa7, 0x89, 0x76, 0x8e, 0xd7, 0xc6, 0xdb, 0xd9, 0xb4, 0xf7, 0x32, 0x12, 0x31,
	0x4a, 0x94, 0x90, 0xc1, 0x25, 0xd7, 0x2d, 0xdb, 0x25, 0x24, 0x09, 0x23, 0x28, 0xb6, 0x36, 0x3e,
	0x22, 0xe4, 0x9c, 0xa4, 0xfd, 0x21, 0x53, 0x2f, 0x24, 0xe1, 0x09, 0x09, 0x15, 0x13, 0xdc, 0x87,
	0xf3, 0x14, 0x12, 0x85, 0x57, 0xd0, 0x6c, 0x9c, 0xf6, 0x83, 0xb7, 0x70, 0xe1, 0x34, 0xbb, 0xcd,
	0xad, 0x96, 0x3f, 0x13, 0xa7, 0xfd, 0x67, 0x70, 0x81, 0x57, 0x51, 0x2b, 0x96, 0x22, 0x26, 0x03,
	0xa2, 0xc0, 0xb9, 0xd6, 0x6d, 0x6e, 0xcd, 0xf9, 0xff, 0x05, 0x7c, 0x84, 0x16, 0x85, 0xa4, 0x20,
	0x83, 0x44, 0x83, 0x93, 0x84, 0x09, 0xee, 0x7c, 0x9f, 0xed, 0x36, 0xb7, 0xe6, 0x1f, 0xde, 0x76,
	0xf3, 0xc3, 0x95, 0xef, 0x97, 0x6d, 0xbb, 0x47, 0xda, 0x7a, 0x52, 0x3a, 0xbd, 0x86, 0xbf, 0x20,
	0xaa, 0x12, 0x3e, 0x45, 0xd8, 0x00, 0x43, 0xc2, 0x43, 0x88, 0x22, 0xa2, 0x5f, 0xd2, 0xf9, 0x61,
	0x90, 0x77, 0xa6, 0x20, 0xf7, 0x2c, 0xaf, 0xd7, 0xf0, 0x97, 0xc4, 0x65, 0x11, 0x1f, 0x22, 0x93,
	0x14, 0x90, 0x21, 0x70, 0x3a, 0x04, 0xae, 0x9c, 0x9f, 0x86, 0xd9, 0x9d, 0xc2, 0xdc, 0x19, 0x1b,
	0xbd, 0x86, 0xdf, 0x16, 0x15, 0x05, 0xbf, 0x46, 0x37, 0x46, 0x4c, 0x9d, 0x51, 0x49, 0x46, 0xf6,
	0xc1, 0x7f, 0x19, 0xe2, 0x66, 0x9d, 0xf8, 0xaa, 0x70, 0x57, 0xce, 0x8e, 0x47, 0x35, 0x55, 0x93,
	0x75, 0x73, 0x45, 0x42, 0x22, 0x9b, 0xfc, 0x7b, 0x2a, 0xf9, 0xb8, 0x70, 0x57, 0xc9, 0x71, 0x4d,
	0xd5, 0x1d, 0xc8, 0x84, 0x02, 0x9b, 0xfa, 0x67, 0x6a, 0x07, 0x5e, 0x0a, 0x05, 0x15, 0x62, 0x3b,
	0xab, 0x28, 0x78, 0x84, 0xd6, 0x22, 0x76, 0x9e, 0x32, 0xca, 0xd4, 0x45, 0x10, 0x4b, 0x91, 0x31,
	0x2d, 0xdb, 0xf0, 0xbf, 0x06, 0xee, 0xd6, 0xe1, 0x87, 0xe3, 0xc2, 0xe3, 0x71, 0x5d, 0x25, 0x6a,
	0x35, 0xba, 0x62, 0x1f, 0xfb, 0x68, 0x89, 0x0b, 0x0a, 0x81, 0x84, 0x01, 0x4b, 0x94, 0x34, 0xe3,
	0xf1, 0x79, 0x21, 0xcf, 0xda, 0xa8, 0x67, 0x3d, 0x17, 0x14, 0x7c, 0xcb, 0xea, 0x35, 0xfc, 0x45,
	0x7e, 0x49, 0xc3, 0x8f, 0x51, 0x2b, 0x67, 0xea, 0x33, 0x3a, 0x5f, 0x0c, 0xab, 0x33, 0x99, 0xa5,
	0x1b, 0xe3, 0x35, 0xfc, 0x39, 0x5e, 0xac, 0xf1, 0x53, 0xd4, 0xce, 0x6b, 0x13, 0x36, 0xe0, 0x44,
	0xa5, 0x12, 0x9c, 0xaf, 0x06, 0xb0, 0x3e, 0x19, 0x70, 0x32, 0xf6, 0x79, 0x0d, 0xff, 0x3a, 0xb7,
	0x05, 0xfc, 0x04, 0xcd, 0x87, 0x67, 0x84, 0xf1, 0x00, 0x32, 0x3d, 0x9f, 0xdf, 0x0c, 0x67, 0xb5,
	0xce, 0xd9, 0xd3, 0xae, 0x83, 0xcc, 0xcc, 0x26, 0x0a, 0xcb, 0x27, 0x1c, 0xa0, 0x65, 0xf3, 0xcb,
	0x07, 0x94, 0x28, 0x62, 0x7f, 0x8d, 0x4f, 0x2b, 0x39, 0xec, 0xee, 0xa4, 0x61, 0xd7, 0x05, 0xfb,
	0x44, 0x91, 0xca, 0x57, 0xb8, 0x29, 0x26, 0xe8, 0x7a, 0x3c, 0x29, 0x44, 0xa0, 0x7f, 0x7d, 0x9b,
	0xfe, 0x7e, 0x7d, 0xda, 0x78, 0xee, 0x17, 0xee, 0xea, 0x78, 0xd2, 0x9a, 0x8a, 0xdf, 0xa1, 0x8d,
	0x94, 0x97, 0x6c, 0xa2, 0x02, 0x88, 0x45, 0x78, 0x16, 0x00, 0xa7, 0x76, 0xd0, 0x07, 0x13, 0xf4,
	0xa0, 0x1e, 0x74, 0x5a, 0x16, 0xef, 0xa8, 0x03, 0x5d, 0x7a, 0xc0, 0x69, 0x25, 0x74, 0x2d, 0xbd,
	0xd2, 0xb1, 0xdb, 0x42, 0xb3, 0x05, 0x6f, 0xf7, 0xfe, 0x9b, 0x7b, 0xa1, 0xa0, 0x90, 0x87, 0xe4,
	0x57, 0x67, 0x28, 0x22, 0x97, 0x89, 0x9e, 0x7e, 0xee, 0xe5, 0x42, 0xaf, 0xbc, 0xcf, 0xfb, 0x33,
	0xb9, 0xf0, 0xe8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xa9, 0x4a, 0xf5, 0xe3, 0x05, 0x00,
	0x00,
}
