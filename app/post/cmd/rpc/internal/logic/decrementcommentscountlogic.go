package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrementCommentsCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrementCommentsCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrementCommentsCountLogic {
	return &DecrementCommentsCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrementCommentsCountLogic) DecrementCommentsCount(in *v1.DecrementCommentsCountRequest) (*v1.DecrementCommentsCountResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.DecrementCommentsCountResponse{}, nil
}
