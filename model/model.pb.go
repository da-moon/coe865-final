// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

package model

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Hash ...
type Hash struct {
	Md5                  string   `protobuf:"bytes,1,opt,name=md5,proto3" json:"md5,omitempty"`
	Sha256               string   `protobuf:"bytes,2,opt,name=sha256,proto3" json:"sha256,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *Hash) Reset() { *m = Hash{} }

// String ...
func (m *Hash) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*Hash) ProtoMessage() {}

// Descriptor ...
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

// XXX_Unmarshal ...
func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(m, src)
}

// XXX_Size ...
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}

// XXX_DiscardUnknown ...
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

// GetMd5 ...
func (m *Hash) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

// GetSha256 ...
func (m *Hash) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

// RouteController ...
type RouteController struct {
	ID                     int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AutonomousSystemNumber int32    `protobuf:"varint,2,opt,name=autonomous_system_number,json=autonomousSystemNumber,proto3" json:"autonomous_system_number,omitempty"`
	IP                     string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

// Reset ...
func (m *RouteController) Reset() { *m = RouteController{} }

// String ...
func (m *RouteController) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*RouteController) ProtoMessage() {}

// Descriptor ...
func (*RouteController) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

// XXX_Unmarshal ...
func (m *RouteController) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteController.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *RouteController) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteController.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *RouteController) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteController.Merge(m, src)
}

// XXX_Size ...
func (m *RouteController) XXX_Size() int {
	return xxx_messageInfo_RouteController.Size(m)
}

// XXX_DiscardUnknown ...
func (m *RouteController) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteController.DiscardUnknown(m)
}

var xxx_messageInfo_RouteController proto.InternalMessageInfo

// GetID ...
func (m *RouteController) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

// GetAutonomousSystemNumber ...
func (m *RouteController) GetAutonomousSystemNumber() int32 {
	if m != nil {
		return m.AutonomousSystemNumber
	}
	return 0
}

// GetIP ...
func (m *RouteController) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

// AutonomousSystem ...
type AutonomousSystem struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	LinkCapacity         int32    `protobuf:"varint,2,opt,name=link_capacity,json=linkCapacity,proto3" json:"link_capacity,omitempty"`
	Cost                 int32    `protobuf:"varint,3,opt,name=cost,proto3" json:"cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *AutonomousSystem) Reset() { *m = AutonomousSystem{} }

// String ...
func (m *AutonomousSystem) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*AutonomousSystem) ProtoMessage() {}

// Descriptor ...
func (*AutonomousSystem) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2}
}

// XXX_Unmarshal ...
func (m *AutonomousSystem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutonomousSystem.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *AutonomousSystem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutonomousSystem.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *AutonomousSystem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutonomousSystem.Merge(m, src)
}

// XXX_Size ...
func (m *AutonomousSystem) XXX_Size() int {
	return xxx_messageInfo_AutonomousSystem.Size(m)
}

// XXX_DiscardUnknown ...
func (m *AutonomousSystem) XXX_DiscardUnknown() {
	xxx_messageInfo_AutonomousSystem.DiscardUnknown(m)
}

var xxx_messageInfo_AutonomousSystem proto.InternalMessageInfo

// GetNumber ...
func (m *AutonomousSystem) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

// GetLinkCapacity ...
func (m *AutonomousSystem) GetLinkCapacity() int32 {
	if m != nil {
		return m.LinkCapacity
	}
	return 0
}

// GetCost ...
func (m *AutonomousSystem) GetCost() int32 {
	if m != nil {
		return m.Cost
	}
	return 0
}

// UpdateRequest ...
type UpdateRequest struct {
	UUID                        string            `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	SourceRouteController       *RouteController  `protobuf:"bytes,2,opt,name=source_route_controller,json=sourceRouteController,proto3" json:"source_route_controller,omitempty"`
	DestinationAutonomousSystem *AutonomousSystem `protobuf:"bytes,3,opt,name=destination_autonomous_system,json=destinationAutonomousSystem,proto3" json:"destination_autonomous_system,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}          `json:"-"`
	XXX_unrecognized            []byte            `json:"-"`
	XXX_sizecache               int32             `json:"-"`
}

// Reset ...
func (m *UpdateRequest) Reset() { *m = UpdateRequest{} }

// String ...
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*UpdateRequest) ProtoMessage() {}

// Descriptor ...
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{3}
}

// XXX_Unmarshal ...
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}

// XXX_Size ...
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}

// XXX_DiscardUnknown ...
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

