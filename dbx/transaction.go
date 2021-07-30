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
	t, err := d.Beginx()
	if err != nil {
		return errors.Wrap(err, "beginx")
	}
	tx := &Tx{Tx: t}
	// recover
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			err = errors.WithMessage(err, "recover")
		}
		if err != nil {
			if e := tx.Rollback(); e != nil {
				err = errors.WithMessagef(err, "rollback %v", e)
			}
			return
		}
		err = errors.Wrap(tx.Commit(), "tx commit")
	}()

	return fn(ctx, tx)
}
