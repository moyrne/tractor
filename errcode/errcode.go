package errcode

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(code int, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}

func ErrCode(err error) int {
	if e, ok := err.(*Error); ok {
		return e.Code
	}
	return http.StatusOK
}

func ErrMsg(err error) string {
	if e, ok := err.(*Error); ok {
		return e.Msg
	}
	return ""
}

func (e *Error) Error() string {
	return "error code: " + strconv.Itoa(e.Code) + ", error msg: " + e.Msg
}

type ErrorInfo struct {
	err      *Error
	Domain   string            `json:"domain"`
	Reason   string            `json:"reason"`
	Metadata map[string]string `json:"metadata"`
}

func (e *ErrorInfo) Error() string {
	return fmt.Sprintf("error: domain = %s reason = %s matedata = %v", e.Domain, e.Reason, e.Metadata)
}

func (e *ErrorInfo) Err() *Error {
	return e.err
}

func (e *ErrorInfo) Is(err error) bool {
	if target := new(ErrorInfo); errors.As(err, &target) {
		return target.Domain == e.Domain && target.Reason == e.Reason
	}
	return false
}

func (e *ErrorInfo) WithMetadata(mp map[string]string) *ErrorInfo {
	err := *e
	err.Metadata = mp
	return &err
}
