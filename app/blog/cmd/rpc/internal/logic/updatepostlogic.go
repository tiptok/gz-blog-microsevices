package logic

import (
	"context"
	postv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	userv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/pkg/interceptor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePostLogic) UpdatePost(req *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {
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
	postID := req.GetPost().GetId()
	postResp, err := l.svcCtx.PostRpc.GetPost(l.ctx, &postv1.GetPostRequest{
		Id: postID,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	// 授权检查，只能删除自己发布的文章
	if userResp.GetUser().GetId() != postResp.GetPost().GetUserId() {
		return nil, status.Error(codes.PermissionDenied, "user not authorized")
	}
	updatedPost := &postv1.Post{
		Id: req.GetPost().GetId(),
	}

	if req.GetPost().GetTitle() != "" {
		updatedPost.Title = req.GetPost().GetTitle()
	}

	if req.GetPost().GetContent() != "" {
		updatedPost.Content = req.GetPost().GetContent()
	}

	updatePostResp, err := l.svcCtx.PostRpc.UpdatePost(l.ctx, &postv1.UpdatePostRequest{
		Post: updatedPost,
	})

	if err != nil || !updatePostResp.GetSuccess() {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.UpdatePostResponse{Success: true}, nil
}
