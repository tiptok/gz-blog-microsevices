package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentsByPostIDCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentsByPostIDCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentsByPostIDCompensateLogic {
	return &DeleteCommentsByPostIDCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentsByPostIDCompensateLogic) DeleteCommentsByPostIDCompensate(in *v1.DeleteCommentsByPostIDRequest) (*v1.DeleteCommentsByPostIDResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.DeleteCommentsByPostIDResponse{}, nil
}
