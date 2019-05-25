// Code generated by protoc-gen-go. DO NOT EDIT.
// source: idl.proto

package idl

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Gender int32

const (
	Gender_MALE   Gender = 0
	Gender_FEMALE Gender = 1
)

var Gender_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
}

var Gender_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0730a0330a34daf1, []int{0}
}

type ResponseStatus struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseStatus) Reset()         { *m = ResponseStatus{} }
func (m *ResponseStatus) String() string { return proto.CompactTextString(m) }
func (*ResponseStatus) ProtoMessage()    {}
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_0730a0330a34daf1, []int{0}
}

func (m *ResponseStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseStatus.Unmarshal(m, b)
}
func (m *ResponseStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseStatus.Marshal(b, m, deterministic)
}
func (m *ResponseStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseStatus.Merge(m, src)
}
func (m *ResponseStatus) XXX_Size() int {
	return xxx_messageInfo_ResponseStatus.Size(m)
}
func (m *ResponseStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseStatus proto.InternalMessageInfo

func (m *ResponseStatus) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type UserBrief struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserBrief) Reset()         { *m = UserBrief{} }
func (m *UserBrief) String() string { return proto.CompactTextString(m) }
func (*UserBrief) ProtoMessage()    {}
func (*UserBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_0730a0330a34daf1, []int{1}
}

func (m *UserBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBrief.Unmarshal(m, b)
}
func (m *UserBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBrief.Marshal(b, m, deterministic)
}
func (m *UserBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBrief.Merge(m, src)
}
func (m *UserBrief) XXX_Size() int {
	return xxx_messageInfo_UserBrief.Size(m)
}
func (m *UserBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBrief.DiscardUnknown(m)
}

var xxx_messageInfo_UserBrief proto.InternalMessageInfo

func (m *UserBrief) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FriendPayload struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FriendId             int32    `protobuf:"varint,2,opt,name=friendId,proto3" json:"friendId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendPayload) Reset()         { *m = FriendPayload{} }
func (m *FriendPayload) String() string { return proto.CompactTextString(m) }
func (*FriendPayload) ProtoMessage()    {}
func (*FriendPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_0730a0330a34daf1, []int{2}
}

func (m *FriendPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendPayload.Unmarshal(m, b)
}
func (m *FriendPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendPayload.Marshal(b, m, deterministic)
}
func (m *FriendPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendPayload.Merge(m, src)
}
func (m *FriendPayload) XXX_Size() int {
	return xxx_messageInfo_FriendPayload.Size(m)
}
func (m *FriendPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendPayload.DiscardUnknown(m)
}

var xxx_messageInfo_FriendPayload proto.InternalMessageInfo

func (m *FriendPayload) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *FriendPayload) GetFriendId() int32 {
	if m != nil {
		return m.FriendId
	}
	return 0
}

type UserInfo struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender               Gender   `protobuf:"varint,3,opt,name=gender,proto3,enum=idl.Gender" json:"gender,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0730a0330a34daf1, []int{3}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_MALE
}

func (m *UserInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type UserDetail struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender               Gender   `protobuf:"varint,3,opt,name=gender,proto3,enum=idl.Gender" json:"gender,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Friends              []int32  `protobuf:"varint,5,rep,packed,name=friends,proto3" json:"friends,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDetail) Reset()         { *m = UserDetail{} }
func (m *UserDetail) String() string { return proto.CompactTextString(m) }
func (*UserDetail) ProtoMessage()    {}
func (*UserDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_0730a0330a34daf1, []int{4}
}

func (m *UserDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetail.Unmarshal(m, b)
}
func (m *UserDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetail.Marshal(b, m, deterministic)
}
func (m *UserDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetail.Merge(m, src)
}
func (m *UserDetail) XXX_Size() int {
	return xxx_messageInfo_UserDetail.Size(m)
}
func (m *UserDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetail.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetail proto.InternalMessageInfo

func (m *UserDetail) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserDetail) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_MALE
}

func (m *UserDetail) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *UserDetail) GetFriends() []int32 {
	if m != nil {
		return m.Friends
	}
	return nil
}

func init() {
	proto.RegisterEnum("idl.Gender", Gender_name, Gender_value)
	proto.RegisterType((*ResponseStatus)(nil), "idl.ResponseStatus")
	proto.RegisterType((*UserBrief)(nil), "idl.UserBrief")
	proto.RegisterType((*FriendPayload)(nil), "idl.FriendPayload")
	proto.RegisterType((*UserInfo)(nil), "idl.UserInfo")
	proto.RegisterType((*UserDetail)(nil), "idl.UserDetail")
}

func init() { proto.RegisterFile("idl.proto", fileDescriptor_0730a0330a34daf1) }

