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

func NewUnknownError() Error {
	return &errorImpl{"An error occurred", http.StatusInternalServerError}
}

func (e *errorImpl) Error() string {
	return e.message
}

func (e *errorImpl) Code() int {
	return e.code
}

func NewSystemError(err error) Error {
	return &errorImpl{err.Error(), http.StatusInternalServerError}
}

func NewBadRequestError(message string) Error {
	return &errorImpl{message, http.StatusBadRequest}
}

func NewNotFoundError(message string) Error {
	return &errorImpl{message, http.StatusNotFound}
}

func NewObjectNotFoundError(request string) Error {
	message := fmt.Sprintf("The object was not found on request: %s", request)
	return &errorImpl{message, http.StatusNotFound}
}
