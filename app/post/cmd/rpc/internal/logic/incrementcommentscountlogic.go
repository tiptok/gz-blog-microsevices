package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncrementCommentsCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncrementCommentsCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncrementCommentsCountLogic {
	return &IncrementCommentsCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncrementCommentsCountLogic) IncrementCommentsCount(in *v1.IncrementCommentsCountRequest) (*v1.IncrementCommentsCountResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.IncrementCommentsCountResponse{}, nil
}
