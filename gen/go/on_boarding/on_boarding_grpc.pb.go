// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: on_boarding/on_boarding.proto

package onboardingpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OtpService_SendOtp_FullMethodName   = "/onboardingpb.OtpService/SendOtp"
	OtpService_VerifyOtp_FullMethodName = "/onboardingpb.OtpService/VerifyOtp"
)

// OtpServiceClient is the client API for OtpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OtpServiceClient interface {
	SendOtp(ctx context.Context, in *SendOtpRequest, opts ...grpc.CallOption) (*SendOtpResponse, error)
	VerifyOtp(ctx context.Context, in *VerifyOtpRequest, opts ...grpc.CallOption) (*VerifyOtpResponse, error)
}

type otpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOtpServiceClient(cc grpc.ClientConnInterface) OtpServiceClient {
	return &otpServiceClient{cc}
}

func (c *otpServiceClient) SendOtp(ctx context.Context, in *SendOtpRequest, opts ...grpc.CallOption) (*SendOtpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendOtpResponse)
	err := c.cc.Invoke(ctx, OtpService_SendOtp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otpServiceClient) VerifyOtp(ctx context.Context, in *VerifyOtpRequest, opts ...grpc.CallOption) (*VerifyOtpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyOtpResponse)
	err := c.cc.Invoke(ctx, OtpService_VerifyOtp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OtpServiceServer is the server API for OtpService service.
// All implementations must embed UnimplementedOtpServiceServer
// for forward compatibility.
type OtpServiceServer interface {
	SendOtp(context.Context, *SendOtpRequest) (*SendOtpResponse, error)
	VerifyOtp(context.Context, *VerifyOtpRequest) (*VerifyOtpResponse, error)
	mustEmbedUnimplementedOtpServiceServer()
}

// UnimplementedOtpServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOtpServiceServer struct{}

func (UnimplementedOtpServiceServer) SendOtp(context.Context, *SendOtpRequest) (*SendOtpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOtp not implemented")
}
func (UnimplementedOtpServiceServer) VerifyOtp(context.Context, *VerifyOtpRequest) (*VerifyOtpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOtp not implemented")
}
func (UnimplementedOtpServiceServer) mustEmbedUnimplementedOtpServiceServer() {}
func (UnimplementedOtpServiceServer) testEmbeddedByValue()                    {}

// UnsafeOtpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OtpServiceServer will
// result in compilation errors.
type UnsafeOtpServiceServer interface {
	mustEmbedUnimplementedOtpServiceServer()
}

func RegisterOtpServiceServer(s grpc.ServiceRegistrar, srv OtpServiceServer) {
	// If the following call pancis, it indicates UnimplementedOtpServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OtpService_ServiceDesc, srv)
}

func _OtpService_SendOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtpServiceServer).SendOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OtpService_SendOtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtpServiceServer).SendOtp(ctx, req.(*SendOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OtpService_VerifyOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtpServiceServer).VerifyOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OtpService_VerifyOtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtpServiceServer).VerifyOtp(ctx, req.(*VerifyOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OtpService_ServiceDesc is the grpc.ServiceDesc for OtpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OtpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "onboardingpb.OtpService",
	HandlerType: (*OtpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOtp",
			Handler:    _OtpService_SendOtp_Handler,
		},
		{
			MethodName: "VerifyOtp",
			Handler:    _OtpService_VerifyOtp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "on_boarding/on_boarding.proto",
}
