package utils

import "github.com/pkg/errors"

func WrapError(err error, message string) error {
	return errors.Wrap(err, message)
}
