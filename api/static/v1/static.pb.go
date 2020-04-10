// Code generated by protoc-gen-go. DO NOT EDIT.
// source: static.proto

package binacs_api_static_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type IndexRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexRequest) Reset()         { *m = IndexRequest{} }
func (m *IndexRequest) String() string { return proto.CompactTextString(m) }
func (*IndexRequest) ProtoMessage()    {}
func (*IndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_800d8474b3bfe66d, []int{0}
}

func (m *IndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexRequest.Unmarshal(m, b)
}
func (m *IndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexRequest.Marshal(b, m, deterministic)
}
func (m *IndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexRequest.Merge(m, src)
}
func (m *IndexRequest) XXX_Size() int {
	return xxx_messageInfo_IndexRequest.Size(m)
}
func (m *IndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexRequest proto.InternalMessageInfo

type IndexResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexResponse) Reset()         { *m = IndexResponse{} }
func (m *IndexResponse) String() string { return proto.CompactTextString(m) }
func (*IndexResponse) ProtoMessage()    {}
func (*IndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_800d8474b3bfe66d, []int{1}
}

func (m *IndexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexResponse.Unmarshal(m, b)
}
func (m *IndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexResponse.Marshal(b, m, deterministic)
}
func (m *IndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexResponse.Merge(m, src)
}
func (m *IndexResponse) XXX_Size() int {
	return xxx_messageInfo_IndexResponse.Size(m)
}
func (m *IndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IndexResponse proto.InternalMessageInfo

func (m *IndexResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*IndexRequest)(nil), "binacs_api_static_v1.IndexRequest")
	proto.RegisterType((*IndexResponse)(nil), "binacs_api_static_v1.IndexResponse")
}

func init() {
	proto.RegisterFile("static.proto", fileDescriptor_800d8474b3bfe66d)
}

var fileDescriptor_800d8474b3bfe66d = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x49, 0x2c,
	0xc9, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x49, 0xca, 0xcc, 0x4b, 0x4c, 0x2e,
	0x8e, 0x4f, 0x2c, 0xc8, 0x8c, 0x87, 0x48, 0xc4, 0x97, 0x19, 0x4a, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7,
	0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x83, 0x24, 0xf2, 0xf3, 0x8a,
	0x21, 0x7a, 0x94, 0xf8, 0xb8, 0x78, 0x3c, 0xf3, 0x52, 0x52, 0x2b, 0x82, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x94, 0x34, 0xb9, 0x78, 0xa1, 0xfc, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x09,
	0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x18, 0xd7, 0x28, 0x87, 0x8b, 0x15, 0xac, 0x54, 0x28, 0x19, 0xc6, 0x50, 0xd2, 0xc3, 0xe6, 0x02,
	0x3d, 0x64, 0x0b, 0xa4, 0x94, 0xf1, 0xaa, 0x81, 0x58, 0xaa, 0x24, 0xd8, 0x74, 0xf9, 0xc9, 0x64,
	0x26, 0x6e, 0x25, 0x36, 0xfd, 0x4c, 0x90, 0xb8, 0x15, 0xa3, 0x56, 0x12, 0x1b, 0xd8, 0xbd, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xc5, 0xd6, 0x01, 0xf3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IndexClient is the client API for Index service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexClient interface {
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error)
}

type indexClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexClient(cc grpc.ClientConnInterface) IndexClient {
	return &indexClient{cc}
}

func (c *indexClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := c.cc.Invoke(ctx, "/binacs_api_static_v1.Index/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServer is the server API for Index service.
type IndexServer interface {
	Index(context.Context, *IndexRequest) (*IndexResponse, error)
}

// UnimplementedIndexServer can be embedded to have forward compatible implementations.
type UnimplementedIndexServer struct {
}

func (*UnimplementedIndexServer) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}

func RegisterIndexServer(s *grpc.Server, srv IndexServer) {
	s.RegisterService(&_Index_serviceDesc, srv)
}

func _Index_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binacs_api_static_v1.Index/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Index_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binacs_api_static_v1.Index",
	HandlerType: (*IndexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _Index_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "static.proto",
}
