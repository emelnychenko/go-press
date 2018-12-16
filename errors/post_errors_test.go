package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPostErrors(t *testing.T) {
	t.Run("PostNotFoundError", func(t *testing.T) {
		err := NewPostNotFoundError("check")
		assert.Equal(t, fmt.Sprintf("The post was not found on request: %s", "check"), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