var fileDescriptor_0730a0330a34daf1 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x4d, 0x6b, 0xea, 0x40,
	0x14, 0xcd, 0x44, 0x8d, 0xf1, 0x8a, 0x79, 0x8f, 0xfb, 0xe0, 0x11, 0xf2, 0xe0, 0x11, 0xa6, 0x9b,
	0xe0, 0x42, 0xd4, 0xae, 0xba, 0xb4, 0x56, 0x45, 0xb0, 0x20, 0x91, 0x2e, 0xbb, 0x48, 0x3b, 0xd7,
	0x32, 0x98, 0x26, 0x32, 0x13, 0x0b, 0xfd, 0x09, 0xfd, 0x8b, 0xfd, 0x35, 0x25, 0x13, 0x23, 0xe8,
	0xaa, 0x9b, 0xee, 0xe6, 0xcc, 0x39, 0xf7, 0xf3, 0x5c, 0xe8, 0x48, 0x91, 0x0e, 0xf6, 0x2a, 0x2f,
	0x72, 0x6c, 0x48, 0x91, 0xf2, 0x10, 0xbc, 0x98, 0xf4, 0x3e, 0xcf, 0x34, 0x6d, 0x8a, 0xa4, 0x38,
	0x68, 0xf4, 0xc0, 0xce, 0x77, 0x3e, 0x0b, 0x59, 0xe4, 0xc6, 0x76, 0xbe, 0xe3, 0xff, 0xa0, 0xf3,
	0xa0, 0x49, 0xdd, 0x2a, 0x49, 0xdb, 0x92, 0x94, 0xc2, 0x90, 0xad, 0xd8, 0x96, 0x82, 0x4f, 0xa1,
	0x37, 0x57, 0x92, 0x32, 0xb1, 0x4e, 0xde, 0xd3, 0x3c, 0x11, 0xf8, 0x17, 0x9c, 0x83, 0x26, 0xb5,
	0xac, 0x45, 0x47, 0x84, 0x01, 0xb8, 0x5b, 0x23, 0x5c, 0x0a, 0xdf, 0x36, 0xcc, 0x09, 0xf3, 0x47,
	0x70, 0xcb, 0x0a, 0xcb, 0x6c, 0x9b, 0x23, 0x42, 0x33, 0x4b, 0x5e, 0xc9, 0x68, 0x3a, 0xb1, 0x79,
	0xe3, 0x15, 0x38, 0x2f, 0x94, 0x09, 0x52, 0x7e, 0x23, 0x64, 0x91, 0x37, 0xee, 0x0e, 0xca, 0x21,
	0x16, 0xe6, 0x2b, 0x3e, 0x52, 0xe8, 0x43, 0x3b, 0x11, 0x42, 0x91, 0xd6, 0x7e, 0xd3, 0xc4, 0xd6,
	0x90, 0x7f, 0x30, 0x80, 0x32, 0xff, 0x1d, 0x15, 0x89, 0x4c, 0x2f, 0x47, 0xf8, 0x81, 0x8a, 0x25,
	0x53, 0x0d, 0xa7, 0xfd, 0x56, 0xd8, 0x88, 0x5a, 0x71, 0x0d, 0xfb, 0xff, 0xc1, 0xa9, 0xb2, 0xa0,
	0x0b, 0xcd, 0xfb, 0xc9, 0x6a, 0xf6, 0xdb, 0x42, 0x00, 0x67, 0x3e, 0x33, 0x6f, 0x36, 0xfe, 0x64,
	0xd0, 0x5d, 0xc4, 0xeb, 0xe9, 0x86, 0xd4, 0x9b, 0x7c, 0x26, 0xec, 0x43, 0x7b, 0x22, 0x44, 0xd9,
	0x3d, 0xf6, 0x4c, 0x0f, 0xf5, 0xa2, 0x02, 0xef, 0x04, 0x8d, 0x33, 0xdc, 0xc2, 0x11, 0xc0, 0x82,
	0x8a, 0xca, 0x0e, 0x8d, 0x17, 0x7c, 0xf0, 0xeb, 0x84, 0xab, 0x3d, 0x70, 0x6b, 0xc8, 0xf0, 0x06,
	0x60, 0x22, 0x44, 0x1d, 0x82, 0x46, 0x72, 0xe6, 0x67, 0xf0, 0xc7, 0xfc, 0x9d, 0x9f, 0x08, 0xb7,
	0x22, 0x86, 0xa3, 0xca, 0xb4, 0x95, 0xd4, 0xc5, 0x37, 0x6a, 0x45, 0x6c, 0xc8, 0x9e, 0x1c, 0x73,
	0x77, 0xd7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0xdb, 0x4a, 0xb6, 0x84, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GRPCServiceClient is the client API for GRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GRPCServiceClient interface {
	AddUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserBrief, error)
	GetFriends(ctx context.Context, in *UserBrief, opts ...grpc.CallOption) (GRPCService_GetFriendsClient, error)
	AddFriends(ctx context.Context, opts ...grpc.CallOption) (GRPCService_AddFriendsClient, error)
	UserList(ctx context.Context, opts ...grpc.CallOption) (GRPCService_UserListClient, error)
}

