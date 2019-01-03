package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/models"
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
		categoryId := new(models.CategoryId)
		err := NewCategoryByIdNotFoundError(categoryId)
		assert.Equal(t, fmt.Sprintf("The category was not found on request: id = %s", categoryId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
