// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package helloworld

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

// ComunicationClient is the client API for Comunication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComunicationClient interface {
	// Sends Comands_Informantes
	Comands_Informantes_Broker(ctx context.Context, in *ComandIBRequest, opts ...grpc.CallOption) (*ComandIBReply, error)
	// Sends Comands_Leia
	Comands_Leia_Broker(ctx context.Context, in *ComandLBRequest, opts ...grpc.CallOption) (*ComandLBReply, error)
	// Sends Comands_Broker_Fulcrum
	Comands_Broker_Fulcrum(ctx context.Context, in *ComandBFRequest, opts ...grpc.CallOption) (*ComandBFReply, error)
	// Sends Comands_Informantes_Fulcrum
	Comands_Informantes_Fulcrum(ctx context.Context, in *ComandIFRequest, opts ...grpc.CallOption) (*ComandIFReply, error)
	// Sends Comands_Request_Hashing
	Comands_Request_Hashing(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*HashRepply, error)
	// Sends Comands_Fulcrum_Fulcrum
	Comands_Request_Files(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*ComandFFFiles, error)
	Comands_Retrieve_Files(ctx context.Context, in *ComandFFFiles, opts ...grpc.CallOption) (*PingMsg, error)
}

type comunicationClient struct {
	cc grpc.ClientConnInterface
}

func NewComunicationClient(cc grpc.ClientConnInterface) ComunicationClient {
	return &comunicationClient{cc}
}

