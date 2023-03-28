package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/auth/v1"
	"github.com/tiptok/gz-blog-microsevices/app/auth/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *v1.RefreshTokenRequest) (*v1.RefreshTokenResponse, error) {
	claims, err := l.svcCtx.JwtManager.Validate(req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	token, err := l.svcCtx.JwtManager.Generate(claims.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}
	return &v1.RefreshTokenResponse{
		Token: token,
	}, nil
}
