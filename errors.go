package gerrors

import "fmt"

type Errors []error

// Error implements the error interface
func (errs Errors) Error() string {
	return fmt.Sprintf("%s", errs)
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
