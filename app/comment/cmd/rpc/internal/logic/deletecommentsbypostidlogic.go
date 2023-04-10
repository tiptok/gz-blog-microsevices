package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/comment/interanl/pkg/db/transaction"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentsByPostIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentsByPostIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentsByPostIDLogic {
	return &DeleteCommentsByPostIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentsByPostIDLogic) DeleteCommentsByPostID(req *v1.DeleteCommentsByPostIDRequest) (*v1.DeleteCommentsByPostIDResponse, error) {
	postID := req.GetPostId()
	conn := l.svcCtx.DefaultDBConn()
	_, comments, err := l.svcCtx.CommentsRepository.Find(l.ctx, conn, map[string]interface{}{
		"postId": postID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not found comments: %v", err)
	}
	err = transaction.UseTrans(l.ctx, l.svcCtx.DB, func(ctx context.Context, conn transaction.Conn) error {
		for _, comment := range comments {
			_, err = l.svcCtx.CommentsRepository.Delete(l.ctx, conn, comment)
			if err != nil {
				return err
			}
		}
		return nil
	}, true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete comments: %v", err)
	}

	return &v1.DeleteCommentsByPostIDResponse{
		Success: true,
	}, nil
}
