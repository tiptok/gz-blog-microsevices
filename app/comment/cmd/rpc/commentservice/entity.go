package commentservice

import (
	"github.com/tiptok/gz-blog-microsevices/api/protobuf/comment/v1"
	"github.com/tiptok/gz-blog-microsevices/app/comment/interanl/pkg/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CommentsEntityToProtobuf(comment *domain.Comments) *v1.Comment {
	return &v1.Comment{
		Id:        uint64(comment.Id),
		Uuid:      comment.Uuid,
		Content:   comment.Content,
		PostId:    uint64(comment.PostId),
		UserId:    uint64(comment.UserId),
		CreatedAt: timestamppb.New(comment.CreatedAt),
		UpdatedAt: timestamppb.New(comment.UpdatedAt),
	}
}
