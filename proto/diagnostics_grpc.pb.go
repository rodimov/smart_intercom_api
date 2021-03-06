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

// DiagnosticsClient is the client API for Diagnostics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiagnosticsClient interface {
	GetDiagnostic(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Diagnostic, error)
}

type diagnosticsClient struct {
	cc grpc.ClientConnInterface
}

func NewDiagnosticsClient(cc grpc.ClientConnInterface) DiagnosticsClient {
	return &diagnosticsClient{cc}
}

func (c *diagnosticsClient) GetDiagnostic(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Diagnostic, error) {
	out := new(Diagnostic)
	err := c.cc.Invoke(ctx, "/Diagnostics/GetDiagnostic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiagnosticsServer is the server API for Diagnostics service.
// All implementations must embed UnimplementedDiagnosticsServer
// for forward compatibility
type DiagnosticsServer interface {
	GetDiagnostic(context.Context, *Empty) (*Diagnostic, error)
	mustEmbedUnimplementedDiagnosticsServer()
}

// UnimplementedDiagnosticsServer must be embedded to have forward compatible implementations.
type UnimplementedDiagnosticsServer struct {
}

func (UnimplementedDiagnosticsServer) GetDiagnostic(context.Context, *Empty) (*Diagnostic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiagnostic not implemented")
}
func (UnimplementedDiagnosticsServer) mustEmbedUnimplementedDiagnosticsServer() {}

// UnsafeDiagnosticsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiagnosticsServer will
// result in compilation errors.
type UnsafeDiagnosticsServer interface {
	mustEmbedUnimplementedDiagnosticsServer()
}

func RegisterDiagnosticsServer(s grpc.ServiceRegistrar, srv DiagnosticsServer) {
	s.RegisterService(&Diagnostics_ServiceDesc, srv)
}

func _Diagnostics_GetDiagnostic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiagnosticsServer).GetDiagnostic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Diagnostics/GetDiagnostic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiagnosticsServer).GetDiagnostic(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Diagnostics_ServiceDesc is the grpc.ServiceDesc for Diagnostics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Diagnostics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Diagnostics",
	HandlerType: (*DiagnosticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDiagnostic",
			Handler:    _Diagnostics_GetDiagnostic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/diagnostics.proto",
}
