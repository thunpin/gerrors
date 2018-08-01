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
