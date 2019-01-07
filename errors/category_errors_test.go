package errors

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCategoryErrors(t *testing.T) {
	t.Run("CategoryNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewCategoryNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The category was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("CategoryNotFoundByIdError", func(t *testing.T) {
		categoryId := new(uuid.UUID)
		err := NewCategoryByIdNotFoundError(categoryId)
		assert.Equal(t, fmt.Sprintf("The category was not found on request: id = %s", categoryId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("CategoryXrefNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewCategoryXrefNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The category reference was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("CategoryXrefNotFoundByReferenceError", func(t *testing.T) {
		categoryId := new(uuid.UUID)
		objectType := "object"
		objectId := new(uuid.UUID)
		err := NewCategoryXrefNotFoundByReferenceError(categoryId, objectType, objectId)
		message := fmt.Sprintf("The category reference was not found on request: categoryId = %s, objectType = %s, objectId = %s", categoryId, objectType, objectId)
		assert.Equal(t, message, err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
