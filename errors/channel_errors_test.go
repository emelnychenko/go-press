package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestChannelErrors(t *testing.T) {
	t.Run("ChannelNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewChannelNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The channel was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("ChannelNotFoundByIdError", func(t *testing.T) {
		channelId := new(models.ChannelId)
		err := NewChannelByIdNotFoundError(channelId)
		assert.Equal(t, fmt.Sprintf("The channel was not found on request: id = %s", channelId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
