package domain

import (
	"context"
	"database/sql"
	"github.com/tiptok/gz-blog-microsevices/app/comment/interanl/pkg/db/transaction"
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

type CommentsRepository interface {
	Insert(ctx context.Context, conn transaction.Conn, dm *Comments) (*Comments, error)
	Update(ctx context.Context, conn transaction.Conn, dm *Comments) (*Comments, error)
	UpdateUnscoped(ctx context.Context, conn transaction.Conn, dm *Comments) (*Comments, error)
	UpdateWithVersion(ctx context.Context, conn transaction.Conn, dm *Comments) (*Comments, error)
	Delete(ctx context.Context, conn transaction.Conn, dm *Comments) (*Comments, error)
	FindOne(ctx context.Context, conn transaction.Conn, id int64) (*Comments, error)
	FindOneUnscoped(ctx context.Context, conn transaction.Conn, id int64) (*Comments, error)
	Find(ctx context.Context, conn transaction.Conn, queryOptions map[string]interface{}) (int64, []*Comments, error)
	FindOneByUuid(ctx context.Context, conn transaction.Conn, uuid string) (*Comments, error)
}

func (m *Comments) Identify() interface{} {
	if m.Id == 0 {
		return nil
	}
	return m.Id
}
