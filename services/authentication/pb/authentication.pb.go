// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication.proto

package pb

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

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *LoginResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type SignUpRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	FirstName            string   `protobuf:"bytes,4,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,5,opt,name=LastName,proto3" json:"LastName,omitempty"`
	BirthDate            string   `protobuf:"bytes,6,opt,name=BirthDate,proto3" json:"BirthDate,omitempty"`
	City                 string   `protobuf:"bytes,7,opt,name=City,proto3" json:"City,omitempty"`
	Country              string   `protobuf:"bytes,8,opt,name=Country,proto3" json:"Country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{2}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *SignUpRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *SignUpRequest) GetBirthDate() string {
	if m != nil {
		return m.BirthDate
	}
	return ""
}

func (m *SignUpRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *SignUpRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type SignUpResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpResponse) Reset()         { *m = SignUpResponse{} }
func (m *SignUpResponse) String() string { return proto.CompactTextString(m) }
func (*SignUpResponse) ProtoMessage()    {}
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{3}
}

func (m *SignUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResponse.Unmarshal(m, b)
}
func (m *SignUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResponse.Marshal(b, m, deterministic)
}
func (m *SignUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResponse.Merge(m, src)
}
func (m *SignUpResponse) XXX_Size() int {
	return xxx_messageInfo_SignUpResponse.Size(m)
}
func (m *SignUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResponse proto.InternalMessageInfo

func (m *SignUpResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "loginRequest")
	proto.RegisterType((*LoginResponse)(nil), "loginResponse")
	proto.RegisterType((*SignUpRequest)(nil), "SignUpRequest")
	proto.RegisterType((*SignUpResponse)(nil), "SignUpResponse")
}

func init() {
	proto.RegisterFile("authentication.proto", fileDescriptor_d0dbc99083440df2)
}

var fileDescriptor_d0dbc99083440df2 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x0d, 0x42, 0xf9, 0x18, 0x01, 0x93, 0x0d, 0x87, 0x0d, 0xf1, 0x40, 0x7a, 0x30, 0xea, 0xa1,
	0x26, 0xfa, 0x07, 0x04, 0xd4, 0x13, 0x1a, 0x83, 0x72, 0xd0, 0xdb, 0x82, 0x63, 0xbb, 0x11, 0x76,
	0xeb, 0xce, 0x36, 0x86, 0xdf, 0xeb, 0x1f, 0x31, 0xdd, 0x76, 0x91, 0x7a, 0xf1, 0xd6, 0xf7, 0x5e,
	0xdf, 0xcb, 0x9b, 0x99, 0x85, 0x81, 0xc8, 0x6c, 0x82, 0xca, 0xca, 0x95, 0xb0, 0x52, 0xab, 0x28,
	0x35, 0xda, 0xea, 0xf0, 0x1a, 0xba, 0x6b, 0x1d, 0x4b, 0x35, 0xc7, 0xcf, 0x0c, 0xc9, 0xb2, 0x01,
	0x04, 0xb7, 0x1b, 0x21, 0xd7, 0xbc, 0x36, 0xaa, 0x9d, 0x76, 0xe6, 0x05, 0x60, 0x43, 0x68, 0x3f,
	0x0a, 0xa2, 0x2f, 0x6d, 0xde, 0xf8, 0x81, 0x13, 0x76, 0x38, 0x5c, 0x40, 0xaf, 0x4c, 0xa0, 0x54,
	0x2b, 0x42, 0x36, 0x82, 0xc3, 0xf1, 0x6a, 0x85, 0x44, 0xcf, 0xfa, 0x03, 0x55, 0x19, 0xb4, 0x4f,
	0xb1, 0x10, 0xba, 0x73, 0x7c, 0x37, 0x48, 0x49, 0xf1, 0x4b, 0x11, 0x59, 0xe1, 0xc2, 0xef, 0x1a,
	0xf4, 0x9e, 0x64, 0xac, 0x16, 0xe9, 0xbf, 0xd5, 0x16, 0x84, 0xe6, 0x41, 0x6c, 0xd0, 0x57, 0xf3,
	0xb8, 0x52, 0xbb, 0x5e, 0xad, 0xcd, 0x8e, 0xa1, 0x73, 0x27, 0x0d, 0x59, 0x67, 0x6c, 0x38, 0xf1,
	0x97, 0xc8, 0x9d, 0x33, 0x51, 0x8a, 0x41, 0xe1, 0xf4, 0x38, 0x77, 0x4e, 0xa4, 0xb1, 0xc9, 0x8d,
	0xb0, 0xc8, 0x9b, 0x85, 0x73, 0x47, 0x30, 0x06, 0x8d, 0xa9, 0xb4, 0x5b, 0xde, 0x72, 0x82, 0xfb,
	0x66, 0x1c, 0x5a, 0x53, 0x9d, 0x29, 0x6b, 0xb6, 0xbc, 0xed, 0x68, 0x0f, 0xc3, 0x73, 0xe8, 0xfb,
	0x21, 0xcb, 0xed, 0x71, 0x68, 0xdd, 0x23, 0x91, 0x88, 0xb1, 0x9c, 0xd3, 0xc3, 0xcb, 0x17, 0x68,
	0x8c, 0x33, 0x9b, 0xb0, 0x33, 0x68, 0xe6, 0x9e, 0x2c, 0x65, 0xfd, 0xa8, 0xb2, 0xa1, 0xe1, 0x51,
	0xf4, 0x27, 0xec, 0x04, 0x82, 0x59, 0x7e, 0x1b, 0xd6, 0x8b, 0xf6, 0xaf, 0x3c, 0xec, 0x47, 0x95,
	0x93, 0x4d, 0x82, 0xd7, 0xfa, 0x45, 0xba, 0x5c, 0x36, 0xdd, 0x9b, 0xb8, 0xfa, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x32, 0xad, 0xed, 0x2e, 0x2b, 0x02, 0x00, 0x00,
}
