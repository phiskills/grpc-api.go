// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package example

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
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
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
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
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

func (m *Response) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "example.Request")
	proto.RegisterType((*Response)(nil), "example.Response")
}

func init() {
	proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6)
}

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xe4, 0xb9,
	0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x05, 0x2e, 0x8e, 0xa0, 0xd4,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xec, 0x2a, 0x8c, 0x9e, 0x31, 0x72, 0xb1, 0x07, 0xa7, 0x16,
	0x95, 0x65, 0x26, 0xa7, 0x0a, 0xe9, 0x72, 0xb1, 0x05, 0x67, 0x82, 0x0c, 0x16, 0x12, 0xd0, 0x83,
	0xd9, 0x08, 0x35, 0x5f, 0x4a, 0x10, 0x49, 0x04, 0x6a, 0xa0, 0x29, 0x17, 0x0f, 0x48, 0x67, 0x6a,
	0x51, 0x70, 0x49, 0x51, 0x6a, 0x62, 0x2e, 0x51, 0x9a, 0x0c, 0x18, 0x41, 0xda, 0x9c, 0x73, 0x32,
	0x53, 0xf3, 0x4a, 0x48, 0xd0, 0xa6, 0xc1, 0x28, 0x64, 0xc7, 0x25, 0xec, 0x94, 0x99, 0x92, 0x59,
	0x94, 0x9a, 0x5c, 0x92, 0x99, 0x9f, 0x97, 0x98, 0x43, 0x92, 0x6e, 0x03, 0xc6, 0x24, 0x36, 0x70,
	0xd8, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x73, 0xf0, 0xac, 0x4c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Simple(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Service_ServerStreamClient, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (Service_ClientStreamClient, error)
	BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (Service_BidirectionalStreamClient, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Simple(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/example.Service/Simple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Service_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[0], "/example.Service/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_ServerStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type serviceServerStreamClient struct {
	grpc.ClientStream
}

func (x *serviceServerStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (Service_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[1], "/example.Service/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceClientStreamClient{stream}
	return x, nil
}

type Service_ClientStreamClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type serviceClientStreamClient struct {
	grpc.ClientStream
}

func (x *serviceClientStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceClientStreamClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (Service_BidirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[2], "/example.Service/BidirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceBidirectionalStreamClient{stream}
	return x, nil
}

type Service_BidirectionalStreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type serviceBidirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *serviceBidirectionalStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceBidirectionalStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Simple(context.Context, *Request) (*Response, error)
	ServerStream(*Request, Service_ServerStreamServer) error
	ClientStream(Service_ClientStreamServer) error
	BidirectionalStream(Service_BidirectionalStreamServer) error
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Simple(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Simple not implemented")
}
func (*UnimplementedServiceServer) ServerStream(req *Request, srv Service_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (*UnimplementedServiceServer) ClientStream(srv Service_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (*UnimplementedServiceServer) BidirectionalStream(srv Service_BidirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStream not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Simple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Simple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Service/Simple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Simple(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).ServerStream(m, &serviceServerStreamServer{stream})
}

type Service_ServerStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type serviceServerStreamServer struct {
	grpc.ServerStream
}

func (x *serviceServerStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).ClientStream(&serviceClientStreamServer{stream})
}

type Service_ClientStreamServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type serviceClientStreamServer struct {
	grpc.ServerStream
}

func (x *serviceClientStreamServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceClientStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_BidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).BidirectionalStream(&serviceBidirectionalStreamServer{stream})
}

type Service_BidirectionalStreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type serviceBidirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *serviceBidirectionalStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceBidirectionalStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Simple",
			Handler:    _Service_Simple_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _Service_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStream",
			Handler:       _Service_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStream",
			Handler:       _Service_BidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "example.proto",
}