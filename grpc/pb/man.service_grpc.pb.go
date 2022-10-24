// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: man.service.proto

package pb

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

// ManServiceClient is the client API for ManService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManServiceClient interface {
	GetMan(ctx context.Context, in *GetManRequest, opts ...grpc.CallOption) (*GetManResponse, error)
}

type manServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManServiceClient(cc grpc.ClientConnInterface) ManServiceClient {
	return &manServiceClient{cc}
}

func (c *manServiceClient) GetMan(ctx context.Context, in *GetManRequest, opts ...grpc.CallOption) (*GetManResponse, error) {
	out := new(GetManResponse)
	err := c.cc.Invoke(ctx, "/pb.ManService/GetMan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManServiceServer is the server API for ManService service.
// All implementations must embed UnimplementedManServiceServer
// for forward compatibility
type ManServiceServer interface {
	GetMan(context.Context, *GetManRequest) (*GetManResponse, error)
	mustEmbedUnimplementedManServiceServer()
}

// UnimplementedManServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManServiceServer struct {
}

func (UnimplementedManServiceServer) GetMan(context.Context, *GetManRequest) (*GetManResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMan not implemented")
}
func (UnimplementedManServiceServer) mustEmbedUnimplementedManServiceServer() {}

// UnsafeManServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManServiceServer will
// result in compilation errors.
type UnsafeManServiceServer interface {
	mustEmbedUnimplementedManServiceServer()
}

func RegisterManServiceServer(s grpc.ServiceRegistrar, srv ManServiceServer) {
	s.RegisterService(&ManService_ServiceDesc, srv)
}

func _ManService_GetMan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManServiceServer).GetMan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ManService/GetMan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManServiceServer).GetMan(ctx, req.(*GetManRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManService_ServiceDesc is the grpc.ServiceDesc for ManService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ManService",
	HandlerType: (*ManServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMan",
			Handler:    _ManService_GetMan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "man.service.proto",
}