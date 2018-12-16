package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
)

type (
	FileNotFoundError struct {
		common.NotFoundError
	}
)

func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("The file was not found on request: %s", string(e.NotFoundError))
}

func NewFileNotFoundError(query string) FileNotFoundError {
	return FileNotFoundError{common.NotFoundError(query)}
}
