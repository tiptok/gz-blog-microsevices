package models

import (
	"fmt"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/domain"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Shop struct {
	Id           int64
	Name         string
	Introduction string
	Address      *domain.Address `gorm:"serializer:json"`
	Context      *domain.Context `gorm:"serializer:json"`
	Rank         int
	CreatedAt    int64
	UpdatedAt    int64
	DeletedAt    int64
	IsDel        soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
	Version      int
}

func (m *Shop) TableName() string {
	return "skt_shop"
}

func (m *Shop) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (m *Shop) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (m *Shop) CacheKeyFunc() string {
	if m.Id == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:%v:id:%v", "project", m.TableName(), m.Id)
}

func (m *Shop) CacheKeyFuncByObject(obj interface{}) string {
	if v, ok := obj.(*Shop); ok {
		return v.CacheKeyFunc()
	}
	return ""
}

func (m *Shop) CachePrimaryKeyFunc() string {
	if len("") == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:%v:primarykey:%v", "project", m.TableName(), "key")
}
