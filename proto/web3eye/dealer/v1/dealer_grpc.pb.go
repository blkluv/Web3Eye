// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: web3eye/dealer/v1/dealer.proto

package v1

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
	GetSnapshots(ctx context.Context, in *GetSnapshotsRequest, opts ...grpc.CallOption) (*GetSnapshotsResponse, error)
	CreateBackup(ctx context.Context, in *CreateBackupRequest, opts ...grpc.CallOption) (*CreateBackupResponse, error)
	GetBackups(ctx context.Context, in *GetBackupsRequest, opts ...grpc.CallOption) (*GetBackupsResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateSnapshot(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error) {
	out := new(CreateSnapshotResponse)
	err := c.cc.Invoke(ctx, "/dealer.v1.Manager/CreateSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetSnapshots(ctx context.Context, in *GetSnapshotsRequest, opts ...grpc.CallOption) (*GetSnapshotsResponse, error) {
	out := new(GetSnapshotsResponse)
	err := c.cc.Invoke(ctx, "/dealer.v1.Manager/GetSnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateBackup(ctx context.Context, in *CreateBackupRequest, opts ...grpc.CallOption) (*CreateBackupResponse, error) {
	out := new(CreateBackupResponse)
	err := c.cc.Invoke(ctx, "/dealer.v1.Manager/CreateBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetBackups(ctx context.Context, in *GetBackupsRequest, opts ...grpc.CallOption) (*GetBackupsResponse, error) {
	out := new(GetBackupsResponse)
	err := c.cc.Invoke(ctx, "/dealer.v1.Manager/GetBackups", in, out, opts...)
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
	GetSnapshots(context.Context, *GetSnapshotsRequest) (*GetSnapshotsResponse, error)
	CreateBackup(context.Context, *CreateBackupRequest) (*CreateBackupResponse, error)
	GetBackups(context.Context, *GetBackupsRequest) (*GetBackupsResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateSnapshot(context.Context, *CreateSnapshotRequest) (*CreateSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSnapshot not implemented")
}
func (UnimplementedManagerServer) GetSnapshots(context.Context, *GetSnapshotsRequest) (*GetSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnapshots not implemented")
}
func (UnimplementedManagerServer) CreateBackup(context.Context, *CreateBackupRequest) (*CreateBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBackup not implemented")
}
func (UnimplementedManagerServer) GetBackups(context.Context, *GetBackupsRequest) (*GetBackupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBackups not implemented")
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
		FullMethod: "/dealer.v1.Manager/CreateSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateSnapshot(ctx, req.(*CreateSnapshotRequest))
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
		FullMethod: "/dealer.v1.Manager/GetSnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetSnapshots(ctx, req.(*GetSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dealer.v1.Manager/CreateBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateBackup(ctx, req.(*CreateBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetBackups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBackupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetBackups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dealer.v1.Manager/GetBackups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetBackups(ctx, req.(*GetBackupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dealer.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSnapshot",
			Handler:    _Manager_CreateSnapshot_Handler,
		},
		{
			MethodName: "GetSnapshots",
			Handler:    _Manager_GetSnapshots_Handler,
		},
		{
			MethodName: "CreateBackup",
			Handler:    _Manager_CreateBackup_Handler,
		},
		{
			MethodName: "GetBackups",
			Handler:    _Manager_GetBackups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web3eye/dealer/v1/dealer.proto",
}
