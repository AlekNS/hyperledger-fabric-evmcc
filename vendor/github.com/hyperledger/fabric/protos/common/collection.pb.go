// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/collection.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/collection.proto
	common/common.proto
	common/configtx.proto
	common/configuration.proto
	common/ledger.proto
	common/policies.proto

It has these top-level messages:
	CollectionConfigPackage
	CollectionConfig
	StaticCollectionConfig
	CollectionPolicyConfig
	CollectionCriteria
	LastConfig
	Metadata
	MetadataSignature
	Header
	ChannelHeader
	SignatureHeader
	Payload
	Envelope
	Block
	BlockHeader
	BlockData
	BlockMetadata
	ConfigEnvelope
	ConfigGroupSchema
	ConfigValueSchema
	ConfigPolicySchema
	Config
	ConfigUpdateEnvelope
	ConfigUpdate
	ConfigGroup
	ConfigValue
	ConfigPolicy
	ConfigSignature
	HashingAlgorithm
	BlockDataHashingStructure
	OrdererAddresses
	Consortium
	Capabilities
	Capability
	BlockchainInfo
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
	ImplicitMetaPolicy
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// CollectionConfigPackage represents an array of CollectionConfig
// messages; the extra struct is required because repeated oneof is
// forbidden by the protobuf syntax
type CollectionConfigPackage struct {
	Config []*CollectionConfig `protobuf:"bytes,1,rep,name=config" json:"config,omitempty"`
}

func (m *CollectionConfigPackage) Reset()                    { *m = CollectionConfigPackage{} }
func (m *CollectionConfigPackage) String() string            { return proto.CompactTextString(m) }
func (*CollectionConfigPackage) ProtoMessage()               {}
func (*CollectionConfigPackage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CollectionConfigPackage) GetConfig() []*CollectionConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// CollectionConfig defines the configuration of a collection object;
// it currently contains a single, static type.
// Dynamic collections are deferred.
type CollectionConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionConfig_StaticCollectionConfig
	Payload isCollectionConfig_Payload `protobuf_oneof:"payload"`
}

func (m *CollectionConfig) Reset()                    { *m = CollectionConfig{} }
func (m *CollectionConfig) String() string            { return proto.CompactTextString(m) }
func (*CollectionConfig) ProtoMessage()               {}
func (*CollectionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isCollectionConfig_Payload interface {
	isCollectionConfig_Payload()
}

type CollectionConfig_StaticCollectionConfig struct {
	StaticCollectionConfig *StaticCollectionConfig `protobuf:"bytes,1,opt,name=static_collection_config,json=staticCollectionConfig,oneof"`
}

func (*CollectionConfig_StaticCollectionConfig) isCollectionConfig_Payload() {}

func (m *CollectionConfig) GetPayload() isCollectionConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CollectionConfig) GetStaticCollectionConfig() *StaticCollectionConfig {
	if x, ok := m.GetPayload().(*CollectionConfig_StaticCollectionConfig); ok {
		return x.StaticCollectionConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CollectionConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CollectionConfig_OneofMarshaler, _CollectionConfig_OneofUnmarshaler, _CollectionConfig_OneofSizer, []interface{}{
		(*CollectionConfig_StaticCollectionConfig)(nil),
	}
}

func _CollectionConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CollectionConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionConfig_StaticCollectionConfig:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StaticCollectionConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CollectionConfig.Payload has unexpected type %T", x)
	}
	return nil
}

func _CollectionConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CollectionConfig)
	switch tag {
	case 1: // payload.static_collection_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StaticCollectionConfig)
		err := b.DecodeMessage(msg)
		m.Payload = &CollectionConfig_StaticCollectionConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CollectionConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CollectionConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionConfig_StaticCollectionConfig:
		s := proto.Size(x.StaticCollectionConfig)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// StaticCollectionConfig constitutes the configuration parameters of a
// static collection object. Static collections are collections that are
// known at chaincode instantiation time, and that cannot be changed.
// Dynamic collections are deferred.
type StaticCollectionConfig struct {
	// the name of the collection inside the denoted chaincode
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// a reference to a policy residing / managed in the config block
	// to define which orgs have access to this collection’s private data
	MemberOrgsPolicy *CollectionPolicyConfig `protobuf:"bytes,2,opt,name=member_orgs_policy,json=memberOrgsPolicy" json:"member_orgs_policy,omitempty"`
	// the minimum number of internal/external peers required to be sent
	// private data to
	RequiredInternalPeerCount int32 `protobuf:"varint,3,opt,name=required_internal_peer_count,json=requiredInternalPeerCount" json:"required_internal_peer_count,omitempty"`
	RequiredExternalPeerCount int32 `protobuf:"varint,4,opt,name=required_external_peer_count,json=requiredExternalPeerCount" json:"required_external_peer_count,omitempty"`
}

func (m *StaticCollectionConfig) Reset()                    { *m = StaticCollectionConfig{} }
func (m *StaticCollectionConfig) String() string            { return proto.CompactTextString(m) }
func (*StaticCollectionConfig) ProtoMessage()               {}
func (*StaticCollectionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StaticCollectionConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StaticCollectionConfig) GetMemberOrgsPolicy() *CollectionPolicyConfig {
	if m != nil {
		return m.MemberOrgsPolicy
	}
	return nil
}

func (m *StaticCollectionConfig) GetRequiredInternalPeerCount() int32 {
	if m != nil {
		return m.RequiredInternalPeerCount
	}
	return 0
}

func (m *StaticCollectionConfig) GetRequiredExternalPeerCount() int32 {
	if m != nil {
		return m.RequiredExternalPeerCount
	}
	return 0
}

// Collection policy configuration. Initially, the configuration can only
// contain a SignaturePolicy. In the future, the SignaturePolicy may be a
// more general Policy. Instead of containing the actual policy, the
// configuration may in the future contain a string reference to a policy.
type CollectionPolicyConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionPolicyConfig_SignaturePolicy
	Payload isCollectionPolicyConfig_Payload `protobuf_oneof:"payload"`
}

