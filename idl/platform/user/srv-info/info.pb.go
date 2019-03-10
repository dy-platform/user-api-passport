// Code generated by protoc-gen-go. DO NOT EDIT.
// source: info.proto

package platform_user_srv_info

import (
	fmt "fmt"
	idl "github.com/dy-platform/user-api-passport/idl"
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

//获取用户信息请求参数
type GetUserInfoReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Fields               []string `protobuf:"bytes,2,rep,name=Fields,proto3" json:"Fields,omitempty"`
	AppId                string   `protobuf:"bytes,3,opt,name=AppId,proto3" json:"AppId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoReq) Reset()         { *m = GetUserInfoReq{} }
func (m *GetUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoReq) ProtoMessage()    {}
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{0}
}

func (m *GetUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoReq.Unmarshal(m, b)
}
func (m *GetUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoReq.Marshal(b, m, deterministic)
}
func (m *GetUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoReq.Merge(m, src)
}
func (m *GetUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoReq.Size(m)
}
func (m *GetUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoReq proto.InternalMessageInfo

func (m *GetUserInfoReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetUserInfoReq) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *GetUserInfoReq) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type GetUserInfoResp struct {
	BaseResp             *idl.Resp `protobuf:"bytes,1,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
	Data                 string    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetUserInfoResp) Reset()         { *m = GetUserInfoResp{} }
func (m *GetUserInfoResp) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoResp) ProtoMessage()    {}
func (*GetUserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{1}
}

