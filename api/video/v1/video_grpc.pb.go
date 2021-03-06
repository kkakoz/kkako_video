// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	GetVideo(ctx context.Context, in *GetVideoReq, opts ...grpc.CallOption) (*GetVideoRes, error)
	GetVideos(ctx context.Context, in *VideoNewsReq, opts ...grpc.CallOption) (*VideoNewsRes, error)
	AddVideo(ctx context.Context, in *VideoNewsReq, opts ...grpc.CallOption) (*VideoNewsRes, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) GetVideo(ctx context.Context, in *GetVideoReq, opts ...grpc.CallOption) (*GetVideoRes, error) {
	out := new(GetVideoRes)
	err := c.cc.Invoke(ctx, "/VideoService/GetVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetVideos(ctx context.Context, in *VideoNewsReq, opts ...grpc.CallOption) (*VideoNewsRes, error) {
	out := new(VideoNewsRes)
	err := c.cc.Invoke(ctx, "/VideoService/GetVideos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) AddVideo(ctx context.Context, in *VideoNewsReq, opts ...grpc.CallOption) (*VideoNewsRes, error) {
	out := new(VideoNewsRes)
	err := c.cc.Invoke(ctx, "/VideoService/AddVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations must embed UnimplementedVideoServiceServer
// for forward compatibility
type VideoServiceServer interface {
	GetVideo(context.Context, *GetVideoReq) (*GetVideoRes, error)
	GetVideos(context.Context, *VideoNewsReq) (*VideoNewsRes, error)
	AddVideo(context.Context, *VideoNewsReq) (*VideoNewsRes, error)
	mustEmbedUnimplementedVideoServiceServer()
}

// UnimplementedVideoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (UnimplementedVideoServiceServer) GetVideo(context.Context, *GetVideoReq) (*GetVideoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}
func (UnimplementedVideoServiceServer) GetVideos(context.Context, *VideoNewsReq) (*VideoNewsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideos not implemented")
}
func (UnimplementedVideoServiceServer) AddVideo(context.Context, *VideoNewsReq) (*VideoNewsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVideo not implemented")
}
func (UnimplementedVideoServiceServer) mustEmbedUnimplementedVideoServiceServer() {}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VideoService/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideo(ctx, req.(*GetVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoNewsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VideoService/GetVideos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideos(ctx, req.(*VideoNewsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_AddVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoNewsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).AddVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VideoService/AddVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).AddVideo(ctx, req.(*VideoNewsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideo",
			Handler:    _VideoService_GetVideo_Handler,
		},
		{
			MethodName: "GetVideos",
			Handler:    _VideoService_GetVideos_Handler,
		},
		{
			MethodName: "AddVideo",
			Handler:    _VideoService_AddVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/video/v1/video.proto",
}
