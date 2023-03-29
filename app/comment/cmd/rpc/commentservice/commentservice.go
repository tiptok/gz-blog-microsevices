// Code generated by goctl. DO NOT EDIT!
// Source: comment.proto

package commentservice

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment                        = v1.Comment
	CreateCommentRequest           = v1.CreateCommentRequest
	CreateCommentResponse          = v1.CreateCommentResponse
	DeleteCommentRequest           = v1.DeleteCommentRequest
	DeleteCommentResponse          = v1.DeleteCommentResponse
	DeleteCommentsByPostIDRequest  = v1.DeleteCommentsByPostIDRequest
	DeleteCommentsByPostIDResponse = v1.DeleteCommentsByPostIDResponse
	GetCommentByUUIDRequest        = v1.GetCommentByUUIDRequest
	GetCommentByUUIDResponse       = v1.GetCommentByUUIDResponse
	GetCommentRequest              = v1.GetCommentRequest
	GetCommentResponse             = v1.GetCommentResponse
	ListCommentsByPostIDRequest    = v1.ListCommentsByPostIDRequest
	ListCommentsByPostIDResponse   = v1.ListCommentsByPostIDResponse
	UpdateCommentRequest           = v1.UpdateCommentRequest
	UpdateCommentResponse          = v1.UpdateCommentResponse

	CommentService interface {
		CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
		CreateCommentCompensate(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
		GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error)
		GetCommentByUUID(ctx context.Context, in *GetCommentByUUIDRequest, opts ...grpc.CallOption) (*GetCommentByUUIDResponse, error)
		UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
		DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
		DeleteCommentCompensate(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
		DeleteCommentsByPostID(ctx context.Context, in *DeleteCommentsByPostIDRequest, opts ...grpc.CallOption) (*DeleteCommentsByPostIDResponse, error)
		DeleteCommentsByPostIDCompensate(ctx context.Context, in *DeleteCommentsByPostIDRequest, opts ...grpc.CallOption) (*DeleteCommentsByPostIDResponse, error)
		ListCommentsByPostID(ctx context.Context, in *ListCommentsByPostIDRequest, opts ...grpc.CallOption) (*ListCommentsByPostIDResponse, error)
	}

	defaultCommentService struct {
		cli zrpc.Client
	}
)

func NewCommentService(cli zrpc.Client) CommentService {
	return &defaultCommentService{
		cli: cli,
	}
}

func (m *defaultCommentService) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	client := v1.NewCommentServiceClient(m.cli.Conn())
	return client.CreateComment(ctx, in, opts...)
}

func (m *defaultCommentService) CreateCommentCompensate(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	client := v1.NewCommentServiceClient(m.cli.Conn())
	return client.CreateCommentCompensate(ctx, in, opts...)
}

func (m *defaultCommentService) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	client := v1.NewCommentServiceClient(m.cli.Conn())
	return client.GetComment(ctx, in, opts...)
}

func (m *defaultCommentService) GetCommentByUUID(ctx context.Context, in *GetCommentByUUIDRequest, opts ...grpc.CallOption) (*GetCommentByUUIDResponse, error) {
	client := v1.NewCommentServiceClient(m.cli.Conn())
	return client.GetCommentByUUID(ctx, in, opts...)
}

func (m *defaultCommentService) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	client := v1.NewCommentServiceClient(m.cli.Conn())
	return client.UpdateComment(ctx, in, opts...)
}

func (m *defaultCommentService) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	client := v1.NewCommentServiceClient(m.cli.Conn())
	return client.DeleteComment(ctx, in, opts...)
}

func (m *defaultCommentService) DeleteCommentCompensate(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	client := v1.NewCommentServiceClient(m.cli.Conn())
	return client.DeleteCommentCompensate(ctx, in, opts...)
}

func (m *defaultCommentService) DeleteCommentsByPostID(ctx context.Context, in *DeleteCommentsByPostIDRequest, opts ...grpc.CallOption) (*DeleteCommentsByPostIDResponse, error) {
	client := v1.NewCommentServiceClient(m.cli.Conn())
	return client.DeleteCommentsByPostID(ctx, in, opts...)
}

func (m *defaultCommentService) DeleteCommentsByPostIDCompensate(ctx context.Context, in *DeleteCommentsByPostIDRequest, opts ...grpc.CallOption) (*DeleteCommentsByPostIDResponse, error) {
	client := v1.NewCommentServiceClient(m.cli.Conn())
	return client.DeleteCommentsByPostIDCompensate(ctx, in, opts...)
}

func (m *defaultCommentService) ListCommentsByPostID(ctx context.Context, in *ListCommentsByPostIDRequest, opts ...grpc.CallOption) (*ListCommentsByPostIDResponse, error) {
	client := v1.NewCommentServiceClient(m.cli.Conn())
	return client.ListCommentsByPostID(ctx, in, opts...)
}