package errors

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFileErrors(t *testing.T) {
	t.Run("FileNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewFileNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The file was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("FileNotFoundByIdError", func(t *testing.T) {
		fileId := new(uuid.UUID)
		err := NewFileByIdNotFoundError(fileId)
		assert.Equal(t, fmt.Sprintf("The file was not found on request: id = %s", fileId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