func (m *CollectionPolicyConfig) Reset()                    { *m = CollectionPolicyConfig{} }
func (m *CollectionPolicyConfig) String() string            { return proto.CompactTextString(m) }
func (*CollectionPolicyConfig) ProtoMessage()               {}
func (*CollectionPolicyConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isCollectionPolicyConfig_Payload interface {
	isCollectionPolicyConfig_Payload()
}

type CollectionPolicyConfig_SignaturePolicy struct {
	SignaturePolicy *SignaturePolicyEnvelope `protobuf:"bytes,1,opt,name=signature_policy,json=signaturePolicy,oneof"`
}

func (*CollectionPolicyConfig_SignaturePolicy) isCollectionPolicyConfig_Payload() {}

func (m *CollectionPolicyConfig) GetPayload() isCollectionPolicyConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CollectionPolicyConfig) GetSignaturePolicy() *SignaturePolicyEnvelope {
	if x, ok := m.GetPayload().(*CollectionPolicyConfig_SignaturePolicy); ok {
		return x.SignaturePolicy
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CollectionPolicyConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CollectionPolicyConfig_OneofMarshaler, _CollectionPolicyConfig_OneofUnmarshaler, _CollectionPolicyConfig_OneofSizer, []interface{}{
		(*CollectionPolicyConfig_SignaturePolicy)(nil),
	}
}

func _CollectionPolicyConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CollectionPolicyConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionPolicyConfig_SignaturePolicy:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SignaturePolicy); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CollectionPolicyConfig.Payload has unexpected type %T", x)
	}
	return nil
}

