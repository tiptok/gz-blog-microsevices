package models

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Posts struct {
	Id            int64
	Uuid          string
	UserId        int64
	Title         string
	Content       string
	CommentsCount int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"softDelete"`
	Version       int64
}

func (m *Posts) TableName() string {
	return "posts"
}

func (m *Posts) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	// m.DeleteTime = time.Now()
	return
}

func (m *Posts) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Posts) CacheKeyFunc() string {
	if m.Id == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:%v:id:%v", "project", m.TableName(), m.Id)
}

func (m *Posts) CacheKeyFuncByObject(obj interface{}) string {
	if v, ok := obj.(*Posts); ok {
		return v.CacheKeyFunc()
	}
	return ""
}

func (m *Posts) CachePrimaryKeyFunc() string {
	if len("") == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:%v:primarykey:%v", "project", m.TableName(), "key")
}
