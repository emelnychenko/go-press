package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTagErrors(t *testing.T) {
	t.Run("TagNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewTagNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The tag was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("TagNotFoundByIdError", func(t *testing.T) {
		tagId := new(models.TagId)
		err := NewTagByIdNotFoundError(tagId)
		assert.Equal(t, fmt.Sprintf("The tag was not found on request: id = %s", tagId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
