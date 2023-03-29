package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrementCommentsCountCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrementCommentsCountCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrementCommentsCountCompensateLogic {
	return &DecrementCommentsCountCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrementCommentsCountCompensateLogic) DecrementCommentsCountCompensate(in *v1.DecrementCommentsCountRequest) (*v1.DecrementCommentsCountResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.DecrementCommentsCountResponse{}, nil
}
