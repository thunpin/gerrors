package gerrors

import "net/http"

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

func Unauthorized() ModelError {
	return httpErr{http.StatusUnauthorized, "err_unauthorazied", nil}
}

func Forbidden() ModelError {
	return httpErr{http.StatusForbidden, "err_forbidden", nil}
}

func NotFound(msg string, obj interface{}) ModelError {
	return httpErr{http.StatusNotFound, msg, obj}
}

func PreconditionFailed(msg string) ModelError {
	return httpErr{http.StatusPreconditionFailed, msg, nil}
}
func InternalServerError(err error) ModelError {
	return httpErr{
		http.StatusInternalServerError,
		"err_internal_server_error",
		err,
	}
}
