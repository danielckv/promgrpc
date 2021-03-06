// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v4/pb/private/test/test.proto

package test // import "github.com/piotrkowalczuk/promgrpc/v4/pb/private/test"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_8175f87157c5e8bc, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Response struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_8175f87157c5e8bc, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "piotrkowalczuk.promgrpc.v4.test.Request")
	proto.RegisterType((*Response)(nil), "piotrkowalczuk.promgrpc.v4.test.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	Unary(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ServerSide(ctx context.Context, in *Request, opts ...grpc.CallOption) (TestService_ServerSideClient, error)
	ClientSide(ctx context.Context, opts ...grpc.CallOption) (TestService_ClientSideClient, error)
	Bidirectional(ctx context.Context, opts ...grpc.CallOption) (TestService_BidirectionalClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Unary(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/piotrkowalczuk.promgrpc.v4.test.TestService/Unary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ServerSide(ctx context.Context, in *Request, opts ...grpc.CallOption) (TestService_ServerSideClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/piotrkowalczuk.promgrpc.v4.test.TestService/ServerSide", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceServerSideClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_ServerSideClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type testServiceServerSideClient struct {
	grpc.ClientStream
}

func (x *testServiceServerSideClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) ClientSide(ctx context.Context, opts ...grpc.CallOption) (TestService_ClientSideClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[1], "/piotrkowalczuk.promgrpc.v4.test.TestService/ClientSide", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceClientSideClient{stream}
	return x, nil
}

type TestService_ClientSideClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type testServiceClientSideClient struct {
	grpc.ClientStream
}

func (x *testServiceClientSideClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceClientSideClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) Bidirectional(ctx context.Context, opts ...grpc.CallOption) (TestService_BidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[2], "/piotrkowalczuk.promgrpc.v4.test.TestService/Bidirectional", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceBidirectionalClient{stream}
	return x, nil
}

type TestService_BidirectionalClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type testServiceBidirectionalClient struct {
	grpc.ClientStream
}

func (x *testServiceBidirectionalClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceBidirectionalClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	Unary(context.Context, *Request) (*Response, error)
	ServerSide(*Request, TestService_ServerSideServer) error
	ClientSide(TestService_ClientSideServer) error
	Bidirectional(TestService_BidirectionalServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_Unary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Unary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/piotrkowalczuk.promgrpc.v4.test.TestService/Unary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Unary(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ServerSide_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).ServerSide(m, &testServiceServerSideServer{stream})
}

type TestService_ServerSideServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type testServiceServerSideServer struct {
	grpc.ServerStream
}

func (x *testServiceServerSideServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_ClientSide_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).ClientSide(&testServiceClientSideServer{stream})
}

type TestService_ClientSideServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type testServiceClientSideServer struct {
	grpc.ServerStream
}

func (x *testServiceClientSideServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceClientSideServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_Bidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).Bidirectional(&testServiceBidirectionalServer{stream})
}

type TestService_BidirectionalServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type testServiceBidirectionalServer struct {
	grpc.ServerStream
}

func (x *testServiceBidirectionalServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceBidirectionalServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "piotrkowalczuk.promgrpc.v4.test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unary",
			Handler:    _TestService_Unary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerSide",
			Handler:       _TestService_ServerSide_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientSide",
			Handler:       _TestService_ClientSide_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Bidirectional",
			Handler:       _TestService_Bidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v4/pb/private/test/test.proto",
}

func init() { proto.RegisterFile("v4/pb/private/test/test.proto", fileDescriptor_test_8175f87157c5e8bc) }

var fileDescriptor_test_8175f87157c5e8bc = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0xd2, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x70, 0x82, 0xd4, 0x3f, 0x23, 0x5e, 0x82, 0x07, 0x11, 0xa4, 0xa5, 0xa7, 0xf5, 0x92,
	0x14, 0xad, 0x78, 0xaf, 0x6f, 0xb0, 0xd5, 0x8b, 0x78, 0xc9, 0xa6, 0x43, 0x3b, 0x74, 0xbb, 0x89,
	0x93, 0xd9, 0x88, 0xbe, 0x8c, 0xaf, 0x2a, 0x5b, 0xeb, 0x41, 0x14, 0xf4, 0xb2, 0x97, 0x40, 0xc2,
	0xf7, 0xf1, 0xcb, 0xe1, 0x83, 0x8b, 0x3c, 0xb5, 0xb1, 0xb2, 0x91, 0x29, 0x3b, 0x41, 0x2b, 0x98,
	0x64, 0x7b, 0x98, 0xc8, 0x41, 0x82, 0x1e, 0x46, 0x0a, 0xc2, 0xeb, 0xf0, 0xe2, 0x6a, 0xff, 0xd6,
	0xae, 0xbb, 0xd7, 0xcd, 0x92, 0xa3, 0x37, 0x79, 0x6a, 0xba, 0xd8, 0x78, 0x08, 0x07, 0x25, 0x3e,
	0xb7, 0x98, 0x44, 0x9f, 0xc2, 0x20, 0xbb, 0xba, 0xc5, 0x33, 0x35, 0x52, 0xc5, 0x51, 0xf9, 0x79,
	0x19, 0x8f, 0xe0, 0xb0, 0xc4, 0x14, 0x43, 0x93, 0xf0, 0xf7, 0xc4, 0xd5, 0xfb, 0x1e, 0x1c, 0xdf,
	0x63, 0x92, 0x39, 0x72, 0x26, 0x8f, 0xfa, 0x09, 0x06, 0x0f, 0x8d, 0xe3, 0x57, 0x5d, 0x98, 0x3f,
	0x74, 0xb3, 0xa3, 0xcf, 0x2f, 0xff, 0x91, 0xdc, 0xfd, 0xc1, 0x03, 0x74, 0x10, 0xf2, 0x9c, 0x16,
	0xd8, 0x0b, 0x31, 0x51, 0x1d, 0x72, 0x57, 0x13, 0x36, 0xd2, 0x1b, 0x52, 0x28, 0xbd, 0x82, 0x93,
	0x19, 0x2d, 0x88, 0xd1, 0x0b, 0x85, 0xc6, 0xd5, 0x3d, 0x39, 0x13, 0x35, 0xbb, 0x7d, 0xbc, 0x59,
	0x92, 0xac, 0xda, 0xca, 0xf8, 0xb0, 0xb1, 0xdf, 0xab, 0xf6, 0xab, 0x6a, 0x7f, 0x2e, 0xa9, 0xda,
	0xdf, 0xae, 0xe8, 0xfa, 0x23, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xf1, 0x90, 0xd8, 0x66, 0x02, 0x00,
	0x00,
}
