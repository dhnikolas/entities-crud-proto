// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.13.0
// source: entities.proto

package entities_crud

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

// EntitiesCrudClient is the client API for EntitiesCrud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntitiesCrudClient interface {
	ListByFilter(ctx context.Context, in *EventsFilter, opts ...grpc.CallOption) (*EventsResponse, error)
	Info(ctx context.Context, in *ConstructionsRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	Meta(ctx context.Context, in *MetaRequest, opts ...grpc.CallOption) (*MetaResponse, error)
}

type entitiesCrudClient struct {
	cc grpc.ClientConnInterface
}

func NewEntitiesCrudClient(cc grpc.ClientConnInterface) EntitiesCrudClient {
	return &entitiesCrudClient{cc}
}

func (c *entitiesCrudClient) ListByFilter(ctx context.Context, in *EventsFilter, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/entities_crud.EntitiesCrud/ListByFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entitiesCrudClient) Info(ctx context.Context, in *ConstructionsRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/entities_crud.EntitiesCrud/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entitiesCrudClient) Meta(ctx context.Context, in *MetaRequest, opts ...grpc.CallOption) (*MetaResponse, error) {
	out := new(MetaResponse)
	err := c.cc.Invoke(ctx, "/entities_crud.EntitiesCrud/Meta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntitiesCrudServer is the server API for EntitiesCrud service.
// All implementations must embed UnimplementedEntitiesCrudServer
// for forward compatibility
type EntitiesCrudServer interface {
	ListByFilter(context.Context, *EventsFilter) (*EventsResponse, error)
	Info(context.Context, *ConstructionsRequest) (*InfoResponse, error)
	Meta(context.Context, *MetaRequest) (*MetaResponse, error)
	mustEmbedUnimplementedEntitiesCrudServer()
}

// UnimplementedEntitiesCrudServer must be embedded to have forward compatible implementations.
type UnimplementedEntitiesCrudServer struct {
}

func (UnimplementedEntitiesCrudServer) ListByFilter(context.Context, *EventsFilter) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByFilter not implemented")
}
func (UnimplementedEntitiesCrudServer) Info(context.Context, *ConstructionsRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedEntitiesCrudServer) Meta(context.Context, *MetaRequest) (*MetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Meta not implemented")
}
func (UnimplementedEntitiesCrudServer) mustEmbedUnimplementedEntitiesCrudServer() {}

// UnsafeEntitiesCrudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntitiesCrudServer will
// result in compilation errors.
type UnsafeEntitiesCrudServer interface {
	mustEmbedUnimplementedEntitiesCrudServer()
}

func RegisterEntitiesCrudServer(s grpc.ServiceRegistrar, srv EntitiesCrudServer) {
	s.RegisterService(&EntitiesCrud_ServiceDesc, srv)
}

func _EntitiesCrud_ListByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntitiesCrudServer).ListByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entities_crud.EntitiesCrud/ListByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntitiesCrudServer).ListByFilter(ctx, req.(*EventsFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntitiesCrud_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConstructionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntitiesCrudServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entities_crud.EntitiesCrud/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntitiesCrudServer).Info(ctx, req.(*ConstructionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntitiesCrud_Meta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntitiesCrudServer).Meta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entities_crud.EntitiesCrud/Meta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntitiesCrudServer).Meta(ctx, req.(*MetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntitiesCrud_ServiceDesc is the grpc.ServiceDesc for EntitiesCrud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntitiesCrud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entities_crud.EntitiesCrud",
	HandlerType: (*EntitiesCrudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListByFilter",
			Handler:    _EntitiesCrud_ListByFilter_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _EntitiesCrud_Info_Handler,
		},
		{
			MethodName: "Meta",
			Handler:    _EntitiesCrud_Meta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entities.proto",
}
