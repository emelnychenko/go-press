package errors

import (
	"fmt"
)

func NewPollNotFoundError(request string) Error {
	message := fmt.Sprintf("The poll was not found on request: %s", request)
	return NewNotFoundError(message)
}

func NewPollByIdNotFoundError(pollId fmt.Stringer) Error {
	request := fmt.Sprintf("id = %s", pollId)
	return NewPollNotFoundError(request)
}
