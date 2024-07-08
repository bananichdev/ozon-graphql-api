package errors

import "errors"

var (
	InternalError = GenerateError("Internal server error")
)

func GenerateError(err string) error {
	return errors.New(err)
}
