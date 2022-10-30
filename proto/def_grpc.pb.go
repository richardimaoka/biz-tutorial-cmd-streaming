// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/def.proto

package proto

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

// CommandServiceClient is the client API for CommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandServiceClient interface {
	// Sends a greeting
	StreamCommands(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (CommandService_StreamCommandsClient, error)
}

type commandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandServiceClient(cc grpc.ClientConnInterface) CommandServiceClient {
	return &commandServiceClient{cc}
}

func (c *commandServiceClient) StreamCommands(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (CommandService_StreamCommandsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CommandService_ServiceDesc.Streams[0], "/CommandService/StreamCommands", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandServiceStreamCommandsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CommandService_StreamCommandsClient interface {
	Recv() (*CommandReply, error)
	grpc.ClientStream
}

type commandServiceStreamCommandsClient struct {
	grpc.ClientStream
}

func (x *commandServiceStreamCommandsClient) Recv() (*CommandReply, error) {
	m := new(CommandReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommandServiceServer is the server API for CommandService service.
// All implementations must embed UnimplementedCommandServiceServer
// for forward compatibility
type CommandServiceServer interface {
	// Sends a greeting
	StreamCommands(*CommandRequest, CommandService_StreamCommandsServer) error
	mustEmbedUnimplementedCommandServiceServer()
}

// UnimplementedCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommandServiceServer struct {
}

func (UnimplementedCommandServiceServer) StreamCommands(*CommandRequest, CommandService_StreamCommandsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCommands not implemented")
}
func (UnimplementedCommandServiceServer) mustEmbedUnimplementedCommandServiceServer() {}

// UnsafeCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandServiceServer will
// result in compilation errors.
type UnsafeCommandServiceServer interface {
	mustEmbedUnimplementedCommandServiceServer()
}

func RegisterCommandServiceServer(s grpc.ServiceRegistrar, srv CommandServiceServer) {
	s.RegisterService(&CommandService_ServiceDesc, srv)
}

func _CommandService_StreamCommands_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommandServiceServer).StreamCommands(m, &commandServiceStreamCommandsServer{stream})
}

type CommandService_StreamCommandsServer interface {
	Send(*CommandReply) error
	grpc.ServerStream
}

type commandServiceStreamCommandsServer struct {
	grpc.ServerStream
}

func (x *commandServiceStreamCommandsServer) Send(m *CommandReply) error {
	return x.ServerStream.SendMsg(m)
}

// CommandService_ServiceDesc is the grpc.ServiceDesc for CommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CommandService",
	HandlerType: (*CommandServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCommands",
			Handler:       _CommandService_StreamCommands_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/def.proto",
}
