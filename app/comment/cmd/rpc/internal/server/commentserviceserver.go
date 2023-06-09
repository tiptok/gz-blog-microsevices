// Code generated by goctl. DO NOT EDIT!
// Source: comment.proto

package server

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/logic"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"
)

type CommentServiceServer struct {
	svcCtx *svc.ServiceContext
	v1.UnimplementedCommentServiceServer
}

func NewCommentServiceServer(svcCtx *svc.ServiceContext) *CommentServiceServer {
	return &CommentServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentServiceServer) CreateComment(ctx context.Context, in *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	l := logic.NewCreateCommentLogic(ctx, s.svcCtx)
	return l.CreateComment(in)
}

func (s *CommentServiceServer) CreateCommentCompensate(ctx context.Context, in *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	l := logic.NewCreateCommentCompensateLogic(ctx, s.svcCtx)
	return l.CreateCommentCompensate(in)
}

func (s *CommentServiceServer) GetComment(ctx context.Context, in *v1.GetCommentRequest) (*v1.GetCommentResponse, error) {
	l := logic.NewGetCommentLogic(ctx, s.svcCtx)
	return l.GetComment(in)
}

func (s *CommentServiceServer) GetCommentByUUID(ctx context.Context, in *v1.GetCommentByUUIDRequest) (*v1.GetCommentByUUIDResponse, error) {
	l := logic.NewGetCommentByUUIDLogic(ctx, s.svcCtx)
	return l.GetCommentByUUID(in)
}

func (s *CommentServiceServer) UpdateComment(ctx context.Context, in *v1.UpdateCommentRequest) (*v1.UpdateCommentResponse, error) {
	l := logic.NewUpdateCommentLogic(ctx, s.svcCtx)
	return l.UpdateComment(in)
}

func (s *CommentServiceServer) DeleteComment(ctx context.Context, in *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	l := logic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}

func (s *CommentServiceServer) DeleteCommentCompensate(ctx context.Context, in *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	l := logic.NewDeleteCommentCompensateLogic(ctx, s.svcCtx)
	return l.DeleteCommentCompensate(in)
}

func (s *CommentServiceServer) DeleteCommentsByPostID(ctx context.Context, in *v1.DeleteCommentsByPostIDRequest) (*v1.DeleteCommentsByPostIDResponse, error) {
	l := logic.NewDeleteCommentsByPostIDLogic(ctx, s.svcCtx)
	return l.DeleteCommentsByPostID(in)
}

func (s *CommentServiceServer) DeleteCommentsByPostIDCompensate(ctx context.Context, in *v1.DeleteCommentsByPostIDRequest) (*v1.DeleteCommentsByPostIDResponse, error) {
	l := logic.NewDeleteCommentsByPostIDCompensateLogic(ctx, s.svcCtx)
	return l.DeleteCommentsByPostIDCompensate(in)
}

func (s *CommentServiceServer) ListCommentsByPostID(ctx context.Context, in *v1.ListCommentsByPostIDRequest) (*v1.ListCommentsByPostIDResponse, error) {
	l := logic.NewListCommentsByPostIDLogic(ctx, s.svcCtx)
	return l.ListCommentsByPostID(in)
}
