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
