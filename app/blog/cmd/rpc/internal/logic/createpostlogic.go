package logic

import (
	"context"
	"github.com/google/uuid"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/postservice"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/userservice"
	"github.com/tiptok/gz-blog-microsevices/pkg/interceptor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePostLogic) CreatePost(in *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	userID, ok := l.ctx.Value(interceptor.ContextKeyID).(uint64)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user not authenticated")
	}
	userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userservice.GetUserRequest{
		Id: userID,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	title := in.GetPost().GetTitle()
	content := in.GetPost().GetContent()

	postResp, err := l.svcCtx.PostRpc.CreatePost(l.ctx, &postservice.CreatePostRequest{
		Post: &postservice.Post{
			Uuid:    uuid.New().String(),
			Title:   title,
			Content: content,
			UserId:  userResp.GetUser().GetId(),
		},
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.CreatePostResponse{
		Post: &v1.Post{
			Id:            postResp.GetPost().GetId(),
			Title:         postResp.GetPost().GetTitle(),
			Content:       postResp.GetPost().GetContent(),
			UserId:        postResp.GetPost().GetUserId(),
			CommentsCount: postResp.GetPost().GetCommentsCount(),
			CreatedAt:     postResp.GetPost().GetCreatedAt(),
			UpdatedAt:     postResp.GetPost().GetUpdatedAt(),
			User: &v1.User{
				Id:       userResp.GetUser().GetId(),
				Username: userResp.GetUser().GetUsername(),
				Avatar:   userResp.GetUser().GetAvatar(),
			},
		},
	}, nil
}
