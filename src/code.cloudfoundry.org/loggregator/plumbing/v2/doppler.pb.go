// Code generated by protoc-gen-go.
// source: doppler.proto
// DO NOT EDIT!

/*
Package loggregator_v2 is a generated protocol buffer package.

It is generated from these files:
	doppler.proto
	egress.proto
	egress_query.proto
	envelope.proto
	ingress.proto

It has these top-level messages:
	SenderResponse
	EgressRequest
	Filter
	LogFilter
	ContainerMetricRequest
	QueryResponse
	Envelope
	Value
	Log
	Counter
	Gauge
	GaugeValue
	Timer
	IngressResponse
	EnvelopeBatch
	BatchSenderResponse
*/
package loggregator_v2

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

type SenderResponse struct {
}

func (m *SenderResponse) Reset()                    { *m = SenderResponse{} }
func (m *SenderResponse) String() string            { return proto.CompactTextString(m) }
func (*SenderResponse) ProtoMessage()               {}
func (*SenderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*SenderResponse)(nil), "loggregator.v2.SenderResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DopplerIngress service

type DopplerIngressClient interface {
	Sender(ctx context.Context, opts ...grpc.CallOption) (DopplerIngress_SenderClient, error)
	BatchSender(ctx context.Context, opts ...grpc.CallOption) (DopplerIngress_BatchSenderClient, error)
}

type dopplerIngressClient struct {
	cc *grpc.ClientConn
}

func NewDopplerIngressClient(cc *grpc.ClientConn) DopplerIngressClient {
	return &dopplerIngressClient{cc}
}

func (c *dopplerIngressClient) Sender(ctx context.Context, opts ...grpc.CallOption) (DopplerIngress_SenderClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DopplerIngress_serviceDesc.Streams[0], c.cc, "/loggregator.v2.DopplerIngress/Sender", opts...)
	if err != nil {
		return nil, err
	}
	x := &dopplerIngressSenderClient{stream}
	return x, nil
}

type DopplerIngress_SenderClient interface {
	Send(*Envelope) error
	CloseAndRecv() (*SenderResponse, error)
	grpc.ClientStream
}

type dopplerIngressSenderClient struct {
	grpc.ClientStream
}

func (x *dopplerIngressSenderClient) Send(m *Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dopplerIngressSenderClient) CloseAndRecv() (*SenderResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SenderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dopplerIngressClient) BatchSender(ctx context.Context, opts ...grpc.CallOption) (DopplerIngress_BatchSenderClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DopplerIngress_serviceDesc.Streams[1], c.cc, "/loggregator.v2.DopplerIngress/BatchSender", opts...)
	if err != nil {
		return nil, err
	}
	x := &dopplerIngressBatchSenderClient{stream}
	return x, nil
}

type DopplerIngress_BatchSenderClient interface {
	Send(*EnvelopeBatch) error
	CloseAndRecv() (*BatchSenderResponse, error)
	grpc.ClientStream
}

type dopplerIngressBatchSenderClient struct {
	grpc.ClientStream
}

func (x *dopplerIngressBatchSenderClient) Send(m *EnvelopeBatch) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dopplerIngressBatchSenderClient) CloseAndRecv() (*BatchSenderResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BatchSenderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DopplerIngress service

type DopplerIngressServer interface {
	Sender(DopplerIngress_SenderServer) error
	BatchSender(DopplerIngress_BatchSenderServer) error
}

func RegisterDopplerIngressServer(s *grpc.Server, srv DopplerIngressServer) {
	s.RegisterService(&_DopplerIngress_serviceDesc, srv)
}

func _DopplerIngress_Sender_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DopplerIngressServer).Sender(&dopplerIngressSenderServer{stream})
}

type DopplerIngress_SenderServer interface {
	SendAndClose(*SenderResponse) error
	Recv() (*Envelope, error)
	grpc.ServerStream
}

type dopplerIngressSenderServer struct {
	grpc.ServerStream
}

func (x *dopplerIngressSenderServer) SendAndClose(m *SenderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dopplerIngressSenderServer) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DopplerIngress_BatchSender_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DopplerIngressServer).BatchSender(&dopplerIngressBatchSenderServer{stream})
}

type DopplerIngress_BatchSenderServer interface {
	SendAndClose(*BatchSenderResponse) error
	Recv() (*EnvelopeBatch, error)
	grpc.ServerStream
}

type dopplerIngressBatchSenderServer struct {
	grpc.ServerStream
}

func (x *dopplerIngressBatchSenderServer) SendAndClose(m *BatchSenderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dopplerIngressBatchSenderServer) Recv() (*EnvelopeBatch, error) {
	m := new(EnvelopeBatch)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DopplerIngress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggregator.v2.DopplerIngress",
	HandlerType: (*DopplerIngressServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sender",
			Handler:       _DopplerIngress_Sender_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BatchSender",
			Handler:       _DopplerIngress_BatchSender_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "doppler.proto",
}

func init() { proto.RegisterFile("doppler.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xc9, 0x2f, 0x28,
	0xc8, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xc9, 0x4f, 0x4f, 0x2f,
	0x4a, 0x4d, 0x4f, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x33, 0x92, 0xe2, 0x4b, 0xcd, 0x2b, 0x4b, 0xcd,
	0xc9, 0x2f, 0x48, 0x85, 0xc8, 0x4b, 0xf1, 0x66, 0xe6, 0xa5, 0x17, 0xa5, 0x16, 0x17, 0x43, 0xb8,
	0x4a, 0x02, 0x5c, 0x7c, 0xc1, 0xa9, 0x79, 0x29, 0xa9, 0x45, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0x46, 0xeb, 0x19, 0xb9, 0xf8, 0x5c, 0x20, 0x46, 0x7a, 0x42, 0x94, 0x0a, 0xb9, 0x71,
	0xb1, 0x41, 0x14, 0x09, 0x49, 0xe8, 0xa1, 0x1a, 0xaf, 0xe7, 0x0a, 0x35, 0x5d, 0x4a, 0x0e, 0x5d,
	0x06, 0xd5, 0x58, 0x25, 0x06, 0x0d, 0x46, 0xa1, 0x50, 0x2e, 0x6e, 0xa7, 0xc4, 0x92, 0xe4, 0x0c,
	0xa8, 0x61, 0xb2, 0xb8, 0x0c, 0x03, 0x2b, 0x92, 0x52, 0x46, 0x97, 0x46, 0xd2, 0x8b, 0x6c, 0x6c,
	0x12, 0x1b, 0xd8, 0x2b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xc2, 0x57, 0xb9, 0x0a,
	0x01, 0x00, 0x00,
}
