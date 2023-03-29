package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentsByPostIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCommentsByPostIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentsByPostIDLogic {
	return &ListCommentsByPostIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCommentsByPostIDLogic) ListCommentsByPostID(in *v1.ListCommentsByPostIDRequest) (*v1.ListCommentsByPostIDResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.ListCommentsByPostIDResponse{}, nil
}
