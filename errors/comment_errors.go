package errors

import (
	"fmt"
)

func NewCommentNotFoundError(request string) Error {
	message := fmt.Sprintf("The comment was not found on request: %s", request)
	return NewNotFoundError(message)
}

func NewCommentByIdNotFoundError(commentId fmt.Stringer) Error {
	request := fmt.Sprintf("id = %s", commentId)
	return NewCommentNotFoundError(request)
}
