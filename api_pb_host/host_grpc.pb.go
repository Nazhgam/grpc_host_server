// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: protos/host.proto

package api_pb_host

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

const (
	Hosts_AddQuestion_FullMethodName = "/mypackage.Hosts/AddQuestion"
)

// HostsClient is the client API for Hosts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostsClient interface {
	AddQuestion(ctx context.Context, in *HostReqeust, opts ...grpc.CallOption) (*HostResponse, error)
}

type hostsClient struct {
	cc grpc.ClientConnInterface
}

func NewHostsClient(cc grpc.ClientConnInterface) HostsClient {
	return &hostsClient{cc}
}

func (c *hostsClient) AddQuestion(ctx context.Context, in *HostReqeust, opts ...grpc.CallOption) (*HostResponse, error) {
	out := new(HostResponse)
	err := c.cc.Invoke(ctx, Hosts_AddQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostsServer is the server API for Hosts service.
// All implementations must embed UnimplementedHostsServer
// for forward compatibility
type HostsServer interface {
	AddQuestion(context.Context, *HostReqeust) (*HostResponse, error)
	mustEmbedUnimplementedHostsServer()
}

// UnimplementedHostsServer must be embedded to have forward compatible implementations.
type UnimplementedHostsServer struct {
}

func (UnimplementedHostsServer) AddQuestion(context.Context, *HostReqeust) (*HostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddQuestion not implemented")
}
func (UnimplementedHostsServer) mustEmbedUnimplementedHostsServer() {}

// UnsafeHostsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostsServer will
// result in compilation errors.
type UnsafeHostsServer interface {
	mustEmbedUnimplementedHostsServer()
}

func RegisterHostsServer(s grpc.ServiceRegistrar, srv HostsServer) {
	s.RegisterService(&Hosts_ServiceDesc, srv)
}

func _Hosts_AddQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostReqeust)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostsServer).AddQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hosts_AddQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostsServer).AddQuestion(ctx, req.(*HostReqeust))
	}
	return interceptor(ctx, in, info, handler)
}

// Hosts_ServiceDesc is the grpc.ServiceDesc for Hosts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hosts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mypackage.Hosts",
	HandlerType: (*HostsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddQuestion",
			Handler:    _Hosts_AddQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/host.proto",
}
