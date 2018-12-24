package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

func NewUserNotFoundError(request string) common.Error {
	message := fmt.Sprintf("The UserEntity was not found on request: %s", request)
	return common.NewNotFoundError(message)
}

func NewUserByIdNotFoundError(fileId *models.UserId) common.Error {
	request := fmt.Sprintf("Id = %s", fileId)
	return NewUserNotFoundError(request)
}
