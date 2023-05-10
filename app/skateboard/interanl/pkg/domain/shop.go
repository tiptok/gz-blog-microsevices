package domain

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/db/transaction"
)

type Shop struct {
	Id           int64
	Name         string
	Introduction string
	Address      *Address
	Context      *Context
	Rank         int
	CreatedAt    int64
	UpdatedAt    int64
	DeletedAt    int64
	Version      int
}

type ShopRepository interface {
	Insert(ctx context.Context, conn transaction.Conn, dm *Shop) (*Shop, error)
	Update(ctx context.Context, conn transaction.Conn, dm *Shop) (*Shop, error)
	UpdateWithVersion(ctx context.Context, conn transaction.Conn, dm *Shop) (*Shop, error)
	Delete(ctx context.Context, conn transaction.Conn, dm *Shop) (*Shop, error)
	FindOne(ctx context.Context, conn transaction.Conn, id int64) (*Shop, error)
	Find(ctx context.Context, conn transaction.Conn, queryOptions map[string]interface{}) (int64, []*Shop, error)
}

func (m *Shop) Identify() interface{} {
	if m.Id == 0 {
		return nil
	}
	return m.Id
}
