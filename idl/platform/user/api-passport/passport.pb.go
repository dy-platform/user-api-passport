// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passport.proto

package platform_user_api_passport

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

type SignUpReq struct {
	DeviceID             int64    `protobuf:"varint,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	MsgCode              string   `protobuf:"bytes,6,opt,name=MsgCode,proto3" json:"MsgCode,omitempty"`
	ActiveType           int32    `protobuf:"varint,7,opt,name=ActiveType,proto3" json:"ActiveType,omitempty"`
	AppID                int32    `protobuf:"varint,8,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Version              string   `protobuf:"bytes,9,opt,name=Version,proto3" json:"Version,omitempty"`
	ExtraInfo            string   `protobuf:"bytes,10,opt,name=ExtraInfo,proto3" json:"ExtraInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpReq) Reset()         { *m = SignUpReq{} }
func (m *SignUpReq) String() string { return proto.CompactTextString(m) }
func (*SignUpReq) ProtoMessage()    {}
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4affa6d033a78188, []int{0}
}

func (m *SignUpReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpReq.Unmarshal(m, b)
}
func (m *SignUpReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpReq.Marshal(b, m, deterministic)
}
func (m *SignUpReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpReq.Merge(m, src)
}
func (m *SignUpReq) XXX_Size() int {
	return xxx_messageInfo_SignUpReq.Size(m)
}
func (m *SignUpReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpReq proto.InternalMessageInfo

func (m *SignUpReq) GetDeviceID() int64 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

func (m *SignUpReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignUpReq) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *SignUpReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpReq) GetMsgCode() string {
	if m != nil {
		return m.MsgCode
	}
	return ""
}

func (m *SignUpReq) GetActiveType() int32 {
	if m != nil {
		return m.ActiveType
	}
	return 0
}

func (m *SignUpReq) GetAppID() int32 {
	if m != nil {
		return m.AppID
	}
	return 0
}

func (m *SignUpReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SignUpReq) GetExtraInfo() string {
	if m != nil {
		return m.ExtraInfo
	}
	return ""
}

type SignUpResp struct {
	UserID               int64     `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Mobile               string    `protobuf:"bytes,2,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	BaseResp             *idl.Resp `protobuf:"bytes,3,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SignUpResp) Reset()         { *m = SignUpResp{} }
func (m *SignUpResp) String() string { return proto.CompactTextString(m) }
func (*SignUpResp) ProtoMessage()    {}
func (*SignUpResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4affa6d033a78188, []int{1}
}

func (m *SignUpResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResp.Unmarshal(m, b)
}
func (m *SignUpResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResp.Marshal(b, m, deterministic)
}
func (m *SignUpResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResp.Merge(m, src)
}
func (m *SignUpResp) XXX_Size() int {
	return xxx_messageInfo_SignUpResp.Size(m)
}
func (m *SignUpResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResp proto.InternalMessageInfo

func (m *SignUpResp) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *SignUpResp) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *SignUpResp) GetBaseResp() *idl.Resp {
	if m != nil {
		return m.BaseResp
	}
	return nil
}

