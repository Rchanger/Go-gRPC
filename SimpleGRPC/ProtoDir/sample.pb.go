// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample.proto

package ProtoDir

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

// to accept the request
type RequestData struct {
	FirstNumber          int64    `protobuf:"varint,1,opt,name=firstNumber,proto3" json:"firstNumber,omitempty"`
	SecondNumber         int64    `protobuf:"varint,2,opt,name=SecondNumber,proto3" json:"SecondNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestData) Reset()         { *m = RequestData{} }
func (m *RequestData) String() string { return proto.CompactTextString(m) }
func (*RequestData) ProtoMessage()    {}
func (*RequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{0}
}

func (m *RequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestData.Unmarshal(m, b)
}
func (m *RequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestData.Marshal(b, m, deterministic)
}
func (m *RequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestData.Merge(m, src)
}
func (m *RequestData) XXX_Size() int {
	return xxx_messageInfo_RequestData.Size(m)
}
func (m *RequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestData.DiscardUnknown(m)
}

var xxx_messageInfo_RequestData proto.InternalMessageInfo

func (m *RequestData) GetFirstNumber() int64 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *RequestData) GetSecondNumber() int64 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

// to send the response to request
type ResponseData struct {
	Result               int64    `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseData) Reset()         { *m = ResponseData{} }
func (m *ResponseData) String() string { return proto.CompactTextString(m) }
func (*ResponseData) ProtoMessage()    {}
func (*ResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{1}
}

func (m *ResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseData.Unmarshal(m, b)
}
func (m *ResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseData.Marshal(b, m, deterministic)
}
func (m *ResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseData.Merge(m, src)
}
func (m *ResponseData) XXX_Size() int {
	return xxx_messageInfo_ResponseData.Size(m)
}
func (m *ResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseData proto.InternalMessageInfo

func (m *ResponseData) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestData)(nil), "ProtoDir.RequestData")
	proto.RegisterType((*ResponseData)(nil), "ProtoDir.ResponseData")
}

func init() {
	proto.RegisterFile("sample.proto", fileDescriptor_2141552de9bf11d0)
}

var fileDescriptor_2141552de9bf11d0 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x08, 0x00, 0x51, 0x2e, 0x99, 0x45,
	0x4a, 0xc1, 0x5c, 0xdc, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x2e, 0x89, 0x25, 0x89, 0x42,
	0x0a, 0x5c, 0xdc, 0x69, 0x99, 0x45, 0xc5, 0x25, 0x7e, 0xa5, 0xb9, 0x49, 0xa9, 0x45, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0xc8, 0x42, 0x42, 0x4a, 0x5c, 0x3c, 0xc1, 0xa9, 0xc9, 0xf9, 0x79,
	0x29, 0x50, 0x25, 0x4c, 0x60, 0x25, 0x28, 0x62, 0x4a, 0x6a, 0x5c, 0x3c, 0x41, 0xa9, 0xc5, 0x05,
	0xf9, 0x79, 0xc5, 0xa9, 0x60, 0x53, 0xc5, 0xb8, 0xd8, 0x82, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0xa0,
	0x06, 0x42, 0x79, 0x46, 0x8d, 0x8c, 0x5c, 0x5c, 0xce, 0x89, 0x39, 0xc1, 0xa9, 0x45, 0x65, 0x99,
	0xc9, 0xa9, 0x42, 0x66, 0x5c, 0xcc, 0x8e, 0x29, 0x29, 0x42, 0xa2, 0x7a, 0x30, 0xd7, 0xe9, 0x21,
	0x39, 0x4d, 0x4a, 0x0c, 0x59, 0x18, 0x61, 0xb8, 0x12, 0x83, 0x90, 0x35, 0x17, 0x47, 0x70, 0x69,
	0x52, 0x49, 0x51, 0x62, 0x72, 0x09, 0xc9, 0x9a, 0x93, 0xd8, 0xc0, 0x21, 0x62, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x9a, 0x1f, 0x6e, 0x30, 0x21, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalServiceClient is the client API for CalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalServiceClient interface {
	Add(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error)
	Subtract(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error)
}

type calServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalServiceClient(cc grpc.ClientConnInterface) CalServiceClient {
	return &calServiceClient{cc}
}

func (c *calServiceClient) Add(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error) {
	out := new(ResponseData)
	err := c.cc.Invoke(ctx, "/ProtoDir.CalService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calServiceClient) Subtract(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error) {
	out := new(ResponseData)
	err := c.cc.Invoke(ctx, "/ProtoDir.CalService/Subtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalServiceServer is the server API for CalService service.
type CalServiceServer interface {
	Add(context.Context, *RequestData) (*ResponseData, error)
	Subtract(context.Context, *RequestData) (*ResponseData, error)
}

// UnimplementedCalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalServiceServer struct {
}

func (*UnimplementedCalServiceServer) Add(ctx context.Context, req *RequestData) (*ResponseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCalServiceServer) Subtract(ctx context.Context, req *RequestData) (*ResponseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}

func RegisterCalServiceServer(s *grpc.Server, srv CalServiceServer) {
	s.RegisterService(&_CalService_serviceDesc, srv)
}

func _CalService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProtoDir.CalService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalServiceServer).Add(ctx, req.(*RequestData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalService_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalServiceServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProtoDir.CalService/Subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalServiceServer).Subtract(ctx, req.(*RequestData))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProtoDir.CalService",
	HandlerType: (*CalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalService_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _CalService_Subtract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample.proto",
}