package user_domain

import (
	"../common"
	"fmt"
)

type (
	UserNotFoundError struct {
		common.NotFoundError
	}
)

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("The user was not found on request: %s", string(e.NotFoundError))
}

func NewUserNotFoundError(query string) UserNotFoundError {
	return UserNotFoundError{common.NotFoundError(query)}
}
