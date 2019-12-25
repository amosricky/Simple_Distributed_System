// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetScoreRequest struct {
	GameName             string   `protobuf:"bytes,1,opt,name=gameName,proto3" json:"gameName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetScoreRequest) Reset()         { *m = GetScoreRequest{} }
func (m *GetScoreRequest) String() string { return proto.CompactTextString(m) }
func (*GetScoreRequest) ProtoMessage()    {}
func (*GetScoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

func (m *GetScoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetScoreRequest.Unmarshal(m, b)
}
func (m *GetScoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetScoreRequest.Marshal(b, m, deterministic)
}
func (m *GetScoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetScoreRequest.Merge(m, src)
}
func (m *GetScoreRequest) XXX_Size() int {
	return xxx_messageInfo_GetScoreRequest.Size(m)
}
func (m *GetScoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetScoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetScoreRequest proto.InternalMessageInfo

func (m *GetScoreRequest) GetGameName() string {
	if m != nil {
		return m.GameName
	}
	return ""
}

type GetScoreReply struct {
	Home                 int64    `protobuf:"varint,1,opt,name=Home,json=home,proto3" json:"Home,omitempty"`
	Visitor              int64    `protobuf:"varint,2,opt,name=Visitor,json=visitor,proto3" json:"Visitor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetScoreReply) Reset()         { *m = GetScoreReply{} }
func (m *GetScoreReply) String() string { return proto.CompactTextString(m) }
func (*GetScoreReply) ProtoMessage()    {}
func (*GetScoreReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}

func (m *GetScoreReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetScoreReply.Unmarshal(m, b)
}
func (m *GetScoreReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetScoreReply.Marshal(b, m, deterministic)
}
func (m *GetScoreReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetScoreReply.Merge(m, src)
}
func (m *GetScoreReply) XXX_Size() int {
	return xxx_messageInfo_GetScoreReply.Size(m)
}
func (m *GetScoreReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetScoreReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetScoreReply proto.InternalMessageInfo

func (m *GetScoreReply) GetHome() int64 {
	if m != nil {
		return m.Home
	}
	return 0
}

func (m *GetScoreReply) GetVisitor() int64 {
	if m != nil {
		return m.Visitor
	}
	return 0
}

func init() {
	proto.RegisterType((*GetScoreRequest)(nil), "pb.GetScoreRequest")
	proto.RegisterType((*GetScoreReply)(nil), "pb.GetScoreReply")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769) }

var fileDescriptor_38fc58335341d769 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4f, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe5, 0xe2, 0x77, 0x4f,
	0x2d, 0x09, 0x4e, 0xce, 0x2f, 0x4a, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2,
	0xe2, 0x00, 0x29, 0xf2, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3,
	0x95, 0x6c, 0xb9, 0x78, 0x11, 0xca, 0x0b, 0x72, 0x2a, 0x85, 0x84, 0xb8, 0x58, 0x3c, 0xf2, 0xa1,
	0x0a, 0x99, 0x83, 0x58, 0x32, 0xf2, 0x73, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0xc3, 0x32, 0x8b, 0x33,
	0x4b, 0xf2, 0x8b, 0x24, 0x98, 0xc0, 0xc2, 0xec, 0x65, 0x10, 0xae, 0x91, 0x03, 0x17, 0x07, 0x4c,
	0xbb, 0x90, 0x09, 0x12, 0x5b, 0x58, 0xaf, 0x20, 0x49, 0x0f, 0xcd, 0x1d, 0x52, 0x82, 0xa8, 0x82,
	0x05, 0x39, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xa7, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0x86, 0x93, 0x1c, 0xc8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GetScoreClient is the client API for GetScore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetScoreClient interface {
	GetScore(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreReply, error)
}

type getScoreClient struct {
	cc *grpc.ClientConn
}

func NewGetScoreClient(cc *grpc.ClientConn) GetScoreClient {
	return &getScoreClient{cc}
}

func (c *getScoreClient) GetScore(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreReply, error) {
	out := new(GetScoreReply)
	err := c.cc.Invoke(ctx, "/pb.GetScore/GetScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetScoreServer is the server API for GetScore service.
type GetScoreServer interface {
	GetScore(context.Context, *GetScoreRequest) (*GetScoreReply, error)
}

// UnimplementedGetScoreServer can be embedded to have forward compatible implementations.
type UnimplementedGetScoreServer struct {
}

func (*UnimplementedGetScoreServer) GetScore(ctx context.Context, req *GetScoreRequest) (*GetScoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScore not implemented")
}

func RegisterGetScoreServer(s *grpc.Server, srv GetScoreServer) {
	s.RegisterService(&_GetScore_serviceDesc, srv)
}

func _GetScore_GetScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetScoreServer).GetScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GetScore/GetScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetScoreServer).GetScore(ctx, req.(*GetScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetScore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GetScore",
	HandlerType: (*GetScoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScore",
			Handler:    _GetScore_GetScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game.proto",
}
