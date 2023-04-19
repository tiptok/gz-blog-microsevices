package blog

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/api/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogic) SignUp(req *types.SignUpRequest) (resp *types.SignUpResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
