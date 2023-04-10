package logic

import (
	"context"
	postv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	userv1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPostsLogic {
	return &ListPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPostsLogic) ListPosts(req *v1.ListPostsRequest) (*v1.ListPostsResponse, error) {
	offset := req.GetOffset()
	limit := req.GetLimit()
	postResp, err := l.svcCtx.PostRpc.ListPosts(l.ctx, &postv1.ListPostsRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var posts []*v1.Post
	var postUserIDs []uint64

	for _, post := range postResp.GetPosts() {
		postUserIDs = append(postUserIDs, post.GetUserId())
	}

	postUserResp, err := l.svcCtx.UserRpc.ListUsersByIDs(l.ctx, &userv1.ListUsersByIDsRequest{
		Ids: postUserIDs,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	for _, post := range postResp.GetPosts() {
		for _, postUser := range postUserResp.GetUsers() {
			if post.GetUserId() == postUser.GetId() {
				posts = append(posts, &v1.Post{
					Id:            post.GetId(),
					Title:         post.GetTitle(),
					Content:       post.GetContent(),
					UserId:        post.GetUserId(),
					CommentsCount: post.GetCommentsCount(),
					CreatedAt:     post.GetCreatedAt(),
					UpdatedAt:     post.GetUpdatedAt(),
					User: &v1.User{
						Id:       postUser.GetId(),
						Username: postUser.GetUsername(),
						Avatar:   postUser.GetAvatar(),
					},
				})
			}
		}
	}

	return &v1.ListPostsResponse{
		Posts: posts,
		Total: postResp.GetCount(),
	}, nil
}
