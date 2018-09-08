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
			message = fmt.Sprintf("%s,%+v", message, err)
		}
	}

	return message
}

func (errs Errors) Contains(err error) bool {
	if err == nil {
		return false
	}

	for _, currentError := range errs {
		if currentError == err {
			return true
		}
	}

	return false
}

func New(errs ...error) Errors {
	if len(errs) == 0 {
		return nil
	}
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
