package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
)

type (
	UserNotFoundError struct {
		common.NotFoundError
	}
)

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("The aggregator was not found on request: %s", string(e.NotFoundError))
}

func NewUserNotFoundError(query string) UserNotFoundError {
	return UserNotFoundError{common.NotFoundError(query)}
}
