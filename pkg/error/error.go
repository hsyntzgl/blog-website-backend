package errorGenerator

import "errors"

func GenerateError(message string) error {
	err := errors.New(message)
	return err
}
