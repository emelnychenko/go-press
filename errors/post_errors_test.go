package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPostErrors(t *testing.T) {
	t.Run("PostNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewPostNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The post was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("PostNotFoundByIdError", func(t *testing.T) {
		fileId := new(models.PostId)
		err := NewPostByIdNotFoundError(fileId)
		assert.Equal(t, fmt.Sprintf("The post was not found on request: id = %s", fileId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
