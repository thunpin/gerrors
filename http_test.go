package gerrors

import (
	"net/http"
	"testing"
)

func TestUnauthorizedCode(t *testing.T) {
	err := Unauthorized()
	if err == nil || err.(HttpErr).Code() != http.StatusUnauthorized {
		t.Fail()
	}
}

func TestForbidenCode(t *testing.T) {
	err := Forbidden()
	if err == nil || err.(HttpErr).Code() != http.StatusForbidden {
		t.Fail()
	}
}
