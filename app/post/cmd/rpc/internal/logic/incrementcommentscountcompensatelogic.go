package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncrementCommentsCountCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncrementCommentsCountCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncrementCommentsCountCompensateLogic {
	return &IncrementCommentsCountCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncrementCommentsCountCompensateLogic) IncrementCommentsCountCompensate(in *v1.IncrementCommentsCountRequest) (*v1.IncrementCommentsCountResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	id := in.GetId()
	post, err := l.svcCtx.PostsRepository.FindOne(l.ctx, conn, int64(id))
	if err != nil {
		return nil, err
	}

	post.CommentsCount--
	_, err = l.svcCtx.PostsRepository.UpdateWithVersion(l.ctx, conn, post)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update post: %v", err)
	}

	return &v1.IncrementCommentsCountResponse{}, nil
}
