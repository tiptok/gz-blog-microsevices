package logic

import (
	"context"
	"fmt"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/commentservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentByUUIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentByUUIDLogic {
	return &GetCommentByUUIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentByUUIDLogic) GetCommentByUUID(req *v1.GetCommentByUUIDRequest) (*v1.GetCommentByUUIDResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	find, err := l.svcCtx.CommentsRepository.FindOneByUuid(l.ctx, conn, req.GetUuid())
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("not found comment :%v", err.Error()))
	}
	return &v1.GetCommentByUUIDResponse{
		Comment: commentservice.CommentsEntityToProtobuf(find),
	}, nil
}
