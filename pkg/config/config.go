package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type JWT struct {
	Secret  string        `json:",optional"`
	Expires time.Duration `json:",optional"`
}

type Config struct {
	UserRpc    zrpc.RpcClientConf `json:",optional"`
	AuthRpc    zrpc.RpcClientConf `json:",optional"`
	PostRpc    zrpc.RpcClientConf `json:",optional"`
	CommentRpc zrpc.RpcClientConf `json:",optional"`
	JWT        JWT                `json:",optional"`
	DB         struct {
		DataSource string
	} `json:",optional"`
	Cache cache.CacheConf `json:",optional"`
	DTM   DTM             `json:",optional"`
}

type DTM struct {
	Server Server
}

type Server struct {
	Name    string
	Host    string
	GRPC    GRPC
	HTTP    HTTP
	Metrics Metrics
}

type HTTP struct {
	Port string
}

type GRPC struct {
	Port string
}

type Metrics struct {
	Port string
}
