package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentCompensateLogic {
	return &DeleteCommentCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentCompensateLogic) DeleteCommentCompensate(in *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.DeleteCommentResponse{}, nil
}
