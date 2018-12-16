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

	ServerError   string
	NotFoundError string
)

func (e ServerError) Error() string {
	return string(e)
}

func (ServerError) Code() int {
	return http.StatusInternalServerError
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("The object was not found on request: %s", string(e))
}

func (NotFoundError) Code() int {
	return http.StatusNotFound
}

func NewServerError(err error) ServerError {
	return ServerError(err.Error())
}
