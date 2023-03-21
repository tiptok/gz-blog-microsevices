package userservice

import (
	v1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/app/user/interanl/pkg/domain"
)

func UserEntityToProtobuf(user *domain.Users) *v1.User {
	return &v1.User{
		Id:       uint64(user.Id),
		Uuid:     user.Uuid,
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		//CreatedAt: timestamppb.New(user.CreatedAt),
		//UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func UsersEntityToProtobuf(users []*domain.Users) []*v1.User {
	var ret = make([]*v1.User, 0)
	for _, user := range users {
		ret = append(ret, UserEntityToProtobuf(user))
	}
	return ret
}
