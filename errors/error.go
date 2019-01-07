package errors

import (
	"fmt"
	"net/http"
)

type (
	Error interface {
		Error() string
		Code() int
	}

	errorImpl struct {
		message string
		code    int
	}
)

//NewError
func NewError(message string, code int) Error {
	return &errorImpl{message, code}
}

//NewErrorFromBuiltin
func NewErrorFromBuiltin(err error, code int) Error {
	return NewError(err.Error(), code)
}

func (e *errorImpl) Error() string {
	return e.message
}

//Code
func (e *errorImpl) Code() int {
	return e.code
}

//NewUnknownError
func NewUnknownError() Error {
	return NewError("An error occurred", http.StatusInternalServerError)
}

//NewSystemError
func NewSystemError(message string) Error {
	return NewError(message, http.StatusInternalServerError)
}

//NewSystemErrorFromBuiltin
func NewSystemErrorFromBuiltin(err error) Error {
	message := err.Error()
	return NewSystemError(message)
}

//NewBadRequestError
func NewBadRequestError(message string) Error {
	return NewError(message, http.StatusBadRequest)
}

//NewBadRequestErrorFromBuiltin
func NewBadRequestErrorFromBuiltin(err error) Error {
	message := err.Error()
	return NewBadRequestError(message)
}

//NewNotFoundError
func NewNotFoundError(message string) Error {
	return NewError(message, http.StatusNotFound)
}

//NewObjectNotFoundError
func NewObjectNotFoundError(request string) Error {
	message := fmt.Sprintf("The object was not found on request: %s", request)
	return NewNotFoundError(message)
}
