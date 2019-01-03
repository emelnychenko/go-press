package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

func NewFileNotFoundError(request string) common.Error {
	message := fmt.Sprintf("The file was not found on request: %s", request)
	return common.NewNotFoundError(message)
}

func NewFileByIdNotFoundError(fileId *models.FileId) common.Error {
	request := fmt.Sprintf("id = %s", fileId)
	return NewFileNotFoundError(request)
}
