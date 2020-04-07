// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/trace/v2/trace_service.proto

package envoy_service_trace_v2

import (
	context "context"
	fmt "fmt"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type StreamTracesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamTracesResponse) Reset()         { *m = StreamTracesResponse{} }
func (m *StreamTracesResponse) String() string { return proto.CompactTextString(m) }
func (*StreamTracesResponse) ProtoMessage()    {}
func (*StreamTracesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6feca8f22ae39b94, []int{0}
}

func (m *StreamTracesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesResponse.Unmarshal(m, b)
}
func (m *StreamTracesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesResponse.Marshal(b, m, deterministic)
}
func (m *StreamTracesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesResponse.Merge(m, src)
}
func (m *StreamTracesResponse) XXX_Size() int {
	return xxx_messageInfo_StreamTracesResponse.Size(m)
}
func (m *StreamTracesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesResponse proto.InternalMessageInfo

type StreamTracesMessage struct {
	Identifier           *StreamTracesMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Spans                []*v1.Span                      `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *StreamTracesMessage) Reset()         { *m = StreamTracesMessage{} }
func (m *StreamTracesMessage) String() string { return proto.CompactTextString(m) }
func (*StreamTracesMessage) ProtoMessage()    {}
func (*StreamTracesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6feca8f22ae39b94, []int{1}
}

func (m *StreamTracesMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesMessage.Unmarshal(m, b)
}
func (m *StreamTracesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesMessage.Marshal(b, m, deterministic)
}
func (m *StreamTracesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesMessage.Merge(m, src)
}
func (m *StreamTracesMessage) XXX_Size() int {
	return xxx_messageInfo_StreamTracesMessage.Size(m)
}
func (m *StreamTracesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesMessage proto.InternalMessageInfo

func (m *StreamTracesMessage) GetIdentifier() *StreamTracesMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamTracesMessage) GetSpans() []*v1.Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type StreamTracesMessage_Identifier struct {
	Node                 *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamTracesMessage_Identifier) Reset()         { *m = StreamTracesMessage_Identifier{} }
func (m *StreamTracesMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamTracesMessage_Identifier) ProtoMessage()    {}
func (*StreamTracesMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_6feca8f22ae39b94, []int{1, 0}
}

func (m *StreamTracesMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Unmarshal(m, b)
}
func (m *StreamTracesMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamTracesMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesMessage_Identifier.Merge(m, src)
}
func (m *StreamTracesMessage_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Size(m)
}
func (m *StreamTracesMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesMessage_Identifier proto.InternalMessageInfo

func (m *StreamTracesMessage_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamTracesResponse)(nil), "envoy.service.trace.v2.StreamTracesResponse")
	proto.RegisterType((*StreamTracesMessage)(nil), "envoy.service.trace.v2.StreamTracesMessage")
	proto.RegisterType((*StreamTracesMessage_Identifier)(nil), "envoy.service.trace.v2.StreamTracesMessage.Identifier")
}

func init() {
	proto.RegisterFile("envoy/service/trace/v2/trace_service.proto", fileDescriptor_6feca8f22ae39b94)
}

var fileDescriptor_6feca8f22ae39b94 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0xef, 0xf4, 0xde, 0x96, 0xcb, 0xb4, 0x8b, 0x1a, 0xa5, 0x2d, 0xa1, 0x62, 0x29, 0x0a,
	0x45, 0x65, 0x42, 0x23, 0xfa, 0x00, 0x11, 0x84, 0x2e, 0x94, 0x92, 0x8a, 0x5b, 0x99, 0x26, 0xc7,
	0x32, 0xd0, 0xce, 0x0c, 0x99, 0x69, 0xb0, 0x0b, 0xc1, 0xa5, 0xf8, 0x48, 0x3e, 0x81, 0x5b, 0x1f,
	0xc3, 0x57, 0x70, 0x25, 0x99, 0x49, 0x6a, 0x16, 0x15, 0x74, 0x77, 0x38, 0xe7, 0xff, 0x7f, 0xbe,
	0xf3, 0xe3, 0x43, 0xe0, 0xa9, 0x58, 0x79, 0x0a, 0x92, 0x94, 0x45, 0xe0, 0xe9, 0x84, 0x46, 0xe0,
	0xa5, 0xbe, 0x1d, 0x6e, 0xf3, 0x35, 0x91, 0x89, 0xd0, 0xc2, 0x69, 0x19, 0x2d, 0x29, 0x96, 0x46,
	0x42, 0x52, 0xdf, 0xed, 0xda, 0x0c, 0x2a, 0x59, 0xe6, 0x8c, 0x44, 0x02, 0xde, 0x94, 0xaa, 0xdc,
	0xe5, 0x76, 0x67, 0x42, 0xcc, 0xe6, 0x60, 0xce, 0x94, 0x73, 0xa1, 0xa9, 0x66, 0x82, 0xab, 0xfc,
	0x7a, 0x20, 0x24, 0xf0, 0x08, 0xb8, 0x5a, 0x2a, 0xcf, 0x6c, 0x0a, 0x84, 0xa1, 0x1d, 0x72, 0xd9,
	0xee, 0x32, 0x96, 0xb4, 0x6c, 0xf7, 0x94, 0xa6, 0x7a, 0x59, 0xa4, 0xb4, 0x53, 0x3a, 0x67, 0x31,
	0xd5, 0xe0, 0x15, 0x83, 0x3d, 0xf4, 0x5b, 0x78, 0x67, 0xa2, 0x13, 0xa0, 0x8b, 0xeb, 0x2c, 0x4c,
	0x85, 0xa0, 0xa4, 0xe0, 0x0a, 0xfa, 0xef, 0x08, 0x6f, 0x97, 0x0f, 0x97, 0xa0, 0x14, 0x9d, 0x81,
	0x73, 0x83, 0x31, 0x8b, 0x81, 0x6b, 0x76, 0xc7, 0x20, 0xe9, 0xa0, 0x1e, 0x1a, 0xd4, 0xfd, 0x33,
	0xb2, 0xf9, 0x6f, 0xb2, 0x21, 0x80, 0x8c, 0xd6, 0xee, 0xb0, 0x94, 0xe4, 0x9c, 0xe2, 0xaa, 0x92,
	0x94, 0xab, 0x4e, 0xa5, 0xf7, 0x77, 0x50, 0xf7, 0xf7, 0xc8, 0xd7, 0xdb, 0x96, 0xb4, 0x48, 0x1d,
	0x92, 0x89, 0xa4, 0x3c, 0xb4, 0x6a, 0xf7, 0x1c, 0xe3, 0x51, 0x39, 0xe4, 0x1f, 0x17, 0x31, 0xe4,
	0x58, 0xed, 0x1c, 0x8b, 0x4a, 0x96, 0xc1, 0x64, 0xb5, 0x93, 0x2b, 0x11, 0x43, 0xf0, 0xff, 0x23,
	0xa8, 0x3e, 0xa3, 0x4a, 0x13, 0x85, 0x46, 0xee, 0x3f, 0xe0, 0x86, 0x61, 0x9c, 0x58, 0x7e, 0x67,
	0x81, 0x1b, 0x65, 0x72, 0xe7, 0xe8, 0x17, 0xff, 0xb9, 0xc7, 0x3f, 0x11, 0xaf, 0x6b, 0xfe, 0x33,
	0x40, 0xc1, 0xc5, 0xcb, 0xe3, 0xeb, 0x5b, 0xad, 0xd2, 0x44, 0x78, 0x9f, 0x09, 0xeb, 0x96, 0x89,
	0xb8, 0x5f, 0x7d, 0x13, 0x14, 0x6c, 0x95, 0x61, 0xc7, 0x59, 0x37, 0x63, 0xf4, 0x84, 0xd0, 0xb4,
	0x66, 0x7a, 0x3a, 0xf9, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x38, 0x9d, 0x0d, 0xe6, 0xb2, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceServiceClient interface {
	StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error)
}

type traceServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceServiceClient(cc *grpc.ClientConn) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceService_serviceDesc.Streams[0], "/envoy.service.trace.v2.TraceService/StreamTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceServiceStreamTracesClient{stream}
	return x, nil
}

type TraceService_StreamTracesClient interface {
	Send(*StreamTracesMessage) error
	CloseAndRecv() (*StreamTracesResponse, error)
	grpc.ClientStream
}

type traceServiceStreamTracesClient struct {
	grpc.ClientStream
}

func (x *traceServiceStreamTracesClient) Send(m *StreamTracesMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceServiceStreamTracesClient) CloseAndRecv() (*StreamTracesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamTracesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceServiceServer is the server API for TraceService service.
type TraceServiceServer interface {
	StreamTraces(TraceService_StreamTracesServer) error
}

// UnimplementedTraceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (*UnimplementedTraceServiceServer) StreamTraces(srv TraceService_StreamTracesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTraces not implemented")
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_StreamTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceServiceServer).StreamTraces(&traceServiceStreamTracesServer{stream})
}

type TraceService_StreamTracesServer interface {
	SendAndClose(*StreamTracesResponse) error
	Recv() (*StreamTracesMessage, error)
	grpc.ServerStream
}

type traceServiceStreamTracesServer struct {
	grpc.ServerStream
}

func (x *traceServiceStreamTracesServer) SendAndClose(m *StreamTracesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceServiceStreamTracesServer) Recv() (*StreamTracesMessage, error) {
	m := new(StreamTracesMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.trace.v2.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTraces",
			Handler:       _TraceService_StreamTraces_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/trace/v2/trace_service.proto",
}
