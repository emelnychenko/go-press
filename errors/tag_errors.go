package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

func NewTagNotFoundError(request string) common.Error {
	message := fmt.Sprintf("The tag was not found on request: %s", request)
	return common.NewNotFoundError(message)
}

func NewTagByIdNotFoundError(categoryId *models.TagId) common.Error {
	request := fmt.Sprintf("id = %s", categoryId)
	return NewTagNotFoundError(request)
}
