package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Http struct {
		Port string
	}
	UserRpc zrpc.RpcClientConf
}
