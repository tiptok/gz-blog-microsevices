package logic

import (
	"context"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/google/uuid"
	"github.com/tiptok/gz-blog-microsevices/pkg/interceptor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	commentv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	postv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	userv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"

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
	postID := req.GetComment().GetPostId()
	postResp, err := l.svcCtx.PostRpc.GetPost(l.ctx, &postv1.GetPostRequest{
		Id: postID,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	comment := &commentv1.Comment{
		Uuid:    uuid.New().String(),
		PostId:  postResp.GetPost().GetId(),
		UserId:  userResp.GetUser().GetId(),
		Content: req.GetComment().GetContent(),
	}
	dtmGRPCServerAddr := l.svcCtx.Config.DTM.Server.GRPC.Port // "etcd://localhost:2379/dtmservice"
	gid := dtmgrpc.MustGenGid(dtmGRPCServerAddr)
	logx.Info("gid:", gid)
	commentTarget, err := l.svcCtx.Config.CommentRpc.BuildTarget()
	if err != nil {
		logx.Error(err.Error())
		return nil, status.Error(codes.Internal, "build comment target failed")
	}
	postRpcTarget, err := l.svcCtx.Config.PostRpc.BuildTarget()
	if err != nil {
		logx.Error(err.Error())
		return nil, status.Error(codes.Internal, "build post target failed")
	}
	saga := dtmgrpc.NewSagaGrpc(dtmGRPCServerAddr, gid).Add(
		commentTarget+commentv1.CommentService_CreateComment_FullMethodName,
		commentTarget+commentv1.CommentService_CreateCommentCompensate_FullMethodName,
		&commentv1.CreateCommentRequest{
			Comment: comment,
		},
	).Add(
		postRpcTarget+postv1.PostService_IncrementCommentsCount_FullMethodName,
		postRpcTarget+postv1.PostService_DecrementCommentsCountCompensate_FullMethodName,
		&postv1.IncrementCommentsCountRequest{
			Id: postID,
		})
	saga.WaitResult = true
	err = saga.Submit()
	if err != nil {
		logx.Error("saga submit error:", err)
		return nil, status.Error(codes.Internal, "saga submit failed")
	}

	postUserResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userv1.GetUserRequest{
		Id: postResp.GetPost().GetUserId(),
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	commentResp, err := l.svcCtx.CommentRpc.GetCommentByUUID(l.ctx, &commentv1.GetCommentByUUIDRequest{
		Uuid: comment.GetUuid(),
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &v1.CreateCommentResponse{
		Comment: &v1.Comment{
			Id:        commentResp.GetComment().GetId(),
			Content:   commentResp.GetComment().GetContent(),
			PostId:    commentResp.GetComment().GetPostId(),
			UserId:    commentResp.GetComment().GetUserId(),
			CreatedAt: commentResp.GetComment().GetCreatedAt(),
			UpdatedAt: commentResp.GetComment().GetUpdatedAt(),
			Post: &v1.Post{
				Id:            postResp.GetPost().GetId(),
				Title:         postResp.GetPost().GetTitle(),
				Content:       postResp.GetPost().GetContent(),
				UserId:        postResp.GetPost().GetUserId(),
				CommentsCount: postResp.GetPost().GetCommentsCount(),
				CreatedAt:     postResp.GetPost().GetCreatedAt(),
				UpdatedAt:     postResp.GetPost().GetUpdatedAt(),
				User: &v1.User{
					Id:       postUserResp.GetUser().GetId(),
					Username: postUserResp.GetUser().GetUsername(),
					Avatar:   postUserResp.GetUser().GetAvatar(),
				},
			},
			User: &v1.User{
				Id:       userResp.GetUser().GetId(),
				Username: userResp.GetUser().GetUsername(),
				Avatar:   userResp.GetUser().GetAvatar(),
			},
		},
	}, nil
}
