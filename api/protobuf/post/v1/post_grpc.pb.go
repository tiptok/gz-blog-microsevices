// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: post/v1/post.proto

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

const (
	PostService_GetPost_FullMethodName                          = "/api.protobuf.post.v1.PostService/GetPost"
	PostService_CreatePost_FullMethodName                       = "/api.protobuf.post.v1.PostService/CreatePost"
	PostService_UpdatePost_FullMethodName                       = "/api.protobuf.post.v1.PostService/UpdatePost"
	PostService_DeletePost_FullMethodName                       = "/api.protobuf.post.v1.PostService/DeletePost"
	PostService_DeletePostCompensate_FullMethodName             = "/api.protobuf.post.v1.PostService/DeletePostCompensate"
	PostService_ListPosts_FullMethodName                        = "/api.protobuf.post.v1.PostService/ListPosts"
	PostService_IncrementCommentsCount_FullMethodName           = "/api.protobuf.post.v1.PostService/IncrementCommentsCount"
	PostService_IncrementCommentsCountCompensate_FullMethodName = "/api.protobuf.post.v1.PostService/IncrementCommentsCountCompensate"
	PostService_DecrementCommentsCount_FullMethodName           = "/api.protobuf.post.v1.PostService/DecrementCommentsCount"
	PostService_DecrementCommentsCountCompensate_FullMethodName = "/api.protobuf.post.v1.PostService/DecrementCommentsCountCompensate"
)

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error)
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	DeletePostCompensate(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	ListPosts(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error)
	IncrementCommentsCount(ctx context.Context, in *IncrementCommentsCountRequest, opts ...grpc.CallOption) (*IncrementCommentsCountResponse, error)
	IncrementCommentsCountCompensate(ctx context.Context, in *IncrementCommentsCountRequest, opts ...grpc.CallOption) (*IncrementCommentsCountResponse, error)
	DecrementCommentsCount(ctx context.Context, in *DecrementCommentsCountRequest, opts ...grpc.CallOption) (*DecrementCommentsCountResponse, error)
	DecrementCommentsCountCompensate(ctx context.Context, in *DecrementCommentsCountRequest, opts ...grpc.CallOption) (*DecrementCommentsCountResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error) {
	out := new(GetPostResponse)
	err := c.cc.Invoke(ctx, PostService_GetPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, PostService_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error) {
	out := new(UpdatePostResponse)
	err := c.cc.Invoke(ctx, PostService_UpdatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, PostService_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePostCompensate(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, PostService_DeletePostCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ListPosts(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error) {
	out := new(ListPostsResponse)
	err := c.cc.Invoke(ctx, PostService_ListPosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) IncrementCommentsCount(ctx context.Context, in *IncrementCommentsCountRequest, opts ...grpc.CallOption) (*IncrementCommentsCountResponse, error) {
	out := new(IncrementCommentsCountResponse)
	err := c.cc.Invoke(ctx, PostService_IncrementCommentsCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) IncrementCommentsCountCompensate(ctx context.Context, in *IncrementCommentsCountRequest, opts ...grpc.CallOption) (*IncrementCommentsCountResponse, error) {
	out := new(IncrementCommentsCountResponse)
	err := c.cc.Invoke(ctx, PostService_IncrementCommentsCountCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DecrementCommentsCount(ctx context.Context, in *DecrementCommentsCountRequest, opts ...grpc.CallOption) (*DecrementCommentsCountResponse, error) {
	out := new(DecrementCommentsCountResponse)
	err := c.cc.Invoke(ctx, PostService_DecrementCommentsCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DecrementCommentsCountCompensate(ctx context.Context, in *DecrementCommentsCountRequest, opts ...grpc.CallOption) (*DecrementCommentsCountResponse, error) {
	out := new(DecrementCommentsCountResponse)
	err := c.cc.Invoke(ctx, PostService_DecrementCommentsCountCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error)
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	DeletePostCompensate(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	ListPosts(context.Context, *ListPostsRequest) (*ListPostsResponse, error)
	IncrementCommentsCount(context.Context, *IncrementCommentsCountRequest) (*IncrementCommentsCountResponse, error)
	IncrementCommentsCountCompensate(context.Context, *IncrementCommentsCountRequest) (*IncrementCommentsCountResponse, error)
	DecrementCommentsCount(context.Context, *DecrementCommentsCountRequest) (*DecrementCommentsCountResponse, error)
	DecrementCommentsCountCompensate(context.Context, *DecrementCommentsCountRequest) (*DecrementCommentsCountResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostServiceServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServiceServer) UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedPostServiceServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServiceServer) DeletePostCompensate(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePostCompensate not implemented")
}
func (UnimplementedPostServiceServer) ListPosts(context.Context, *ListPostsRequest) (*ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
func (UnimplementedPostServiceServer) IncrementCommentsCount(context.Context, *IncrementCommentsCountRequest) (*IncrementCommentsCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementCommentsCount not implemented")
}
func (UnimplementedPostServiceServer) IncrementCommentsCountCompensate(context.Context, *IncrementCommentsCountRequest) (*IncrementCommentsCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementCommentsCountCompensate not implemented")
}
func (UnimplementedPostServiceServer) DecrementCommentsCount(context.Context, *DecrementCommentsCountRequest) (*DecrementCommentsCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrementCommentsCount not implemented")
}
func (UnimplementedPostServiceServer) DecrementCommentsCountCompensate(context.Context, *DecrementCommentsCountRequest) (*DecrementCommentsCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrementCommentsCountCompensate not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePostCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePostCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeletePostCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePostCompensate(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ListPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ListPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_ListPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ListPosts(ctx, req.(*ListPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_IncrementCommentsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementCommentsCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).IncrementCommentsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_IncrementCommentsCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).IncrementCommentsCount(ctx, req.(*IncrementCommentsCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_IncrementCommentsCountCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementCommentsCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).IncrementCommentsCountCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_IncrementCommentsCountCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).IncrementCommentsCountCompensate(ctx, req.(*IncrementCommentsCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DecrementCommentsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrementCommentsCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DecrementCommentsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DecrementCommentsCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DecrementCommentsCount(ctx, req.(*DecrementCommentsCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DecrementCommentsCountCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrementCommentsCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DecrementCommentsCountCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DecrementCommentsCountCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DecrementCommentsCountCompensate(ctx, req.(*DecrementCommentsCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.protobuf.post.v1.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPost",
			Handler:    _PostService_GetPost_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _PostService_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
		{
			MethodName: "DeletePostCompensate",
			Handler:    _PostService_DeletePostCompensate_Handler,
		},
		{
			MethodName: "ListPosts",
			Handler:    _PostService_ListPosts_Handler,
		},
		{
			MethodName: "IncrementCommentsCount",
			Handler:    _PostService_IncrementCommentsCount_Handler,
		},
		{
			MethodName: "IncrementCommentsCountCompensate",
			Handler:    _PostService_IncrementCommentsCountCompensate_Handler,
		},
		{
			MethodName: "DecrementCommentsCount",
			Handler:    _PostService_DecrementCommentsCount_Handler,
		},
		{
			MethodName: "DecrementCommentsCountCompensate",
			Handler:    _PostService_DecrementCommentsCountCompensate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post/v1/post.proto",
}
