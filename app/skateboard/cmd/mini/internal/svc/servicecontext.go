package svc

import (
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/config"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/db/repository"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/db/transaction"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/domain"
	"github.com/tiptok/gz-blog-microsevices/pkg/cache"
	"github.com/tiptok/gz-blog-microsevices/pkg/database"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	DB             *gorm.DB
	ShopRepository domain.ShopRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := database.OpenGormPGDB(c.DB.DataSource)
	mlCache := cache.NewMultiLevelCache([]string{c.Redis.Host}, c.Redis.Pass)
	return &ServiceContext{
		Config:         c,
		DB:             db,
		ShopRepository: repository.NewShopRepository(cache.NewCachedRepository(mlCache)),
	}
}

func (svc *ServiceContext) DefaultDBConn() transaction.Conn {
	return transaction.NewTransactionContext(svc.DB)
}
