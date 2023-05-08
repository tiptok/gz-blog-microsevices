package svc

import (
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
