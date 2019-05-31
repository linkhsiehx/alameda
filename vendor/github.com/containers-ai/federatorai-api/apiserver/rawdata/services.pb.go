// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apiserver/rawdata/services.proto

package rawdata // import "github.com/containers-ai/federatorai-api/apiserver/rawdata"

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
	return fileDescriptor_services_0f5d858991c20b9a, []int{0}
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
	return fileDescriptor_services_0f5d858991c20b9a, []int{1}
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
	Rawdata              []*common.WriteRawdata `protobuf:"bytes,2,rep,name=rawdata,proto3" json:"rawdata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *WriteRawdataRequest) Reset()         { *m = WriteRawdataRequest{} }
func (m *WriteRawdataRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRawdataRequest) ProtoMessage()    {}
func (*WriteRawdataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_0f5d858991c20b9a, []int{2}
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
	proto.RegisterType((*ReadRawdataRequest)(nil), "containersai.apiserver.rawdata.ReadRawdataRequest")
	proto.RegisterType((*ReadRawdataResponse)(nil), "containersai.apiserver.rawdata.ReadRawdataResponse")
	proto.RegisterType((*WriteRawdataRequest)(nil), "containersai.apiserver.rawdata.WriteRawdataRequest")
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
	err := c.cc.Invoke(ctx, "/containersai.apiserver.rawdata.RawdataService/ReadRawdata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rawdataServiceClient) WriteRawdata(ctx context.Context, in *WriteRawdataRequest, opts ...grpc.CallOption) (*status.Status, error) {
	out := new(status.Status)
	err := c.cc.Invoke(ctx, "/containersai.apiserver.rawdata.RawdataService/WriteRawdata", in, out, opts...)
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
		FullMethod: "/containersai.apiserver.rawdata.RawdataService/ReadRawdata",
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
		FullMethod: "/containersai.apiserver.rawdata.RawdataService/WriteRawdata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawdataServiceServer).WriteRawdata(ctx, req.(*WriteRawdataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RawdataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containersai.apiserver.rawdata.RawdataService",
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
	Metadata: "apiserver/rawdata/services.proto",
}

func init() {
	proto.RegisterFile("apiserver/rawdata/services.proto", fileDescriptor_services_0f5d858991c20b9a)
}

var fileDescriptor_services_0f5d858991c20b9a = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0x9b, 0xf7, 0x95, 0x5a, 0xc9, 0x45, 0x0c, 0xee, 0x40, 0x95, 0x01, 0x85, 0x48, 0x48,
	0x15, 0x52, 0x6d, 0x29, 0x65, 0x2a, 0x4c, 0x8c, 0x6c, 0xa4, 0x03, 0x82, 0xed, 0x9a, 0x1e, 0xc5,
	0x12, 0x8d, 0xdd, 0xb3, 0xc3, 0x9f, 0x81, 0x8f, 0xcb, 0xf7, 0x40, 0xad, 0x5d, 0x68, 0xda, 0x08,
	0xc4, 0x14, 0xc5, 0xbe, 0xfb, 0x3d, 0xe7, 0xe7, 0x1e, 0x96, 0x80, 0x51, 0x16, 0xe9, 0x19, 0x49,
	0x12, 0xbc, 0xcc, 0xc0, 0x81, 0x5c, 0xfd, 0xaa, 0x02, 0xad, 0x30, 0xa4, 0x9d, 0xe6, 0xc7, 0x85,
	0x2e, 0x1d, 0xa8, 0x12, 0xc9, 0x82, 0x12, 0x5f, 0xe5, 0x22, 0x94, 0xc7, 0xa7, 0xdf, 0xf7, 0x43,
	0x50, 0x12, 0x8c, 0x92, 0x85, 0x5e, 0x2c, 0x74, 0x19, 0x3e, 0x1e, 0x13, 0x1f, 0xcd, 0xb5, 0x9e,
	0x3f, 0xa1, 0x24, 0x53, 0x48, 0xeb, 0xc0, 0x55, 0x81, 0x9f, 0x5e, 0x33, 0x9e, 0x23, 0xcc, 0x72,
	0x8f, 0xcb, 0x71, 0x59, 0xa1, 0x75, 0xfc, 0x9c, 0x75, 0x96, 0x15, 0x92, 0x42, 0xdb, 0x8f, 0x92,
	0xff, 0x83, 0x6e, 0x16, 0x8b, 0xda, 0x1c, 0x81, 0x7d, 0x53, 0x21, 0xbd, 0xe5, 0x9b, 0xd2, 0xf4,
	0x9d, 0xf5, 0x6a, 0x2c, 0x6b, 0x74, 0x69, 0x91, 0x9f, 0xb1, 0xb6, 0x97, 0xec, 0x47, 0x49, 0x34,
	0xe8, 0x66, 0x5c, 0xf8, 0x61, 0x04, 0x99, 0x42, 0x4c, 0xd6, 0x37, 0x79, 0xa8, 0xe0, 0x63, 0xd6,
	0x09, 0x2f, 0xeb, 0xff, 0x5b, 0x0b, 0x27, 0x8d, 0xc2, 0xdb, 0x32, 0x9b, 0x86, 0x34, 0x67, 0xbd,
	0x5b, 0x52, 0x0e, 0x77, 0xde, 0x72, 0xb1, 0x8b, 0x3c, 0x69, 0x44, 0xd6, 0x5a, 0x37, 0x1d, 0xd9,
	0x47, 0xc4, 0x0e, 0xc3, 0xe1, 0xc4, 0x2f, 0x86, 0xbf, 0xb2, 0xee, 0x96, 0x3c, 0xcf, 0xc4, 0xcf,
	0x1b, 0x12, 0xfb, 0xf6, 0xc6, 0xa3, 0x3f, 0xf5, 0x78, 0x1b, 0xd3, 0x16, 0xbf, 0x63, 0x07, 0xdb,
	0x53, 0xf2, 0x5f, 0x31, 0x0d, 0x76, 0xc4, 0x0d, 0xee, 0xa7, 0xad, 0xab, 0xcb, 0xfb, 0xf1, 0x5c,
	0xb9, 0xc7, 0x6a, 0xba, 0xb2, 0x43, 0xd6, 0x33, 0xf5, 0x80, 0x33, 0x24, 0x70, 0x9a, 0x40, 0x0d,
	0x57, 0xf9, 0xda, 0xcb, 0xec, 0xb4, 0xbd, 0xce, 0xd2, 0xe8, 0x33, 0x00, 0x00, 0xff, 0xff, 0x46,
	0x81, 0xdc, 0xbe, 0xcf, 0x02, 0x00, 0x00,
}
