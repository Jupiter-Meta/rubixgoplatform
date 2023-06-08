// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: rubix-external.proto

package protos

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

const (
	RubixExternal_ApproveAuthRequest_FullMethodName       = "/protos.RubixExternal/ApproveAuthRequest"
	RubixExternal_ApproveOrgAuthRequest_FullMethodName    = "/protos.RubixExternal/ApproveOrgAuthRequest"
	RubixExternal_StreamTransactionRequest_FullMethodName = "/protos.RubixExternal/StreamTransactionRequest"
)

// RubixExternalClient is the client API for RubixExternal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RubixExternalClient interface {
	ApproveAuthRequest(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ApproveOrgAuthRequest(ctx context.Context, in *OrgAuthRequest, opts ...grpc.CallOption) (*OrgStatus, error)
	StreamTransactionRequest(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RubixExternal_StreamTransactionRequestClient, error)
}

type rubixExternalClient struct {
	cc grpc.ClientConnInterface
}

func NewRubixExternalClient(cc grpc.ClientConnInterface) RubixExternalClient {
	return &rubixExternalClient{cc}
}

func (c *rubixExternalClient) ApproveAuthRequest(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RubixExternal_ApproveAuthRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rubixExternalClient) ApproveOrgAuthRequest(ctx context.Context, in *OrgAuthRequest, opts ...grpc.CallOption) (*OrgStatus, error) {
	out := new(OrgStatus)
	err := c.cc.Invoke(ctx, RubixExternal_ApproveOrgAuthRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rubixExternalClient) StreamTransactionRequest(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RubixExternal_StreamTransactionRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &RubixExternal_ServiceDesc.Streams[0], RubixExternal_StreamTransactionRequest_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &rubixExternalStreamTransactionRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RubixExternal_StreamTransactionRequestClient interface {
	Recv() (*TxnRequest, error)
	grpc.ClientStream
}

type rubixExternalStreamTransactionRequestClient struct {
	grpc.ClientStream
}

func (x *rubixExternalStreamTransactionRequestClient) Recv() (*TxnRequest, error) {
	m := new(TxnRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RubixExternalServer is the server API for RubixExternal service.
// All implementations must embed UnimplementedRubixExternalServer
// for forward compatibility
type RubixExternalServer interface {
	ApproveAuthRequest(context.Context, *AuthRequest) (*emptypb.Empty, error)
	ApproveOrgAuthRequest(context.Context, *OrgAuthRequest) (*OrgStatus, error)
	StreamTransactionRequest(*emptypb.Empty, RubixExternal_StreamTransactionRequestServer) error
	mustEmbedUnimplementedRubixExternalServer()
}

// UnimplementedRubixExternalServer must be embedded to have forward compatible implementations.
type UnimplementedRubixExternalServer struct {
}

func (UnimplementedRubixExternalServer) ApproveAuthRequest(context.Context, *AuthRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveAuthRequest not implemented")
}
func (UnimplementedRubixExternalServer) ApproveOrgAuthRequest(context.Context, *OrgAuthRequest) (*OrgStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveOrgAuthRequest not implemented")
}
func (UnimplementedRubixExternalServer) StreamTransactionRequest(*emptypb.Empty, RubixExternal_StreamTransactionRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTransactionRequest not implemented")
}
func (UnimplementedRubixExternalServer) mustEmbedUnimplementedRubixExternalServer() {}

// UnsafeRubixExternalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RubixExternalServer will
// result in compilation errors.
type UnsafeRubixExternalServer interface {
	mustEmbedUnimplementedRubixExternalServer()
}

func RegisterRubixExternalServer(s grpc.ServiceRegistrar, srv RubixExternalServer) {
	s.RegisterService(&RubixExternal_ServiceDesc, srv)
}

func _RubixExternal_ApproveAuthRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RubixExternalServer).ApproveAuthRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RubixExternal_ApproveAuthRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RubixExternalServer).ApproveAuthRequest(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RubixExternal_ApproveOrgAuthRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrgAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RubixExternalServer).ApproveOrgAuthRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RubixExternal_ApproveOrgAuthRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RubixExternalServer).ApproveOrgAuthRequest(ctx, req.(*OrgAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RubixExternal_StreamTransactionRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RubixExternalServer).StreamTransactionRequest(m, &rubixExternalStreamTransactionRequestServer{stream})
}

type RubixExternal_StreamTransactionRequestServer interface {
	Send(*TxnRequest) error
	grpc.ServerStream
}

type rubixExternalStreamTransactionRequestServer struct {
	grpc.ServerStream
}

func (x *rubixExternalStreamTransactionRequestServer) Send(m *TxnRequest) error {
	return x.ServerStream.SendMsg(m)
}

// RubixExternal_ServiceDesc is the grpc.ServiceDesc for RubixExternal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RubixExternal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.RubixExternal",
	HandlerType: (*RubixExternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApproveAuthRequest",
			Handler:    _RubixExternal_ApproveAuthRequest_Handler,
		},
		{
			MethodName: "ApproveOrgAuthRequest",
			Handler:    _RubixExternal_ApproveOrgAuthRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTransactionRequest",
			Handler:       _RubixExternal_StreamTransactionRequest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rubix-external.proto",
}