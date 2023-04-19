package blog

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/api/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostsLogic {
	return &CreatePostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostsLogic) CreatePosts(req *types.CreatePostRequest) (resp *types.CreatePostResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
