package errcode

import (
	"fmt"
	"net/http"
	"strconv"
)

var (
	Success     = NewError(http.StatusOK, "success")
	ServerError = NewError(http.StatusInternalServerError, "Internal Server Error")
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ErrorInfo struct {
	err      *Error
	Domain   string            `json:"domain"`
	Reason   string            `json:"reason"`
	Metadata map[string]string `json:"metadata"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("error code %d is already exist", code))
	}
	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Error() string {
	return "error code: " + strconv.Itoa(e.Code) + ", error msg: " + e.Msg
}
