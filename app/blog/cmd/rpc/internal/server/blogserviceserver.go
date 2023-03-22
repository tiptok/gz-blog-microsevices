// Code generated by goctl. DO NOT EDIT!
// Source: blog.proto

package server

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/logic"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"
)

type BlogServiceServer struct {
	svcCtx *svc.ServiceContext
	v1.UnimplementedBlogServiceServer
}

func NewBlogServiceServer(svcCtx *svc.ServiceContext) *BlogServiceServer {
	return &BlogServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *BlogServiceServer) SignUp(ctx context.Context, in *v1.SignUpRequest) (*v1.SignUpResponse, error) {
	l := logic.NewSignUpLogic(ctx, s.svcCtx)
	return l.SignUp(in)
}

func (s *BlogServiceServer) SignIn(ctx context.Context, in *v1.SignInRequest) (*v1.SignInResponse, error) {
	l := logic.NewSignInLogic(ctx, s.svcCtx)
	return l.SignIn(in)
}

func (s *BlogServiceServer) CreatePost(ctx context.Context, in *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	l := logic.NewCreatePostLogic(ctx, s.svcCtx)
	return l.CreatePost(in)
}

func (s *BlogServiceServer) GetPost(ctx context.Context, in *v1.GetPostRequest) (*v1.GetPostResponse, error) {
	l := logic.NewGetPostLogic(ctx, s.svcCtx)
	return l.GetPost(in)
}

func (s *BlogServiceServer) ListPosts(ctx context.Context, in *v1.ListPostsRequest) (*v1.ListPostsResponse, error) {
	l := logic.NewListPostsLogic(ctx, s.svcCtx)
	return l.ListPosts(in)
}

func (s *BlogServiceServer) UpdatePost(ctx context.Context, in *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {
	l := logic.NewUpdatePostLogic(ctx, s.svcCtx)
	return l.UpdatePost(in)
}

func (s *BlogServiceServer) DeletePost(ctx context.Context, in *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	l := logic.NewDeletePostLogic(ctx, s.svcCtx)
	return l.DeletePost(in)
}

func (s *BlogServiceServer) CreateComment(ctx context.Context, in *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	l := logic.NewCreateCommentLogic(ctx, s.svcCtx)
	return l.CreateComment(in)
}

func (s *BlogServiceServer) DeleteComment(ctx context.Context, in *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	l := logic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}

func (s *BlogServiceServer) UpdateComment(ctx context.Context, in *v1.UpdateCommentRequest) (*v1.UpdateCommentResponse, error) {
	l := logic.NewUpdateCommentLogic(ctx, s.svcCtx)
	return l.UpdateComment(in)
}

func (s *BlogServiceServer) ListCommentsByPostID(ctx context.Context, in *v1.ListCommentsByPostIDRequest) (*v1.ListCommentsByPostIDResponse, error) {
	l := logic.NewListCommentsByPostIDLogic(ctx, s.svcCtx)
	return l.ListCommentsByPostID(in)
}
