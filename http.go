package gerrors

import "net/http"

type HttpErr interface {
	Code() uint
	error
}

type httpErr struct {
	code    uint
	message string
}

func (h *httpErr) Code() uint {
	return h.code
}
func (h *httpErr) Error() string {
	return h.message
}

func Unauthorized() error {
	return &httpErr{http.StatusUnauthorized, "err_unauthorazied"}
}

func Forbidden() error {
	return &httpErr{http.StatusForbidden, "err_forbidden"}
}
