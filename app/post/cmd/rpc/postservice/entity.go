package postservice

import (
	"github.com/tiptok/gz-blog-microsevices/app/post/interanl/pkg/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PostEntityToProtobuf(post *domain.Posts) *Post {
	return &Post{
		Id:            uint64(post.Id),
		Uuid:          post.Uuid,
		Title:         post.Title,
		Content:       post.Content,
		CommentsCount: uint32(post.CommentsCount),
		UserId:        uint64(post.UserId),
		CreatedAt:     timestamppb.New(post.CreatedAt),
		UpdatedAt:     timestamppb.New(post.UpdatedAt),
	}
}
