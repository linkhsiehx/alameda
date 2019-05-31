// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datahub/rawdata/services.proto

package rawdata // import "github.com/containers-ai/api/datahub/rawdata"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/containers-ai/api/common"
import status "google.golang.org/genproto/googleapis/rpc/status"

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

// Represents a request for reading rawdata from database
type ReadRawdataRequest struct {
	Queries              []*common.Query `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReadRawdataRequest) Reset()         { *m = ReadRawdataRequest{} }
func (m *ReadRawdataRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRawdataRequest) ProtoMessage()    {}
func (*ReadRawdataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_eb1efe62dfc246fa, []int{0}
}
func (m *ReadRawdataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRawdataRequest.Unmarshal(m, b)
}
func (m *ReadRawdataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRawdataRequest.Marshal(b, m, deterministic)
}
func (dst *ReadRawdataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRawdataRequest.Merge(dst, src)
}
func (m *ReadRawdataRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRawdataRequest.Size(m)
}
func (m *ReadRawdataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRawdataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRawdataRequest proto.InternalMessageInfo

func (m *ReadRawdataRequest) GetQueries() []*common.Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

// Represents a response for listing rawdata from database
type ReadRawdataResponse struct {
	Status               *status.Status        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Rawdata              []*common.ReadRawdata `protobuf:"bytes,2,rep,name=rawdata,proto3" json:"rawdata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ReadRawdataResponse) Reset()         { *m = ReadRawdataResponse{} }
func (m *ReadRawdataResponse) String() string { return proto.CompactTextString(m) }
func (*ReadRawdataResponse) ProtoMessage()    {}
func (*ReadRawdataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_eb1efe62dfc246fa, []int{1}
}
func (m *ReadRawdataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRawdataResponse.Unmarshal(m, b)
}
func (m *ReadRawdataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRawdataResponse.Marshal(b, m, deterministic)
}
func (dst *ReadRawdataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRawdataResponse.Merge(dst, src)
}
func (m *ReadRawdataResponse) XXX_Size() int {
	return xxx_messageInfo_ReadRawdataResponse.Size(m)
}
func (m *ReadRawdataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRawdataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRawdataResponse proto.InternalMessageInfo

func (m *ReadRawdataResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReadRawdataResponse) GetRawdata() []*common.ReadRawdata {
	if m != nil {
		return m.Rawdata
	}
	return nil
}

// Represents a request for writing rawdata to database
type WriteRawdataRequest struct {
	Rawdata              []*common.WriteRawdata `protobuf:"bytes,1,rep,name=rawdata,proto3" json:"rawdata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *WriteRawdataRequest) Reset()         { *m = WriteRawdataRequest{} }
func (m *WriteRawdataRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRawdataRequest) ProtoMessage()    {}
func (*WriteRawdataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_eb1efe62dfc246fa, []int{2}
}
func (m *WriteRawdataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRawdataRequest.Unmarshal(m, b)
}
func (m *WriteRawdataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRawdataRequest.Marshal(b, m, deterministic)
}
func (dst *WriteRawdataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRawdataRequest.Merge(dst, src)
}
func (m *WriteRawdataRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRawdataRequest.Size(m)
}
func (m *WriteRawdataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRawdataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRawdataRequest proto.InternalMessageInfo

func (m *WriteRawdataRequest) GetRawdata() []*common.WriteRawdata {
	if m != nil {
		return m.Rawdata
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadRawdataRequest)(nil), "containersai.datahub.rawdata.ReadRawdataRequest")
	proto.RegisterType((*ReadRawdataResponse)(nil), "containersai.datahub.rawdata.ReadRawdataResponse")
	proto.RegisterType((*WriteRawdataRequest)(nil), "containersai.datahub.rawdata.WriteRawdataRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RawdataServiceClient is the client API for RawdataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RawdataServiceClient interface {
	// Used to read rawdata from database
	ReadRawdata(ctx context.Context, in *ReadRawdataRequest, opts ...grpc.CallOption) (*ReadRawdataResponse, error)
	// Used to write rawdata to database
	WriteRawdata(ctx context.Context, in *WriteRawdataRequest, opts ...grpc.CallOption) (*status.Status, error)
}

type rawdataServiceClient struct {
	cc *grpc.ClientConn
}

func NewRawdataServiceClient(cc *grpc.ClientConn) RawdataServiceClient {
	return &rawdataServiceClient{cc}
}

func (c *rawdataServiceClient) ReadRawdata(ctx context.Context, in *ReadRawdataRequest, opts ...grpc.CallOption) (*ReadRawdataResponse, error) {
	out := new(ReadRawdataResponse)
	err := c.cc.Invoke(ctx, "/containersai.datahub.rawdata.RawdataService/ReadRawdata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rawdataServiceClient) WriteRawdata(ctx context.Context, in *WriteRawdataRequest, opts ...grpc.CallOption) (*status.Status, error) {
	out := new(status.Status)
	err := c.cc.Invoke(ctx, "/containersai.datahub.rawdata.RawdataService/WriteRawdata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RawdataServiceServer is the server API for RawdataService service.
type RawdataServiceServer interface {
	// Used to read rawdata from database
	ReadRawdata(context.Context, *ReadRawdataRequest) (*ReadRawdataResponse, error)
	// Used to write rawdata to database
	WriteRawdata(context.Context, *WriteRawdataRequest) (*status.Status, error)
}

func RegisterRawdataServiceServer(s *grpc.Server, srv RawdataServiceServer) {
	s.RegisterService(&_RawdataService_serviceDesc, srv)
}

func _RawdataService_ReadRawdata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRawdataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawdataServiceServer).ReadRawdata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containersai.datahub.rawdata.RawdataService/ReadRawdata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawdataServiceServer).ReadRawdata(ctx, req.(*ReadRawdataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RawdataService_WriteRawdata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRawdataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawdataServiceServer).WriteRawdata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containersai.datahub.rawdata.RawdataService/WriteRawdata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawdataServiceServer).WriteRawdata(ctx, req.(*WriteRawdataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RawdataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containersai.datahub.rawdata.RawdataService",
	HandlerType: (*RawdataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadRawdata",
			Handler:    _RawdataService_ReadRawdata_Handler,
		},
		{
			MethodName: "WriteRawdata",
			Handler:    _RawdataService_WriteRawdata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datahub/rawdata/services.proto",
}

func init() {
	proto.RegisterFile("datahub/rawdata/services.proto", fileDescriptor_services_eb1efe62dfc246fa)
}

var fileDescriptor_services_eb1efe62dfc246fa = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x4b, 0xc3, 0x30,
	0x14, 0x5f, 0x15, 0x36, 0xc8, 0xc4, 0x43, 0x76, 0x70, 0x04, 0x91, 0xd9, 0xd3, 0x10, 0x4d, 0x74,
	0x7a, 0xd2, 0x9b, 0x47, 0x6f, 0x66, 0x87, 0x81, 0xb7, 0x2c, 0x7b, 0xd4, 0x80, 0x6d, 0xb2, 0x24,
	0x55, 0x76, 0xf0, 0xb3, 0xfa, 0x55, 0x64, 0x4d, 0xca, 0xda, 0x59, 0x14, 0x4f, 0x2d, 0x79, 0xbf,
	0x3f, 0xef, 0xfd, 0xde, 0x43, 0x67, 0x2b, 0xe1, 0xc5, 0x6b, 0xb9, 0x64, 0x56, 0x7c, 0x6c, 0x7f,
	0x99, 0x03, 0xfb, 0xae, 0x24, 0x38, 0x6a, 0xac, 0xf6, 0x1a, 0x9f, 0x4a, 0x5d, 0x78, 0xa1, 0x0a,
	0xb0, 0x4e, 0x28, 0x1a, 0xc1, 0x34, 0x82, 0xc9, 0x48, 0xea, 0x3c, 0xd7, 0x05, 0x0b, 0x9f, 0x40,
	0x21, 0x38, 0x3e, 0xfa, 0x8d, 0xa9, 0x65, 0xc8, 0x49, 0xa6, 0x75, 0xf6, 0x06, 0xcc, 0x1a, 0xc9,
	0x9c, 0x17, 0xbe, 0x8c, 0x85, 0xf4, 0x09, 0x61, 0x0e, 0x62, 0xc5, 0x83, 0x20, 0x87, 0x75, 0x09,
	0xce, 0xe3, 0x3b, 0x34, 0x58, 0x97, 0x60, 0x15, 0xb8, 0x71, 0x32, 0x39, 0x9c, 0x0e, 0x67, 0x84,
	0xb6, 0xfa, 0x88, 0x7e, 0xcf, 0x25, 0xd8, 0x0d, 0xaf, 0xa1, 0xe9, 0x27, 0x1a, 0xb5, 0xb4, 0x9c,
	0xd1, 0x85, 0x03, 0x7c, 0x81, 0xfa, 0xc1, 0x72, 0x9c, 0x4c, 0x92, 0xe9, 0x70, 0x86, 0x69, 0x68,
	0x86, 0x5a, 0x23, 0xe9, 0xbc, 0xaa, 0xf0, 0x88, 0xc0, 0xf7, 0x68, 0x10, 0x67, 0x1b, 0x1f, 0x54,
	0xc6, 0x93, 0x4e, 0xe3, 0xa6, 0x4d, 0x4d, 0x48, 0x39, 0x1a, 0x2d, 0xac, 0xf2, 0xb0, 0x37, 0xcb,
	0xc3, 0x4e, 0x32, 0xcc, 0x72, 0xde, 0x29, 0xd9, 0xa2, 0xd6, 0x8c, 0xd9, 0x57, 0x82, 0x8e, 0xe3,
	0xe3, 0x3c, 0x2c, 0x06, 0x7b, 0x34, 0x6c, 0xd8, 0xe3, 0x6b, 0xfa, 0xdb, 0x86, 0xe8, 0xcf, 0x70,
	0xc9, 0xcd, 0x3f, 0x18, 0x21, 0xc2, 0xb4, 0x87, 0x17, 0xe8, 0xa8, 0xd9, 0x21, 0xfe, 0x43, 0xa4,
	0x23, 0x08, 0xd2, 0x91, 0x7b, 0xda, 0x7b, 0xa4, 0x2f, 0x97, 0x99, 0xf2, 0x5b, 0xae, 0xd4, 0x39,
	0xdb, 0x89, 0x5e, 0x09, 0xc5, 0x84, 0x51, 0x6c, 0xef, 0x3e, 0x97, 0xfd, 0xea, 0x6e, 0x6e, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x73, 0x32, 0x71, 0xb6, 0xb9, 0x02, 0x00, 0x00,
}
