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

func (d *DBX) NewTransaction(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	var tx *sqlx.Tx
	tx, err = d.Beginx()
	if err != nil {
		return errors.Wrap(err, "beginx")
	}
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
			if e := tx.Rollback(); e != nil {
				err = errors.Wrapf(err, "rollback %v", e)
			}
			return
		}
		err = errors.Wrap(tx.Commit(), "tx commit")
	}()

	return fn(ctx)
}
