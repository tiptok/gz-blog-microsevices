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

type DeleteCommentsByPostIDCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentsByPostIDCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentsByPostIDCompensateLogic {
	return &DeleteCommentsByPostIDCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentsByPostIDCompensateLogic) DeleteCommentsByPostIDCompensate(req *v1.DeleteCommentsByPostIDRequest) (*v1.DeleteCommentsByPostIDResponse, error) {
	postID := req.GetPostId()
	conn := l.svcCtx.DefaultDBConn()
	_, comments, err := l.svcCtx.CommentsRepository.Find(l.ctx, conn, map[string]interface{}{
		"postId":   postID,
		"unscoped": true,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not found comments: %v", err)
	}
	for _, comment := range comments {
		comment.DeletedAt = sql.NullTime{}
		_, err = l.svcCtx.CommentsRepository.UpdateUnscoped(l.ctx, conn, comment)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "could not delete comments: %v", err)
		}
	}

	return &v1.DeleteCommentsByPostIDResponse{
		Success: true,
	}, nil
}
