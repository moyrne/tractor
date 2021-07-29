package safe_test

import (
	"github.com/moyrne/tractor/safe"
	"github.com/pkg/errors"
	"testing"
)

func TestNewRoutineGroup(t *testing.T) {
	group := safe.NewRoutineGroup()
	group.Run(func() {
		panic(errors.New("panic"))
	})
	group.Wait()
}
