package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type JWT struct {
	Secret  string        `json:",optional"`
	Expires time.Duration `json:",optional"`
}

type Config struct {
	UserRpc zrpc.RpcClientConf `json:",optional"`
	AuthRpc zrpc.RpcClientConf `json:",optional"`
	JWT     JWT                `json:",optional"`
}
