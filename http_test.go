package gerrors

import (
	"errors"
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
	if !testIsMsg(err, msg) {
		t.Fail()
	}
}

func TestNotFoundObj(t *testing.T) {
	obj := "obj"
	err := NotFound("", obj)
	if !testIsObj(err, obj) {
		t.Fail()
	}
}

func TestPreconditionFailedCode(t *testing.T) {
	err := PreconditionFailed("")
	if !testIsCode(err, http.StatusPreconditionFailed) {
		t.Fail()
	}
}

func TestPreconditionFailedMsg(t *testing.T) {
	msg := "oops"
	err := PreconditionFailed(msg)
	if !testIsMsg(err, msg) {
		t.Fail()
	}
}

func TestInternalServerErrorCode(t *testing.T) {
	origError := errors.New("oops")
	err := InternalServerError(origError)
	if !testIsCode(err, http.StatusInternalServerError) {
		t.Fail()
	}
}

func TestInternalServerErrorObj(t *testing.T) {
	origError := errors.New("oops")
	err := InternalServerError(origError)
	if !testIsObj(err, origError) {
		t.Fail()
	}
}

func testIsCode(err ModelError, statusCode uint) bool {
	return err != nil && err.Code() == statusCode
}
func testIsObj(err ModelError, obj interface{}) bool {
	return err != nil && err.Obj() == obj
}
func testIsMsg(err ModelError, msg string) bool {
	return err != nil && err.Error() == msg
}
