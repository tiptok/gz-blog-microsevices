package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrementCommentsCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrementCommentsCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrementCommentsCountLogic {
	return &DecrementCommentsCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrementCommentsCountLogic) DecrementCommentsCount(in *v1.DecrementCommentsCountRequest) (*v1.DecrementCommentsCountResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	id := in.GetId()
	post, err := l.svcCtx.PostsRepository.FindOneUnscoped(l.ctx, conn, int64(id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "post %d not found", in.GetId())
	}
	post.CommentsCount--
	_, err = l.svcCtx.PostsRepository.UpdateUnscoped(l.ctx, conn, post)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to decrement post comments count : %v", err)
	}
	return &v1.DecrementCommentsCountResponse{}, nil
}
