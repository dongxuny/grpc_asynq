// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/after.proto

package demo

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

// AfterClient is the client API for After service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AfterClient interface {
	CreateAfter(ctx context.Context, in *CreateAfterReq, opts ...grpc.CallOption) (*CreateAfterResp, error)
	UpdateAfter(ctx context.Context, in *UpdateAfterReq, opts ...grpc.CallOption) (*UpdateAfterResp, error)
}

type afterClient struct {
	cc grpc.ClientConnInterface
}

func NewAfterClient(cc grpc.ClientConnInterface) AfterClient {
	return &afterClient{cc}
}

func (c *afterClient) CreateAfter(ctx context.Context, in *CreateAfterReq, opts ...grpc.CallOption) (*CreateAfterResp, error) {
	out := new(CreateAfterResp)
	err := c.cc.Invoke(ctx, "/api.v1.After/CreateAfter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *afterClient) UpdateAfter(ctx context.Context, in *UpdateAfterReq, opts ...grpc.CallOption) (*UpdateAfterResp, error) {
	out := new(UpdateAfterResp)
	err := c.cc.Invoke(ctx, "/api.v1.After/UpdateAfter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AfterServer is the server API for After service.
// All implementations should embed UnimplementedAfterServer
// for forward compatibility
type AfterServer interface {
	CreateAfter(context.Context, *CreateAfterReq) (*CreateAfterResp, error)
	UpdateAfter(context.Context, *UpdateAfterReq) (*UpdateAfterResp, error)
}

// UnimplementedAfterServer should be embedded to have forward compatible implementations.
type UnimplementedAfterServer struct {
}

func (UnimplementedAfterServer) CreateAfter(context.Context, *CreateAfterReq) (*CreateAfterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAfter not implemented")
}
func (UnimplementedAfterServer) UpdateAfter(context.Context, *UpdateAfterReq) (*UpdateAfterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAfter not implemented")
}

// UnsafeAfterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AfterServer will
// result in compilation errors.
type UnsafeAfterServer interface {
	mustEmbedUnimplementedAfterServer()
}

func RegisterAfterServer(s grpc.ServiceRegistrar, srv AfterServer) {
	s.RegisterService(&After_ServiceDesc, srv)
}

func _After_CreateAfter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAfterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterServer).CreateAfter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.After/CreateAfter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterServer).CreateAfter(ctx, req.(*CreateAfterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _After_UpdateAfter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAfterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AfterServer).UpdateAfter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.After/UpdateAfter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AfterServer).UpdateAfter(ctx, req.(*UpdateAfterReq))
	}
	return interceptor(ctx, in, info, handler)
}

// After_ServiceDesc is the grpc.ServiceDesc for After service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var After_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.After",
	HandlerType: (*AfterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAfter",
			Handler:    _After_CreateAfter_Handler,
		},
		{
			MethodName: "UpdateAfter",
			Handler:    _After_UpdateAfter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/after.proto",
}