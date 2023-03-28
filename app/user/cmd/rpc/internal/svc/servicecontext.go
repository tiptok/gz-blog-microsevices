package svc

import (
	"fmt"
	"github.com/tiptok/gocomm/pkg/cache"
	"github.com/tiptok/gocomm/pkg/cache/gzcache"
	gocommlog "github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gz-blog-microsevices/app/user/cmd/rpc/internal/config"
	"github.com/tiptok/gz-blog-microsevices/app/user/interanl/pkg/db/repository"
	"github.com/tiptok/gz-blog-microsevices/app/user/interanl/pkg/db/transaction"
	"github.com/tiptok/gz-blog-microsevices/app/user/interanl/pkg/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type ServiceContext struct {
	Config         config.Config
	UserRepository domain.UsersRepository
	DB             *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	fmt.Println("starting db...")
	db, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("starting multi level cache...")
	mlCache := cache.NewMultiLevelCacheNew(cache.WithDebugLog(true, func() gocommlog.Log {
		return gocommlog.DefaultLog
	}))
	mlCache.RegisterCache(gzcache.NewClusterCache([]string{c.Redis.Host}, c.Redis.Pass))

	return &ServiceContext{
		Config:         c,
		DB:             db,
		UserRepository: repository.NewUsersRepository(cache.NewCachedRepository(mlCache)), //repository.NewUsersRepository(nil),
	}
}

func (svc *ServiceContext) DefaultDBConn() transaction.Conn {
	return transaction.NewTransactionContext(svc.DB)
}
