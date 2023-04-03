package domain

import (
	"context"
	"database/sql"
	"github.com/tiptok/gz-blog-microsevices/app/post/interanl/pkg/db/transaction"
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
	DeletedAt     sql.NullTime
	Version       int64
}

type PostsRepository interface {
	Insert(ctx context.Context, conn transaction.Conn, dm *Posts) (*Posts, error)
	Update(ctx context.Context, conn transaction.Conn, dm *Posts) (*Posts, error)
	UpdateWithVersion(ctx context.Context, conn transaction.Conn, dm *Posts) (*Posts, error)
	Delete(ctx context.Context, conn transaction.Conn, dm *Posts) (*Posts, error)
	FindOne(ctx context.Context, conn transaction.Conn, id int64) (*Posts, error)
	Find(ctx context.Context, conn transaction.Conn, queryOptions map[string]interface{}) (int64, []*Posts, error)

	FindOneUnscoped(ctx context.Context, conn transaction.Conn, id int64) (*Posts, error)
	UpdateUnscoped(ctx context.Context, conn transaction.Conn, dm *Posts) (*Posts, error)
}

func (m *Posts) Identify() interface{} {
	if m.Id == 0 {
		return nil
	}
	return m.Id
}
