package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/commentservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentsByPostIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCommentsByPostIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentsByPostIDLogic {
	return &ListCommentsByPostIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCommentsByPostIDLogic) ListCommentsByPostID(req *v1.ListCommentsByPostIDRequest) (*v1.ListCommentsByPostIDResponse, error) {
	postID := req.GetPostId()
	offset := req.GetOffset()
	limit := req.GetLimit()

	conn := l.svcCtx.DefaultDBConn()
	count, list, err := l.svcCtx.CommentsRepository.Find(l.ctx, conn, map[string]interface{}{
		"limit":  int(limit),
		"offset": int(offset),
		"postId": postID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get comments: %v", err)
	}

	var comments []*v1.Comment
	for _, comment := range list {
		comments = append(comments, commentservice.CommentsEntityToProtobuf(comment))
	}
	return &v1.ListCommentsByPostIDResponse{
		Comments: comments,
		Total:    uint64(count),
	}, nil
}