// GetUUID ...
func (m *UpdateRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

// GetSourceRouteController ...
func (m *UpdateRequest) GetSourceRouteController() *RouteController {
	if m != nil {
		return m.SourceRouteController
	}
	return nil
}

// GetDestinationAutonomousSystem ...
func (m *UpdateRequest) GetDestinationAutonomousSystem() *AutonomousSystem {
	if m != nil {
		return m.DestinationAutonomousSystem
	}
	return nil
}

// UpdateResponse ...
type UpdateResponse struct {
	UUID                        string            `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	DestinationAutonomousSystem *AutonomousSystem `protobuf:"bytes,2,opt,name=destination_autonomous_system,json=destinationAutonomousSystem,proto3" json:"destination_autonomous_system,omitempty"`
	Path                        []int32           `protobuf:"varint,3,rep,packed,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}          `json:"-"`
	XXX_unrecognized            []byte            `json:"-"`
	XXX_sizecache               int32             `json:"-"`
}

// Reset ...
func (m *UpdateResponse) Reset() { *m = UpdateResponse{} }

// String ...
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*UpdateResponse) ProtoMessage() {}

// Descriptor ...
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{4}
}

// XXX_Unmarshal ...
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}

// XXX_Size ...
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}

// XXX_DiscardUnknown ...
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

// GetUUID ...
func (m *UpdateResponse) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

// GetDestinationAutonomousSystem ...
func (m *UpdateResponse) GetDestinationAutonomousSystem() *AutonomousSystem {
	if m != nil {
		return m.DestinationAutonomousSystem
	}
	return nil
}

// GetPath ...
func (m *UpdateResponse) GetPath() []int32 {
	if m != nil {
		return m.Path
	}
	return nil
}

// KeyExchangeRequest ...
type KeyExchangeRequest struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Nonce                string   `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *KeyExchangeRequest) Reset() { *m = KeyExchangeRequest{} }

// String ...
func (m *KeyExchangeRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*KeyExchangeRequest) ProtoMessage() {}

// Descriptor ...
func (*KeyExchangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{5}
}

// XXX_Unmarshal ...
func (m *KeyExchangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyExchangeRequest.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *KeyExchangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyExchangeRequest.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *KeyExchangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyExchangeRequest.Merge(m, src)
}

// XXX_Size ...
func (m *KeyExchangeRequest) XXX_Size() int {
	return xxx_messageInfo_KeyExchangeRequest.Size(m)
}

// XXX_DiscardUnknown ...
func (m *KeyExchangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyExchangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyExchangeRequest proto.InternalMessageInfo

// GetUUID ...
func (m *KeyExchangeRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

// GetNonce ...
func (m *KeyExchangeRequest) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

// GetKey ...
func (m *KeyExchangeRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// KeyExchangeResponse ...
type KeyExchangeResponse struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	IsOk                 bool     `protobuf:"varint,2,opt,name=is_ok,json=isOk,proto3" json:"is_ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *KeyExchangeResponse) Reset() { *m = KeyExchangeResponse{} }

// String ...
func (m *KeyExchangeResponse) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*KeyExchangeResponse) ProtoMessage() {}

// Descriptor ...
func (*KeyExchangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{6}
}

// XXX_Unmarshal ...
func (m *KeyExchangeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyExchangeResponse.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *KeyExchangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyExchangeResponse.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *KeyExchangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyExchangeResponse.Merge(m, src)
}

// XXX_Size ...
func (m *KeyExchangeResponse) XXX_Size() int {
	return xxx_messageInfo_KeyExchangeResponse.Size(m)
}

