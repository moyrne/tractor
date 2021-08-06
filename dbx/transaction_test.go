package dbx

import (
	_ "unsafe"

	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	ctx := context.Background()
	mockDB := &MockDB{}
	t.Run("success", func(t *testing.T) {
		defer resetTrans()
		assert.Equal(t, nil, NewTransaction(ctx, mockDB, func(ctx context.Context, tx Transaction) error {
			return nil
		}))
		assert.Equal(t, false, rollback)
		assert.Equal(t, true, commit)
	})
	t.Run("return error", func(t *testing.T) {
		defer resetTrans()
		returnErr := errors.New("return error")
		err := NewTransaction(ctx, mockDB, func(ctx context.Context, tx Transaction) error {
			return errors.WithStack(returnErr)
		})
		assert.Equal(t, true, rollback)
		assert.Equal(t, false, commit)
		assert.Equal(t, returnErr, errors.Cause(err))
	})
	t.Run("panic error", func(t *testing.T) {
		defer resetTrans()
		panicErr := errors.New("panic error")
		err := NewTransaction(ctx, mockDB, func(ctx context.Context, tx Transaction) error {
			panic(panicErr)
		})
		assert.Equal(t, true, rollback)
		assert.Equal(t, false, commit)
		assert.Equal(t, panicErr, errors.Cause(err))
	})

}
