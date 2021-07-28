package logx

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	"time"
)

var logs = &log.Logger{}

func Init(out ...io.Writer) {
	// 输出到文件
	logs.SetOutput(io.MultiWriter(append(out, os.Stdout)...))
}

func FileWriter(filename string) (*os.File, error) {
	if filename == "" {
		filename = "logs/default.log"
		d, e := os.Open("logs")
		if e != nil {
			if err := os.MkdirAll("logs", 0666); err != nil {
				return nil, err
			}
		} else {
			_ = d.Close()
		}
	}
	return os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
}

type LogValue struct {
	Level  string                 `json:"level"`
	Time   time.Time              `json:"time"`
	Value  string                 `json:"value"`
	Detail map[string]interface{} `json:"detail"`
}

const (
	LevelInfo  = "info"
	LevelDebug = "debug"
	LevelError = "error"
	LevelPanic = "panic"
)

func Info(v string, kv ...interface{}) {
	logs.Println(logValue(&LogValue{Level: LevelInfo, Value: v}, kv...))
}

func Debug(v string, kv ...interface{}) {
	logs.Println(logValue(&LogValue{Level: LevelDebug, Value: v}, kv...))
}

func Error(v string, kv ...interface{}) {
	logs.Println(logValue(&LogValue{Level: LevelError, Value: v}, kv...))
}

func ErrorStack(v string, kv ...interface{}) {
	logs.Println(logValue(&LogValue{Level: LevelError, Value: v}, append(kv, "stack", string(debug.Stack()))))
}

func Panic(v string, kv ...interface{}) {
	logs.Panicln(logValue(&LogValue{Level: LevelPanic, Value: v}, kv...))
}

func logValue(value *LogValue, kv ...interface{}) string {
	kvl := len(kv)
	if kvl%2 != 0 {
		return logKVV(value, kv)
	}
	value.Time = time.Now()
	value.Detail = map[string]interface{}{}
	for i := 0; i < kvl-1; i += 2 {
		key, ok := kv[i].(string)
		if !ok {
			return logKVV(value, kv)
		}
		value.Detail[key] = fmt.Sprintf("%+v", kv[i+1])
	}
	return marshalLog(value)
}

func logKVV(value *LogValue, kv ...interface{}) string {
	value.Detail = map[string]interface{}{
		"kvv": kv,
	}
	return marshalLog(value)
}

func marshalLog(logValue *LogValue) string {
	value, err := json.Marshal(logValue)
	if err != nil {
		logs.Printf("marshalLog error %v\n", err)
		return err.Error()
	}
	return string(value)
}
