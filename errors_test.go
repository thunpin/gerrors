package gerrors

import (
	"errors"
	"strings"
	"testing"
)

func TestNilMustReturnNil(t *testing.T) {
	err := New(nil)
	if err != nil {
		t.Fail()
	}
}

func TestNilWithValidsMustReturnOnlyValids(t *testing.T) {
	err := New(nil, errors.New("err"), nil)
	if len(err.(Errors)) != 1 {
		t.Fail()
	}
}

func TestNewWithErrorAndErrorsMustAppend(t *testing.T) {
	errs := New(errors.New("err1"), errors.New("err2"))
	err := New(errs, errors.New("err3"))

	if len(err.(Errors)) != 3 {
		t.Fail()
	}
}

func TestNewWithInvalidErrorsAndError(t *testing.T) {
	errs := New(nil, nil, nil)
	err := New(errs, nil)

	if err != nil {
		t.Fail()
	}
}

func TestNewWithPartialInvalidErrorsAndInvalidError(t *testing.T) {
	errs := New(nil, nil, errors.New("error"))
	err := New(errs, nil)

	if len(err.(Errors)) != 1 {
		t.Fail()
	}
}

func TestNewWithInvalidErrorsAndPartialInvalidError(t *testing.T) {
	errs := New(nil, nil, nil)
	err := New(errs, errors.New("error"))

	if len(err.(Errors)) != 1 {
		t.Fail()
	}
}

func TestContainsAllErrorsMsg(t *testing.T) {
	errs := New(errors.New("err1"), errors.New("err2"))
	err := New(errs, errors.New("err3"))

	containsErr1 := strings.Contains(err.Error(), "err1")
	containsErr2 := strings.Contains(err.Error(), "err2")
	containsErr3 := strings.Contains(err.Error(), "err3")

	if !(containsErr1 && containsErr2 && containsErr3) {
		t.Fail()
	}
}

func TestVerifyIfContainsTheError(t *testing.T) {
	err := errors.New("test")
	errs := New(err)
	if !errs.(Errors).Contains(err) {
		t.Fail()
	}

	if errs.(Errors).Contains(nil) {
		t.Fail()
	}

	err = errors.New("test2")
	if errs.(Errors).Contains(err) {
		t.Fail()
	}
}
