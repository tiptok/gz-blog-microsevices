package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/commentservice"
	"github.com/tiptok/gz-blog-microsevices/app/comment/interanl/pkg/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(req *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	find, err := l.svcCtx.CommentsRepository.FindOneByUuid(l.ctx, conn, req.GetComment().GetUuid())
	if err == nil {
		return &v1.CreateCommentResponse{
			Comment: commentservice.CommentsEntityToProtobuf(find),
		}, nil
	}

	comment := &domain.Comments{
		Uuid:    req.GetComment().GetUuid(),
		Content: req.GetComment().GetContent(),
		PostId:  int64(req.GetComment().GetPostId()),
		UserId:  int64(req.GetComment().GetUserId()),
	}

	comment, err = l.svcCtx.CommentsRepository.Insert(l.ctx, conn, comment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create comment: %v", err)
	}
	return &v1.CreateCommentResponse{
		Comment: commentservice.CommentsEntityToProtobuf(comment),
	}, nil
}
