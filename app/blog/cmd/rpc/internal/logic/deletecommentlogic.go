package logic

import (
	"context"
	"github.com/dtm-labs/client/dtmgrpc"
	commentv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	postv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	userv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/pkg/interceptor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
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

	commentID := req.GetId()
	commentResp, err := l.svcCtx.CommentRpc.GetComment(l.ctx, &commentv1.GetCommentRequest{
		Id: commentID,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	postID := commentResp.GetComment().GetPostId()
	postResp, err := l.svcCtx.PostRpc.GetPost(l.ctx, &postv1.GetPostRequest{
		Id: postID,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if userResp.GetUser().GetId() != commentResp.GetComment().GetUserId() || userResp.GetUser().GetId() != postResp.GetPost().GetUserId() {
		return nil, status.Error(codes.PermissionDenied, "user not authorized")
	}

	dtmGRPCServerAddr := l.svcCtx.Config.DTM.Server.GRPC.Port
	gid := dtmgrpc.MustGenGid(dtmGRPCServerAddr)
	logx.Info("gid:", gid)
	postTarget, err := l.svcCtx.Config.PostRpc.BuildTarget()
	if err != nil {
		logx.Error(err.Error())
		return nil, status.Error(codes.Internal, "build post target failed")
	}

	commentTarget, err := l.svcCtx.Config.CommentRpc.BuildTarget()
	if err != nil {
		logx.Error(err.Error())
		return nil, status.Error(codes.Internal, "build comment target failed")
	}

	saga := dtmgrpc.NewSagaGrpc(dtmGRPCServerAddr, gid).Add(
		commentTarget+commentv1.CommentService_DeleteComment_FullMethodName,
		commentTarget+commentv1.CommentService_DeleteCommentCompensate_FullMethodName,
		&commentv1.DeleteCommentRequest{
			Id: commentID,
		}).Add(
		postTarget+postv1.PostService_DecrementCommentsCount_FullMethodName,
		postTarget+postv1.PostService_DecrementCommentsCountCompensate_FullMethodName,
		&postv1.DecrementCommentsCountRequest{
			Id: postID,
		})

	saga.WaitResult = true
	err = saga.Submit()

	if err != nil {
		logx.Error("saga submit error:", err)
		return nil, status.Error(codes.Internal, "saga submit failed")
	}

	return &v1.DeleteCommentResponse{Success: true}, nil
}
