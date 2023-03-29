package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentByUUIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentByUUIDLogic {
	return &GetCommentByUUIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentByUUIDLogic) GetCommentByUUID(in *v1.GetCommentByUUIDRequest) (*v1.GetCommentByUUIDResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.GetCommentByUUIDResponse{}, nil
}
