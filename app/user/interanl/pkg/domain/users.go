package domain

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/user/interanl/pkg/db/transaction"
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
	DeletedAt gorm.DeletedAt `gorm:"softDelete:milli"`
	Version   int
}

type UsersRepository interface {
	Insert(ctx context.Context, conn transaction.Conn, dm *Users) (*Users, error)
	Update(ctx context.Context, conn transaction.Conn, dm *Users) (*Users, error)
	UpdateWithVersion(ctx context.Context, conn transaction.Conn, dm *Users) (*Users, error)
	Delete(ctx context.Context, conn transaction.Conn, dm *Users) (*Users, error)
	FindOne(ctx context.Context, conn transaction.Conn, id int64) (*Users, error)
	Find(ctx context.Context, conn transaction.Conn, queryOptions map[string]interface{}) (int64, []*Users, error)
	FindOneByEmail(ctx context.Context, conn transaction.Conn, email string) (*Users, error)
	FindOneByUsername(ctx context.Context, conn transaction.Conn, username string) (*Users, error)
	FindByIDs(ctx context.Context, conn transaction.Conn, ids []int64) (int64, []*Users, error)
}

func (m *Users) Identify() interface{} {
	if m.Id == 0 {
		return nil
	}
	return m.Id
}
