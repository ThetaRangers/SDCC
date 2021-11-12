// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package registerServer

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

// OperationsClient is the client API for Operations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationsClient interface {
	Register(ctx context.Context, in *RegisterMessage, opts ...grpc.CallOption) (*Cluster, error)
	GetAllNodes(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*NodesIp, error)
}

type operationsClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationsClient(cc grpc.ClientConnInterface) OperationsClient {
	return &operationsClient{cc}
}

func (c *operationsClient) Register(ctx context.Context, in *RegisterMessage, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/registerServer.Operations/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) GetAllNodes(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*NodesIp, error) {
	out := new(NodesIp)
	err := c.cc.Invoke(ctx, "/registerServer.Operations/GetAllNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationsServer is the server API for Operations service.
// All implementations must embed UnimplementedOperationsServer
// for forward compatibility
type OperationsServer interface {
	Register(context.Context, *RegisterMessage) (*Cluster, error)
	GetAllNodes(context.Context, *EmptyMessage) (*NodesIp, error)
	mustEmbedUnimplementedOperationsServer()
}

// UnimplementedOperationsServer must be embedded to have forward compatible implementations.
type UnimplementedOperationsServer struct {
}

func (UnimplementedOperationsServer) Register(context.Context, *RegisterMessage) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedOperationsServer) GetAllNodes(context.Context, *EmptyMessage) (*NodesIp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNodes not implemented")
}
func (UnimplementedOperationsServer) mustEmbedUnimplementedOperationsServer() {}

// UnsafeOperationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationsServer will
// result in compilation errors.
type UnsafeOperationsServer interface {
	mustEmbedUnimplementedOperationsServer()
}

func RegisterOperationsServer(s grpc.ServiceRegistrar, srv OperationsServer) {
	s.RegisterService(&Operations_ServiceDesc, srv)
}

func _Operations_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registerServer.Operations/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).Register(ctx, req.(*RegisterMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_GetAllNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).GetAllNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registerServer.Operations/GetAllNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).GetAllNodes(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Operations_ServiceDesc is the grpc.ServiceDesc for Operations service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Operations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registerServer.Operations",
	HandlerType: (*OperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Operations_Register_Handler,
		},
		{
			MethodName: "GetAllNodes",
			Handler:    _Operations_GetAllNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operations.proto",
}
