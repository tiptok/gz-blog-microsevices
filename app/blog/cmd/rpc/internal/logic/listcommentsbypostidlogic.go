package logic

import (
	"context"
	commentv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	userv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"

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

	commentResp, err := l.svcCtx.CommentRpc.ListCommentsByPostID(l.ctx, &commentv1.ListCommentsByPostIDRequest{
		PostId: postID,
		Offset: int32(offset),
		Limit:  int32(limit),
	})

	var commentUserIDs []uint64
	for _, post := range commentResp.GetComments() {
		commentUserIDs = append(commentUserIDs, post.GetUserId())
	}

	commentUserResp, err := l.svcCtx.UserRpc.ListUsersByIDs(l.ctx, &userv1.ListUsersByIDsRequest{
		Ids: commentUserIDs,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var comments []*v1.Comment
	for _, comment := range commentResp.GetComments() {
		for _, user := range commentUserResp.GetUsers() {
			if user.GetId() == comment.GetUserId() {
				comments = append(comments, &v1.Comment{
					Id:        comment.GetId(),
					Content:   comment.GetContent(),
					PostId:    comment.GetPostId(),
					UserId:    comment.GetUserId(),
					CreatedAt: comment.GetCreatedAt(),
					UpdatedAt: comment.GetUpdatedAt(),
					User: &v1.User{
						Id:       user.GetId(),
						Username: user.GetUsername(),
						Avatar:   user.GetAvatar(),
					},
				})
			}
		}
	}
	return &v1.ListCommentsByPostIDResponse{
		Comments: comments,
		Total:    commentResp.GetTotal(),
	}, nil
}
