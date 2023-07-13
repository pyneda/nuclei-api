// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/service/service.proto

package service

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

// NucleiApiClient is the client API for NucleiApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NucleiApiClient interface {
	Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (NucleiApi_ScanClient, error)
}

type nucleiApiClient struct {
	cc grpc.ClientConnInterface
}

func NewNucleiApiClient(cc grpc.ClientConnInterface) NucleiApiClient {
	return &nucleiApiClient{cc}
}

func (c *nucleiApiClient) Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (NucleiApi_ScanClient, error) {
	stream, err := c.cc.NewStream(ctx, &NucleiApi_ServiceDesc.Streams[0], "/NucleiApi/Scan", opts...)
	if err != nil {
		return nil, err
	}
	x := &nucleiApiScanClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NucleiApi_ScanClient interface {
	Recv() (*ScanResult, error)
	grpc.ClientStream
}

type nucleiApiScanClient struct {
	grpc.ClientStream
}

func (x *nucleiApiScanClient) Recv() (*ScanResult, error) {
	m := new(ScanResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NucleiApiServer is the server API for NucleiApi service.
// All implementations must embed UnimplementedNucleiApiServer
// for forward compatibility
type NucleiApiServer interface {
	Scan(*ScanRequest, NucleiApi_ScanServer) error
	mustEmbedUnimplementedNucleiApiServer()
}

// UnimplementedNucleiApiServer must be embedded to have forward compatible implementations.
type UnimplementedNucleiApiServer struct {
}

func (UnimplementedNucleiApiServer) Scan(*ScanRequest, NucleiApi_ScanServer) error {
	return status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (UnimplementedNucleiApiServer) mustEmbedUnimplementedNucleiApiServer() {}

// UnsafeNucleiApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NucleiApiServer will
// result in compilation errors.
type UnsafeNucleiApiServer interface {
	mustEmbedUnimplementedNucleiApiServer()
}

func RegisterNucleiApiServer(s grpc.ServiceRegistrar, srv NucleiApiServer) {
	s.RegisterService(&NucleiApi_ServiceDesc, srv)
}

func _NucleiApi_Scan_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ScanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NucleiApiServer).Scan(m, &nucleiApiScanServer{stream})
}

type NucleiApi_ScanServer interface {
	Send(*ScanResult) error
	grpc.ServerStream
}

type nucleiApiScanServer struct {
	grpc.ServerStream
}

func (x *nucleiApiScanServer) Send(m *ScanResult) error {
	return x.ServerStream.SendMsg(m)
}

// NucleiApi_ServiceDesc is the grpc.ServiceDesc for NucleiApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NucleiApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NucleiApi",
	HandlerType: (*NucleiApiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Scan",
			Handler:       _NucleiApi_Scan_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/service/service.proto",
}