package errors

import (
	"fmt"
)

func NewUserNotFoundError(request string) Error {
	message := fmt.Sprintf("The user was not found on request: %s", request)
	return NewNotFoundError(message)
}

func NewUserByIdNotFoundError(userId fmt.Stringer) Error {
	request := fmt.Sprintf("id = %s", userId)
	return NewUserNotFoundError(request)
}
