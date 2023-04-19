package svc

import (
	"context"
	"fmt"
	"github.com/tiptok/gz-blog-microsevices/app/auth/cmd/rpc/authservice"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/api/internal/config"
	"github.com/tiptok/gz-blog-microsevices/app/comment/cmd/rpc/commentservice"
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/postservice"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/userservice"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"time"
)

type ServiceContext struct {
	Config     config.Config
	UserRpc    userservice.UserService
	AuthRpc    authservice.AuthService
	PostRpc    postservice.PostService
	CommentRpc commentservice.CommentService
}

func timeInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	stime := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	fmt.Printf("调用 %s 方法 耗时: %v\n", method, time.Now().Sub(stime))
	return nil
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRpc:    userservice.NewUserService(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor))),
		AuthRpc:    authservice.NewAuthService(zrpc.MustNewClient(c.AuthRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor))),
		PostRpc:    postservice.NewPostService(zrpc.MustNewClient(c.PostRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor))),
		CommentRpc: commentservice.NewCommentService(zrpc.MustNewClient(c.CommentRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor))),
	}
}
