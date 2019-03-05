// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snowflake.proto

package platform_id_srv_snowflake

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

type GetIDReq struct {
	Num                  int64    `protobuf:"varint,1,opt,name=Num,proto3" json:"Num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIDReq) Reset()         { *m = GetIDReq{} }
func (m *GetIDReq) String() string { return proto.CompactTextString(m) }
func (*GetIDReq) ProtoMessage()    {}
func (*GetIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa8bfc3cc8f3970, []int{0}
}

func (m *GetIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDReq.Unmarshal(m, b)
}
func (m *GetIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDReq.Marshal(b, m, deterministic)
}
func (m *GetIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDReq.Merge(m, src)
}
func (m *GetIDReq) XXX_Size() int {
	return xxx_messageInfo_GetIDReq.Size(m)
}
func (m *GetIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDReq proto.InternalMessageInfo

func (m *GetIDReq) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type GetIDResp struct {
	IDs                  []int64  `protobuf:"varint,1,rep,packed,name=IDs,proto3" json:"IDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIDResp) Reset()         { *m = GetIDResp{} }
func (m *GetIDResp) String() string { return proto.CompactTextString(m) }
func (*GetIDResp) ProtoMessage()    {}
func (*GetIDResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa8bfc3cc8f3970, []int{1}
}

func (m *GetIDResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDResp.Unmarshal(m, b)
}
func (m *GetIDResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDResp.Marshal(b, m, deterministic)
}
func (m *GetIDResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDResp.Merge(m, src)
}
func (m *GetIDResp) XXX_Size() int {
	return xxx_messageInfo_GetIDResp.Size(m)
}
func (m *GetIDResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDResp proto.InternalMessageInfo

func (m *GetIDResp) GetIDs() []int64 {
	if m != nil {
		return m.IDs
	}
	return nil
}

func init() {
	proto.RegisterType((*GetIDReq)(nil), "platform.id.srv.snowflake.GetIDReq")
	proto.RegisterType((*GetIDResp)(nil), "platform.id.srv.snowflake.GetIDResp")
}

func init() { proto.RegisterFile("snowflake.proto", fileDescriptor_aaa8bfc3cc8f3970) }

var fileDescriptor_aaa8bfc3cc8f3970 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xce, 0xcb, 0x2f,
	0x4f, 0xcb, 0x49, 0xcc, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2c, 0xc8, 0x49,
	0x2c, 0x49, 0xcb, 0x2f, 0xca, 0xd5, 0xcb, 0x4c, 0xd1, 0x2b, 0x2e, 0x2a, 0xd3, 0x83, 0x2b, 0x50,
	0x92, 0xe1, 0xe2, 0x70, 0x4f, 0x2d, 0xf1, 0x74, 0x09, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0xf6,
	0x2b, 0xcd, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0x31, 0x95, 0x64, 0xb9, 0x38, 0xa1,
	0xb2, 0xc5, 0x05, 0x20, 0x69, 0x4f, 0x97, 0x62, 0x09, 0x46, 0x05, 0x66, 0x90, 0xb4, 0xa7, 0x4b,
	0xb1, 0x51, 0x3c, 0x17, 0x67, 0x70, 0x5e, 0x7e, 0xb9, 0x1b, 0xc8, 0x24, 0xa1, 0x20, 0x2e, 0x56,
	0xb0, 0x5a, 0x21, 0x65, 0x3d, 0x9c, 0xd6, 0xe9, 0xc1, 0xec, 0x92, 0x52, 0x21, 0xac, 0xa8, 0xb8,
	0x20, 0x89, 0x0d, 0xec, 0x7e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x32, 0xbc, 0xba,
	0xd2, 0x00, 0x00, 0x00,
}
