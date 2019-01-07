package errors

import (
	"fmt"
	"github.com/google/uuid"
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
		pollId := new(uuid.UUID)
		err := NewPollByIdNotFoundError(pollId)
		assert.Equal(t, fmt.Sprintf("The poll was not found on request: id = %s", pollId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
