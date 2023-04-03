package svc

import (
	"github.com/tiptok/gz-blog-microsevices/app/post/cmd/rpc/internal/config"
	"github.com/tiptok/gz-blog-microsevices/app/post/interanl/pkg/db/repository"
	"github.com/tiptok/gz-blog-microsevices/app/post/interanl/pkg/db/transaction"
	"github.com/tiptok/gz-blog-microsevices/app/post/interanl/pkg/domain"
	"github.com/tiptok/gz-blog-microsevices/pkg/cache"
	"github.com/tiptok/gz-blog-microsevices/pkg/database"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	PostsRepository domain.PostsRepository
	DB              *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := database.OpenGormDB(c.DB.DataSource)
	mlCache := cache.NewMultiLevelCache([]string{c.Redis.Host}, c.Redis.Pass)
	return &ServiceContext{
		Config:          c,
		DB:              db,
		PostsRepository: repository.NewPostsRepository(cache.NewCachedRepository(mlCache)),
	}
}

func (svc *ServiceContext) DefaultDBConn() transaction.Conn {
	return transaction.NewTransactionContext(svc.DB)
}
