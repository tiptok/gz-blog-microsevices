package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/api/protobuf/post/v1"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/postservice"
	"github.com/tiptok/gz-blog-microsevices/app/post/interanl/pkg/domain"
	"github.com/tiptok/gz-blog-microsevices/app/user/interanl/pkg/db/transaction"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePostLogic) CreatePost(req *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	post := &domain.Posts{
		Uuid:    req.GetPost().GetUuid(),
		Title:   req.GetPost().GetTitle(),
		Content: req.GetPost().GetContent(),
		UserId:  int64(req.GetPost().GetUserId()),
	}
	var err error
	err = transaction.UseTrans(l.ctx, l.svcCtx.DB, func(ctx context.Context, conn transaction.Conn) error {
		post, err = l.svcCtx.PostsRepository.Insert(l.ctx, conn, post)
		if err != nil {
			return err
		}
		return nil
	}, true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create post: %v", err)
	}

	return &v1.CreatePostResponse{
		Post: postservice.PostEntityToProtobuf(post),
	}, nil
}
