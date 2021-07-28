package safe

import (
	"github.com/moyrne/tractor/logx"
)

// Go run the fn. recover fn, when it panic
func Go(fn func()) {
	defer Recover()
	go fn()
}

// Recover cleanup on panic
func Recover() {
	if r := recover(); r != nil {
		logx.Error("recover", "r", r)
	}
}
