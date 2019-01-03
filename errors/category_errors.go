package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

func NewCategoryNotFoundError(request string) common.Error {
	message := fmt.Sprintf("The category was not found on request: %s", request)
	return common.NewNotFoundError(message)
}

func NewCategoryByIdNotFoundError(categoryId *models.CategoryId) common.Error {
	request := fmt.Sprintf("id = %s", categoryId)
	return NewCategoryNotFoundError(request)
}
