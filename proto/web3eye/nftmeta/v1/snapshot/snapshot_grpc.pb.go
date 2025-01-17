// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: web3eye/nftmeta/v1/snapshot/snapshot.proto

package snapshot

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
	CreateSnapshot(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error)
	CreateSnapshots(ctx context.Context, in *CreateSnapshotsRequest, opts ...grpc.CallOption) (*CreateSnapshotsResponse, error)
	UpdateSnapshot(ctx context.Context, in *UpdateSnapshotRequest, opts ...grpc.CallOption) (*UpdateSnapshotResponse, error)
	GetSnapshot(ctx context.Context, in *GetSnapshotRequest, opts ...grpc.CallOption) (*GetSnapshotResponse, error)
	GetSnapshotOnly(ctx context.Context, in *GetSnapshotOnlyRequest, opts ...grpc.CallOption) (*GetSnapshotOnlyResponse, error)
	GetSnapshots(ctx context.Context, in *GetSnapshotsRequest, opts ...grpc.CallOption) (*GetSnapshotsResponse, error)
	ExistSnapshot(ctx context.Context, in *ExistSnapshotRequest, opts ...grpc.CallOption) (*ExistSnapshotResponse, error)
	ExistSnapshotConds(ctx context.Context, in *ExistSnapshotCondsRequest, opts ...grpc.CallOption) (*ExistSnapshotCondsResponse, error)
	CountSnapshots(ctx context.Context, in *CountSnapshotsRequest, opts ...grpc.CallOption) (*CountSnapshotsResponse, error)
	DeleteSnapshot(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*DeleteSnapshotResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateSnapshot(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error) {
	out := new(CreateSnapshotResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.snapshot.Manager/CreateSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateSnapshots(ctx context.Context, in *CreateSnapshotsRequest, opts ...grpc.CallOption) (*CreateSnapshotsResponse, error) {
	out := new(CreateSnapshotsResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.snapshot.Manager/CreateSnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateSnapshot(ctx context.Context, in *UpdateSnapshotRequest, opts ...grpc.CallOption) (*UpdateSnapshotResponse, error) {
	out := new(UpdateSnapshotResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.snapshot.Manager/UpdateSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetSnapshot(ctx context.Context, in *GetSnapshotRequest, opts ...grpc.CallOption) (*GetSnapshotResponse, error) {
	out := new(GetSnapshotResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.snapshot.Manager/GetSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetSnapshotOnly(ctx context.Context, in *GetSnapshotOnlyRequest, opts ...grpc.CallOption) (*GetSnapshotOnlyResponse, error) {
	out := new(GetSnapshotOnlyResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.snapshot.Manager/GetSnapshotOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetSnapshots(ctx context.Context, in *GetSnapshotsRequest, opts ...grpc.CallOption) (*GetSnapshotsResponse, error) {
	out := new(GetSnapshotsResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.snapshot.Manager/GetSnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistSnapshot(ctx context.Context, in *ExistSnapshotRequest, opts ...grpc.CallOption) (*ExistSnapshotResponse, error) {
	out := new(ExistSnapshotResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.snapshot.Manager/ExistSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistSnapshotConds(ctx context.Context, in *ExistSnapshotCondsRequest, opts ...grpc.CallOption) (*ExistSnapshotCondsResponse, error) {
	out := new(ExistSnapshotCondsResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.snapshot.Manager/ExistSnapshotConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountSnapshots(ctx context.Context, in *CountSnapshotsRequest, opts ...grpc.CallOption) (*CountSnapshotsResponse, error) {
	out := new(CountSnapshotsResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.snapshot.Manager/CountSnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteSnapshot(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*DeleteSnapshotResponse, error) {
	out := new(DeleteSnapshotResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.snapshot.Manager/DeleteSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateSnapshot(context.Context, *CreateSnapshotRequest) (*CreateSnapshotResponse, error)
	CreateSnapshots(context.Context, *CreateSnapshotsRequest) (*CreateSnapshotsResponse, error)
	UpdateSnapshot(context.Context, *UpdateSnapshotRequest) (*UpdateSnapshotResponse, error)
	GetSnapshot(context.Context, *GetSnapshotRequest) (*GetSnapshotResponse, error)
	GetSnapshotOnly(context.Context, *GetSnapshotOnlyRequest) (*GetSnapshotOnlyResponse, error)
	GetSnapshots(context.Context, *GetSnapshotsRequest) (*GetSnapshotsResponse, error)
	ExistSnapshot(context.Context, *ExistSnapshotRequest) (*ExistSnapshotResponse, error)
	ExistSnapshotConds(context.Context, *ExistSnapshotCondsRequest) (*ExistSnapshotCondsResponse, error)
	CountSnapshots(context.Context, *CountSnapshotsRequest) (*CountSnapshotsResponse, error)
	DeleteSnapshot(context.Context, *DeleteSnapshotRequest) (*DeleteSnapshotResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateSnapshot(context.Context, *CreateSnapshotRequest) (*CreateSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSnapshot not implemented")
}
func (UnimplementedManagerServer) CreateSnapshots(context.Context, *CreateSnapshotsRequest) (*CreateSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSnapshots not implemented")
}
func (UnimplementedManagerServer) UpdateSnapshot(context.Context, *UpdateSnapshotRequest) (*UpdateSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSnapshot not implemented")
}
func (UnimplementedManagerServer) GetSnapshot(context.Context, *GetSnapshotRequest) (*GetSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnapshot not implemented")
}
func (UnimplementedManagerServer) GetSnapshotOnly(context.Context, *GetSnapshotOnlyRequest) (*GetSnapshotOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnapshotOnly not implemented")
}
func (UnimplementedManagerServer) GetSnapshots(context.Context, *GetSnapshotsRequest) (*GetSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnapshots not implemented")
}
func (UnimplementedManagerServer) ExistSnapshot(context.Context, *ExistSnapshotRequest) (*ExistSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistSnapshot not implemented")
}
func (UnimplementedManagerServer) ExistSnapshotConds(context.Context, *ExistSnapshotCondsRequest) (*ExistSnapshotCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistSnapshotConds not implemented")
}
func (UnimplementedManagerServer) CountSnapshots(context.Context, *CountSnapshotsRequest) (*CountSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountSnapshots not implemented")
}
func (UnimplementedManagerServer) DeleteSnapshot(context.Context, *DeleteSnapshotRequest) (*DeleteSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSnapshot not implemented")
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

func _Manager_CreateSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.snapshot.Manager/CreateSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateSnapshot(ctx, req.(*CreateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.snapshot.Manager/CreateSnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateSnapshots(ctx, req.(*CreateSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.snapshot.Manager/UpdateSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateSnapshot(ctx, req.(*UpdateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.snapshot.Manager/GetSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetSnapshot(ctx, req.(*GetSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetSnapshotOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnapshotOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetSnapshotOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.snapshot.Manager/GetSnapshotOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetSnapshotOnly(ctx, req.(*GetSnapshotOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.snapshot.Manager/GetSnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetSnapshots(ctx, req.(*GetSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.snapshot.Manager/ExistSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistSnapshot(ctx, req.(*ExistSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistSnapshotConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistSnapshotCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistSnapshotConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.snapshot.Manager/ExistSnapshotConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistSnapshotConds(ctx, req.(*ExistSnapshotCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.snapshot.Manager/CountSnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountSnapshots(ctx, req.(*CountSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.snapshot.Manager/DeleteSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteSnapshot(ctx, req.(*DeleteSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nftmeta.v1.snapshot.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSnapshot",
			Handler:    _Manager_CreateSnapshot_Handler,
		},
		{
			MethodName: "CreateSnapshots",
			Handler:    _Manager_CreateSnapshots_Handler,
		},
		{
			MethodName: "UpdateSnapshot",
			Handler:    _Manager_UpdateSnapshot_Handler,
		},
		{
			MethodName: "GetSnapshot",
			Handler:    _Manager_GetSnapshot_Handler,
		},
		{
			MethodName: "GetSnapshotOnly",
			Handler:    _Manager_GetSnapshotOnly_Handler,
		},
		{
			MethodName: "GetSnapshots",
			Handler:    _Manager_GetSnapshots_Handler,
		},
		{
			MethodName: "ExistSnapshot",
			Handler:    _Manager_ExistSnapshot_Handler,
		},
		{
			MethodName: "ExistSnapshotConds",
			Handler:    _Manager_ExistSnapshotConds_Handler,
		},
		{
			MethodName: "CountSnapshots",
			Handler:    _Manager_CountSnapshots_Handler,
		},
		{
			MethodName: "DeleteSnapshot",
			Handler:    _Manager_DeleteSnapshot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web3eye/nftmeta/v1/snapshot/snapshot.proto",
}
