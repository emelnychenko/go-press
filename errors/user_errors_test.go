package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserErrors(t *testing.T) {
	t.Run("UserNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewUserNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The UserEntity was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("UserNotFoundByIdError", func(t *testing.T) {
		fileId := new(models.UserId)
		err := NewUserByIdNotFoundError(fileId)
		assert.Equal(t, fmt.Sprintf("The UserEntity was not found on request: Id = %s", fileId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}

