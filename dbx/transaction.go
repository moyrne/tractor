package dbx

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DBX struct {
	*sqlx.DB
}
type Tx struct {
	*sqlx.Tx
}

//go:noinline
func (t *Tx) Rollback() error {
	return t.Tx.Rollback()
}

func (d *DBX) NewTransaction(ctx context.Context, fn func(ctx context.Context, tx *Tx) error) (err error) {
	tx, err := d.Beginx()
	if err != nil {
		return errors.Wrap(err, "beginx")
	}
	txx := &Tx{Tx: tx}
	// recover
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
		if err != nil {
			if e := txx.Rollback(); e != nil {
				err = errors.WithMessagef(err, "rollback %v", e)
			}
			return
		}
		err = errors.Wrap(tx.Commit(), "tx commit")
	}()

	return fn(ctx, &Tx{Tx: tx})
}
