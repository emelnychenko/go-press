package errors

import (
	"fmt"
)

func NewPostNotFoundError(request string) Error {
	message := fmt.Sprintf("The post was not found on request: %s", request)
	return NewNotFoundError(message)
}

func NewPostByIdNotFoundError(postId fmt.Stringer) Error {
	request := fmt.Sprintf("id = %s", postId)
	return NewPostNotFoundError(request)
}
