package logic

import (
	"context"
	"database/sql"
	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostCompensateLogic {
	return &DeletePostCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostCompensateLogic) DeletePostCompensate(in *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	id := in.GetId()
	post, err := l.svcCtx.PostsRepository.FindOneUnscoped(l.ctx, conn, int64(id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "post %d not found", in.GetId())
	}
	post.DeletedAt = sql.NullTime{}
	_, err = l.svcCtx.PostsRepository.UpdateUnscoped(l.ctx, conn, post)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to restore post: %v", err)
	}
	return &v1.DeletePostResponse{Success: true}, nil
}
