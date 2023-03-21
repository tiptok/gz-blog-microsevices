package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/userservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByEmailLogic {
	return &GetUserByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByEmailLogic) GetUserByEmail(in *v1.GetUserByEmailRequest) (*v1.GetUserResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	user, err := l.svcCtx.UserRepository.FindOneByEmail(l.ctx, conn, in.Email)
	if err != nil {
		return nil, err
	}
	if in.GetPassword() != "" {
		ok := isCorrectPassword(user.Password, in.GetPassword())
		if !ok {
			return nil, status.Errorf(codes.Internal, "incorrect password")
		}
	}
	return &v1.GetUserResponse{
		User: userservice.UserEntityToProtobuf(user),
	}, nil
}
