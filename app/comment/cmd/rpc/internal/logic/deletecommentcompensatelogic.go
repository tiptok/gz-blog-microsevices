package logic

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentCompensateLogic {
	return &DeleteCommentCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentCompensateLogic) DeleteCommentCompensate(req *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	commentID := int64(req.GetId())
	find, err := l.svcCtx.CommentsRepository.FindOneUnscoped(l.ctx, conn, commentID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "comment %v not fount : %v", commentID, err)
	}
	find.DeletedAt = sql.NullTime{}
	find, err = l.svcCtx.CommentsRepository.UpdateUnscoped(l.ctx, conn, find)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to restore comment: %v", err)
	}
	return &v1.DeleteCommentResponse{}, nil
}