type WeChatSignInReq struct {
	DeviceID             int64    `protobuf:"varint,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	AppID                string   `protobuf:"bytes,3,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Secret               string   `protobuf:"bytes,4,opt,name=Secret,proto3" json:"Secret,omitempty"`
	ExtraInfo            string   `protobuf:"bytes,5,opt,name=ExtraInfo,proto3" json:"ExtraInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeChatSignInReq) Reset()         { *m = WeChatSignInReq{} }
func (m *WeChatSignInReq) String() string { return proto.CompactTextString(m) }
func (*WeChatSignInReq) ProtoMessage()    {}
func (*WeChatSignInReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4affa6d033a78188, []int{2}
}

func (m *WeChatSignInReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeChatSignInReq.Unmarshal(m, b)
}
func (m *WeChatSignInReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeChatSignInReq.Marshal(b, m, deterministic)
}
func (m *WeChatSignInReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeChatSignInReq.Merge(m, src)
}
func (m *WeChatSignInReq) XXX_Size() int {
	return xxx_messageInfo_WeChatSignInReq.Size(m)
}
func (m *WeChatSignInReq) XXX_DiscardUnknown() {
	xxx_messageInfo_WeChatSignInReq.DiscardUnknown(m)
}

var xxx_messageInfo_WeChatSignInReq proto.InternalMessageInfo

func (m *WeChatSignInReq) GetDeviceID() int64 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

func (m *WeChatSignInReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *WeChatSignInReq) GetAppID() string {
	if m != nil {
		return m.AppID
	}
	return ""
}

func (m *WeChatSignInReq) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *WeChatSignInReq) GetExtraInfo() string {
	if m != nil {
		return m.ExtraInfo
	}
	return ""
}

type WeChatSignInResp struct {
	UserID               int64     `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	BaseResp             *idl.Resp `protobuf:"bytes,2,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *WeChatSignInResp) Reset()         { *m = WeChatSignInResp{} }
func (m *WeChatSignInResp) String() string { return proto.CompactTextString(m) }
func (*WeChatSignInResp) ProtoMessage()    {}
func (*WeChatSignInResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4affa6d033a78188, []int{3}
}

func (m *WeChatSignInResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeChatSignInResp.Unmarshal(m, b)
}
func (m *WeChatSignInResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeChatSignInResp.Marshal(b, m, deterministic)
}
func (m *WeChatSignInResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeChatSignInResp.Merge(m, src)
}
func (m *WeChatSignInResp) XXX_Size() int {
	return xxx_messageInfo_WeChatSignInResp.Size(m)
}
func (m *WeChatSignInResp) XXX_DiscardUnknown() {
	xxx_messageInfo_WeChatSignInResp.DiscardUnknown(m)
}

var xxx_messageInfo_WeChatSignInResp proto.InternalMessageInfo

func (m *WeChatSignInResp) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *WeChatSignInResp) GetBaseResp() *idl.Resp {
	if m != nil {
		return m.BaseResp
	}
	return nil
}

func init() {
	proto.RegisterType((*SignUpReq)(nil), "platform.user.api.passport.SignUpReq")
	proto.RegisterType((*SignUpResp)(nil), "platform.user.api.passport.SignUpResp")
	proto.RegisterType((*WeChatSignInReq)(nil), "platform.user.api.passport.WeChatSignInReq")
	proto.RegisterType((*WeChatSignInResp)(nil), "platform.user.api.passport.WeChatSignInResp")
}

func init() { proto.RegisterFile("passport.proto", fileDescriptor_4affa6d033a78188) }

var fileDescriptor_4affa6d033a78188 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x8a, 0xda, 0x40,
	0x14, 0x86, 0x49, 0x34, 0xd1, 0x9c, 0x96, 0xb6, 0x0c, 0x45, 0x86, 0x50, 0x8a, 0x08, 0x15, 0xa1,
	0x25, 0x17, 0xf6, 0x09, 0xac, 0x7a, 0x91, 0x0b, 0x4b, 0x89, 0x75, 0xbd, 0x1e, 0xf5, 0xe8, 0x0e,
	0x68, 0x66, 0x76, 0x26, 0xeb, 0xee, 0x3e, 0x84, 0xaf, 0xb6, 0xcf, 0xb4, 0xcc, 0x24, 0x66, 0xa3,
	0xa0, 0x78, 0xe7, 0x7f, 0xfe, 0x33, 0x67, 0xe6, 0xff, 0x3c, 0x81, 0x4f, 0x92, 0x69, 0x2d, 0x85,
	0xca, 0x22, 0xa9, 0x44, 0x26, 0x48, 0x28, 0xb7, 0x2c, 0x5b, 0x0b, 0xb5, 0x8b, 0x1e, 0x35, 0xaa,
	0x88, 0x49, 0x1e, 0x1d, 0x3b, 0x42, 0x58, 0x30, 0x8d, 0x79, 0x5f, 0xe7, 0xe0, 0x42, 0x30, 0xe5,
	0x9b, 0x74, 0x26, 0x13, 0x7c, 0x20, 0x21, 0x34, 0x47, 0xb8, 0xe7, 0x4b, 0x8c, 0x47, 0xd4, 0x69,
	0x3b, 0xbd, 0x5a, 0x52, 0x6a, 0x42, 0xa0, 0xfe, 0x97, 0xed, 0x90, 0xba, 0x6d, 0xa7, 0x17, 0x24,
	0xf6, 0x37, 0x69, 0x81, 0x3f, 0x11, 0x0b, 0xbe, 0x45, 0x5a, 0xb3, 0xd5, 0x42, 0x91, 0xaf, 0xe0,
	0x8d, 0x77, 0x8c, 0x6f, 0x69, 0xdd, 0x96, 0x73, 0x61, 0xa6, 0xff, 0x63, 0x5a, 0x3f, 0x09, 0xb5,
	0xa2, 0x9e, 0x35, 0x4a, 0x4d, 0x28, 0x34, 0x26, 0x7a, 0x33, 0x14, 0x2b, 0xa4, 0xbe, 0xb5, 0x8e,
	0x92, 0x7c, 0x07, 0x18, 0x2c, 0x33, 0xbe, 0xc7, 0xff, 0x2f, 0x12, 0x69, 0xa3, 0xed, 0xf4, 0xbc,
	0xa4, 0x52, 0x31, 0x77, 0x0d, 0xa4, 0x8c, 0x47, 0xb4, 0x69, 0xad, 0x5c, 0x98, 0x79, 0x77, 0xa8,
	0x34, 0x17, 0x29, 0x0d, 0xf2, 0x79, 0x85, 0x24, 0xdf, 0x20, 0x18, 0x3f, 0x67, 0x8a, 0xc5, 0xe9,
	0x5a, 0x50, 0xb0, 0xde, 0x7b, 0xa1, 0xb3, 0x02, 0x38, 0xe2, 0xd0, 0xd2, 0xe4, 0x9b, 0x69, 0x54,
	0x25, 0x8d, 0x42, 0x55, 0x72, 0xbb, 0x27, 0xb9, 0xbb, 0xd0, 0xfc, 0xc3, 0x34, 0x9a, 0xb3, 0x96,
	0xc8, 0x87, 0x3e, 0x44, 0x16, 0xb6, 0xa9, 0x24, 0xa5, 0xd7, 0x39, 0x38, 0xf0, 0x79, 0x8e, 0xc3,
	0x7b, 0x96, 0x99, 0xcb, 0xe2, 0xf4, 0x06, 0xf6, 0x16, 0x4d, 0xc1, 0xde, 0x72, 0x29, 0x73, 0xe7,
	0xe8, 0x8b, 0xdc, 0x2d, 0xf0, 0xa7, 0xb8, 0x54, 0x98, 0x15, 0xe8, 0x0b, 0x75, 0x9a, 0xda, 0x3b,
	0x4f, 0x9d, 0xc0, 0x97, 0xd3, 0xe7, 0x5c, 0xc9, 0x5e, 0xcd, 0xe8, 0x5e, 0xce, 0xd8, 0x7f, 0x75,
	0xf2, 0xbf, 0xdb, 0xac, 0x1c, 0x99, 0x83, 0x9f, 0x63, 0x25, 0x3f, 0xa2, 0xcb, 0x9b, 0x19, 0x95,
	0x9b, 0x18, 0x76, 0x6f, 0x69, 0xd3, 0x92, 0x70, 0xf8, 0x58, 0x7d, 0x39, 0xf9, 0x79, 0xed, 0xdc,
	0x19, 0xf2, 0xf0, 0xd7, 0xed, 0xcd, 0x5a, 0x2e, 0x7c, 0xfb, 0xc5, 0xfc, 0x7e, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0xd9, 0x69, 0x2d, 0x40, 0x6b, 0x03, 0x00, 0x00,
}
