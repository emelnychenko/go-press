package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
)

type (
	PostNotFoundError struct {
		common.NotFoundError
	}
)

func (e PostNotFoundError) Error() string {
	return fmt.Sprintf("The post was not found on request: %s", string(e.NotFoundError))
}

func NewPostNotFoundError(query string) PostNotFoundError {
	return PostNotFoundError{common.NotFoundError(query)}
}
