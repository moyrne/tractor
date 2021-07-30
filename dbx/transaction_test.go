package dbx

import (
	_ "unsafe"

	"context"
	"database/sql"
	. "github.com/agiledragon/gomonkey"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	var rollback, committed bool
	reset := func() {
		rollback = false
		committed = false
	}
	ApplyMethod(reflect.TypeOf(&sqlx.DB{}), "Beginx", func(db *sqlx.DB) (*sqlx.Tx, error) {
		return &sqlx.Tx{}, nil
	})
	ApplyMethod(reflect.TypeOf(&Tx{}), "Rollback", func(tx *Tx) error {
		rollback = true
		return nil
	})
	ApplyMethod(reflect.TypeOf(&sql.Tx{}), "Commit", func(tx *sql.Tx) error {
		committed = true
		return nil
	})

	ctx := context.Background()
	DB = &DBX{}
	t.Run("success", func(t *testing.T) {
		defer reset()
		assert.Equal(t, nil, DB.NewTransaction(ctx, func(ctx context.Context, tx *Tx) error {
			return nil
		}))
		assert.Equal(t, false, rollback)
		assert.Equal(t, true, committed)
	})
	t.Run("return error", func(t *testing.T) {
		defer reset()
		returnErr := errors.New("return error")
		err := DB.NewTransaction(ctx, func(ctx context.Context, tx *Tx) error {
			return errors.WithStack(returnErr)
		})
		assert.Equal(t, true, rollback)
		assert.Equal(t, false, committed)
		assert.Equal(t, returnErr, errors.Cause(err))
	})
	t.Run("panic error", func(t *testing.T) {
		defer reset()
		panicErr := errors.New("panic error")
		assert.Equal(t, panicErr, DB.NewTransaction(ctx, func(ctx context.Context, tx *Tx) error {
			panic(panicErr)
		}))
		assert.Equal(t, true, rollback)
		assert.Equal(t, false, committed)
	})

}
