package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/postservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"

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

func (l *ListPostsLogic) ListPosts(in *v1.ListPostsRequest) (*v1.ListPostsResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	count, list, err := l.svcCtx.PostsRepository.Find(l.ctx, conn, map[string]interface{}{"offset": int(in.Offset), "limit": int(in.Limit)})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list posts: %v", err)
	}

	var posts []*postservice.Post
	for _, post := range list {
		posts = append(posts, postservice.PostEntityToProtobuf(post))
	}

	return &v1.ListPostsResponse{
		Posts: posts,
		Count: uint64(count),
	}, nil
}
