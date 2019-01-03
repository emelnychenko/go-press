package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

func NewCommentNotFoundError(request string) common.Error {
	message := fmt.Sprintf("The comment was not found on request: %s", request)
	return common.NewNotFoundError(message)
}

func NewCommentByIdNotFoundError(categoryId *models.CommentId) common.Error {
	request := fmt.Sprintf("id = %s", categoryId)
	return NewCommentNotFoundError(request)
}
