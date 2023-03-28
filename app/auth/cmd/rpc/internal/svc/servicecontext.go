package svc

import (
	"github.com/tiptok/gz-blog-microsevices/app/auth/cmd/rpc/internal/config"
	"github.com/tiptok/gz-blog-microsevices/pkg/jwt"
)

type ServiceContext struct {
	Config     config.Config
	JwtManager *jwt.Manager
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		JwtManager: jwt.NewManager(nil, &c.Config),
	}
}