func (c *comunicationClient) Comands_Informantes_Broker(ctx context.Context, in *ComandIBRequest, opts ...grpc.CallOption) (*ComandIBReply, error) {
	out := new(ComandIBReply)
	err := c.cc.Invoke(ctx, "/helloworld.Comunication/Comands_Informantes_Broker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunicationClient) Comands_Leia_Broker(ctx context.Context, in *ComandLBRequest, opts ...grpc.CallOption) (*ComandLBReply, error) {
	out := new(ComandLBReply)
	err := c.cc.Invoke(ctx, "/helloworld.Comunication/Comands_Leia_Broker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunicationClient) Comands_Broker_Fulcrum(ctx context.Context, in *ComandBFRequest, opts ...grpc.CallOption) (*ComandBFReply, error) {
	out := new(ComandBFReply)
	err := c.cc.Invoke(ctx, "/helloworld.Comunication/Comands_Broker_Fulcrum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunicationClient) Comands_Informantes_Fulcrum(ctx context.Context, in *ComandIFRequest, opts ...grpc.CallOption) (*ComandIFReply, error) {
	out := new(ComandIFReply)
	err := c.cc.Invoke(ctx, "/helloworld.Comunication/Comands_Informantes_Fulcrum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunicationClient) Comands_Request_Hashing(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*HashRepply, error) {
	out := new(HashRepply)
	err := c.cc.Invoke(ctx, "/helloworld.Comunication/Comands_Request_Hashing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunicationClient) Comands_Request_Files(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*ComandFFFiles, error) {
	out := new(ComandFFFiles)
	err := c.cc.Invoke(ctx, "/helloworld.Comunication/Comands_Request_Files", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunicationClient) Comands_Retrieve_Files(ctx context.Context, in *ComandFFFiles, opts ...grpc.CallOption) (*PingMsg, error) {
	out := new(PingMsg)
	err := c.cc.Invoke(ctx, "/helloworld.Comunication/Comands_Retrieve_Files", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComunicationServer is the server API for Comunication service.
// All implementations must embed UnimplementedComunicationServer
// for forward compatibility
type ComunicationServer interface {
	// Sends Comands_Informantes
	Comands_Informantes_Broker(context.Context, *ComandIBRequest) (*ComandIBReply, error)
	// Sends Comands_Leia
	Comands_Leia_Broker(context.Context, *ComandLBRequest) (*ComandLBReply, error)
	// Sends Comands_Broker_Fulcrum
	Comands_Broker_Fulcrum(context.Context, *ComandBFRequest) (*ComandBFReply, error)
	// Sends Comands_Informantes_Fulcrum
	Comands_Informantes_Fulcrum(context.Context, *ComandIFRequest) (*ComandIFReply, error)
	// Sends Comands_Request_Hashing
	Comands_Request_Hashing(context.Context, *PingMsg) (*HashRepply, error)
	// Sends Comands_Fulcrum_Fulcrum
	Comands_Request_Files(context.Context, *PingMsg) (*ComandFFFiles, error)
	Comands_Retrieve_Files(context.Context, *ComandFFFiles) (*PingMsg, error)
	mustEmbedUnimplementedComunicationServer()
}

// UnimplementedComunicationServer must be embedded to have forward compatible implementations.
type UnimplementedComunicationServer struct {
}

func (UnimplementedComunicationServer) Comands_Informantes_Broker(context.Context, *ComandIBRequest) (*ComandIBReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comands_Informantes_Broker not implemented")
}
func (UnimplementedComunicationServer) Comands_Leia_Broker(context.Context, *ComandLBRequest) (*ComandLBReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comands_Leia_Broker not implemented")
}
func (UnimplementedComunicationServer) Comands_Broker_Fulcrum(context.Context, *ComandBFRequest) (*ComandBFReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comands_Broker_Fulcrum not implemented")
}
func (UnimplementedComunicationServer) Comands_Informantes_Fulcrum(context.Context, *ComandIFRequest) (*ComandIFReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comands_Informantes_Fulcrum not implemented")
}
func (UnimplementedComunicationServer) Comands_Request_Hashing(context.Context, *PingMsg) (*HashRepply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comands_Request_Hashing not implemented")
}
func (UnimplementedComunicationServer) Comands_Request_Files(context.Context, *PingMsg) (*ComandFFFiles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comands_Request_Files not implemented")
}
func (UnimplementedComunicationServer) Comands_Retrieve_Files(context.Context, *ComandFFFiles) (*PingMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comands_Retrieve_Files not implemented")
}
func (UnimplementedComunicationServer) mustEmbedUnimplementedComunicationServer() {}

// UnsafeComunicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComunicationServer will
// result in compilation errors.
type UnsafeComunicationServer interface {
	mustEmbedUnimplementedComunicationServer()
}

func RegisterComunicationServer(s grpc.ServiceRegistrar, srv ComunicationServer) {
	s.RegisterService(&Comunication_ServiceDesc, srv)
}

func _Comunication_Comands_Informantes_Broker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComandIBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunicationServer).Comands_Informantes_Broker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Comunication/Comands_Informantes_Broker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunicationServer).Comands_Informantes_Broker(ctx, req.(*ComandIBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comunication_Comands_Leia_Broker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComandLBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunicationServer).Comands_Leia_Broker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Comunication/Comands_Leia_Broker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunicationServer).Comands_Leia_Broker(ctx, req.(*ComandLBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comunication_Comands_Broker_Fulcrum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComandBFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunicationServer).Comands_Broker_Fulcrum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Comunication/Comands_Broker_Fulcrum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunicationServer).Comands_Broker_Fulcrum(ctx, req.(*ComandBFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comunication_Comands_Informantes_Fulcrum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComandIFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunicationServer).Comands_Informantes_Fulcrum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Comunication/Comands_Informantes_Fulcrum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunicationServer).Comands_Informantes_Fulcrum(ctx, req.(*ComandIFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comunication_Comands_Request_Hashing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunicationServer).Comands_Request_Hashing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Comunication/Comands_Request_Hashing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunicationServer).Comands_Request_Hashing(ctx, req.(*PingMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comunication_Comands_Request_Files_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunicationServer).Comands_Request_Files(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Comunication/Comands_Request_Files",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunicationServer).Comands_Request_Files(ctx, req.(*PingMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comunication_Comands_Retrieve_Files_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComandFFFiles)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunicationServer).Comands_Retrieve_Files(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Comunication/Comands_Retrieve_Files",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunicationServer).Comands_Retrieve_Files(ctx, req.(*ComandFFFiles))
	}
	return interceptor(ctx, in, info, handler)
}

// Comunication_ServiceDesc is the grpc.ServiceDesc for Comunication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comunication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Comunication",
	HandlerType: (*ComunicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Comands_Informantes_Broker",
			Handler:    _Comunication_Comands_Informantes_Broker_Handler,
		},
		{
			MethodName: "Comands_Leia_Broker",
			Handler:    _Comunication_Comands_Leia_Broker_Handler,
		},
		{
			MethodName: "Comands_Broker_Fulcrum",
			Handler:    _Comunication_Comands_Broker_Fulcrum_Handler,
		},
		{
			MethodName: "Comands_Informantes_Fulcrum",
			Handler:    _Comunication_Comands_Informantes_Fulcrum_Handler,
		},
		{
			MethodName: "Comands_Request_Hashing",
			Handler:    _Comunication_Comands_Request_Hashing_Handler,
		},
		{
			MethodName: "Comands_Request_Files",
			Handler:    _Comunication_Comands_Request_Files_Handler,
		},
		{
			MethodName: "Comands_Retrieve_Files",
			Handler:    _Comunication_Comands_Retrieve_Files_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworld.proto",
}
