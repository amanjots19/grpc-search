// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/service.proto

package grpc_search

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
	TransacUser_AddUser_FullMethodName = "/proto.TransacUser/AddUser"
	TransacUser_GetUser_FullMethodName = "/proto.TransacUser/GetUser"
)

// TransacUserClient is the client API for TransacUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransacUserClient interface {
	AddUser(ctx context.Context, in *UserModel, opts ...grpc.CallOption) (*APIResp, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type transacUserClient struct {
	cc grpc.ClientConnInterface
}

func NewTransacUserClient(cc grpc.ClientConnInterface) TransacUserClient {
	return &transacUserClient{cc}
}

func (c *transacUserClient) AddUser(ctx context.Context, in *UserModel, opts ...grpc.CallOption) (*APIResp, error) {
	out := new(APIResp)
	err := c.cc.Invoke(ctx, TransacUser_AddUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transacUserClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, TransacUser_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransacUserServer is the server API for TransacUser service.
// All implementations must embed UnimplementedTransacUserServer
// for forward compatibility
type TransacUserServer interface {
	AddUser(context.Context, *UserModel) (*APIResp, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	mustEmbedUnimplementedTransacUserServer()
}

// UnimplementedTransacUserServer must be embedded to have forward compatible implementations.
type UnimplementedTransacUserServer struct {
}

func (UnimplementedTransacUserServer) AddUser(context.Context, *UserModel) (*APIResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedTransacUserServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedTransacUserServer) mustEmbedUnimplementedTransacUserServer() {}

// UnsafeTransacUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransacUserServer will
// result in compilation errors.
type UnsafeTransacUserServer interface {
	mustEmbedUnimplementedTransacUserServer()
}

func RegisterTransacUserServer(s grpc.ServiceRegistrar, srv TransacUserServer) {
	s.RegisterService(&TransacUser_ServiceDesc, srv)
}

func _TransacUser_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransacUserServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransacUser_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransacUserServer).AddUser(ctx, req.(*UserModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransacUser_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransacUserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransacUser_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransacUserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransacUser_ServiceDesc is the grpc.ServiceDesc for TransacUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransacUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TransacUser",
	HandlerType: (*TransacUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _TransacUser_AddUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _TransacUser_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
