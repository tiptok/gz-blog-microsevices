package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentCompensateLogic {
	return &CreateCommentCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentCompensateLogic) CreateCommentCompensate(in *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.CreateCommentResponse{}, nil
}
