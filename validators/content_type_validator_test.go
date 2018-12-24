package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContentTypeValidator(t *testing.T) {
	t.Run("NewContentTypeValidator", func(t *testing.T) {
		_, isContentTypeValidator := NewContentTypeValidator().(*contentTypeValidatorImpl)

		assert.True(t, isContentTypeValidator)
	})

	t.Run("ValidateImage", func(t *testing.T) {
		contentTypeValidator := new(contentTypeValidatorImpl)

		contentType := "image/png"
		err := contentTypeValidator.ValidateImage(contentType)

		assert.Nil(t, err)
	})

	t.Run("ValidateImage:Error", func(t *testing.T) {
		contentTypeValidator := new(contentTypeValidatorImpl)

		contentType := "video/3gp"
		err := contentTypeValidator.ValidateImage(contentType)

		assert.Error(t, err)
	})

	t.Run("ValidateVideo", func(t *testing.T) {
		contentTypeValidator := new(contentTypeValidatorImpl)

		contentType := "video/3gp"
		err := contentTypeValidator.ValidateVideo(contentType)

		assert.Nil(t, err)
	})

	t.Run("ValidateVideo:Error", func(t *testing.T) {
		contentTypeValidator := new(contentTypeValidatorImpl)

		contentType := "text/plain"
		err := contentTypeValidator.ValidateVideo(contentType)

		assert.Error(t, err)
	})

	t.Run("ValidateAudio", func(t *testing.T) {
		contentTypeValidator := new(contentTypeValidatorImpl)

		contentType := "audio/mp3"
		err := contentTypeValidator.ValidateAudio(contentType)

		assert.Nil(t, err)
	})

	t.Run("ValidateAudio:Error", func(t *testing.T) {
		contentTypeValidator := new(contentTypeValidatorImpl)

		contentType := "text/plain"
		err := contentTypeValidator.ValidateAudio(contentType)

		assert.Error(t, err)
	})
}

