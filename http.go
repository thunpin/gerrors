package gerrors

import "net/http"

type HttpErr interface {
	Code() uint
	Obj() interface{}
	error
}

type httpErr struct {
	code    uint
	message string
	obj     interface{}
}

func (h httpErr) Code() uint {
	return h.code
}
func (h httpErr) Error() string {
	return h.message
}
func (h httpErr) Obj() interface{} {
	return h.obj
}
func Unauthorized() error {
	return httpErr{http.StatusUnauthorized, "err_unauthorazied", nil}
}

func Forbidden() error {
	return httpErr{http.StatusForbidden, "err_forbidden", nil}
}

func NotFound(msg string, obj interface{}) error {
	return httpErr{http.StatusNotFound, msg, obj}
}
