package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentsByPostIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentsByPostIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentsByPostIDLogic {
	return &DeleteCommentsByPostIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentsByPostIDLogic) DeleteCommentsByPostID(in *v1.DeleteCommentsByPostIDRequest) (*v1.DeleteCommentsByPostIDResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.DeleteCommentsByPostIDResponse{}, nil
}
