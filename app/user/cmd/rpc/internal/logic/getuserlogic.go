package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/userservice"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	id := in.GetId()
	conn := l.svcCtx.DefaultDBConn()
	user, err := l.svcCtx.UserRepository.FindOne(l.ctx, conn, int64(id))
	if err != nil {
		return nil, err
	}
	return &v1.GetUserResponse{
		User: userservice.UserEntityToProtobuf(user),
	}, nil
}
