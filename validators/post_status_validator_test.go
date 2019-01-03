package validators

import (
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"testing"
	"time"
)

func TestPostStatusValidator(t *testing.T) {
	t.Run("NewPostStatusValidator", func(t *testing.T) {
		postStatusValidator, isPostStatusValidator := NewPostStatusValidator().(*postStatusValidatorImpl)

		assert.True(t, isPostStatusValidator)
		assert.NotNil(t, postStatusValidator.validate)
	})

	t.Run("ValidatePostCreate", func(t *testing.T) {
		validate := validator.New()
		postStatusValidator := &postStatusValidatorImpl{validate: validate}

		data := &models.PostCreate{Status: enums.PostScheduledStatus}

		err := postStatusValidator.ValidatePostCreate(data)
		assert.Error(t, err)
	})

	t.Run("ValidatePostCreate:PublishedExists", func(t *testing.T) {
		validate := validator.New()
		postStatusValidator := &postStatusValidatorImpl{validate: validate}

		postPublished := time.Now().UTC().Add(time.Hour)
		data := &models.PostCreate{Status: enums.PostScheduledStatus, Published: &postPublished}

		err := postStatusValidator.ValidatePostCreate(data)
		assert.Nil(t, err)
	})

	t.Run("ValidatePostCreate:NotPostScheduledStatus", func(t *testing.T) {
		validate := validator.New()
		postStatusValidator := &postStatusValidatorImpl{validate: validate}

		data := &models.PostCreate{Status: enums.PostPublishedStatus}

		err := postStatusValidator.ValidatePostCreate(data)
		assert.Nil(t, err)
	})

	t.Run("ValidatePostUpdate", func(t *testing.T) {
		validate := validator.New()
		postStatusValidator := &postStatusValidatorImpl{validate: validate}

		data := &models.PostUpdate{Status: enums.PostScheduledStatus}

		err := postStatusValidator.ValidatePostUpdate(data)
		assert.Error(t, err)
	})

	t.Run("ValidatePostUpdate:PublishedExists", func(t *testing.T) {
		validate := validator.New()
		postStatusValidator := &postStatusValidatorImpl{validate: validate}

		postPublished := time.Now().UTC().Add(time.Hour)
		data := &models.PostUpdate{Status: enums.PostScheduledStatus, Published: &postPublished}

		err := postStatusValidator.ValidatePostUpdate(data)
		assert.Nil(t, err)
	})

	t.Run("ValidatePostUpdate:NotPostScheduledStatus", func(t *testing.T) {
		validate := validator.New()
		postStatusValidator := &postStatusValidatorImpl{validate: validate}

		data := &models.PostUpdate{Status: enums.PostPublishedStatus}

		err := postStatusValidator.ValidatePostUpdate(data)
		assert.Nil(t, err)
	})
}
