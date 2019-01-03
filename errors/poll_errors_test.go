package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPollErrors(t *testing.T) {
	t.Run("PollNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewPollNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The poll was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("PollNotFoundByIdError", func(t *testing.T) {
		pollId := new(models.PollId)
		err := NewPollByIdNotFoundError(pollId)
		assert.Equal(t, fmt.Sprintf("The poll was not found on request: id = %s", pollId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
