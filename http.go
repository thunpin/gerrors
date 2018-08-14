package gerrors

import "net/http"

func Unauthorized() Error {
	return Error{http.StatusUnauthorized, "err_unauthorazied", nil}
}

func Forbidden() Error {
	return Error{http.StatusForbidden, "err_forbidden", nil}
}

func NotFound(msg string, obj interface{}) Error {
	return Error{http.StatusNotFound, msg, obj}
}

func PreconditionFailed(msg string) Error {
	return Error{http.StatusPreconditionFailed, msg, nil}
}
func InternalServerError(err error) Error {
	return Error{
		http.StatusInternalServerError,
		"err_internal_server_error",
		err,
	}
}
