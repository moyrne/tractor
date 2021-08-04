package dbx

import (
	"context"
	"database/sql"
)

type (
	Conn interface {
		Begin() error
	}
	Transaction interface {
		Session
		Commit() error
		Rollback() error
	}
	Session interface {
		Get(dst interface{}, query string, args ...interface{}) error
		Select(dst interface{}, query string, args ...interface{}) error
		Exec(query string, args ...interface{}) (sql.Result, error)
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
		GetContext(ctx context.Context, dst interface{}, query string, args ...interface{}) error
		SelectContext(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	}
)
