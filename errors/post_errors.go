package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

func NewPostNotFoundError(request string) common.Error {
	message := fmt.Sprintf("The post was not found on request: %s", request)
	return common.NewNotFoundError(message)
}

func NewPostByIdNotFoundError(fileId *models.PostId) common.Error {
	request := fmt.Sprintf("id = %s", fileId)
	return NewPostNotFoundError(request)
}
