package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/postservice"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"

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
	conn := l.svcCtx.DefaultDBConn()
	id := in.GetId()
	post, err := l.svcCtx.PostsRepository.FindOne(l.ctx, conn, int64(id))
	if err != nil {
		return nil, err
	}
	return &v1.GetPostResponse{
		Post: postservice.PostEntityToProtobuf(post),
	}, nil
}
