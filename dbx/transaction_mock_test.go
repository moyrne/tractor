package dbx

import (
	"context"
	"database/sql"
)

var (
	commit, rollback bool
)

func resetTrans() {
	commit = false
	rollback = false
}

type MockDB struct{}

func (t MockDB) Begin() (Transaction, error) {
	return &MockTrans{}, nil
}

func (t MockDB) Ping() error {
	panic("implement me")
}

type MockTrans struct{}

func (t MockTrans) Commit() error {
	commit = true
	return nil
}

func (t MockTrans) Rollback() error {
	rollback = true
	return nil
}

func (t MockTrans) Get(dst interface{}, query string, args ...interface{}) error {
	panic("implement me")
}

func (t MockTrans) Select(dst interface{}, query string, args ...interface{}) error {
	panic("implement me")
}

func (t MockTrans) Exec(query string, args ...interface{}) (sql.Result, error) {
	panic("implement me")
}

func (t MockTrans) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	panic("implement me")
}

func (t MockTrans) GetContext(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	panic("implement me")
}

func (t MockTrans) SelectContext(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	panic("implement me")
}