// XXX_DiscardUnknown ...
func (m *KeyExchangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyExchangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeyExchangeResponse proto.InternalMessageInfo

// GetUUID ...
func (m *KeyExchangeResponse) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

// GetIsOk ...
func (m *KeyExchangeResponse) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func init() {
	proto.RegisterType((*Hash)(nil), "model.Hash")
	proto.RegisterType((*RouteController)(nil), "model.RouteController")
	proto.RegisterType((*AutonomousSystem)(nil), "model.AutonomousSystem")
	proto.RegisterType((*UpdateRequest)(nil), "model.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "model.UpdateResponse")
	proto.RegisterType((*KeyExchangeRequest)(nil), "model.KeyExchangeRequest")
	proto.RegisterType((*KeyExchangeResponse)(nil), "model.KeyExchangeResponse")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4f, 0x6f, 0xd4, 0x3e,
	0x10, 0xfd, 0x6d, 0x76, 0xb3, 0x6a, 0x67, 0xdb, 0xfe, 0x2a, 0xb7, 0xdd, 0x86, 0x05, 0xd4, 0x2a,
	0x5c, 0xb8, 0xb0, 0x45, 0x8b, 0x8a, 0xb8, 0x70, 0xa0, 0x7f, 0x50, 0x2b, 0xa4, 0x2d, 0x32, 0xda,
	0x0b, 0x1c, 0x22, 0x6f, 0x62, 0x36, 0x56, 0x12, 0x3b, 0xc4, 0x36, 0x10, 0xf1, 0x41, 0x38, 0xf2,
	0xbd, 0x38, 0xf4, 0xd0, 0x4f, 0x82, 0xec, 0xb8, 0xd0, 0xa6, 0x48, 0xad, 0xb8, 0xcd, 0xcc, 0xcb,
	0xbc, 0x79, 0x7e, 0x33, 0x81, 0x41, 0x21, 0x12, 0x9a, 0x8f, 0xcb, 0x4a, 0x28, 0x81, 0x7c, 0x9b,
	0x8c, 0x9e, 0x2c, 0x98, 0x4a, 0xf5, 0x7c, 0x1c, 0x8b, 0x62, 0x6f, 0x21, 0x16, 0x62, 0xcf, 0xa2,
	0x73, 0xfd, 0xd1, 0x66, 0x36, 0xb1, 0x51, 0xd3, 0x15, 0x3e, 0x85, 0xde, 0x09, 0x91, 0x29, 0x5a,
	0x87, 0x6e, 0x91, 0xec, 0x07, 0x9d, 0xdd, 0xce, 0xe3, 0x65, 0x6c, 0x42, 0x34, 0x84, 0xbe, 0x4c,
	0xc9, 0x64, 0xff, 0x79, 0xe0, 0xd9, 0xa2, 0xcb, 0xc2, 0x6f, 0xf0, 0x3f, 0x16, 0x5a, 0xd1, 0x43,
	0xc1, 0x55, 0x25, 0xf2, 0x9c, 0x56, 0x68, 0x08, 0x1e, 0x4b, 0x6c, 0xaf, 0x7f, 0xd0, 0xbf, 0x38,
	0xdf, 0xf1, 0x4e, 0x8f, 0xb0, 0xc7, 0x12, 0xf4, 0x02, 0x02, 0xa2, 0x95, 0xe0, 0xa2, 0x10, 0x5a,
	0x46, 0xb2, 0x96, 0x8a, 0x16, 0x11, 0xd7, 0xc5, 0x9c, 0x56, 0x96, 0xd4, 0xc7, 0xc3, 0x3f, 0xf8,
	0x3b, 0x0b, 0x4f, 0x2d, 0x6a, 0x19, 0xcb, 0xa0, 0x6b, 0x06, 0x3b, 0xc6, 0xb7, 0xd8, 0x63, 0x65,
	0x18, 0xc3, 0xfa, 0xab, 0x56, 0x87, 0x11, 0xea, 0x38, 0xad, 0x02, 0xec, 0x32, 0xf4, 0x08, 0x56,
	0x73, 0xc6, 0xb3, 0x28, 0x26, 0x25, 0x89, 0x99, 0xaa, 0xdd, 0xc8, 0x15, 0x53, 0x3c, 0x74, 0x35,
	0x84, 0xa0, 0x17, 0x0b, 0xa9, 0xec, 0x28, 0x1f, 0xdb, 0x38, 0xfc, 0xd9, 0x81, 0xd5, 0x59, 0x99,
	0x10, 0x45, 0x31, 0xfd, 0xa4, 0xa9, 0x54, 0xe8, 0x01, 0xf4, 0xb4, 0x76, 0x4f, 0x5c, 0x3e, 0x58,
	0xba, 0x38, 0xdf, 0xe9, 0xcd, 0x66, 0xa7, 0x47, 0xd8, 0x56, 0xd1, 0x14, 0xb6, 0xa5, 0xd0, 0x55,
	0x4c, 0xa3, 0xca, 0x18, 0x13, 0xc5, 0xbf, 0x9d, 0xb1, 0x23, 0x07, 0x93, 0xe1, 0xb8, 0x59, 0x54,
	0xcb, 0x37, 0xbc, 0xd5, 0xb4, 0xb5, 0xed, 0xfc, 0x00, 0x0f, 0x13, 0x2a, 0x15, 0xe3, 0x44, 0x31,
	0xc1, 0xa3, 0x1b, 0x16, 0x5a, 0xb1, 0x83, 0xc9, 0xb6, 0x63, 0x6d, 0x1b, 0x82, 0xef, 0x5f, 0xe9,
	0x6e, 0x83, 0xe1, 0x8f, 0x0e, 0xac, 0x5d, 0x3e, 0x4e, 0x96, 0x82, 0x4b, 0x7a, 0xcb, 0xeb, 0x6e,
	0x55, 0xe3, 0xfd, 0xbb, 0x1a, 0x63, 0x7f, 0x49, 0x54, 0x1a, 0x74, 0x77, 0xbb, 0xc6, 0x7e, 0x13,
	0x87, 0xef, 0x01, 0xbd, 0xa1, 0xf5, 0xf1, 0xd7, 0x38, 0x25, 0x7c, 0x71, 0xc7, 0x15, 0x6c, 0x82,
	0xcf, 0x05, 0x8f, 0xa9, 0xbb, 0xd5, 0x26, 0x31, 0x47, 0x9d, 0xd1, 0xba, 0x39, 0x23, 0x6c, 0xc2,
	0xf0, 0x04, 0x36, 0xae, 0x71, 0xdf, 0xc9, 0x81, 0x0d, 0xf0, 0x99, 0x8c, 0x44, 0x66, 0xc9, 0x97,
	0x70, 0x8f, 0xc9, 0xb3, 0x6c, 0xf2, 0xbd, 0x03, 0x6b, 0x67, 0x9f, 0x69, 0x95, 0x93, 0x7a, 0x4a,
	0xd5, 0x17, 0x51, 0x65, 0xe8, 0x25, 0xac, 0x1c, 0x4b, 0xc5, 0x0a, 0x62, 0xb6, 0x29, 0x15, 0xda,
	0x74, 0x96, 0x5c, 0xbb, 0xa5, 0xd1, 0x56, 0xab, 0xda, 0x48, 0x08, 0xff, 0x43, 0xaf, 0x61, 0x70,
	0x45, 0x1b, 0xba, 0xe7, 0xbe, 0xbb, 0xe9, 0xc5, 0x68, 0xf4, 0x37, 0xe8, 0x92, 0x67, 0xde, 0xb7,
	0x7f, 0xf6, 0xb3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xe5, 0x7b, 0x82, 0x1e, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OverlayNetworkClient is the client API for OverlayNetwork service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OverlayNetworkClient interface {
	EstimateCost(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	KeyExchange(ctx context.Context, in *KeyExchangeRequest, opts ...grpc.CallOption) (*KeyExchangeResponse, error)
}

type overlayNetworkClient struct {
	cc *grpc.ClientConn
}

// NewOverlayNetworkClient ...
func NewOverlayNetworkClient(cc *grpc.ClientConn) OverlayNetworkClient {
	return &overlayNetworkClient{cc}
}

// EstimateCost ...
func (c *overlayNetworkClient) EstimateCost(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/model.OverlayNetwork/EstimateCost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyExchange ...
func (c *overlayNetworkClient) KeyExchange(ctx context.Context, in *KeyExchangeRequest, opts ...grpc.CallOption) (*KeyExchangeResponse, error) {
	out := new(KeyExchangeResponse)
	err := c.cc.Invoke(ctx, "/model.OverlayNetwork/KeyExchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OverlayNetworkServer is the server API for OverlayNetwork service.
type OverlayNetworkServer interface {
	EstimateCost(context.Context, *UpdateRequest) (*UpdateResponse, error)
	KeyExchange(context.Context, *KeyExchangeRequest) (*KeyExchangeResponse, error)
}

// UnimplementedOverlayNetworkServer can be embedded to have forward compatible implementations.
type UnimplementedOverlayNetworkServer struct {
}

// EstimateCost ...
func (*UnimplementedOverlayNetworkServer) EstimateCost(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateCost not implemented")
}

// KeyExchange ...
func (*UnimplementedOverlayNetworkServer) KeyExchange(ctx context.Context, req *KeyExchangeRequest) (*KeyExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyExchange not implemented")
}

// RegisterOverlayNetworkServer ...
func RegisterOverlayNetworkServer(s *grpc.Server, srv OverlayNetworkServer) {
	s.RegisterService(&_OverlayNetwork_serviceDesc, srv)
}

func _OverlayNetwork_EstimateCost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayNetworkServer).EstimateCost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.OverlayNetwork/EstimateCost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayNetworkServer).EstimateCost(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OverlayNetwork_KeyExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayNetworkServer).KeyExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.OverlayNetwork/KeyExchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayNetworkServer).KeyExchange(ctx, req.(*KeyExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OverlayNetwork_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.OverlayNetwork",
	HandlerType: (*OverlayNetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EstimateCost",
			Handler:    _OverlayNetwork_EstimateCost_Handler,
		},
		{
			MethodName: "KeyExchange",
			Handler:    _OverlayNetwork_KeyExchange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model.proto",
}
