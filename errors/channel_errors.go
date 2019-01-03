package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

func NewChannelNotFoundError(request string) common.Error {
	message := fmt.Sprintf("The channel was not found on request: %s", request)
	return common.NewNotFoundError(message)
}

func NewChannelByIdNotFoundError(categoryId *models.ChannelId) common.Error {
	request := fmt.Sprintf("id = %s", categoryId)
	return NewChannelNotFoundError(request)
}
