package syncx

import (
	"reflect"
	"runtime"
)

func FunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
