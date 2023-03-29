package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncrementCommentsCountCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncrementCommentsCountCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncrementCommentsCountCompensateLogic {
	return &IncrementCommentsCountCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncrementCommentsCountCompensateLogic) IncrementCommentsCountCompensate(in *v1.IncrementCommentsCountRequest) (*v1.IncrementCommentsCountResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.IncrementCommentsCountResponse{}, nil
}
