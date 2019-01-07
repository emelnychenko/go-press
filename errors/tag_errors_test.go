package errors

import (
	"fmt"
	"github.com/google/uuid"
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
		tagId := new(uuid.UUID)
		err := NewTagByIdNotFoundError(tagId)
		assert.Equal(t, fmt.Sprintf("The tag was not found on request: id = %s", tagId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("TagXrefNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewTagXrefNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The tag reference was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("TagXrefNotFoundByReferenceError", func(t *testing.T) {
		tagId := new(uuid.UUID)
		objectType := "object"
		objectId := new(uuid.UUID)
		err := NewTagXrefNotFoundByReferenceError(tagId, objectType, objectId)
		message := fmt.Sprintf("The tag reference was not found on request: tagId = %s, objectType = %s, objectId = %s", tagId, objectType, objectId)
		assert.Equal(t, message, err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
