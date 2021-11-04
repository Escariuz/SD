// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ComunicacionJLClient is the client API for ComunicacionJL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComunicacionJLClient interface {
	// rpc GetRespuesta(Request del cliente) returns (respuesta del servidor)
	GetRespuesta(ctx context.Context, in *SolicitudParticipacion, opts ...grpc.CallOption) (*RespuestaLider, error)
}

type comunicacionJLClient struct {
	cc grpc.ClientConnInterface
}

func NewComunicacionJLClient(cc grpc.ClientConnInterface) ComunicacionJLClient {
	return &comunicacionJLClient{cc}
}

func (c *comunicacionJLClient) GetRespuesta(ctx context.Context, in *SolicitudParticipacion, opts ...grpc.CallOption) (*RespuestaLider, error) {
	out := new(RespuestaLider)
	err := c.cc.Invoke(ctx, "/proto.ComunicacionJL/GetRespuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComunicacionJLServer is the server API for ComunicacionJL service.
// All implementations must embed UnimplementedComunicacionJLServer
// for forward compatibility
type ComunicacionJLServer interface {
	// rpc GetRespuesta(Request del cliente) returns (respuesta del servidor)
	GetRespuesta(context.Context, *SolicitudParticipacion) (*RespuestaLider, error)
	mustEmbedUnimplementedComunicacionJLServer()
}

// UnimplementedComunicacionJLServer must be embedded to have forward compatible implementations.
type UnimplementedComunicacionJLServer struct {
}

func (UnimplementedComunicacionJLServer) GetRespuesta(context.Context, *SolicitudParticipacion) (*RespuestaLider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRespuesta not implemented")
}
func (UnimplementedComunicacionJLServer) mustEmbedUnimplementedComunicacionJLServer() {}

// UnsafeComunicacionJLServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComunicacionJLServer will
// result in compilation errors.
type UnsafeComunicacionJLServer interface {
	mustEmbedUnimplementedComunicacionJLServer()
}

func RegisterComunicacionJLServer(s grpc.ServiceRegistrar, srv ComunicacionJLServer) {
	s.RegisterService(&ComunicacionJL_ServiceDesc, srv)
}

func _ComunicacionJL_GetRespuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudParticipacion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunicacionJLServer).GetRespuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ComunicacionJL/GetRespuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunicacionJLServer).GetRespuesta(ctx, req.(*SolicitudParticipacion))
	}
	return interceptor(ctx, in, info, handler)
}

// ComunicacionJL_ServiceDesc is the grpc.ServiceDesc for ComunicacionJL service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComunicacionJL_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ComunicacionJL",
	HandlerType: (*ComunicacionJLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRespuesta",
			Handler:    _ComunicacionJL_GetRespuesta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}
