package validators

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"testing"
)

func TestModelValidator(t *testing.T) {
	t.Run("NewModelValidator", func(t *testing.T) {
		modelValidator, isModelValidator := NewModelValidator().(*modelValidatorImpl)

		assert.True(t, isModelValidator)
		assert.NotNil(t, modelValidator.validate)
	})

	t.Run("ValidateModel", func(t *testing.T) {
		type TestModel struct {
			Name string `validate:"required"`
		}

		model := &TestModel{}
		validate := validator.New()
		postStatusValidator := &modelValidatorImpl{validate: validate}
		err := postStatusValidator.ValidateModel(model)
		assert.Error(t, err)
	})
}
