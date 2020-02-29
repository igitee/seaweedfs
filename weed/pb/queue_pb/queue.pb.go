// Code generated by protoc-gen-go.
// source: queue.proto
// DO NOT EDIT!

/*
Package queue_pb is a generated protocol buffer package.

It is generated from these files:
	queue.proto

It has these top-level messages:
	WriteMessageRequest
	WriteMessageResponse
	ReadMessageRequest
	ReadMessageResponse
	ConfigureTopicRequest
	ConfigureTopicResponse
	DeleteTopicRequest
	DeleteTopicResponse
*/
package queue_pb

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

type WriteMessageRequest struct {
	Topic   string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	EventNs int64  `protobuf:"varint,2,opt,name=event_ns,json=eventNs" json:"event_ns,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *WriteMessageRequest) Reset()                    { *m = WriteMessageRequest{} }
func (m *WriteMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteMessageRequest) ProtoMessage()               {}
func (*WriteMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WriteMessageRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *WriteMessageRequest) GetEventNs() int64 {
	if m != nil {
		return m.EventNs
	}
	return 0
}

func (m *WriteMessageRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type WriteMessageResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	AckNs int64  `protobuf:"varint,2,opt,name=ack_ns,json=ackNs" json:"ack_ns,omitempty"`
}

func (m *WriteMessageResponse) Reset()                    { *m = WriteMessageResponse{} }
func (m *WriteMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteMessageResponse) ProtoMessage()               {}
func (*WriteMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WriteMessageResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *WriteMessageResponse) GetAckNs() int64 {
	if m != nil {
		return m.AckNs
	}
	return 0
}

type ReadMessageRequest struct {
	Topic   string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	StartNs int64  `protobuf:"varint,2,opt,name=start_ns,json=startNs" json:"start_ns,omitempty"`
}

func (m *ReadMessageRequest) Reset()                    { *m = ReadMessageRequest{} }
func (m *ReadMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadMessageRequest) ProtoMessage()               {}
func (*ReadMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadMessageRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ReadMessageRequest) GetStartNs() int64 {
	if m != nil {
		return m.StartNs
	}
	return 0
}

type ReadMessageResponse struct {
	Error   string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	EventNs int64  `protobuf:"varint,2,opt,name=event_ns,json=eventNs" json:"event_ns,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ReadMessageResponse) Reset()                    { *m = ReadMessageResponse{} }
func (m *ReadMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadMessageResponse) ProtoMessage()               {}
func (*ReadMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadMessageResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ReadMessageResponse) GetEventNs() int64 {
	if m != nil {
		return m.EventNs
	}
	return 0
}

func (m *ReadMessageResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ConfigureTopicRequest struct {
	Topic      string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	TtlSeconds int64  `protobuf:"varint,2,opt,name=ttl_seconds,json=ttlSeconds" json:"ttl_seconds,omitempty"`
}

func (m *ConfigureTopicRequest) Reset()                    { *m = ConfigureTopicRequest{} }
func (m *ConfigureTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigureTopicRequest) ProtoMessage()               {}
func (*ConfigureTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ConfigureTopicRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ConfigureTopicRequest) GetTtlSeconds() int64 {
	if m != nil {
		return m.TtlSeconds
	}
	return 0
}

type ConfigureTopicResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *ConfigureTopicResponse) Reset()                    { *m = ConfigureTopicResponse{} }
func (m *ConfigureTopicResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigureTopicResponse) ProtoMessage()               {}
func (*ConfigureTopicResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ConfigureTopicResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DeleteTopicRequest struct {
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
}

func (m *DeleteTopicRequest) Reset()                    { *m = DeleteTopicRequest{} }
func (m *DeleteTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTopicRequest) ProtoMessage()               {}
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteTopicRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type DeleteTopicResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *DeleteTopicResponse) Reset()                    { *m = DeleteTopicResponse{} }
func (m *DeleteTopicResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteTopicResponse) ProtoMessage()               {}
func (*DeleteTopicResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteTopicResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*WriteMessageRequest)(nil), "queue_pb.WriteMessageRequest")
	proto.RegisterType((*WriteMessageResponse)(nil), "queue_pb.WriteMessageResponse")
	proto.RegisterType((*ReadMessageRequest)(nil), "queue_pb.ReadMessageRequest")
	proto.RegisterType((*ReadMessageResponse)(nil), "queue_pb.ReadMessageResponse")
	proto.RegisterType((*ConfigureTopicRequest)(nil), "queue_pb.ConfigureTopicRequest")
	proto.RegisterType((*ConfigureTopicResponse)(nil), "queue_pb.ConfigureTopicResponse")
	proto.RegisterType((*DeleteTopicRequest)(nil), "queue_pb.DeleteTopicRequest")
	proto.RegisterType((*DeleteTopicResponse)(nil), "queue_pb.DeleteTopicResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SeaweedQueue service

type SeaweedQueueClient interface {
	StreamWrite(ctx context.Context, opts ...grpc.CallOption) (SeaweedQueue_StreamWriteClient, error)
	StreamRead(ctx context.Context, in *ReadMessageRequest, opts ...grpc.CallOption) (SeaweedQueue_StreamReadClient, error)
	ConfigureTopic(ctx context.Context, in *ConfigureTopicRequest, opts ...grpc.CallOption) (*ConfigureTopicResponse, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicResponse, error)
}

type seaweedQueueClient struct {
	cc *grpc.ClientConn
}

func NewSeaweedQueueClient(cc *grpc.ClientConn) SeaweedQueueClient {
	return &seaweedQueueClient{cc}
}

func (c *seaweedQueueClient) StreamWrite(ctx context.Context, opts ...grpc.CallOption) (SeaweedQueue_StreamWriteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SeaweedQueue_serviceDesc.Streams[0], c.cc, "/queue_pb.SeaweedQueue/StreamWrite", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedQueueStreamWriteClient{stream}
	return x, nil
}

type SeaweedQueue_StreamWriteClient interface {
	Send(*WriteMessageRequest) error
	Recv() (*WriteMessageResponse, error)
	grpc.ClientStream
}

type seaweedQueueStreamWriteClient struct {
	grpc.ClientStream
}

func (x *seaweedQueueStreamWriteClient) Send(m *WriteMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedQueueStreamWriteClient) Recv() (*WriteMessageResponse, error) {
	m := new(WriteMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedQueueClient) StreamRead(ctx context.Context, in *ReadMessageRequest, opts ...grpc.CallOption) (SeaweedQueue_StreamReadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SeaweedQueue_serviceDesc.Streams[1], c.cc, "/queue_pb.SeaweedQueue/StreamRead", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedQueueStreamReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SeaweedQueue_StreamReadClient interface {
	Recv() (*ReadMessageResponse, error)
	grpc.ClientStream
}

type seaweedQueueStreamReadClient struct {
	grpc.ClientStream
}

func (x *seaweedQueueStreamReadClient) Recv() (*ReadMessageResponse, error) {
	m := new(ReadMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedQueueClient) ConfigureTopic(ctx context.Context, in *ConfigureTopicRequest, opts ...grpc.CallOption) (*ConfigureTopicResponse, error) {
	out := new(ConfigureTopicResponse)
	err := grpc.Invoke(ctx, "/queue_pb.SeaweedQueue/ConfigureTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedQueueClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicResponse, error) {
	out := new(DeleteTopicResponse)
	err := grpc.Invoke(ctx, "/queue_pb.SeaweedQueue/DeleteTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SeaweedQueue service

type SeaweedQueueServer interface {
	StreamWrite(SeaweedQueue_StreamWriteServer) error
	StreamRead(*ReadMessageRequest, SeaweedQueue_StreamReadServer) error
	ConfigureTopic(context.Context, *ConfigureTopicRequest) (*ConfigureTopicResponse, error)
	DeleteTopic(context.Context, *DeleteTopicRequest) (*DeleteTopicResponse, error)
}

func RegisterSeaweedQueueServer(s *grpc.Server, srv SeaweedQueueServer) {
	s.RegisterService(&_SeaweedQueue_serviceDesc, srv)
}

func _SeaweedQueue_StreamWrite_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedQueueServer).StreamWrite(&seaweedQueueStreamWriteServer{stream})
}

type SeaweedQueue_StreamWriteServer interface {
	Send(*WriteMessageResponse) error
	Recv() (*WriteMessageRequest, error)
	grpc.ServerStream
}

type seaweedQueueStreamWriteServer struct {
	grpc.ServerStream
}

func (x *seaweedQueueStreamWriteServer) Send(m *WriteMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedQueueStreamWriteServer) Recv() (*WriteMessageRequest, error) {
	m := new(WriteMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SeaweedQueue_StreamRead_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SeaweedQueueServer).StreamRead(m, &seaweedQueueStreamReadServer{stream})
}

type SeaweedQueue_StreamReadServer interface {
	Send(*ReadMessageResponse) error
	grpc.ServerStream
}

type seaweedQueueStreamReadServer struct {
	grpc.ServerStream
}

func (x *seaweedQueueStreamReadServer) Send(m *ReadMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SeaweedQueue_ConfigureTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedQueueServer).ConfigureTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queue_pb.SeaweedQueue/ConfigureTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedQueueServer).ConfigureTopic(ctx, req.(*ConfigureTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedQueue_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedQueueServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queue_pb.SeaweedQueue/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedQueueServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SeaweedQueue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "queue_pb.SeaweedQueue",
	HandlerType: (*SeaweedQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureTopic",
			Handler:    _SeaweedQueue_ConfigureTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _SeaweedQueue_DeleteTopic_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamWrite",
			Handler:       _SeaweedQueue_StreamWrite_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamRead",
			Handler:       _SeaweedQueue_StreamRead_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "queue.proto",
}

func init() { proto.RegisterFile("queue.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xd1, 0x6e, 0xda, 0x30,
	0x14, 0x86, 0x09, 0x0c, 0xc6, 0x4e, 0xd0, 0x34, 0x19, 0x98, 0x18, 0x1a, 0x10, 0xf9, 0x2a, 0xda,
	0xa4, 0x08, 0x6d, 0x6f, 0x00, 0xed, 0x5d, 0x89, 0xda, 0xd0, 0xaa, 0x52, 0x6f, 0x90, 0x49, 0x0e,
	0x28, 0x22, 0x4d, 0x82, 0xed, 0xb4, 0x6f, 0xda, 0xe7, 0xa9, 0xe2, 0x28, 0x22, 0x29, 0x10, 0xd1,
	0xbb, 0xfc, 0xb1, 0xfd, 0x9d, 0xdf, 0xff, 0x39, 0x06, 0x7d, 0x9f, 0x60, 0x82, 0x56, 0xcc, 0x23,
	0x19, 0x91, 0xb6, 0x12, 0xab, 0x78, 0x4d, 0x9f, 0xa0, 0xfb, 0xc8, 0x7d, 0x89, 0x0b, 0x14, 0x82,
	0x6d, 0xd1, 0xc1, 0x7d, 0x82, 0x42, 0x92, 0x1e, 0x34, 0x65, 0x14, 0xfb, 0xee, 0x40, 0x33, 0x34,
	0xf3, 0x9b, 0x93, 0x09, 0xf2, 0x0b, 0xda, 0xf8, 0x82, 0xa1, 0x5c, 0x85, 0x62, 0x50, 0x37, 0x34,
	0xb3, 0xe1, 0x7c, 0x55, 0xda, 0x16, 0x84, 0xc0, 0x17, 0x8f, 0x49, 0x36, 0x68, 0x18, 0x9a, 0xd9,
	0x71, 0xd4, 0x37, 0x9d, 0x43, 0xaf, 0xcc, 0x16, 0x71, 0x14, 0x0a, 0x4c, 0xe1, 0xc8, 0x79, 0xc4,
	0x73, 0xb8, 0x12, 0xa4, 0x0f, 0x2d, 0xe6, 0xee, 0x0e, 0xe8, 0x26, 0x73, 0x77, 0xb6, 0xa0, 0xd7,
	0x40, 0x1c, 0x64, 0xde, 0xa5, 0xfe, 0x84, 0x64, 0xbc, 0xe8, 0x4f, 0x69, 0x5b, 0xa4, 0xf7, 0x2c,
	0x61, 0x2a, 0xad, 0x7c, 0xf2, 0x9e, 0x36, 0xf4, 0xe7, 0x51, 0xb8, 0xf1, 0xb7, 0x09, 0xc7, 0xfb,
	0xd4, 0x48, 0xb5, 0xcb, 0x09, 0xe8, 0x52, 0x06, 0x2b, 0x81, 0x6e, 0x14, 0x7a, 0x79, 0x01, 0x90,
	0x32, 0x58, 0x66, 0x7f, 0xa8, 0x05, 0x3f, 0x3f, 0xf2, 0xaa, 0xec, 0xd2, 0x3f, 0x40, 0xae, 0x30,
	0x40, 0x79, 0x41, 0x71, 0xfa, 0x17, 0xba, 0xa5, 0xbd, 0x55, 0xe0, 0x7f, 0x6f, 0x75, 0xe8, 0x2c,
	0x91, 0xbd, 0x22, 0x7a, 0x77, 0xe9, 0xc0, 0x10, 0x07, 0xf4, 0xa5, 0xe4, 0xc8, 0x9e, 0x55, 0x5f,
	0xc9, 0xc8, 0xca, 0xe7, 0xc8, 0x3a, 0x31, 0x44, 0xc3, 0xf1, 0xb9, 0xe5, 0xac, 0x28, 0xad, 0x99,
	0xda, 0x54, 0x23, 0x0b, 0x80, 0x8c, 0x99, 0xf6, 0x87, 0xfc, 0x3e, 0x9c, 0x39, 0x6e, 0xfb, 0x70,
	0x74, 0x66, 0x35, 0x07, 0x4e, 0x35, 0xf2, 0x00, 0xdf, 0xcb, 0xe1, 0x91, 0xc9, 0xe1, 0xd0, 0xc9,
	0x36, 0x0d, 0x8d, 0xf3, 0x1b, 0x72, 0x30, 0xb9, 0x01, 0xbd, 0x90, 0x5b, 0xd1, 0xe6, 0x71, 0xf4,
	0x45, 0x9b, 0x27, 0xc2, 0xa6, 0xb5, 0xd9, 0x18, 0x7e, 0x88, 0x2c, 0xd7, 0x8d, 0xb0, 0xdc, 0xc0,
	0xc7, 0x50, 0xce, 0x40, 0x45, 0x7c, 0x9b, 0xbe, 0xcf, 0x75, 0x4b, 0x3d, 0xd3, 0xff, 0xef, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x34, 0x84, 0x96, 0x74, 0xb5, 0x03, 0x00, 0x00,
}
