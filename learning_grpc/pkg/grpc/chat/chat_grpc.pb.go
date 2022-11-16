// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: api/chat.proto

package chat

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	// 服务器流服务,服务器可以多次发送信息
	ChatServerStream(ctx context.Context, in *ChatReq, opts ...grpc.CallOption) (Chat_ChatServerStreamClient, error)
	// 客户端流服务,客户端可以多次发送信息
	ChatClientStream(ctx context.Context, opts ...grpc.CallOption) (Chat_ChatClientStreamClient, error)
	// 双向流服务,服务端、客户端都可以多次发送信息
	ChatBiStreams(ctx context.Context, opts ...grpc.CallOption) (Chat_ChatBiStreamsClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) ChatServerStream(ctx context.Context, in *ChatReq, opts ...grpc.CallOption) (Chat_ChatServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], "/chat.Chat/ChatServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatChatServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_ChatServerStreamClient interface {
	Recv() (*ChatResp, error)
	grpc.ClientStream
}

type chatChatServerStreamClient struct {
	grpc.ClientStream
}

func (x *chatChatServerStreamClient) Recv() (*ChatResp, error) {
	m := new(ChatResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) ChatClientStream(ctx context.Context, opts ...grpc.CallOption) (Chat_ChatClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[1], "/chat.Chat/ChatClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatChatClientStreamClient{stream}
	return x, nil
}

type Chat_ChatClientStreamClient interface {
	Send(*ChatReq) error
	CloseAndRecv() (*ChatResp, error)
	grpc.ClientStream
}

type chatChatClientStreamClient struct {
	grpc.ClientStream
}

func (x *chatChatClientStreamClient) Send(m *ChatReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatChatClientStreamClient) CloseAndRecv() (*ChatResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ChatResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) ChatBiStreams(ctx context.Context, opts ...grpc.CallOption) (Chat_ChatBiStreamsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[2], "/chat.Chat/ChatBiStreams", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatChatBiStreamsClient{stream}
	return x, nil
}

type Chat_ChatBiStreamsClient interface {
	Send(*ChatReq) error
	Recv() (*ChatResp, error)
	grpc.ClientStream
}

type chatChatBiStreamsClient struct {
	grpc.ClientStream
}

func (x *chatChatBiStreamsClient) Send(m *ChatReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatChatBiStreamsClient) Recv() (*ChatResp, error) {
	m := new(ChatResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	// 服务器流服务,服务器可以多次发送信息
	ChatServerStream(*ChatReq, Chat_ChatServerStreamServer) error
	// 客户端流服务,客户端可以多次发送信息
	ChatClientStream(Chat_ChatClientStreamServer) error
	// 双向流服务,服务端、客户端都可以多次发送信息
	ChatBiStreams(Chat_ChatBiStreamsServer) error
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) ChatServerStream(*ChatReq, Chat_ChatServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatServerStream not implemented")
}
func (UnimplementedChatServer) ChatClientStream(Chat_ChatClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatClientStream not implemented")
}
func (UnimplementedChatServer) ChatBiStreams(Chat_ChatBiStreamsServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatBiStreams not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_ChatServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChatReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).ChatServerStream(m, &chatChatServerStreamServer{stream})
}

type Chat_ChatServerStreamServer interface {
	Send(*ChatResp) error
	grpc.ServerStream
}

type chatChatServerStreamServer struct {
	grpc.ServerStream
}

func (x *chatChatServerStreamServer) Send(m *ChatResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Chat_ChatClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).ChatClientStream(&chatChatClientStreamServer{stream})
}

type Chat_ChatClientStreamServer interface {
	SendAndClose(*ChatResp) error
	Recv() (*ChatReq, error)
	grpc.ServerStream
}

type chatChatClientStreamServer struct {
	grpc.ServerStream
}

func (x *chatChatClientStreamServer) SendAndClose(m *ChatResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatChatClientStreamServer) Recv() (*ChatReq, error) {
	m := new(ChatReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_ChatBiStreams_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).ChatBiStreams(&chatChatBiStreamsServer{stream})
}

type Chat_ChatBiStreamsServer interface {
	Send(*ChatResp) error
	Recv() (*ChatReq, error)
	grpc.ServerStream
}

type chatChatBiStreamsServer struct {
	grpc.ServerStream
}

func (x *chatChatBiStreamsServer) Send(m *ChatResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatChatBiStreamsServer) Recv() (*ChatReq, error) {
	m := new(ChatReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatServerStream",
			Handler:       _Chat_ChatServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ChatClientStream",
			Handler:       _Chat_ChatClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ChatBiStreams",
			Handler:       _Chat_ChatBiStreams_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/chat.proto",
}
