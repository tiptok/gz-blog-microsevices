package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/userservice"

	"github.com/samber/lo"
	"github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListUsersByIDsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUsersByIDsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUsersByIDsLogic {
	return &ListUsersByIDsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUsersByIDsLogic) ListUsersByIDs(in *v1.ListUsersByIDsRequest) (*v1.ListUsersByIDsResponse, error) {
	ids := lo.Map[uint64, int64](in.GetIds(), func(item uint64, index int) int64 {
		return int64(item)
	})
	conn := l.svcCtx.DefaultDBConn()
	_, users, err := l.svcCtx.UserRepository.FindByIDs(l.ctx, conn, ids)
	if err != nil {
		return nil, err
	}
	return &v1.ListUsersByIDsResponse{
		Users: userservice.UsersEntityToProtobuf(users),
	}, nil
}
