// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/logic"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/svc"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	v1.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) ListUsersByIDs(ctx context.Context, in *v1.ListUsersByIDsRequest) (*v1.ListUsersByIDsResponse, error) {
	l := logic.NewListUsersByIDsLogic(ctx, s.svcCtx)
	return l.ListUsersByIDs(in)
}

func (s *UserServiceServer) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UserServiceServer) GetUserByEmail(ctx context.Context, in *v1.GetUserByEmailRequest) (*v1.GetUserResponse, error) {
	l := logic.NewGetUserByEmailLogic(ctx, s.svcCtx)
	return l.GetUserByEmail(in)
}

func (s *UserServiceServer) GetUserByUsername(ctx context.Context, in *v1.GetUserByUsernameRequest) (*v1.GetUserResponse, error) {
	l := logic.NewGetUserByUsernameLogic(ctx, s.svcCtx)
	return l.GetUserByUsername(in)
}

func (s *UserServiceServer) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *UserServiceServer) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	l := logic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}
