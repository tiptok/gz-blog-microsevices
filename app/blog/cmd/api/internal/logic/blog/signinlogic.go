package blog

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/auth/cmd/rpc/authservice"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/userservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/api/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInRequest) (resp *types.SignInResponse, err error) {
	email := req.Email
	username := req.Username
	password := req.Password
	if email == "" && username == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email and username cannot be empty")
	}
	var userID uint64
	if email != "" {
		resp, err := l.svcCtx.UserRpc.GetUserByEmail(l.ctx, &userservice.GetUserByEmailRequest{
			Email:    email,
			Password: password,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get user by email: %v", err)
		}
		user := resp.GetUser()
		userID = user.GetId()
	} else {
		resp, err := l.svcCtx.UserRpc.GetUserByUsername(l.ctx, &userservice.GetUserByUsernameRequest{
			Username: username,
			Password: password,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get user by email: %v", err)
		}
		user := resp.GetUser()
		userID = user.GetId()
	}

	authResp, err := l.svcCtx.AuthRpc.GenerateToken(l.ctx, &authservice.GenerateTokenRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return &types.SignInResponse{
		Token: authResp.GetToken(),
	}, nil
}
