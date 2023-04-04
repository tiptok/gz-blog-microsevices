package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentCompensateLogic {
	return &CreateCommentCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentCompensateLogic) CreateCommentCompensate(req *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	comment, err := l.svcCtx.CommentsRepository.FindOneByUuid(l.ctx, conn, req.GetComment().GetUuid())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not find comment: %v", err)
	}
	_, err = l.svcCtx.CommentsRepository.Delete(l.ctx, conn, comment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create comment: %v", err)
	}
	return &v1.CreateCommentResponse{}, nil
}
