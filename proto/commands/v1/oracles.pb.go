// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands/v1/oracles.proto

package v1

import (
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

// The supported Oracle sources
type OracleDataSubmission_OracleSource int32

const (
	// The default value
	OracleDataSubmission_ORACLE_SOURCE_UNSPECIFIED OracleDataSubmission_OracleSource = 0
	// Support for Open Oracle standard
	OracleDataSubmission_ORACLE_SOURCE_OPEN_ORACLE OracleDataSubmission_OracleSource = 1
)

var OracleDataSubmission_OracleSource_name = map[int32]string{
	0: "ORACLE_SOURCE_UNSPECIFIED",
	1: "ORACLE_SOURCE_OPEN_ORACLE",
}

var OracleDataSubmission_OracleSource_value = map[string]int32{
	"ORACLE_SOURCE_UNSPECIFIED": 0,
	"ORACLE_SOURCE_OPEN_ORACLE": 1,
}

func (x OracleDataSubmission_OracleSource) String() string {
	return proto.EnumName(OracleDataSubmission_OracleSource_name, int32(x))
}

func (OracleDataSubmission_OracleSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_836c9d5292a73d41, []int{0, 0}
}

// Command to submit new Oracle data from third party providers
type OracleDataSubmission struct {
	// The source from which the data is coming from
	Source OracleDataSubmission_OracleSource `protobuf:"varint,1,opt,name=source,proto3,enum=vega.commands.v1.OracleDataSubmission_OracleSource" json:"source,omitempty"`
	// The data provided by the third party provider
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OracleDataSubmission) Reset()         { *m = OracleDataSubmission{} }
func (m *OracleDataSubmission) String() string { return proto.CompactTextString(m) }
func (*OracleDataSubmission) ProtoMessage()    {}
func (*OracleDataSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_836c9d5292a73d41, []int{0}
}

func (m *OracleDataSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleDataSubmission.Unmarshal(m, b)
}
func (m *OracleDataSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleDataSubmission.Marshal(b, m, deterministic)
}
func (m *OracleDataSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleDataSubmission.Merge(m, src)
}
func (m *OracleDataSubmission) XXX_Size() int {
	return xxx_messageInfo_OracleDataSubmission.Size(m)
}
func (m *OracleDataSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleDataSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_OracleDataSubmission proto.InternalMessageInfo

func (m *OracleDataSubmission) GetSource() OracleDataSubmission_OracleSource {
	if m != nil {
		return m.Source
	}
	return OracleDataSubmission_ORACLE_SOURCE_UNSPECIFIED
}

func (m *OracleDataSubmission) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("vega.commands.v1.OracleDataSubmission_OracleSource", OracleDataSubmission_OracleSource_name, OracleDataSubmission_OracleSource_value)
	proto.RegisterType((*OracleDataSubmission)(nil), "vega.commands.v1.OracleDataSubmission")
}

func init() { proto.RegisterFile("commands/v1/oracles.proto", fileDescriptor_836c9d5292a73d41) }

var fileDescriptor_836c9d5292a73d41 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0x29, 0xd6, 0x2f, 0x33, 0xd4, 0xcf, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0x2d, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0x4b, 0x4d, 0x4f, 0xd4, 0x83, 0xc9, 0xeb, 0x95,
	0x19, 0x2a, 0x9d, 0x66, 0xe4, 0x12, 0xf1, 0x07, 0xab, 0x71, 0x49, 0x2c, 0x49, 0x0c, 0x2e, 0x4d,
	0xca, 0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0x13, 0xf2, 0xe6, 0x62, 0x2b, 0xce, 0x2f, 0x2d, 0x4a,
	0x4e, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x32, 0xd6, 0x43, 0xd7, 0xab, 0x87, 0x4d, 0x1f,
	0x54, 0x30, 0x18, 0xac, 0x35, 0x08, 0x6a, 0x84, 0x90, 0x04, 0x17, 0x7b, 0x41, 0x62, 0x65, 0x4e,
	0x7e, 0x62, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x8c, 0xab, 0xe4, 0xc3, 0xc5, 0x83,
	0xac, 0x43, 0x48, 0x96, 0x4b, 0xd2, 0x3f, 0xc8, 0xd1, 0xd9, 0xc7, 0x35, 0x3e, 0xd8, 0x3f, 0x34,
	0xc8, 0xd9, 0x35, 0x3e, 0xd4, 0x2f, 0x38, 0xc0, 0xd5, 0xd9, 0xd3, 0xcd, 0xd3, 0xd5, 0x45, 0x80,
	0x01, 0x53, 0xda, 0x3f, 0xc0, 0xd5, 0x2f, 0x1e, 0x22, 0x24, 0xc0, 0xe8, 0xa4, 0x1b, 0xa5, 0x9d,
	0x9c, 0x9f, 0x92, 0x0a, 0x76, 0x2a, 0xd8, 0xc7, 0xc9, 0xf9, 0x39, 0x7a, 0x99, 0xf9, 0xfa, 0x20,
	0xbe, 0x3e, 0x58, 0x40, 0x1f, 0x29, 0x70, 0x92, 0xd8, 0xc0, 0x42, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xba, 0x33, 0x8d, 0xb2, 0x32, 0x01, 0x00, 0x00,
}
