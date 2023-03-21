package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/user/interanl/pkg/db/transaction"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	id := in.GetUser().GetId()
	conn := l.svcCtx.DefaultDBConn()
	user, err := l.svcCtx.UserRepository.FindOne(l.ctx, conn, int64(id))
	if err != nil {
		return nil, err
	}
	if err := transaction.UseTrans(l.ctx, l.svcCtx.DB, func(ctx context.Context, conn transaction.Conn) error {
		if in.GetUser().GetUsername() != "" {
			user.Username = in.GetUser().GetUsername()
		}
		if in.GetUser().Email != "" {
			user.Email = in.GetUser().Email
		}
		if in.GetUser().Avatar != "" {
			user.Avatar = in.GetUser().Avatar
		}
		if in.GetUser().GetPassword() != "" {
			password, err := generateFromPassword(in.GetUser().GetPassword())
			if err != nil {
				return status.Errorf(codes.Internal, "failed to bcrypt generate password: %v", err)
			}
			user.Password = password
		}
		if _, err := l.svcCtx.UserRepository.Update(l.ctx, conn, user); err != nil {
			return err
		}
		return nil
	}, true); err != nil {
		return nil, err
	}

	return &v1.UpdateUserResponse{Success: true}, nil
}

func generateFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func isCorrectPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
