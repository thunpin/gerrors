package gerrors

import (
	"net/http"
	"testing"
)

func TestUnauthorizedCode(t *testing.T) {
	err := Unauthorized()
	if !testHasStatusCode(err, http.StatusUnauthorized) {
		t.Fail()
	}
}

func TestForbidenCode(t *testing.T) {
	err := Forbidden()
	if !testHasStatusCode(err, http.StatusForbidden) {
		t.Fail()
	}
}

func TestNotFoundCode(t *testing.T) {
	err := NotFound("", nil)
	if !testHasStatusCode(err, http.StatusNotFound) {
		t.Fail()
	}
}

func TestNotFoundMessage(t *testing.T) {
	msg := "oops"
	err := NotFound(msg, nil)
	if !testHasStatusCode(err, http.StatusNotFound) || err.Error() != msg {
		t.Fail()
	}
}

func TestNotFoundObj(t *testing.T) {
	obj := "obj"
	err := NotFound("", obj)
	if !testHasStatusCode(err, http.StatusNotFound) ||
		err.(HttpErr).Obj() != obj {
		t.Fail()
	}
}

func testHasStatusCode(err error, statusCode uint) bool {
	if err == nil {
		return false
	}

	switch err.(type) {
	case HttpErr:
		return err.(HttpErr).Code() == statusCode
	default:
		return false
	}
}
