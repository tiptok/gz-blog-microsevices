package logic

import (
	"context"
	"github.com/google/uuid"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/userservice"
	"github.com/tiptok/gz-blog-microsevices/app/user/interanl/pkg/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	password, err := generateFromPassword(in.GetUser().GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to bcrypt generate password: %v", err)
	}

	user := &domain.Users{
		Uuid:     uuid.New().String(),
		Username: in.GetUser().GetUsername(),
		Email:    in.GetUser().GetEmail(),
		Avatar:   in.GetUser().GetAvatar(),
		Password: password,
	}
	conn := l.svcCtx.DefaultDBConn()
	user, err = l.svcCtx.UserRepository.Insert(l.ctx, conn, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	return &v1.CreateUserResponse{
		User: userservice.UserEntityToProtobuf(user),
	}, nil
}
