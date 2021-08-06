package dbx

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type (
	Database interface {
		Begin() (Transaction, error)
		Ping() error
	}
	Transaction interface {
		Commit() error
		Rollback() error

		Get(dst interface{}, query string, args ...interface{}) error
		Select(dst interface{}, query string, args ...interface{}) error
		Exec(query string, args ...interface{}) (sql.Result, error)
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
		GetContext(ctx context.Context, dst interface{}, query string, args ...interface{}) error
		SelectContext(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	}
)

type DatabaseSQLX struct {
	*sqlx.DB
}

func (db *DatabaseSQLX) Begin() (Transaction, error) {
	tx, err := db.Beginx()
	return &Tx{Tx: tx}, err
}

type Tx struct {
	*sqlx.Tx
}

func (t *Tx) Commit() error {
	return t.Tx.Commit()
}

func (t *Tx) Rollback() error {
	return t.Tx.Rollback()
}
