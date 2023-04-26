package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Products struct {
	Id          int64
	Name        string
	Description string
	Price       float32
	Unit        string
	Stock       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"softDelete"`
	Version     int64
}

func (m *Products) TableName() string {
	return "products"
}

func (m *Products) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Products) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Products) CacheKeyFunc() string {
	if m.Id == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:%v:id:%v", "project", m.TableName(), m.Id)
}

func (m *Products) CacheKeyFuncByObject(obj interface{}) string {
	if v, ok := obj.(*Products); ok {
		return v.CacheKeyFunc()
	}
	return ""
}

func (m *Products) CachePrimaryKeyFunc() string {
	if len("") == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:%v:primarykey:%v", "project", m.TableName(), "key")
}
