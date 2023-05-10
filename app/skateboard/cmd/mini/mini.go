package main

import (
	"flag"
	"fmt"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/db"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/config"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/handler"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/mini.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	db.Migrate(ctx.DB)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
