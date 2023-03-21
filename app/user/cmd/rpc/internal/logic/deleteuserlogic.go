package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	id := in.GetId()
	user, err := l.svcCtx.UserRepository.FindOne(l.ctx, conn, int64(id))
	if err != nil {
		return nil, err
	}
	if _, err := l.svcCtx.UserRepository.Delete(l.ctx, conn, user); err != nil {
		return nil, err
	}
	return &v1.DeleteUserResponse{}, nil
}
