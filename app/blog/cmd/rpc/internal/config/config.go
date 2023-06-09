package config

import (
	"github.com/tiptok/gz-blog-microsevices/pkg/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Http struct {
		Port string
	}
	//UserRpc zrpc.RpcClientConf
	config.Config
}
