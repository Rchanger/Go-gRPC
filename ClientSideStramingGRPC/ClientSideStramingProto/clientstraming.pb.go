// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clientstraming.proto

package ClientSideStramingProto

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
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe3c92eb860d1834, []int{0}
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

func (m *Request) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type Response struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe3c92eb860d1834, []int{1}
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

func (m *Response) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "ClientSideStramingProto.Request")
	proto.RegisterType((*Response)(nil), "ClientSideStramingProto.Response")
}

func init() {
	proto.RegisterFile("clientstraming.proto", fileDescriptor_fe3c92eb860d1834)
}

var fileDescriptor_fe3c92eb860d1834 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0x29, 0x2e, 0x29, 0x4a, 0xcc, 0xcd, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x77, 0x06, 0x8b, 0x06, 0x67, 0xa6, 0xa4, 0x06, 0x43, 0x65, 0x02, 0x40, 0x12, 0x4a,
	0xd2, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0x79, 0xa5,
	0xb9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x20, 0xa6, 0x92, 0x12, 0x17, 0x47, 0x50, 0x6a,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18, 0x17, 0x5b, 0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09,
	0x54, 0x01, 0x94, 0x67, 0x94, 0xcc, 0xc5, 0x1b, 0x9c, 0x99, 0x5b, 0x90, 0x93, 0x1a, 0x9c, 0x5a,
	0x54, 0x96, 0x99, 0x9c, 0x2a, 0x14, 0xc4, 0xc5, 0xee, 0x96, 0x99, 0x97, 0xe2, 0x9b, 0x58, 0x21,
	0xa4, 0xa0, 0x87, 0xc3, 0x5a, 0x3d, 0xa8, 0x9d, 0x52, 0x8a, 0x78, 0x54, 0x40, 0x2c, 0x56, 0x62,
	0xd0, 0x60, 0x4c, 0x62, 0x03, 0xfb, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x73, 0x8c, 0x04,
	0x5a, 0xdd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SimpleServiceClient is the client API for SimpleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimpleServiceClient interface {
	FindMax(ctx context.Context, opts ...grpc.CallOption) (SimpleService_FindMaxClient, error)
}

type simpleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleServiceClient(cc grpc.ClientConnInterface) SimpleServiceClient {
	return &simpleServiceClient{cc}
}

func (c *simpleServiceClient) FindMax(ctx context.Context, opts ...grpc.CallOption) (SimpleService_FindMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SimpleService_serviceDesc.Streams[0], "/ClientSideStramingProto.SimpleService/FindMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleServiceFindMaxClient{stream}
	return x, nil
}

type SimpleService_FindMaxClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type simpleServiceFindMaxClient struct {
	grpc.ClientStream
}

func (x *simpleServiceFindMaxClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *simpleServiceFindMaxClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimpleServiceServer is the server API for SimpleService service.
type SimpleServiceServer interface {
	FindMax(SimpleService_FindMaxServer) error
}

// UnimplementedSimpleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSimpleServiceServer struct {
}

func (*UnimplementedSimpleServiceServer) FindMax(srv SimpleService_FindMaxServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMax not implemented")
}

func RegisterSimpleServiceServer(s *grpc.Server, srv SimpleServiceServer) {
	s.RegisterService(&_SimpleService_serviceDesc, srv)
}

func _SimpleService_FindMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SimpleServiceServer).FindMax(&simpleServiceFindMaxServer{stream})
}

type SimpleService_FindMaxServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type simpleServiceFindMaxServer struct {
	grpc.ServerStream
}

func (x *simpleServiceFindMaxServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *simpleServiceFindMaxServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SimpleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ClientSideStramingProto.SimpleService",
	HandlerType: (*SimpleServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindMax",
			Handler:       _SimpleService_FindMax_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "clientstraming.proto",
}
