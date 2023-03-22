package logic

import (
	"context"
	"github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/userservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUpLogic) SignUp(req *v1.SignUpRequest) (*v1.SignUpResponse, error) {
	username := req.GetUsername()
	email := req.GetEmail()
	password := req.GetPassword()

	emailResp, err := l.svcCtx.UserRpc.GetUserByEmail(l.ctx, &userservice.GetUserByEmailRequest{
		Email: email,
	})

	if err == nil && emailResp.GetUser().GetId() != 0 {
		return nil, status.Error(codes.AlreadyExists, "email already exists")
	}

	userResp, err := l.svcCtx.UserRpc.CreateUser(l.ctx, &userservice.CreateUserRequest{
		User: &userservice.User{
			Username: username,
			Email:    email,
			Password: password,
		},
	})
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &v1.SignUpResponse{
		Token: userResp.User.Uuid,
	}, nil
}
