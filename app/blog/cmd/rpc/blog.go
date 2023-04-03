package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/config"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/server"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/rpc/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/pkg/interceptor"
	"github.com/tiptok/gz-blog-microsevices/pkg/jwt"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/blog.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewBlogServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		v1.RegisterBlogServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(interceptor.NewAuthInterceptor(jwt.NewManager(nil, &c.Config), server.AuthMethods).Unary())
	defer s.Stop()

	go func() {
		httpMux := runtime.NewServeMux()
		//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		//v1.RegisterBlogServiceHandlerFromEndpoint(context.Background(), httpMux, c.ListenOn, opts)
		if err := v1.RegisterBlogServiceHandlerServer(context.Background(), httpMux, svr); err != nil {
			log.Fatal(err)
		}
		httpServer := &http.Server{
			Addr:    c.Http.Port,
			Handler: interceptor.WithAuthInterceptor(httpMux, jwt.NewManager(nil, &c.Config), server.AuthHttpMethods),
		}
		fmt.Printf("Starting api server at %s...\n", c.Http.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
