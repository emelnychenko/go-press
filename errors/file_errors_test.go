package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFileErrors(t *testing.T) {
	t.Run("FileNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewFileNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The File was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("FileNotFoundByIdError", func(t *testing.T) {
		fileId := new(models.FileId)
		err := NewFileByIdNotFoundError(fileId)
		assert.Equal(t, fmt.Sprintf("The File was not found on request: Id = %s", fileId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
