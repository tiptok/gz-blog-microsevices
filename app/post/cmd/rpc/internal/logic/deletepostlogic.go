package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostLogic) DeletePost(in *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	id := in.GetId()
	post, err := l.svcCtx.PostsRepository.FindOne(l.ctx, conn, int64(id))
	if err != nil {
		return nil, err
	}
	if _, err := l.svcCtx.PostsRepository.Delete(l.ctx, conn, post); err != nil {
		return nil, err
	}
	return &v1.DeletePostResponse{
		Success: true,
	}, nil
}
