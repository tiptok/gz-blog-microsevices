package models

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id        int64
	Uuid      string
	Username  string
	Email     string
	Avatar    string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"softDelete"`
	Version   int
}

func (m *Users) TableName() string {
	return "users"
}

func (m *Users) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	//m.DeleteTime = time.Now()
	return
}

func (m *Users) BeforeUpdate(tx *gorm.DB) (err error) {
	// m.UpdateTime = time.Now()
	return
}

func (m *Users) CacheKeyFunc() string {
	if m.Id == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:%v:id:%v", "project", m.TableName(), m.Id)
}

func (m *Users) CacheKeyFuncByObject(obj interface{}) string {
	if v, ok := obj.(*Users); ok {
		return v.CacheKeyFunc()
	}
	return ""
}

func (m *Users) CachePrimaryKeyFunc() string {
	if len("") == 0 {
		return ""
	}
	return fmt.Sprintf("%v:cache:%v:primarykey:%v", "project", m.TableName(), "key")
}
