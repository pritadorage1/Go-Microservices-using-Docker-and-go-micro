// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/employee/employee.proto

package employee

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
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone                int32      `protobuf:"varint,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Addresses            []*Address `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	EmailId              string     `protobuf:"bytes,5,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d3d908e03c98bb8, []int{0}
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

func (m *User) GetPhone() int32 {
	if m != nil {
		return m.Phone
	}
	return 0
}

func (m *User) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *User) GetEmailId() string {
	if m != nil {
		return m.EmailId
	}
	return ""
}

type Address struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AddressLine          string   `protobuf:"bytes,2,opt,name=address_line,json=addressLine,proto3" json:"address_line,omitempty"`
	Landmark             string   `protobuf:"bytes,3,opt,name=landmark,proto3" json:"landmark,omitempty"`
	Pincode              string   `protobuf:"bytes,4,opt,name=pincode,proto3" json:"pincode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d3d908e03c98bb8, []int{1}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Address) GetAddressLine() string {
	if m != nil {
		return m.AddressLine
	}
	return ""
}

func (m *Address) GetLandmark() string {
	if m != nil {
		return m.Landmark
	}
	return ""
}

func (m *Address) GetPincode() string {
	if m != nil {
		return m.Pincode
	}
	return ""
}

type CreateUserResponse struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d3d908e03c98bb8, []int{2}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// Created a blank get request
type GetUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersRequest) Reset()         { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()    {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d3d908e03c98bb8, []int{3}
}

func (m *GetUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersRequest.Unmarshal(m, b)
}
func (m *GetUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersRequest.Marshal(b, m, deterministic)
}
func (m *GetUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersRequest.Merge(m, src)
}
func (m *GetUsersRequest) XXX_Size() int {
	return xxx_messageInfo_GetUsersRequest.Size(m)
}
func (m *GetUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersRequest proto.InternalMessageInfo

type GetUsersResponse struct {
	Created bool  `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	User    *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// Added a pluralised consignment to our generic response message
	Users                []*User  `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersResponse) Reset()         { *m = GetUsersResponse{} }
func (m *GetUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetUsersResponse) ProtoMessage()    {}
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d3d908e03c98bb8, []int{4}
}

func (m *GetUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersResponse.Unmarshal(m, b)
}
func (m *GetUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersResponse.Marshal(b, m, deterministic)
}
func (m *GetUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersResponse.Merge(m, src)
}
func (m *GetUsersResponse) XXX_Size() int {
	return xxx_messageInfo_GetUsersResponse.Size(m)
}
func (m *GetUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersResponse proto.InternalMessageInfo

func (m *GetUsersResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *GetUsersResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "employee.User")
	proto.RegisterType((*Address)(nil), "employee.Address")
	proto.RegisterType((*CreateUserResponse)(nil), "employee.CreateUserResponse")
	proto.RegisterType((*GetUsersRequest)(nil), "employee.GetUsersRequest")
	proto.RegisterType((*GetUsersResponse)(nil), "employee.GetUsersResponse")
}

func init() { proto.RegisterFile("proto/employee/employee.proto", fileDescriptor_7d3d908e03c98bb8) }

var fileDescriptor_7d3d908e03c98bb8 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xfd, 0xd2, 0x26, 0x5f, 0xd3, 0x69, 0x69, 0xed, 0xe0, 0x61, 0x1b, 0x14, 0xea, 0xe2, 0xa1,
	0xa7, 0x16, 0xea, 0x1f, 0x50, 0x44, 0x8a, 0xe0, 0x69, 0xc5, 0x73, 0x89, 0xdd, 0x01, 0x17, 0x93,
	0x4d, 0xdc, 0x4d, 0x0b, 0xfe, 0x07, 0x7f, 0x82, 0x3f, 0x56, 0xb2, 0x49, 0x1a, 0x68, 0x3d, 0x7a,
	0xca, 0xbe, 0xf7, 0xe6, 0xcd, 0xbc, 0x19, 0x02, 0x97, 0xb9, 0xc9, 0x8a, 0x6c, 0x49, 0x69, 0x9e,
	0x64, 0x9f, 0x44, 0x87, 0xc7, 0xc2, 0xf1, 0x18, 0x36, 0x98, 0x7f, 0x79, 0xe0, 0xbf, 0x58, 0x32,
	0x38, 0x82, 0x8e, 0x92, 0xcc, 0x9b, 0x79, 0xf3, 0xbe, 0xe8, 0x28, 0x89, 0x08, 0xbe, 0x8e, 0x53,
	0x62, 0x1d, 0xc7, 0xb8, 0x37, 0x9e, 0x43, 0x90, 0xbf, 0x65, 0x9a, 0x58, 0x77, 0xe6, 0xcd, 0x03,
	0x51, 0x01, 0x5c, 0x42, 0x3f, 0x96, 0xd2, 0x90, 0xb5, 0x64, 0x99, 0x3f, 0xeb, 0xce, 0x07, 0xab,
	0xc9, 0xe2, 0x30, 0xf0, 0xae, 0x92, 0x44, 0x5b, 0x83, 0x53, 0x08, 0x29, 0x8d, 0x55, 0xb2, 0x51,
	0x92, 0x05, 0xae, 0x7d, 0xcf, 0xe1, 0x47, 0xc9, 0x0d, 0xf4, 0x6a, 0xc3, 0x49, 0xa0, 0x2b, 0x18,
	0xd6, 0x2d, 0x36, 0x89, 0xd2, 0x4d, 0xb0, 0x41, 0xcd, 0x3d, 0x29, 0x4d, 0x18, 0x41, 0x98, 0xc4,
	0x5a, 0xa6, 0xb1, 0x79, 0x77, 0x11, 0xfb, 0xe2, 0x80, 0x91, 0x41, 0x2f, 0x57, 0x7a, 0x9b, 0x49,
	0x62, 0x7e, 0x35, 0xb3, 0x86, 0x5c, 0x00, 0xde, 0x1b, 0x8a, 0x0b, 0x2a, 0xef, 0x20, 0xc8, 0xe6,
	0x99, 0xb6, 0x54, 0xd6, 0x6f, 0x1d, 0x5b, 0x65, 0x08, 0x45, 0x03, 0x91, 0x83, 0xbf, 0xb3, 0x64,
	0x5c, 0x80, 0xc1, 0x6a, 0xd4, 0xae, 0xea, 0xfc, 0x4e, 0xe3, 0x13, 0x18, 0xaf, 0xa9, 0x28, 0x09,
	0x2b, 0xe8, 0x63, 0x47, 0xb6, 0xe0, 0x7b, 0x38, 0x6b, 0xa9, 0xbf, 0x18, 0x82, 0xd7, 0x10, 0x94,
	0x5f, 0xcb, 0xba, 0xee, 0xe8, 0xc7, 0x45, 0x95, 0xb8, 0xfa, 0xf6, 0x60, 0xfc, 0x50, 0x0b, 0xcf,
	0x64, 0xf6, 0x6a, 0x4b, 0x78, 0x0b, 0xa3, 0x6a, 0xe5, 0x46, 0xc0, 0x23, 0x73, 0x74, 0xd1, 0xe2,
	0xd3, 0xe3, 0xf0, 0x7f, 0xb8, 0x86, 0xe1, 0x9a, 0x8a, 0xc6, 0x6e, 0x71, 0xda, 0xd6, 0x1f, 0x2d,
	0x1e, 0x45, 0xbf, 0x49, 0x4d, 0xa3, 0xd7, 0xff, 0xee, 0x8f, 0xbc, 0xf9, 0x09, 0x00, 0x00, 0xff,
	0xff, 0xcf, 0x14, 0xee, 0xd5, 0xb2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for EmployeeService service

type EmployeeServiceClient interface {
	CreateEmployee(ctx context.Context, in *User, opts ...client.CallOption) (*CreateUserResponse, error)
	// Created a new method
	GetEmployees(ctx context.Context, in *GetUsersRequest, opts ...client.CallOption) (*GetUsersResponse, error)
}

type employeeServiceClient struct {
	c           client.Client
	serviceName string
}

func NewEmployeeServiceClient(serviceName string, c client.Client) EmployeeServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "employee"
	}
	return &employeeServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *employeeServiceClient) CreateEmployee(ctx context.Context, in *User, opts ...client.CallOption) (*CreateUserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "EmployeeService.CreateEmployee", in)
	out := new(CreateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetEmployees(ctx context.Context, in *GetUsersRequest, opts ...client.CallOption) (*GetUsersResponse, error) {
	req := c.c.NewRequest(c.serviceName, "EmployeeService.GetEmployees", in)
	out := new(GetUsersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EmployeeService service

type EmployeeServiceHandler interface {
	CreateEmployee(context.Context, *User, *CreateUserResponse) error
	// Created a new method
	GetEmployees(context.Context, *GetUsersRequest, *GetUsersResponse) error
}

func RegisterEmployeeServiceHandler(s server.Server, hdlr EmployeeServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&EmployeeService{hdlr}, opts...))
}

type EmployeeService struct {
	EmployeeServiceHandler
}

func (h *EmployeeService) CreateEmployee(ctx context.Context, in *User, out *CreateUserResponse) error {
	return h.EmployeeServiceHandler.CreateEmployee(ctx, in, out)
}

func (h *EmployeeService) GetEmployees(ctx context.Context, in *GetUsersRequest, out *GetUsersResponse) error {
	return h.EmployeeServiceHandler.GetEmployees(ctx, in, out)
}
