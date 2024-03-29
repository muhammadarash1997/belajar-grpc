// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: pb/mychatting.proto

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

// MyChattingClient is the client API for MyChatting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyChattingClient interface {
	ChatService(ctx context.Context, opts ...grpc.CallOption) (MyChatting_ChatServiceClient, error)
}

type myChattingClient struct {
	cc grpc.ClientConnInterface
}

func NewMyChattingClient(cc grpc.ClientConnInterface) MyChattingClient {
	return &myChattingClient{cc}
}

func (c *myChattingClient) ChatService(ctx context.Context, opts ...grpc.CallOption) (MyChatting_ChatServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &MyChatting_ServiceDesc.Streams[0], "/proto.MyChatting/ChatService", opts...)
	if err != nil {
		return nil, err
	}
	x := &myChattingChatServiceClient{stream}
	return x, nil
}

type MyChatting_ChatServiceClient interface {
	Send(*ClientRequest) error
	Recv() (*ClientResponse, error)
	grpc.ClientStream
}

type myChattingChatServiceClient struct {
	grpc.ClientStream
}

func (x *myChattingChatServiceClient) Send(m *ClientRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myChattingChatServiceClient) Recv() (*ClientResponse, error) {
	m := new(ClientResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyChattingServer is the server API for MyChatting service.
// All implementations must embed UnimplementedMyChattingServer
// for forward compatibility
type MyChattingServer interface {
	ChatService(MyChatting_ChatServiceServer) error
	mustEmbedUnimplementedMyChattingServer()
}

// UnimplementedMyChattingServer must be embedded to have forward compatible implementations.
type UnimplementedMyChattingServer struct {
}

func (UnimplementedMyChattingServer) ChatService(MyChatting_ChatServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatService not implemented")
}
func (UnimplementedMyChattingServer) mustEmbedUnimplementedMyChattingServer() {}

// UnsafeMyChattingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyChattingServer will
// result in compilation errors.
type UnsafeMyChattingServer interface {
	mustEmbedUnimplementedMyChattingServer()
}

func RegisterMyChattingServer(s grpc.ServiceRegistrar, srv MyChattingServer) {
	s.RegisterService(&MyChatting_ServiceDesc, srv)
}

func _MyChatting_ChatService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyChattingServer).ChatService(&myChattingChatServiceServer{stream})
}

type MyChatting_ChatServiceServer interface {
	Send(*ClientResponse) error
	Recv() (*ClientRequest, error)
	grpc.ServerStream
}

type myChattingChatServiceServer struct {
	grpc.ServerStream
}

func (x *myChattingChatServiceServer) Send(m *ClientResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myChattingChatServiceServer) Recv() (*ClientRequest, error) {
	m := new(ClientRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyChatting_ServiceDesc is the grpc.ServiceDesc for MyChatting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyChatting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MyChatting",
	HandlerType: (*MyChattingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatService",
			Handler:       _MyChatting_ChatService_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/mychatting.proto",
}
