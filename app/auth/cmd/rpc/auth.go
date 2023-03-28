package main

import (
	"flag"
	"fmt"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/auth/v1"
	"github.com/tiptok/gz-blog-microsevices/app/auth/cmd/rpc/internal/config"
	"github.com/tiptok/gz-blog-microsevices/app/auth/cmd/rpc/internal/server"
	"github.com/tiptok/gz-blog-microsevices/app/auth/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewAuthServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		v1.RegisterAuthServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
