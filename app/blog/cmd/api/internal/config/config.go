package config

import (
	"github.com/tiptok/gz-blog-microsevices/pkg/config"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	config.Config
}
