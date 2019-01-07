package errors

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBannerErrors(t *testing.T) {
	t.Run("BannerNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewBannerNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The banner was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("BannerNotFoundByIdError", func(t *testing.T) {
		bannerId := new(uuid.UUID)
		err := NewBannerByIdNotFoundError(bannerId)
		assert.Equal(t, fmt.Sprintf("The banner was not found on request: id = %s", bannerId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
