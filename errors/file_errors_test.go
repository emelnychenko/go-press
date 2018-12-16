package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFileErrors(t *testing.T) {
	t.Run("FileNotFoundError", func(t *testing.T) {
		err := NewFileNotFoundError("check")
		assert.Equal(t, fmt.Sprintf("The file was not found on request: %s", "check"), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
