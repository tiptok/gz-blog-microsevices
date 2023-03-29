package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostCompensateLogic {
	return &DeletePostCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostCompensateLogic) DeletePostCompensate(in *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.DeletePostResponse{}, nil
}
