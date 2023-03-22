package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignInLogic) SignIn(in *v1.SignInRequest) (*v1.SignInResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.SignInResponse{}, nil
}
