// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: web3eye/ranker/v1/token/token.proto

package token

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

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
	GetTokenOnly(ctx context.Context, in *GetTokenOnlyRequest, opts ...grpc.CallOption) (*GetTokenOnlyResponse, error)
	GetTokens(ctx context.Context, in *GetTokensRequest, opts ...grpc.CallOption) (*GetTokensResponse, error)
	CountTokens(ctx context.Context, in *CountTokensRequest, opts ...grpc.CallOption) (*CountTokensResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, "/ranker.v1.token.Manager/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetTokenOnly(ctx context.Context, in *GetTokenOnlyRequest, opts ...grpc.CallOption) (*GetTokenOnlyResponse, error) {
	out := new(GetTokenOnlyResponse)
	err := c.cc.Invoke(ctx, "/ranker.v1.token.Manager/GetTokenOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetTokens(ctx context.Context, in *GetTokensRequest, opts ...grpc.CallOption) (*GetTokensResponse, error) {
	out := new(GetTokensResponse)
	err := c.cc.Invoke(ctx, "/ranker.v1.token.Manager/GetTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountTokens(ctx context.Context, in *CountTokensRequest, opts ...grpc.CallOption) (*CountTokensResponse, error) {
	out := new(CountTokensResponse)
	err := c.cc.Invoke(ctx, "/ranker.v1.token.Manager/CountTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
	GetTokenOnly(context.Context, *GetTokenOnlyRequest) (*GetTokenOnlyResponse, error)
	GetTokens(context.Context, *GetTokensRequest) (*GetTokensResponse, error)
	CountTokens(context.Context, *CountTokensRequest) (*CountTokensResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedManagerServer) GetTokenOnly(context.Context, *GetTokenOnlyRequest) (*GetTokenOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenOnly not implemented")
}
func (UnimplementedManagerServer) GetTokens(context.Context, *GetTokensRequest) (*GetTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokens not implemented")
}
func (UnimplementedManagerServer) CountTokens(context.Context, *CountTokensRequest) (*CountTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountTokens not implemented")
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

func _Manager_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ranker.v1.token.Manager/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetTokenOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetTokenOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ranker.v1.token.Manager/GetTokenOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetTokenOnly(ctx, req.(*GetTokenOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ranker.v1.token.Manager/GetTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetTokens(ctx, req.(*GetTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ranker.v1.token.Manager/CountTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountTokens(ctx, req.(*CountTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ranker.v1.token.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _Manager_GetToken_Handler,
		},
		{
			MethodName: "GetTokenOnly",
			Handler:    _Manager_GetTokenOnly_Handler,
		},
		{
			MethodName: "GetTokens",
			Handler:    _Manager_GetTokens_Handler,
		},
		{
			MethodName: "CountTokens",
			Handler:    _Manager_CountTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web3eye/ranker/v1/token/token.proto",
}