func _CollectionPolicyConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CollectionPolicyConfig)
	switch tag {
	case 1: // payload.signature_policy
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SignaturePolicyEnvelope)
		err := b.DecodeMessage(msg)
		m.Payload = &CollectionPolicyConfig_SignaturePolicy{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CollectionPolicyConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CollectionPolicyConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionPolicyConfig_SignaturePolicy:
		s := proto.Size(x.SignaturePolicy)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// CollectionCriteria defines an element of a private data that corresponds
// to a certain transaction and collection
type CollectionCriteria struct {
	Channel    string `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
	TxId       string `protobuf:"bytes,2,opt,name=tx_id,json=txId" json:"tx_id,omitempty"`
	Collection string `protobuf:"bytes,3,opt,name=collection" json:"collection,omitempty"`
	Namespace  string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *CollectionCriteria) Reset()                    { *m = CollectionCriteria{} }
func (m *CollectionCriteria) String() string            { return proto.CompactTextString(m) }
func (*CollectionCriteria) ProtoMessage()               {}
func (*CollectionCriteria) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CollectionCriteria) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *CollectionCriteria) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *CollectionCriteria) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *CollectionCriteria) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*CollectionConfigPackage)(nil), "common.CollectionConfigPackage")
	proto.RegisterType((*CollectionConfig)(nil), "common.CollectionConfig")
	proto.RegisterType((*StaticCollectionConfig)(nil), "common.StaticCollectionConfig")
	proto.RegisterType((*CollectionPolicyConfig)(nil), "common.CollectionPolicyConfig")
	proto.RegisterType((*CollectionCriteria)(nil), "common.CollectionCriteria")
}

func init() { proto.RegisterFile("common/collection.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0xa3, 0xc6, 0x71, 0xd0, 0xe4, 0x50, 0xb3, 0xa5, 0x8e, 0x5a, 0x42, 0x1a, 0x4c, 0x0f,
	0x86, 0x82, 0x54, 0xd2, 0x0f, 0x50, 0x88, 0x09, 0x24, 0x34, 0x50, 0xa3, 0xdc, 0x72, 0x11, 0xeb,
	0xd5, 0x44, 0x5e, 0x2a, 0xed, 0xca, 0xb3, 0xeb, 0x62, 0x1f, 0xfb, 0xbd, 0x7b, 0x28, 0xde, 0x95,
	0xfc, 0x47, 0xf5, 0xcd, 0x33, 0xef, 0x37, 0xcf, 0x33, 0x4f, 0x0b, 0x97, 0x42, 0x57, 0x95, 0x56,
	0x89, 0xd0, 0x65, 0x89, 0xc2, 0x4a, 0xad, 0xe2, 0x9a, 0xb4, 0xd5, 0xac, 0xef, 0x85, 0x8f, 0xef,
	0x1b, 0xa0, 0xd6, 0xa5, 0x14, 0x12, 0x8d, 0x97, 0x47, 0x3f, 0xe0, 0x72, 0xb2, 0x1d, 0x99, 0x68,
	0xf5, 0x2a, 0x8b, 0x29, 0x17, 0xbf, 0x78, 0x81, 0xec, 0x2b, 0xf4, 0x85, 0x6b, 0x44, 0xc1, 0xcd,
	0xe9, 0xf8, 0xe2, 0x36, 0x8a, 0xbd, 0x45, 0xdc, 0x1d, 0x48, 0x1b, 0x6e, 0xb4, 0x86, 0x41, 0x57,
	0x63, 0x2f, 0x10, 0x19, 0xcb, 0xad, 0x14, 0xd9, 0x6e, 0xb5, 0x6c, 0xeb, 0x1b, 0x8c, 0x2f, 0x6e,
	0xaf, 0x5b, 0xdf, 0x67, 0xc7, 0x75, 0x1d, 0x1e, 0x4e, 0xd2, 0xa1, 0x39, 0xaa, 0xdc, 0x85, 0x70,
	0x5e, 0xf3, 0x75, 0xa9, 0x79, 0x3e, 0xfa, 0x1b, 0xc0, 0xf0, 0xf8, 0x3c, 0x63, 0xd0, 0x53, 0xbc,
	0x42, 0xf7, 0x6f, 0x61, 0xea, 0x7e, 0xb3, 0x27, 0x60, 0x15, 0x56, 0x33, 0xa4, 0x4c, 0x53, 0x61,
	0x32, 0x17, 0xca, 0x3a, 0x7a, 0x73, 0xb8, 0xcf, 0xce, 0x69, 0xea, 0xf4, 0xe6, 0xda, 0x81, 0x9f,
	0xfc, 0x49, 0x85, 0xf1, 0x7d, 0xf6, 0x1d, 0xae, 0x08, 0x17, 0x4b, 0x49, 0x98, 0x67, 0x52, 0x59,
	0x24, 0xc5, 0xcb, 0xac, 0x46, 0xa4, 0x4c, 0xe8, 0xa5, 0xb2, 0xd1, 0xe9, 0x4d, 0x30, 0x3e, 0x4b,
	0x3f, 0xb4, 0xcc, 0x63, 0x83, 0x4c, 0x11, 0x69, 0xb2, 0x01, 0x0e, 0x0c, 0x70, 0xf5, 0xbf, 0x41,
	0xef, 0xd0, 0xe0, 0x7e, 0xd5, 0x31, 0x18, 0x2d, 0x60, 0x78, 0x7c, 0x5b, 0xf6, 0x04, 0x03, 0x23,
	0x0b, 0xc5, 0xed, 0x92, 0xb0, 0xbd, 0xd3, 0xe7, 0xfe, 0x69, 0x9b, 0x7b, 0xab, 0xfb, 0xc1, 0x7b,
	0xf5, 0x1b, 0x4b, 0x5d, 0xe3, 0xc3, 0x49, 0xfa, 0xd6, 0x1c, 0x4a, 0xfb, 0x89, 0xff, 0x09, 0x80,
	0xed, 0x65, 0x4d, 0xd2, 0x22, 0x49, 0xce, 0x22, 0x38, 0x17, 0x73, 0xae, 0x14, 0x96, 0x4d, 0xe0,
	0x6d, 0xc9, 0xde, 0xc1, 0x99, 0x5d, 0x65, 0x32, 0x77, 0x31, 0x87, 0x69, 0xcf, 0xae, 0x1e, 0x73,
	0x76, 0x0d, 0xb0, 0x7b, 0x17, 0x2e, 0xa8, 0x30, 0xdd, 0xeb, 0xb0, 0x2b, 0x08, 0x37, 0x1f, 0xcc,
	0xd4, 0x5c, 0xa0, 0x8b, 0x21, 0x4c, 0x77, 0x8d, 0xbb, 0x67, 0xf8, 0xac, 0xa9, 0x88, 0xe7, 0xeb,
	0x1a, 0xa9, 0xc4, 0xbc, 0x40, 0x8a, 0x5f, 0xf9, 0x8c, 0xa4, 0xf0, 0xaf, 0xdb, 0x34, 0x17, 0xbe,
	0x7c, 0x29, 0xa4, 0x9d, 0x2f, 0x67, 0x9b, 0x32, 0xd9, 0x83, 0x13, 0x0f, 0x27, 0x1e, 0x4e, 0x3c,
	0x3c, 0xeb, 0xbb, 0xf2, 0xdb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xbf, 0xe9, 0xaa, 0x53,
	0x03, 0x00, 0x00,
}
