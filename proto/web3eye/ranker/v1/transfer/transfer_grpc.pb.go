// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: web3eye/ranker/v1/transfer/transfer.proto

package transfer

import (
	context "context"
	transfer "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
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
	GetTransfer(ctx context.Context, in *transfer.GetTransferRequest, opts ...grpc.CallOption) (*transfer.GetTransferResponse, error)
	GetTransferOnly(ctx context.Context, in *transfer.GetTransferOnlyRequest, opts ...grpc.CallOption) (*transfer.GetTransferOnlyResponse, error)
	GetTransfers(ctx context.Context, in *transfer.GetTransfersRequest, opts ...grpc.CallOption) (*transfer.GetTransfersResponse, error)
	CountTransfers(ctx context.Context, in *transfer.CountTransfersRequest, opts ...grpc.CallOption) (*transfer.CountTransfersResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) GetTransfer(ctx context.Context, in *transfer.GetTransferRequest, opts ...grpc.CallOption) (*transfer.GetTransferResponse, error) {
	out := new(transfer.GetTransferResponse)
	err := c.cc.Invoke(ctx, "/ranker.v1.transfer.Manager/GetTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetTransferOnly(ctx context.Context, in *transfer.GetTransferOnlyRequest, opts ...grpc.CallOption) (*transfer.GetTransferOnlyResponse, error) {
	out := new(transfer.GetTransferOnlyResponse)
	err := c.cc.Invoke(ctx, "/ranker.v1.transfer.Manager/GetTransferOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetTransfers(ctx context.Context, in *transfer.GetTransfersRequest, opts ...grpc.CallOption) (*transfer.GetTransfersResponse, error) {
	out := new(transfer.GetTransfersResponse)
	err := c.cc.Invoke(ctx, "/ranker.v1.transfer.Manager/GetTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountTransfers(ctx context.Context, in *transfer.CountTransfersRequest, opts ...grpc.CallOption) (*transfer.CountTransfersResponse, error) {
	out := new(transfer.CountTransfersResponse)
	err := c.cc.Invoke(ctx, "/ranker.v1.transfer.Manager/CountTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	GetTransfer(context.Context, *transfer.GetTransferRequest) (*transfer.GetTransferResponse, error)
	GetTransferOnly(context.Context, *transfer.GetTransferOnlyRequest) (*transfer.GetTransferOnlyResponse, error)
	GetTransfers(context.Context, *transfer.GetTransfersRequest) (*transfer.GetTransfersResponse, error)
	CountTransfers(context.Context, *transfer.CountTransfersRequest) (*transfer.CountTransfersResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) GetTransfer(context.Context, *transfer.GetTransferRequest) (*transfer.GetTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransfer not implemented")
}
func (UnimplementedManagerServer) GetTransferOnly(context.Context, *transfer.GetTransferOnlyRequest) (*transfer.GetTransferOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransferOnly not implemented")
}
func (UnimplementedManagerServer) GetTransfers(context.Context, *transfer.GetTransfersRequest) (*transfer.GetTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransfers not implemented")
}
func (UnimplementedManagerServer) CountTransfers(context.Context, *transfer.CountTransfersRequest) (*transfer.CountTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountTransfers not implemented")
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

func _Manager_GetTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transfer.GetTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ranker.v1.transfer.Manager/GetTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetTransfer(ctx, req.(*transfer.GetTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetTransferOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transfer.GetTransferOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetTransferOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ranker.v1.transfer.Manager/GetTransferOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetTransferOnly(ctx, req.(*transfer.GetTransferOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transfer.GetTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ranker.v1.transfer.Manager/GetTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetTransfers(ctx, req.(*transfer.GetTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transfer.CountTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ranker.v1.transfer.Manager/CountTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountTransfers(ctx, req.(*transfer.CountTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ranker.v1.transfer.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransfer",
			Handler:    _Manager_GetTransfer_Handler,
		},
		{
			MethodName: "GetTransferOnly",
			Handler:    _Manager_GetTransferOnly_Handler,
		},
		{
			MethodName: "GetTransfers",
			Handler:    _Manager_GetTransfers_Handler,
		},
		{
			MethodName: "CountTransfers",
			Handler:    _Manager_CountTransfers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web3eye/ranker/v1/transfer/transfer.proto",
}