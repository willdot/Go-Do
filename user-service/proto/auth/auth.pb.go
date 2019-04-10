// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type PasswordChange struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	OldPassword          string   `protobuf:"bytes,2,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
	NewPassword          string   `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordChange) Reset()         { *m = PasswordChange{} }
func (m *PasswordChange) String() string { return proto.CompactTextString(m) }
func (*PasswordChange) ProtoMessage()    {}
func (*PasswordChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{4}
}

func (m *PasswordChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordChange.Unmarshal(m, b)
}
func (m *PasswordChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordChange.Marshal(b, m, deterministic)
}
func (m *PasswordChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordChange.Merge(m, src)
}
func (m *PasswordChange) XXX_Size() int {
	return xxx_messageInfo_PasswordChange.Size(m)
}
func (m *PasswordChange) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordChange.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordChange proto.InternalMessageInfo

func (m *PasswordChange) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *PasswordChange) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *PasswordChange) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{5}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*Request)(nil), "auth.Request")
	proto.RegisterType((*Response)(nil), "auth.Response")
	proto.RegisterType((*Token)(nil), "auth.Token")
	proto.RegisterType((*PasswordChange)(nil), "auth.PasswordChange")
	proto.RegisterType((*Error)(nil), "auth.Error")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x7d, 0xcd, 0xd7, 0xeb, 0xbb, 0xe5, 0x65, 0x31, 0x3c, 0x61, 0xe8, 0x42, 0x62, 0x0a, 0x45,
	0x11, 0x2a, 0x54, 0x5c, 0xba, 0x28, 0x45, 0xba, 0x95, 0x60, 0xc5, 0xed, 0xd8, 0xb9, 0xd8, 0x68,
	0x9a, 0x49, 0x67, 0x26, 0x2d, 0xfe, 0x28, 0xff, 0xa3, 0xcc, 0x9d, 0xa4, 0xc6, 0x22, 0xd4, 0x4d,
	0x72, 0xef, 0xb9, 0x27, 0xf7, 0xe3, 0x1c, 0x02, 0xcf, 0x1a, 0xad, 0xac, 0x7a, 0x23, 0x5a, 0xbb,
	0xa7, 0xc7, 0x82, 0x72, 0x16, 0xb9, 0x38, 0x3f, 0x41, 0xb4, 0x35, 0xa8, 0x59, 0x0a, 0x41, 0x29,
	0xf9, 0x28, 0x1b, 0xbd, 0x7c, 0x28, 0x82, 0x52, 0x32, 0x06, 0x51, 0x2d, 0x0e, 0xc8, 0x03, 0x42,
	0x28, 0x66, 0x1c, 0xee, 0x77, 0xea, 0xd0, 0x88, 0xfa, 0x27, 0x0f, 0x09, 0xee, 0x53, 0xf6, 0x04,
	0x31, 0x1e, 0x44, 0x59, 0xf1, 0x88, 0x70, 0x9f, 0xb0, 0x29, 0x8c, 0x1b, 0x61, 0xcc, 0x59, 0x69,
	0xc9, 0x63, 0x2a, 0x5c, 0xf2, 0xfc, 0x01, 0xee, 0x0b, 0x3c, 0xb6, 0x68, 0x6c, 0x7e, 0x84, 0x71,
	0x81, 0xa6, 0x51, 0xb5, 0x41, 0xf6, 0x1c, 0xa2, 0xd6, 0xa0, 0xa6, 0x45, 0x26, 0x4b, 0x58, 0xd0,
	0xbe, 0x6e, 0xc1, 0x82, 0x70, 0x96, 0x41, 0xec, 0xde, 0x86, 0x07, 0x59, 0x78, 0x45, 0xf0, 0x05,
	0x36, 0x83, 0x04, 0xb5, 0x56, 0xda, 0xf0, 0x90, 0x28, 0x13, 0x4f, 0xf9, 0xe0, 0xb0, 0xa2, 0x2b,
	0xe5, 0x5f, 0x20, 0xfe, 0xa4, 0x7e, 0x60, 0xed, 0x16, 0xb7, 0x2e, 0xe8, 0x2e, 0xf7, 0x89, 0x43,
	0x4f, 0xa2, 0x2a, 0x25, 0x5d, 0x3f, 0x2e, 0x7c, 0xf2, 0x7f, 0x9d, 0xbf, 0x43, 0xfa, 0xb1, 0xbb,
	0x71, 0xbd, 0x17, 0xf5, 0x37, 0xfc, 0xa3, 0xcd, 0x68, 0xa8, 0x4d, 0x06, 0x13, 0x55, 0xc9, 0x9e,
	0xda, 0xc9, 0x3c, 0x84, 0x1c, 0xa3, 0xc6, 0xf3, 0x85, 0xe1, 0x15, 0x1f, 0x42, 0xf9, 0x7b, 0x88,
	0x69, 0xb8, 0x33, 0x6b, 0xa7, 0x24, 0xd2, 0x84, 0xb8, 0xa0, 0xd8, 0x7d, 0x2e, 0xd1, 0xec, 0x74,
	0xd9, 0xd8, 0x52, 0xd5, 0xfd, 0x80, 0x01, 0xb4, 0xfc, 0x15, 0x40, 0xb4, 0x6a, 0xed, 0x9e, 0xcd,
	0x21, 0x59, 0x6b, 0x14, 0x16, 0xd9, 0x40, 0xcf, 0x69, 0xea, 0xe3, 0xde, 0x9a, 0xfc, 0x8e, 0xcd,
	0x20, 0xdc, 0xa0, 0xbd, 0x41, 0x7a, 0x05, 0xc9, 0x06, 0xed, 0xaa, 0xaa, 0xd8, 0x63, 0x5f, 0x23,
	0x9b, 0xff, 0x41, 0x7d, 0xd1, 0xcd, 0x1f, 0x36, 0xec, 0x44, 0x25, 0x77, 0xf2, 0x3b, 0xf6, 0x1a,
	0x1e, 0x3f, 0x3b, 0xf1, 0x85, 0x45, 0x6f, 0xd8, 0xb0, 0x7e, 0x4d, 0x9e, 0x43, 0xb2, 0x6d, 0xe4,
	0xed, 0x3b, 0xde, 0x41, 0xea, 0xbd, 0xb9, 0x68, 0xfd, 0xe4, 0x39, 0x7f, 0x3b, 0x77, 0xd5, 0xfe,
	0x6b, 0x42, 0xff, 0xcd, 0xdb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x65, 0xfe, 0xd9, 0x37, 0x50,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
	Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	ChangePassword(ctx context.Context, in *PasswordChange, opts ...client.CallOption) (*Token, error)
}

type authClient struct {
	c           client.Client
	serviceName string
}

func NewAuthClient(serviceName string, c client.Client) AuthClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ChangePassword(ctx context.Context, in *PasswordChange, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.ChangePassword", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
	Update(context.Context, *User, *Response) error
	ChangePassword(context.Context, *PasswordChange, *Token) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Auth{hdlr}, opts...))
}

type Auth struct {
	AuthHandler
}

func (h *Auth) Create(ctx context.Context, in *User, out *Response) error {
	return h.AuthHandler.Create(ctx, in, out)
}

func (h *Auth) Get(ctx context.Context, in *User, out *Response) error {
	return h.AuthHandler.Get(ctx, in, out)
}

func (h *Auth) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.AuthHandler.GetAll(ctx, in, out)
}

func (h *Auth) Auth(ctx context.Context, in *User, out *Token) error {
	return h.AuthHandler.Auth(ctx, in, out)
}

func (h *Auth) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.AuthHandler.ValidateToken(ctx, in, out)
}

func (h *Auth) Update(ctx context.Context, in *User, out *Response) error {
	return h.AuthHandler.Update(ctx, in, out)
}

func (h *Auth) ChangePassword(ctx context.Context, in *PasswordChange, out *Token) error {
	return h.AuthHandler.ChangePassword(ctx, in, out)
}
