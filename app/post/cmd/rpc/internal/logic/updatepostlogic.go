package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePostLogic) UpdatePost(in *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	id := in.GetPost().GetId()
	post, err := l.svcCtx.PostsRepository.FindOne(l.ctx, conn, int64(id))
	if err != nil {
		return nil, err
	}
	if in.GetPost().GetTitle() != "" {
		post.Title = in.GetPost().GetTitle()
	}
	if in.GetPost().GetContent() != "" {
		post.Content = in.GetPost().GetContent()
	}
	logx.Info("update post", post)
	_, err = l.svcCtx.PostsRepository.UpdateWithVersion(l.ctx, conn, post)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update post: %v", err)
	}

	return &v1.UpdatePostResponse{
		Success: true,
	}, nil
}
