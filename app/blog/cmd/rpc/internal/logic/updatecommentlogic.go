package logic

import (
	"context"
	commentv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	userv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/pkg/interceptor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"

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
	userID, ok := l.ctx.Value(interceptor.ContextKeyID).(uint64)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user not authenticated")
	}
	userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userv1.GetUserRequest{
		Id: userID,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	commentID := req.GetComment().GetId()
	commentResp, err := l.svcCtx.CommentRpc.GetComment(l.ctx, &commentv1.GetCommentRequest{
		Id: commentID,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if userResp.GetUser().GetId() != commentResp.GetComment().GetUserId() {
		return nil, status.Error(codes.PermissionDenied, "user not authorized")
	}
	comment := &commentv1.Comment{
		Id:      commentResp.GetComment().GetId(),
		Content: req.GetComment().GetContent(),
	}

	_, err = l.svcCtx.CommentRpc.UpdateComment(l.ctx, &commentv1.UpdateCommentRequest{
		Comment: comment,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &v1.UpdateCommentResponse{Success: true}, nil
}
