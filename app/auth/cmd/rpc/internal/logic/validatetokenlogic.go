package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/auth/v1"
	"github.com/tiptok/gz-blog-microsevices/app/auth/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateTokenLogic) ValidateToken(req *v1.ValidateTokenRequest) (*v1.ValidateTokenResponse, error) {
	claims, err := l.svcCtx.JwtManager.Validate(req.GetToken())
	if claims.ID == 0 || err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")

	}
	return &v1.ValidateTokenResponse{
		Valid: true,
	}, nil
}
