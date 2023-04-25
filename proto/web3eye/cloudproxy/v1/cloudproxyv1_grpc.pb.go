// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: web3eye/cloudproxy/v1/cloudproxyv1.proto

package v1

import (
	context "context"
	web3eye "github.com/web3eye-io/Web3Eye/proto/web3eye"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*web3eye.VersionResponse, error)
	GrpcProxyChannel(ctx context.Context, opts ...grpc.CallOption) (Manager_GrpcProxyChannelClient, error)
	GrpcProxy(ctx context.Context, in *GrpcProxyRequest, opts ...grpc.CallOption) (*GrpcProxyResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*web3eye.VersionResponse, error) {
	out := new(web3eye.VersionResponse)
	err := c.cc.Invoke(ctx, "/cloudproxy.v1.Manager/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GrpcProxyChannel(ctx context.Context, opts ...grpc.CallOption) (Manager_GrpcProxyChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &Manager_ServiceDesc.Streams[0], "/cloudproxy.v1.Manager/GrpcProxyChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &managerGrpcProxyChannelClient{stream}
	return x, nil
}

type Manager_GrpcProxyChannelClient interface {
	Send(*ToGrpcProxy) error
	Recv() (*FromGrpcProxy, error)
	grpc.ClientStream
}

type managerGrpcProxyChannelClient struct {
	grpc.ClientStream
}

func (x *managerGrpcProxyChannelClient) Send(m *ToGrpcProxy) error {
	return x.ClientStream.SendMsg(m)
}

func (x *managerGrpcProxyChannelClient) Recv() (*FromGrpcProxy, error) {
	m := new(FromGrpcProxy)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *managerClient) GrpcProxy(ctx context.Context, in *GrpcProxyRequest, opts ...grpc.CallOption) (*GrpcProxyResponse, error) {
	out := new(GrpcProxyResponse)
	err := c.cc.Invoke(ctx, "/cloudproxy.v1.Manager/GrpcProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	Version(context.Context, *emptypb.Empty) (*web3eye.VersionResponse, error)
	GrpcProxyChannel(Manager_GrpcProxyChannelServer) error
	GrpcProxy(context.Context, *GrpcProxyRequest) (*GrpcProxyResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) Version(context.Context, *emptypb.Empty) (*web3eye.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedManagerServer) GrpcProxyChannel(Manager_GrpcProxyChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method GrpcProxyChannel not implemented")
}
func (UnimplementedManagerServer) GrpcProxy(context.Context, *GrpcProxyRequest) (*GrpcProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrpcProxy not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudproxy.v1.Manager/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GrpcProxyChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ManagerServer).GrpcProxyChannel(&managerGrpcProxyChannelServer{stream})
}

type Manager_GrpcProxyChannelServer interface {
	Send(*FromGrpcProxy) error
	Recv() (*ToGrpcProxy, error)
	grpc.ServerStream
}

type managerGrpcProxyChannelServer struct {
	grpc.ServerStream
}

func (x *managerGrpcProxyChannelServer) Send(m *FromGrpcProxy) error {
	return x.ServerStream.SendMsg(m)
}

func (x *managerGrpcProxyChannelServer) Recv() (*ToGrpcProxy, error) {
	m := new(ToGrpcProxy)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Manager_GrpcProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GrpcProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudproxy.v1.Manager/GrpcProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GrpcProxy(ctx, req.(*GrpcProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloudproxy.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Manager_Version_Handler,
		},
		{
			MethodName: "GrpcProxy",
			Handler:    _Manager_GrpcProxy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GrpcProxyChannel",
			Handler:       _Manager_GrpcProxyChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "web3eye/cloudproxy/v1/cloudproxyv1.proto",
}