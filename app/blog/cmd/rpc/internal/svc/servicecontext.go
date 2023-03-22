package svc

import (
	"context"
	"fmt"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/config"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/userservice"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"time"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userservice.UserService
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
	ur := userservice.NewUserService(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor)))
	return &ServiceContext{
		Config:  c,
		UserRpc: ur,
	}
}
