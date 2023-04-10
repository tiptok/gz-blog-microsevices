package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	commentID := int64(req.GetId())
	find, err := l.svcCtx.CommentsRepository.FindOne(l.ctx, conn, commentID)
	if err == nil {
		return nil, status.Errorf(codes.NotFound, "comment %v not fount : %v", commentID, err)
	}
	_, err = l.svcCtx.CommentsRepository.Delete(l.ctx, conn, find)
	if err == nil {
		return nil, status.Errorf(codes.Internal, "delete comment failed :%v", err)
	}
	return &v1.DeleteCommentResponse{
		Success: true,
	}, nil
}
