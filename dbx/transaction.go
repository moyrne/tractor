package dbx

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

func NewTransaction(ctx context.Context, db Database, fn func(ctx context.Context, tx Transaction) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "begin")
	}
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
