package db

import (
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/db/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Shop{},
	)
}
