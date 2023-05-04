// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vega/data/v1/data.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ETHAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ETHAddress) Reset() {
	*x = ETHAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ETHAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ETHAddress) ProtoMessage() {}

func (x *ETHAddress) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ETHAddress.ProtoReflect.Descriptor instead.
func (*ETHAddress) Descriptor() ([]byte, []int) {
	return file_vega_data_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *ETHAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// PubKey is the public key that signed this data.
// Different public keys coming from different sources will be further separated.
type PubKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PubKey) Reset() {
	*x = PubKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubKey) ProtoMessage() {}

func (x *PubKey) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubKey.ProtoReflect.Descriptor instead.
func (*PubKey) Descriptor() ([]byte, []int) {
	return file_vega_data_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *PubKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Signer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Signer:
	//
	//	*Signer_PubKey
	//	*Signer_EthAddress
	Signer isSigner_Signer `protobuf_oneof:"signer"`
}

func (x *Signer) Reset() {
	*x = Signer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_v1_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signer) ProtoMessage() {}

func (x *Signer) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_v1_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signer.ProtoReflect.Descriptor instead.
func (*Signer) Descriptor() ([]byte, []int) {
	return file_vega_data_v1_data_proto_rawDescGZIP(), []int{2}
}

func (m *Signer) GetSigner() isSigner_Signer {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (x *Signer) GetPubKey() *PubKey {
	if x, ok := x.GetSigner().(*Signer_PubKey); ok {
		return x.PubKey
	}
	return nil
}

func (x *Signer) GetEthAddress() *ETHAddress {
	if x, ok := x.GetSigner().(*Signer_EthAddress); ok {
		return x.EthAddress
	}
	return nil
}

type isSigner_Signer interface {
	isSigner_Signer()
}

type Signer_PubKey struct {
	// List of authorized public keys that signed the data for this
	// source. All the public keys in the data should be contained in these
	// public keys.
	PubKey *PubKey `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3,oneof"`
}

type Signer_EthAddress struct {
	// In case of an open oracle - Ethereum address will be submitted.
	EthAddress *ETHAddress `protobuf:"bytes,2,opt,name=eth_address,json=ethAddress,proto3,oneof"`
}

func (*Signer_PubKey) isSigner_Signer() {}

func (*Signer_EthAddress) isSigner_Signer() {}

// Property describes one property of data spec with a key with its value.
type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the property.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Value of the property.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_v1_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_v1_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_vega_data_v1_data_proto_rawDescGZIP(), []int{3}
}

func (x *Property) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Property) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Data describes valid source data that has been received by the node.
// It represents both matched and unmatched data.
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signers []*Signer `protobuf:"bytes,1,rep,name=signers,proto3" json:"signers,omitempty"`
	// Data holds all the properties of the data
	Data []*Property `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	// `matched_specs_ids` lists all the specs that matched this data.
	// When the array is empty, it means no spec matched this data.
	MatchedSpecIds []string `protobuf:"bytes,3,rep,name=matched_spec_ids,json=matchedSpecIds,proto3" json:"matched_spec_ids,omitempty"`
	// `broadcast_at` is the time at which the data was broadcast to the markets
	// with a matching spec.
	// It has no value when the date did not match any spec.
	// The value is a Unix timestamp in nanoseconds.
	BroadcastAt int64 `protobuf:"varint,4,opt,name=broadcast_at,json=broadcastAt,proto3" json:"broadcast_at,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_v1_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_v1_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_vega_data_v1_data_proto_rawDescGZIP(), []int{4}
}

func (x *Data) GetSigners() []*Signer {
	if x != nil {
		return x.Signers
	}
	return nil
}

func (x *Data) GetData() []*Property {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Data) GetMatchedSpecIds() []string {
	if x != nil {
		return x.MatchedSpecIds
	}
	return nil
}

func (x *Data) GetBroadcastAt() int64 {
	if x != nil {
		return x.BroadcastAt
	}
	return 0
}

type ExternalData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ExternalData) Reset() {
	*x = ExternalData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_v1_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalData) ProtoMessage() {}

func (x *ExternalData) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_v1_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalData.ProtoReflect.Descriptor instead.
func (*ExternalData) Descriptor() ([]byte, []int) {
	return file_vega_data_v1_data_proto_rawDescGZIP(), []int{5}
}

func (x *ExternalData) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_vega_data_v1_data_proto protoreflect.FileDescriptor

var file_vega_data_v1_data_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x76, 0x65, 0x67, 0x61, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x22, 0x26, 0x0a, 0x0a, 0x45, 0x54, 0x48, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x1a, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x80, 0x01, 0x0a, 0x06,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52,
	0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x65, 0x74, 0x68, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76,
	0x65, 0x67, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x54, 0x48, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x74, 0x68, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x22, 0x34,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a,
	0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x65,
	0x67, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x53, 0x70, 0x65, 0x63,
	0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x41, 0x74, 0x22, 0x36, 0x0a, 0x0c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2f,
	0x5a, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vega_data_v1_data_proto_rawDescOnce sync.Once
	file_vega_data_v1_data_proto_rawDescData = file_vega_data_v1_data_proto_rawDesc
)

func file_vega_data_v1_data_proto_rawDescGZIP() []byte {
	file_vega_data_v1_data_proto_rawDescOnce.Do(func() {
		file_vega_data_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_vega_data_v1_data_proto_rawDescData)
	})
	return file_vega_data_v1_data_proto_rawDescData
}

var file_vega_data_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vega_data_v1_data_proto_goTypes = []interface{}{
	(*ETHAddress)(nil),   // 0: vega.data.v1.ETHAddress
	(*PubKey)(nil),       // 1: vega.data.v1.PubKey
	(*Signer)(nil),       // 2: vega.data.v1.Signer
	(*Property)(nil),     // 3: vega.data.v1.Property
	(*Data)(nil),         // 4: vega.data.v1.Data
	(*ExternalData)(nil), // 5: vega.data.v1.ExternalData
}
var file_vega_data_v1_data_proto_depIdxs = []int32{
	1, // 0: vega.data.v1.Signer.pub_key:type_name -> vega.data.v1.PubKey
	0, // 1: vega.data.v1.Signer.eth_address:type_name -> vega.data.v1.ETHAddress
	2, // 2: vega.data.v1.Data.signers:type_name -> vega.data.v1.Signer
	3, // 3: vega.data.v1.Data.data:type_name -> vega.data.v1.Property
	4, // 4: vega.data.v1.ExternalData.data:type_name -> vega.data.v1.Data
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_vega_data_v1_data_proto_init() }
func file_vega_data_v1_data_proto_init() {
	if File_vega_data_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vega_data_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ETHAddress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_data_v1_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_data_v1_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_data_v1_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_data_v1_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_data_v1_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_vega_data_v1_data_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Signer_PubKey)(nil),
		(*Signer_EthAddress)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vega_data_v1_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vega_data_v1_data_proto_goTypes,
		DependencyIndexes: file_vega_data_v1_data_proto_depIdxs,
		MessageInfos:      file_vega_data_v1_data_proto_msgTypes,
	}.Build()
	File_vega_data_v1_data_proto = out.File
	file_vega_data_v1_data_proto_rawDesc = nil
	file_vega_data_v1_data_proto_goTypes = nil
	file_vega_data_v1_data_proto_depIdxs = nil
}
