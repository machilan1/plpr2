// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: pulumi/resource.proto

package pulumirpc

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

// ResourceMonitorClient is the client API for ResourceMonitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceMonitorClient interface {
	SupportsFeature(ctx context.Context, in *SupportsFeatureRequest, opts ...grpc.CallOption) (*SupportsFeatureResponse, error)
	Invoke(ctx context.Context, in *ResourceInvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error)
	StreamInvoke(ctx context.Context, in *ResourceInvokeRequest, opts ...grpc.CallOption) (ResourceMonitor_StreamInvokeClient, error)
	Call(ctx context.Context, in *ResourceCallRequest, opts ...grpc.CallOption) (*CallResponse, error)
	ReadResource(ctx context.Context, in *ReadResourceRequest, opts ...grpc.CallOption) (*ReadResourceResponse, error)
	RegisterResource(ctx context.Context, in *RegisterResourceRequest, opts ...grpc.CallOption) (*RegisterResourceResponse, error)
	RegisterResourceOutputs(ctx context.Context, in *RegisterResourceOutputsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Register a resource transform for the stack
	RegisterStackTransform(ctx context.Context, in *Callback, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Register an invoke transform for the stack
	RegisterStackInvokeTransform(ctx context.Context, in *Callback, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Registers a package and allocates a packageRef. The same package can be registered multiple times in Pulumi.
	// Multiple requests are idempotent and guaranteed to return the same result.
	RegisterPackage(ctx context.Context, in *RegisterPackageRequest, opts ...grpc.CallOption) (*RegisterPackageResponse, error)
}

type resourceMonitorClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceMonitorClient(cc grpc.ClientConnInterface) ResourceMonitorClient {
	return &resourceMonitorClient{cc}
}

func (c *resourceMonitorClient) SupportsFeature(ctx context.Context, in *SupportsFeatureRequest, opts ...grpc.CallOption) (*SupportsFeatureResponse, error) {
	out := new(SupportsFeatureResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceMonitor/SupportsFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) Invoke(ctx context.Context, in *ResourceInvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error) {
	out := new(InvokeResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceMonitor/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) StreamInvoke(ctx context.Context, in *ResourceInvokeRequest, opts ...grpc.CallOption) (ResourceMonitor_StreamInvokeClient, error) {
	stream, err := c.cc.NewStream(ctx, &ResourceMonitor_ServiceDesc.Streams[0], "/pulumirpc.ResourceMonitor/StreamInvoke", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceMonitorStreamInvokeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ResourceMonitor_StreamInvokeClient interface {
	Recv() (*InvokeResponse, error)
	grpc.ClientStream
}

type resourceMonitorStreamInvokeClient struct {
	grpc.ClientStream
}

func (x *resourceMonitorStreamInvokeClient) Recv() (*InvokeResponse, error) {
	m := new(InvokeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *resourceMonitorClient) Call(ctx context.Context, in *ResourceCallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceMonitor/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) ReadResource(ctx context.Context, in *ReadResourceRequest, opts ...grpc.CallOption) (*ReadResourceResponse, error) {
	out := new(ReadResourceResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceMonitor/ReadResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) RegisterResource(ctx context.Context, in *RegisterResourceRequest, opts ...grpc.CallOption) (*RegisterResourceResponse, error) {
	out := new(RegisterResourceResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceMonitor/RegisterResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) RegisterResourceOutputs(ctx context.Context, in *RegisterResourceOutputsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceMonitor/RegisterResourceOutputs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) RegisterStackTransform(ctx context.Context, in *Callback, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceMonitor/RegisterStackTransform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) RegisterStackInvokeTransform(ctx context.Context, in *Callback, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceMonitor/RegisterStackInvokeTransform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) RegisterPackage(ctx context.Context, in *RegisterPackageRequest, opts ...grpc.CallOption) (*RegisterPackageResponse, error) {
	out := new(RegisterPackageResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceMonitor/RegisterPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceMonitorServer is the server API for ResourceMonitor service.
// All implementations must embed UnimplementedResourceMonitorServer
// for forward compatibility
type ResourceMonitorServer interface {
	SupportsFeature(context.Context, *SupportsFeatureRequest) (*SupportsFeatureResponse, error)
	Invoke(context.Context, *ResourceInvokeRequest) (*InvokeResponse, error)
	StreamInvoke(*ResourceInvokeRequest, ResourceMonitor_StreamInvokeServer) error
	Call(context.Context, *ResourceCallRequest) (*CallResponse, error)
	ReadResource(context.Context, *ReadResourceRequest) (*ReadResourceResponse, error)
	RegisterResource(context.Context, *RegisterResourceRequest) (*RegisterResourceResponse, error)
	RegisterResourceOutputs(context.Context, *RegisterResourceOutputsRequest) (*emptypb.Empty, error)
	// Register a resource transform for the stack
	RegisterStackTransform(context.Context, *Callback) (*emptypb.Empty, error)
	// Register an invoke transform for the stack
	RegisterStackInvokeTransform(context.Context, *Callback) (*emptypb.Empty, error)
	// Registers a package and allocates a packageRef. The same package can be registered multiple times in Pulumi.
	// Multiple requests are idempotent and guaranteed to return the same result.
	RegisterPackage(context.Context, *RegisterPackageRequest) (*RegisterPackageResponse, error)
	mustEmbedUnimplementedResourceMonitorServer()
}

// UnimplementedResourceMonitorServer must be embedded to have forward compatible implementations.
type UnimplementedResourceMonitorServer struct {
}

func (UnimplementedResourceMonitorServer) SupportsFeature(context.Context, *SupportsFeatureRequest) (*SupportsFeatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupportsFeature not implemented")
}
func (UnimplementedResourceMonitorServer) Invoke(context.Context, *ResourceInvokeRequest) (*InvokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (UnimplementedResourceMonitorServer) StreamInvoke(*ResourceInvokeRequest, ResourceMonitor_StreamInvokeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamInvoke not implemented")
}
func (UnimplementedResourceMonitorServer) Call(context.Context, *ResourceCallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedResourceMonitorServer) ReadResource(context.Context, *ReadResourceRequest) (*ReadResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadResource not implemented")
}
func (UnimplementedResourceMonitorServer) RegisterResource(context.Context, *RegisterResourceRequest) (*RegisterResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterResource not implemented")
}
func (UnimplementedResourceMonitorServer) RegisterResourceOutputs(context.Context, *RegisterResourceOutputsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterResourceOutputs not implemented")
}
func (UnimplementedResourceMonitorServer) RegisterStackTransform(context.Context, *Callback) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterStackTransform not implemented")
}
func (UnimplementedResourceMonitorServer) RegisterStackInvokeTransform(context.Context, *Callback) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterStackInvokeTransform not implemented")
}
func (UnimplementedResourceMonitorServer) RegisterPackage(context.Context, *RegisterPackageRequest) (*RegisterPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPackage not implemented")
}
func (UnimplementedResourceMonitorServer) mustEmbedUnimplementedResourceMonitorServer() {}

// UnsafeResourceMonitorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceMonitorServer will
// result in compilation errors.
type UnsafeResourceMonitorServer interface {
	mustEmbedUnimplementedResourceMonitorServer()
}

func RegisterResourceMonitorServer(s grpc.ServiceRegistrar, srv ResourceMonitorServer) {
	s.RegisterService(&ResourceMonitor_ServiceDesc, srv)
}

func _ResourceMonitor_SupportsFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportsFeatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).SupportsFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/SupportsFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).SupportsFeature(ctx, req.(*SupportsFeatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceInvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).Invoke(ctx, req.(*ResourceInvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_StreamInvoke_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResourceInvokeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourceMonitorServer).StreamInvoke(m, &resourceMonitorStreamInvokeServer{stream})
}

type ResourceMonitor_StreamInvokeServer interface {
	Send(*InvokeResponse) error
	grpc.ServerStream
}

type resourceMonitorStreamInvokeServer struct {
	grpc.ServerStream
}

func (x *resourceMonitorStreamInvokeServer) Send(m *InvokeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ResourceMonitor_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).Call(ctx, req.(*ResourceCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_ReadResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).ReadResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/ReadResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).ReadResource(ctx, req.(*ReadResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_RegisterResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).RegisterResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/RegisterResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).RegisterResource(ctx, req.(*RegisterResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_RegisterResourceOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterResourceOutputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).RegisterResourceOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/RegisterResourceOutputs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).RegisterResourceOutputs(ctx, req.(*RegisterResourceOutputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_RegisterStackTransform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Callback)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).RegisterStackTransform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/RegisterStackTransform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).RegisterStackTransform(ctx, req.(*Callback))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_RegisterStackInvokeTransform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Callback)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).RegisterStackInvokeTransform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/RegisterStackInvokeTransform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).RegisterStackInvokeTransform(ctx, req.(*Callback))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_RegisterPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).RegisterPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/RegisterPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).RegisterPackage(ctx, req.(*RegisterPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceMonitor_ServiceDesc is the grpc.ServiceDesc for ResourceMonitor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceMonitor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.ResourceMonitor",
	HandlerType: (*ResourceMonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SupportsFeature",
			Handler:    _ResourceMonitor_SupportsFeature_Handler,
		},
		{
			MethodName: "Invoke",
			Handler:    _ResourceMonitor_Invoke_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _ResourceMonitor_Call_Handler,
		},
		{
			MethodName: "ReadResource",
			Handler:    _ResourceMonitor_ReadResource_Handler,
		},
		{
			MethodName: "RegisterResource",
			Handler:    _ResourceMonitor_RegisterResource_Handler,
		},
		{
			MethodName: "RegisterResourceOutputs",
			Handler:    _ResourceMonitor_RegisterResourceOutputs_Handler,
		},
		{
			MethodName: "RegisterStackTransform",
			Handler:    _ResourceMonitor_RegisterStackTransform_Handler,
		},
		{
			MethodName: "RegisterStackInvokeTransform",
			Handler:    _ResourceMonitor_RegisterStackInvokeTransform_Handler,
		},
		{
			MethodName: "RegisterPackage",
			Handler:    _ResourceMonitor_RegisterPackage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamInvoke",
			Handler:       _ResourceMonitor_StreamInvoke_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pulumi/resource.proto",
}
