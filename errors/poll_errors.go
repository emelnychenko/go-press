package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

func NewPollNotFoundError(request string) common.Error {
	message := fmt.Sprintf("The poll was not found on request: %s", request)
	return common.NewNotFoundError(message)
}

func NewPollByIdNotFoundError(pollId *models.PollId) common.Error {
	request := fmt.Sprintf("id = %s", pollId)
	return NewPollNotFoundError(request)
}
