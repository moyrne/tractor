package syncx_test

import (
	"github.com/moyrne/tractor/syncx"
	"github.com/pkg/errors"
	"testing"
)

func TestNewRoutineGroup(t *testing.T) {
	group := syncx.NewRoutineGroup()
	group.Run(func() {
		panic(errors.New("panic"))
	})
	group.Wait()
}
