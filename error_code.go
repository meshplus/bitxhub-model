package bitxhub_model

import (
	"fmt"
)

type ErrIbtpIndexExists struct {
	Require uint64
	Actual  uint64
}

func (e *ErrIbtpIndexExists) Error() string {
	return fmt.Sprintf(FmtIbtpIndexExists, e.Require, e.Actual)
}

const (
	FmtIbtpIndexExists = "index already exists, required %d, but %d"
)

// NewError creates and returns a new instance of CustomError.
func NewErrIbtpIndexExists(require, actual uint64) error {
	return &ErrIbtpIndexExists{require, actual}
}
