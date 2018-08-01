package gerrors

import "fmt"

type Errors []error

// Error implements the error interface
func (errs Errors) Error() string {
	var message string
	for i, err := range errs {
		if i == 0 {
			message = err.Error()
		} else {
			message = fmt.Sprintf("%s, %s", message, err.Error())
		}
	}

	return message
}

func New(errs ...error) error {
	newErrs := buildErrors(&errs)

	if len(newErrs) == 0 {
		return nil
	}
	return newErrs
}

func buildErrors(errs *[]error) Errors {
	newErrs := make(Errors, 0)
	for _, err := range *errs {
		if err == nil {
			continue
		}
		newErrs = appendError(newErrs, err)
	}

	return newErrs
}

func appendError(errs Errors, err error) Errors {
	switch err.(type) {
	case Errors:
		return append(errs, err.(Errors)...)
	default:
		return append(errs, err)
	}
}
