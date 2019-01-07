package errors

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserErrors(t *testing.T) {
	t.Run("UserNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewUserNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The user was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("UserNotFoundByIdError", func(t *testing.T) {
		userId := new(uuid.UUID)
		err := NewUserByIdNotFoundError(userId)
		assert.Equal(t, fmt.Sprintf("The user was not found on request: id = %s", userId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
