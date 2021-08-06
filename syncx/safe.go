package syncx

import (
	"fmt"
	"github.com/moyrne/tractor/logx"
	"github.com/pkg/errors"
)

// Go run the fn. recover fn, when it panic
func Go(fn func()) {
	go Safe(fn)
}

func Safe(fn func()) {
	defer Recover()
	fn()
}

// Recover cleanup on panic
func Recover() {
	if r := recover(); r != nil {
		logx.Error("recover", "r", r)
	}
}

func Func(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			// panic not contains stack
			err = errors.WithStack(err)
		}
	}()
	return fn()
}
