package errorx

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
