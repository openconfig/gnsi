// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: github.com/openconfig/gnsi/acctz/acctz.proto

package acctz

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AcctzClient is the client API for Acctz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AcctzClient interface {
	RecordSubscribe(ctx context.Context, opts ...grpc.CallOption) (Acctz_RecordSubscribeClient, error)
}

type acctzClient struct {
	cc grpc.ClientConnInterface
}

func NewAcctzClient(cc grpc.ClientConnInterface) AcctzClient {
	return &acctzClient{cc}
}

func (c *acctzClient) RecordSubscribe(ctx context.Context, opts ...grpc.CallOption) (Acctz_RecordSubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Acctz_ServiceDesc.Streams[0], "/gnsi.acctz.v1.Acctz/RecordSubscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &acctzRecordSubscribeClient{stream}
	return x, nil
}

type Acctz_RecordSubscribeClient interface {
	Send(*RecordRequest) error
	Recv() (*RecordResponse, error)
	grpc.ClientStream
}

type acctzRecordSubscribeClient struct {
	grpc.ClientStream
}

func (x *acctzRecordSubscribeClient) Send(m *RecordRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *acctzRecordSubscribeClient) Recv() (*RecordResponse, error) {
	m := new(RecordResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AcctzServer is the server API for Acctz service.
// All implementations must embed UnimplementedAcctzServer
// for forward compatibility
type AcctzServer interface {
	RecordSubscribe(Acctz_RecordSubscribeServer) error
	mustEmbedUnimplementedAcctzServer()
}

// UnimplementedAcctzServer must be embedded to have forward compatible implementations.
type UnimplementedAcctzServer struct {
}

func (UnimplementedAcctzServer) RecordSubscribe(Acctz_RecordSubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordSubscribe not implemented")
}
func (UnimplementedAcctzServer) mustEmbedUnimplementedAcctzServer() {}

// UnsafeAcctzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AcctzServer will
// result in compilation errors.
type UnsafeAcctzServer interface {
	mustEmbedUnimplementedAcctzServer()
}

func RegisterAcctzServer(s grpc.ServiceRegistrar, srv AcctzServer) {
	s.RegisterService(&Acctz_ServiceDesc, srv)
}

func _Acctz_RecordSubscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AcctzServer).RecordSubscribe(&acctzRecordSubscribeServer{stream})
}

type Acctz_RecordSubscribeServer interface {
	Send(*RecordResponse) error
	Recv() (*RecordRequest, error)
	grpc.ServerStream
}

type acctzRecordSubscribeServer struct {
	grpc.ServerStream
}

func (x *acctzRecordSubscribeServer) Send(m *RecordResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *acctzRecordSubscribeServer) Recv() (*RecordRequest, error) {
	m := new(RecordRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Acctz_ServiceDesc is the grpc.ServiceDesc for Acctz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Acctz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnsi.acctz.v1.Acctz",
	HandlerType: (*AcctzServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordSubscribe",
			Handler:       _Acctz_RecordSubscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/openconfig/gnsi/acctz/acctz.proto",
}

// AcctzStreamClient is the client API for AcctzStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AcctzStreamClient interface {
	RecordSubscribe(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (AcctzStream_RecordSubscribeClient, error)
}

type acctzStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewAcctzStreamClient(cc grpc.ClientConnInterface) AcctzStreamClient {
	return &acctzStreamClient{cc}
}

func (c *acctzStreamClient) RecordSubscribe(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (AcctzStream_RecordSubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &AcctzStream_ServiceDesc.Streams[0], "/gnsi.acctz.v1.AcctzStream/RecordSubscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &acctzStreamRecordSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AcctzStream_RecordSubscribeClient interface {
	Recv() (*RecordResponse, error)
	grpc.ClientStream
}

type acctzStreamRecordSubscribeClient struct {
	grpc.ClientStream
}

func (x *acctzStreamRecordSubscribeClient) Recv() (*RecordResponse, error) {
	m := new(RecordResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AcctzStreamServer is the server API for AcctzStream service.
// All implementations must embed UnimplementedAcctzStreamServer
// for forward compatibility
type AcctzStreamServer interface {
	RecordSubscribe(*RecordRequest, AcctzStream_RecordSubscribeServer) error
	mustEmbedUnimplementedAcctzStreamServer()
}

// UnimplementedAcctzStreamServer must be embedded to have forward compatible implementations.
type UnimplementedAcctzStreamServer struct {
}

func (UnimplementedAcctzStreamServer) RecordSubscribe(*RecordRequest, AcctzStream_RecordSubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordSubscribe not implemented")
}
func (UnimplementedAcctzStreamServer) mustEmbedUnimplementedAcctzStreamServer() {}

// UnsafeAcctzStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AcctzStreamServer will
// result in compilation errors.
type UnsafeAcctzStreamServer interface {
	mustEmbedUnimplementedAcctzStreamServer()
}

func RegisterAcctzStreamServer(s grpc.ServiceRegistrar, srv AcctzStreamServer) {
	s.RegisterService(&AcctzStream_ServiceDesc, srv)
}

func _AcctzStream_RecordSubscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RecordRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AcctzStreamServer).RecordSubscribe(m, &acctzStreamRecordSubscribeServer{stream})
}

type AcctzStream_RecordSubscribeServer interface {
	Send(*RecordResponse) error
	grpc.ServerStream
}

type acctzStreamRecordSubscribeServer struct {
	grpc.ServerStream
}

func (x *acctzStreamRecordSubscribeServer) Send(m *RecordResponse) error {
	return x.ServerStream.SendMsg(m)
}

// AcctzStream_ServiceDesc is the grpc.ServiceDesc for AcctzStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AcctzStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnsi.acctz.v1.AcctzStream",
	HandlerType: (*AcctzStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordSubscribe",
			Handler:       _AcctzStream_RecordSubscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/openconfig/gnsi/acctz/acctz.proto",
}
