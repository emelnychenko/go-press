package errors

import (
	"fmt"
)

func NewFileNotFoundError(request string) Error {
	message := fmt.Sprintf("The file was not found on request: %s", request)
	return NewNotFoundError(message)
}

func NewFileByIdNotFoundError(fileId fmt.Stringer) Error {
	request := fmt.Sprintf("id = %s", fileId)
	return NewFileNotFoundError(request)
}
