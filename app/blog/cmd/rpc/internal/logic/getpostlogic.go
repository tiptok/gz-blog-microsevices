package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	postv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	userv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLogic {
	return &GetPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostLogic) GetPost(in *v1.GetPostRequest) (*v1.GetPostResponse, error) {
	postID := in.GetId()
	postResp, err := l.svcCtx.PostRpc.GetPost(l.ctx, &postv1.GetPostRequest{Id: postID})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	postUserResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userv1.GetUserRequest{Id: postResp.GetPost().GetUserId()})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &v1.GetPostResponse{
		Post: &v1.Post{
			Id:            postResp.GetPost().GetId(),
			Title:         postResp.GetPost().GetTitle(),
			Content:       postResp.GetPost().GetContent(),
			UserId:        postResp.GetPost().GetUserId(),
			CommentsCount: postResp.GetPost().GetCommentsCount(),
			CreatedAt:     postResp.GetPost().GetCreatedAt(),
			UpdatedAt:     postResp.GetPost().GetUpdatedAt(),
			User: &v1.User{
				Id:       postUserResp.GetUser().GetId(),
				Username: postUserResp.GetUser().GetUsername(),
				Avatar:   postUserResp.GetUser().GetAvatar(),
			},
		},
	}, nil
}
