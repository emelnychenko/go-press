package common

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestError(t *testing.T) {
	t.Run("NewError", func(t *testing.T) {
		errorMessage := "test"
		errorCode := 99

		err, isError := NewError(errorMessage, errorCode).(*errorImpl)
		assert.True(t, isError)
		assert.Equal(t, errorMessage, err.message)
		assert.Equal(t, errorCode, err.code)
	})

	t.Run("Error", func(t *testing.T) {
		errorMessage := "test"
		err := &errorImpl{message: errorMessage}
		assert.Equal(t, errorMessage, err.Error())
	})

	t.Run("Code", func(t *testing.T) {
		errorCode := 99
		err := &errorImpl{code: errorCode}
		assert.Equal(t, errorCode, err.Code())
	})

	t.Run("NewSystemError", func(t *testing.T) {
		errorMessage := "test"
		goErr := errors.New(errorMessage)

		err, isError := NewSystemError(goErr).(*errorImpl)
		assert.True(t, isError)
		assert.Equal(t, errorMessage, err.message)
		assert.Equal(t, http.StatusInternalServerError, err.code)
	})

	t.Run("NewUnknownError", func(t *testing.T) {
		err, isError := NewUnknownError().(*errorImpl)
		assert.True(t, isError)
		assert.Equal(t, "An error occurred", err.message)
		assert.Equal(t, http.StatusInternalServerError, err.code)
	})

	t.Run("NewNotFoundError", func(t *testing.T) {
		errorMessage := "test"

		err, isError := NewNotFoundError(errorMessage).(*errorImpl)
		assert.True(t, isError)
		assert.Equal(t, errorMessage, err.message)
		assert.Equal(t, http.StatusNotFound, err.code)
	})

	t.Run("NewObjectNotFoundError", func(t *testing.T) {
		errorRequest := "test"

		err, isError := NewObjectNotFoundError(errorRequest).(*errorImpl)
		assert.True(t, isError)
		assert.Equal(t, fmt.Sprintf("The object was not found on request: %s", errorRequest), err.message)
		assert.Equal(t, http.StatusNotFound, err.code)
	})
}
