package gerrors

import (
	"net/http"
	"testing"
)

func TestUnauthorizedCode(t *testing.T) {
	err := Unauthorized()
	if !testIsCode(err, http.StatusUnauthorized) {
		t.Fail()
	}
}

func TestForbidenCode(t *testing.T) {
	err := Forbidden()
	if !testIsCode(err, http.StatusForbidden) {
		t.Fail()
	}
}

func TestNotFoundCode(t *testing.T) {
	err := NotFound("", nil)
	if !testIsCode(err, http.StatusNotFound) {
		t.Fail()
	}
}

func TestNotFoundMessage(t *testing.T) {
	msg := "oops"
	err := NotFound(msg, nil)
	if !testIsCode(err, http.StatusNotFound) || err.Error() != msg {
		t.Fail()
	}
}

func TestNotFoundObj(t *testing.T) {
	obj := "obj"
	err := NotFound("", obj)
	if !testIsCode(err, http.StatusNotFound) && !testIsObj(err, obj) {
		t.Fail()
	}
}

func TestPreconditionFailed(t *testing.T) {
}

func testIsCode(err ModelError, statusCode uint) bool {
	return err != nil && err.Code() == statusCode
}
func testIsObj(err ModelError, obj interface{}) bool {
	return err != nil && err.Obj() == obj
}
