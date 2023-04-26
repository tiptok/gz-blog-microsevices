package domain

import (
	"context"
	"database/sql"
	"github.com/tiptok/gz-blog-microsevices/app/product/interanl/pkg/db/transaction"
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
	DeletedAt   sql.NullTime
	Version     int64
}

type ProductsRepository interface {
	Insert(ctx context.Context, conn transaction.Conn, dm *Products) (*Products, error)
	Update(ctx context.Context, conn transaction.Conn, dm *Products) (*Products, error)
	UpdateWithVersion(ctx context.Context, conn transaction.Conn, dm *Products) (*Products, error)
	Delete(ctx context.Context, conn transaction.Conn, dm *Products) (*Products, error)
	FindOne(ctx context.Context, conn transaction.Conn, id int64) (*Products, error)
	Find(ctx context.Context, conn transaction.Conn, queryOptions map[string]interface{}) (int64, []*Products, error)
}

func (m *Products) Identify() interface{} {
	if m.Id == 0 {
		return nil
	}
	return m.Id
}
