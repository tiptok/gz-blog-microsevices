package models

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Comments struct {
	Id        int64
	Uuid      string
	UserId    int64
	PostId    int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Version   int64
}

func (m *Comments) TableName() string {
	return "comments"
}

func (m *Comments) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	// m.DeleteTime = time.Now()
	return
}

func (m *Comments) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Comments) CacheKeyFunc() string {
	if m.Id == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:%v:id:%v", "project", m.TableName(), m.Id)
}

func (m *Comments) CacheKeyFuncByObject(obj interface{}) string {
	if v, ok := obj.(*Comments); ok {
		return v.CacheKeyFunc()
	}
	return ""
}

func (m *Comments) CachePrimaryKeyFunc() string {
	if len("") == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:%v:primarykey:%v", "project", m.TableName(), "key")
}
