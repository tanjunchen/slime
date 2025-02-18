// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plugin_module.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PluginModule struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginModule) Reset()         { *m = PluginModule{} }
func (m *PluginModule) String() string { return proto.CompactTextString(m) }
func (*PluginModule) ProtoMessage()    {}
func (*PluginModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_c98ebf9b1715be96, []int{0}
}
func (m *PluginModule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginModule.Unmarshal(m, b)
}
func (m *PluginModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginModule.Marshal(b, m, deterministic)
}
func (m *PluginModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginModule.Merge(m, src)
}
func (m *PluginModule) XXX_Size() int {
	return xxx_messageInfo_PluginModule.Size(m)
}
func (m *PluginModule) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginModule.DiscardUnknown(m)
}

var xxx_messageInfo_PluginModule proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PluginModule)(nil), "slime.microservice.plugin.config.PluginModule")
}

func init() { proto.RegisterFile("plugin_module.proto", fileDescriptor_c98ebf9b1715be96) }

var fileDescriptor_c98ebf9b1715be96 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc8, 0x29, 0x4d,
	0xcf, 0xcc, 0x8b, 0xcf, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x28, 0xce, 0xc9, 0xcc, 0x4d, 0xd5, 0xcb, 0xcd, 0x4c, 0x2e, 0xca, 0x2f, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x83, 0xa8, 0xd3, 0x4b, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x57, 0xe2, 0xe3,
	0xe2, 0x09, 0x00, 0x0b, 0xf8, 0x82, 0xf5, 0x39, 0x69, 0x45, 0x69, 0x40, 0xf4, 0x64, 0xe6, 0xeb,
	0x83, 0x19, 0xfa, 0x10, 0x03, 0x8b, 0xf5, 0x21, 0xfa, 0xf4, 0x13, 0x0b, 0x32, 0xf5, 0x21, 0x7a,
	0x93, 0xd8, 0xc0, 0x96, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x03, 0x47, 0x12, 0xff, 0x7b,
	0x00, 0x00, 0x00,
}
