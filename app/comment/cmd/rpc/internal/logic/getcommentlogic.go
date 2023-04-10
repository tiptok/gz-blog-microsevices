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

type GetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentLogic) GetComment(req *v1.GetCommentRequest) (*v1.GetCommentResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	commentID := int64(req.GetId())
	find, err := l.svcCtx.CommentsRepository.FindOne(l.ctx, conn, commentID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "comment %v not fount : %v", commentID, err)
	}
	return &v1.GetCommentResponse{
		Comment: commentservice.CommentsEntityToProtobuf(find),
	}, nil
}
