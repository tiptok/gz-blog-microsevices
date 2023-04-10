package transaction

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type Context struct {
	//启用事务标识
	beginTransFlag bool
	db             *gorm.DB
	session        *gorm.DB
	lock           sync.Mutex
}

func (transactionContext *Context) Begin() error {
	transactionContext.lock.Lock()
	defer transactionContext.lock.Unlock()
	transactionContext.beginTransFlag = true
	tx := transactionContext.db.Begin()
	transactionContext.session = tx
	return nil
}

func (transactionContext *Context) Commit() error {
	transactionContext.lock.Lock()
	defer transactionContext.lock.Unlock()
	if !transactionContext.beginTransFlag {
		return nil
	}
	tx := transactionContext.session.Commit()
	return tx.Error
}

func (transactionContext *Context) Rollback() error {
	transactionContext.lock.Lock()
	defer transactionContext.lock.Unlock()
	if !transactionContext.beginTransFlag {
		return nil
	}
	tx := transactionContext.session.Rollback()
	return tx.Error
}

func (transactionContext *Context) DB() *gorm.DB {
	if transactionContext.beginTransFlag && transactionContext.session != nil {
		return transactionContext.session
	}
	return transactionContext.db
}

func NewTransactionContext(db *gorm.DB) *Context {
	return &Context{
		db: db,
	}
}

type Conn interface {
	Begin() error
	Commit() error
	Rollback() error
	DB() *gorm.DB
}

// UseTrans when beginTrans is true , it will begin a new transaction
// to execute the function, recover when  panic happen
func UseTrans(ctx context.Context,
	db *gorm.DB,
	fn func(context.Context, Conn) error, beginTrans bool) (err error) {
	var tx Conn
	tx = NewTransactionContext(db)
	if beginTrans {
		if err = tx.Begin(); err != nil {
			return
		}
	}
	defer func() {
		if p := recover(); p != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("recover from %#v, rollback failed: %w", p, e)
			} else {
				err = fmt.Errorf("recoveer from %#v", p)
			}
		} else if err != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("transaction failed: %s, rollback failed: %w", err, e)
			}
		} else {
			err = tx.Commit()
		}
	}()

	return fn(ctx, tx)
}

func PaginationAndCount(ctx context.Context, tx *gorm.DB, params map[string]interface{}, v interface{}) (int64, *gorm.DB) {
	var total int64
	var enableCounter bool = true
	if v, ok := params["enableCounter"]; ok {
		enableCounter = v.(bool)
	}
	if enableCounter {
		tx = tx.Count(&total)
		if tx.Error != nil {
			return total, tx
		}
	}
	if v, ok := params["offset"]; ok {
		tx.Offset(v.(int))
	}
	if v, ok := params["limit"]; ok {
		tx.Limit(v.(int))
	}
	//if v,ok:=params["unscoped"];ok && v.(bool){
	//	tx.Unscoped()
	//}
	if tx = tx.Find(v); tx.Error != nil {
		return 0, tx
	}
	return total, tx
}
