package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/auth/v1"
	"github.com/tiptok/gz-blog-microsevices/app/auth/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(req *v1.GenerateTokenRequest) (*v1.GenerateTokenResponse, error) {
	token, err := l.svcCtx.JwtManager.Generate(req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}
	return &v1.GenerateTokenResponse{
		Token: token,
	}, nil
}
