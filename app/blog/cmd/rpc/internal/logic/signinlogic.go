package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/auth/cmd/rpc/authservice"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/userservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (l *SignInLogic) SignIn(req *v1.SignInRequest) (*v1.SignInResponse, error) {
	email := req.GetEmail()
	username := req.GetUsername()
	password := req.GetPassword()
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

	return &v1.SignInResponse{
		Token: authResp.GetToken(),
	}, nil
}
