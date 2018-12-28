package common

import (
	"fmt"
	"net/http"
)

type (
	Error interface {
		error
		Code() int
	}

	errorImpl struct {
		message string
		code int
	}
)

func NewError(message string, code int) Error {
	return &errorImpl{message, code}
}

func NewErrorFromBuiltin(err error, code int) Error {
	return NewError(err.Error(), code)
}

func (e *errorImpl) Error() string {
	return e.message
}

func (e *errorImpl) Code() int {
	return e.code
}

func NewUnknownError() Error {
	return NewError("An error occurred", http.StatusInternalServerError)
}

func NewSystemError(message string) Error {
	return NewError(message, http.StatusInternalServerError)
}

func NewSystemErrorFromBuiltin(err error) Error {
	message := err.Error()
	return NewSystemError(message)
}

func NewBadRequestError(message string) Error {
	return NewError(message, http.StatusBadRequest)
}

func NewBadRequestErrorFromBuiltin(err error) Error {
	message := err.Error()
	return NewBadRequestError(message)
}

func NewNotFoundError(message string) Error {
	return NewError(message, http.StatusNotFound)
}

func NewObjectNotFoundError(request string) Error {
	message := fmt.Sprintf("The object was not found on request: %s", request)
	return NewNotFoundError(message)
}