func (m *GetUserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoResp.Unmarshal(m, b)
}
func (m *GetUserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoResp.Marshal(b, m, deterministic)
}
func (m *GetUserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoResp.Merge(m, src)
}
func (m *GetUserInfoResp) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoResp.Size(m)
}
func (m *GetUserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoResp proto.InternalMessageInfo

func (m *GetUserInfoResp) GetBaseResp() *idl.Resp {
	if m != nil {
		return m.BaseResp
	}
	return nil
}

func (m *GetUserInfoResp) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type CreateUserReq struct {
	UserId               int64      `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	NickName             string     `protobuf:"bytes,2,opt,name=NickName,proto3" json:"NickName,omitempty"`
	Gender               idl.Gender `protobuf:"varint,3,opt,name=Gender,proto3,enum=base.Gender" json:"Gender,omitempty"`
	AvatarUrl            string     `protobuf:"bytes,4,opt,name=AvatarUrl,proto3" json:"AvatarUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateUserReq) Reset()         { *m = CreateUserReq{} }
func (m *CreateUserReq) String() string { return proto.CompactTextString(m) }
func (*CreateUserReq) ProtoMessage()    {}
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{2}
}

func (m *CreateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReq.Unmarshal(m, b)
}
func (m *CreateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReq.Marshal(b, m, deterministic)
}
func (m *CreateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReq.Merge(m, src)
}
func (m *CreateUserReq) XXX_Size() int {
	return xxx_messageInfo_CreateUserReq.Size(m)
}
func (m *CreateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReq proto.InternalMessageInfo

func (m *CreateUserReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateUserReq) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *CreateUserReq) GetGender() idl.Gender {
	if m != nil {
		return m.Gender
	}
	return idl.Gender_Unkonw
}

func (m *CreateUserReq) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

type CreateUserResp struct {
	BaseResp             *idl.Resp `protobuf:"bytes,1,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateUserResp) Reset()         { *m = CreateUserResp{} }
func (m *CreateUserResp) String() string { return proto.CompactTextString(m) }
func (*CreateUserResp) ProtoMessage()    {}
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{3}
}

func (m *CreateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResp.Unmarshal(m, b)
}
func (m *CreateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResp.Marshal(b, m, deterministic)
}
func (m *CreateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResp.Merge(m, src)
}
func (m *CreateUserResp) XXX_Size() int {
	return xxx_messageInfo_CreateUserResp.Size(m)
}
func (m *CreateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResp proto.InternalMessageInfo

func (m *CreateUserResp) GetBaseResp() *idl.Resp {
	if m != nil {
		return m.BaseResp
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUserInfoReq)(nil), "platform.user.srv.info.GetUserInfoReq")
	proto.RegisterType((*GetUserInfoResp)(nil), "platform.user.srv.info.GetUserInfoResp")
	proto.RegisterType((*CreateUserReq)(nil), "platform.user.srv.info.CreateUserReq")
	proto.RegisterType((*CreateUserResp)(nil), "platform.user.srv.info.CreateUserResp")
}

func init() { proto.RegisterFile("info.proto", fileDescriptor_f140d5b28dddb141) }

var fileDescriptor_f140d5b28dddb141 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x25, 0x6d, 0xbf, 0xd2, 0x4e, 0x3f, 0x2b, 0x0c, 0x52, 0x42, 0xf0, 0x10, 0x82, 0xc6, 0x9c,
	0xf6, 0x10, 0x2f, 0x5e, 0xab, 0x60, 0xe9, 0xc1, 0x1e, 0x16, 0xea, 0x45, 0x10, 0xb6, 0x66, 0x02,
	0xc1, 0x34, 0x59, 0x77, 0xd7, 0xfe, 0x05, 0x7f, 0x9a, 0x7f, 0x4b, 0xb2, 0x89, 0x6d, 0x03, 0x2a,
	0xb9, 0xcd, 0x7b, 0x33, 0xfb, 0xf6, 0xed, 0x9b, 0x05, 0xc8, 0x8a, 0xb4, 0x64, 0x52, 0x95, 0xa6,
	0xc4, 0x99, 0xcc, 0x85, 0x49, 0x4b, 0xb5, 0x65, 0xef, 0x9a, 0x14, 0xd3, 0x6a, 0xc7, 0xaa, 0xae,
	0x07, 0x1b, 0xa1, 0xa9, 0x9e, 0x09, 0x1e, 0x61, 0xba, 0x20, 0xb3, 0xd6, 0xa4, 0x96, 0x45, 0x5a,
	0x72, 0x7a, 0xc3, 0x19, 0x0c, 0x2d, 0x4c, 0x5c, 0xc7, 0x77, 0xa2, 0x3e, 0x6f, 0x50, 0xc5, 0xdf,
	0x67, 0x94, 0x27, 0xda, 0xed, 0xf9, 0xfd, 0x68, 0xcc, 0x1b, 0x84, 0x67, 0xf0, 0x6f, 0x2e, 0xe5,
	0x32, 0x71, 0xfb, 0xbe, 0x13, 0x8d, 0x79, 0x0d, 0x82, 0x07, 0x38, 0x6d, 0xe9, 0x6a, 0x89, 0x21,
	0x8c, 0x6e, 0x85, 0xa6, 0xaa, 0xb6, 0xd2, 0x93, 0x18, 0x98, 0x75, 0x52, 0x31, 0x7c, 0xdf, 0x43,
	0x84, 0x41, 0x22, 0x8c, 0x70, 0x7b, 0x56, 0xcf, 0xd6, 0xc1, 0x87, 0x03, 0x27, 0x77, 0x8a, 0x84,
	0xa1, 0x4a, 0xf2, 0x2f, 0x9b, 0x1e, 0x8c, 0x56, 0xd9, 0xcb, 0xeb, 0x4a, 0x6c, 0xa9, 0x51, 0xd8,
	0x63, 0xbc, 0x80, 0xe1, 0x82, 0x8a, 0x84, 0x94, 0xf5, 0x3a, 0x8d, 0xff, 0xd7, 0xf7, 0xd7, 0x1c,
	0x6f, 0x7a, 0x78, 0x0e, 0xe3, 0xf9, 0x4e, 0x18, 0xa1, 0xd6, 0x2a, 0x77, 0x07, 0x56, 0xe2, 0x40,
	0x04, 0x37, 0x30, 0x3d, 0x36, 0xd2, 0xfd, 0x5d, 0xf1, 0xa7, 0x03, 0xa3, 0xef, 0x40, 0xf0, 0x19,
	0x26, 0x47, 0xf9, 0x60, 0xc8, 0x7e, 0xde, 0x15, 0x6b, 0x2f, 0xc7, 0xbb, 0xea, 0x34, 0xa7, 0x25,
	0x3e, 0x01, 0x1c, 0x6c, 0xe2, 0xe5, 0x6f, 0xc7, 0x5a, 0x99, 0x7a, 0x61, 0x97, 0x31, 0x2d, 0x37,
	0x43, 0xfb, 0x77, 0xae, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x47, 0x72, 0xce, 0x6d, 0x02,
	0x00, 0x00,
}
