// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: csrf.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CSRFServiceClient is the client API for CSRFService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CSRFServiceClient interface {
	CreateCSRFToken(ctx context.Context, in *EmailCSRF, opts ...grpc.CallOption) (*TokenCSRF, error)
	InvalidateCSRFToken(ctx context.Context, in *CSRF, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidateCSRFToken(ctx context.Context, in *CSRF, opts ...grpc.CallOption) (*Valid, error)
}

type cSRFServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCSRFServiceClient(cc grpc.ClientConnInterface) CSRFServiceClient {
	return &cSRFServiceClient{cc}
}

func (c *cSRFServiceClient) CreateCSRFToken(ctx context.Context, in *EmailCSRF, opts ...grpc.CallOption) (*TokenCSRF, error) {
	out := new(TokenCSRF)
	err := c.cc.Invoke(ctx, "/csrf.CSRFService/CreateCSRFToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cSRFServiceClient) InvalidateCSRFToken(ctx context.Context, in *CSRF, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/csrf.CSRFService/InvalidateCSRFToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cSRFServiceClient) ValidateCSRFToken(ctx context.Context, in *CSRF, opts ...grpc.CallOption) (*Valid, error) {
	out := new(Valid)
	err := c.cc.Invoke(ctx, "/csrf.CSRFService/ValidateCSRFToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CSRFServiceServer is the server API for CSRFService service.
// All implementations must embed UnimplementedCSRFServiceServer
// for forward compatibility
type CSRFServiceServer interface {
	CreateCSRFToken(context.Context, *EmailCSRF) (*TokenCSRF, error)
	InvalidateCSRFToken(context.Context, *CSRF) (*emptypb.Empty, error)
	ValidateCSRFToken(context.Context, *CSRF) (*Valid, error)
	mustEmbedUnimplementedCSRFServiceServer()
}

// UnimplementedCSRFServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCSRFServiceServer struct {
}

func (UnimplementedCSRFServiceServer) CreateCSRFToken(context.Context, *EmailCSRF) (*TokenCSRF, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCSRFToken not implemented")
}
func (UnimplementedCSRFServiceServer) InvalidateCSRFToken(context.Context, *CSRF) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateCSRFToken not implemented")
}
func (UnimplementedCSRFServiceServer) ValidateCSRFToken(context.Context, *CSRF) (*Valid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCSRFToken not implemented")
}
func (UnimplementedCSRFServiceServer) mustEmbedUnimplementedCSRFServiceServer() {}

// UnsafeCSRFServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CSRFServiceServer will
// result in compilation errors.
type UnsafeCSRFServiceServer interface {
	mustEmbedUnimplementedCSRFServiceServer()
}

func RegisterCSRFServiceServer(s grpc.ServiceRegistrar, srv CSRFServiceServer) {
	s.RegisterService(&CSRFService_ServiceDesc, srv)
}

func _CSRFService_CreateCSRFToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailCSRF)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSRFServiceServer).CreateCSRFToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csrf.CSRFService/CreateCSRFToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSRFServiceServer).CreateCSRFToken(ctx, req.(*EmailCSRF))
	}
	return interceptor(ctx, in, info, handler)
}

func _CSRFService_InvalidateCSRFToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRF)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSRFServiceServer).InvalidateCSRFToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csrf.CSRFService/InvalidateCSRFToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSRFServiceServer).InvalidateCSRFToken(ctx, req.(*CSRF))
	}
	return interceptor(ctx, in, info, handler)
}

func _CSRFService_ValidateCSRFToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRF)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSRFServiceServer).ValidateCSRFToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csrf.CSRFService/ValidateCSRFToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSRFServiceServer).ValidateCSRFToken(ctx, req.(*CSRF))
	}
	return interceptor(ctx, in, info, handler)
}

// CSRFService_ServiceDesc is the grpc.ServiceDesc for CSRFService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CSRFService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "csrf.CSRFService",
	HandlerType: (*CSRFServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCSRFToken",
			Handler:    _CSRFService_CreateCSRFToken_Handler,
		},
		{
			MethodName: "InvalidateCSRFToken",
			Handler:    _CSRFService_InvalidateCSRFToken_Handler,
		},
		{
			MethodName: "ValidateCSRFToken",
			Handler:    _CSRFService_ValidateCSRFToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "csrf.proto",
}