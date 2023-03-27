package main

import (
	"flag"
	"fmt"
	"github.com/tiptok/gz-blog-microsevices/api/protobuf/user/v1"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/config"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/server"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()
	fmt.Println("loading config...")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUserServiceServer(ctx)
	fmt.Println("starting rpc 1...")
	s, err := zrpc.NewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		v1.RegisterUserServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("starting rpc 2...")
	defer s.Stop()
	logx.Info(c)
	logx.Info(fmt.Sprintf("Starting rpc server at %s...\n", c.ListenOn))
	s.Start()
}
