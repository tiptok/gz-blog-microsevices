package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentLogic) UpdateComment(req *v1.UpdateCommentRequest) (*v1.UpdateCommentResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	commentID := int64(req.GetComment().GetId())
	find, err := l.svcCtx.CommentsRepository.FindOne(l.ctx, conn, commentID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "comment %v not fount : %v", commentID, err)
	}
	if req.GetComment().GetContent() != "" {
		find.Content = req.GetComment().GetContent()
	}
	_, err = l.svcCtx.CommentsRepository.Update(l.ctx, conn, find)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not update comment: %v", err)
	}
	return &v1.UpdateCommentResponse{Success: true}, nil
}