type gRPCServiceClient struct {
	cc *grpc.ClientConn
}

func NewGRPCServiceClient(cc *grpc.ClientConn) GRPCServiceClient {
	return &gRPCServiceClient{cc}
}

func (c *gRPCServiceClient) AddUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserBrief, error) {
	out := new(UserBrief)
	err := c.cc.Invoke(ctx, "/idl.GRPCService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCServiceClient) GetFriends(ctx context.Context, in *UserBrief, opts ...grpc.CallOption) (GRPCService_GetFriendsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GRPCService_serviceDesc.Streams[0], "/idl.GRPCService/GetFriends", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCServiceGetFriendsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GRPCService_GetFriendsClient interface {
	Recv() (*UserDetail, error)
	grpc.ClientStream
}

type gRPCServiceGetFriendsClient struct {
	grpc.ClientStream
}

func (x *gRPCServiceGetFriendsClient) Recv() (*UserDetail, error) {
	m := new(UserDetail)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gRPCServiceClient) AddFriends(ctx context.Context, opts ...grpc.CallOption) (GRPCService_AddFriendsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GRPCService_serviceDesc.Streams[1], "/idl.GRPCService/AddFriends", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCServiceAddFriendsClient{stream}
	return x, nil
}

type GRPCService_AddFriendsClient interface {
	Send(*FriendPayload) error
	CloseAndRecv() (*ResponseStatus, error)
	grpc.ClientStream
}

type gRPCServiceAddFriendsClient struct {
	grpc.ClientStream
}

func (x *gRPCServiceAddFriendsClient) Send(m *FriendPayload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCServiceAddFriendsClient) CloseAndRecv() (*ResponseStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResponseStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gRPCServiceClient) UserList(ctx context.Context, opts ...grpc.CallOption) (GRPCService_UserListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GRPCService_serviceDesc.Streams[2], "/idl.GRPCService/UserList", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCServiceUserListClient{stream}
	return x, nil
}

type GRPCService_UserListClient interface {
	Send(*UserBrief) error
	Recv() (*UserDetail, error)
	grpc.ClientStream
}

type gRPCServiceUserListClient struct {
	grpc.ClientStream
}

func (x *gRPCServiceUserListClient) Send(m *UserBrief) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCServiceUserListClient) Recv() (*UserDetail, error) {
	m := new(UserDetail)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GRPCServiceServer is the server API for GRPCService service.
type GRPCServiceServer interface {
	AddUser(context.Context, *UserInfo) (*UserBrief, error)
	GetFriends(*UserBrief, GRPCService_GetFriendsServer) error
	AddFriends(GRPCService_AddFriendsServer) error
	UserList(GRPCService_UserListServer) error
}

func RegisterGRPCServiceServer(s *grpc.Server, srv GRPCServiceServer) {
	s.RegisterService(&_GRPCService_serviceDesc, srv)
}

func _GRPCService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.GRPCService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServiceServer).AddUser(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCService_GetFriends_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserBrief)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GRPCServiceServer).GetFriends(m, &gRPCServiceGetFriendsServer{stream})
}

type GRPCService_GetFriendsServer interface {
	Send(*UserDetail) error
	grpc.ServerStream
}

type gRPCServiceGetFriendsServer struct {
	grpc.ServerStream
}

func (x *gRPCServiceGetFriendsServer) Send(m *UserDetail) error {
	return x.ServerStream.SendMsg(m)
}

func _GRPCService_AddFriends_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCServiceServer).AddFriends(&gRPCServiceAddFriendsServer{stream})
}

type GRPCService_AddFriendsServer interface {
	SendAndClose(*ResponseStatus) error
	Recv() (*FriendPayload, error)
	grpc.ServerStream
}

type gRPCServiceAddFriendsServer struct {
	grpc.ServerStream
}

func (x *gRPCServiceAddFriendsServer) SendAndClose(m *ResponseStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gRPCServiceAddFriendsServer) Recv() (*FriendPayload, error) {
	m := new(FriendPayload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GRPCService_UserList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCServiceServer).UserList(&gRPCServiceUserListServer{stream})
}

type GRPCService_UserListServer interface {
	Send(*UserDetail) error
	Recv() (*UserBrief, error)
	grpc.ServerStream
}

type gRPCServiceUserListServer struct {
	grpc.ServerStream
}

func (x *gRPCServiceUserListServer) Send(m *UserDetail) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gRPCServiceUserListServer) Recv() (*UserBrief, error) {
	m := new(UserBrief)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.GRPCService",
	HandlerType: (*GRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _GRPCService_AddUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFriends",
			Handler:       _GRPCService_GetFriends_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddFriends",
			Handler:       _GRPCService_AddFriends_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UserList",
			Handler:       _GRPCService_UserList_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "idl.proto",
}
