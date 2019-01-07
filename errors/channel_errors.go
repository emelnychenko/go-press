package errors

import (
	"fmt"
)

func NewChannelNotFoundError(request string) Error {
	message := fmt.Sprintf("The channel was not found on request: %s", request)
	return NewNotFoundError(message)
}

func NewChannelByIdNotFoundError(channelId fmt.Stringer) Error {
	request := fmt.Sprintf("id = %s", channelId)
	return NewChannelNotFoundError(request)
}
