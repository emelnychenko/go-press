package common

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModelId(t *testing.T) {
	t.Run("NewModelId", func(t *testing.T) {
		e := NewModelId()
		assert.IsType(t, new(ModelId), e)
	})

	t.Run("ParseModelId", func(t *testing.T) {
		userId, err := ParseModelId(uuid.New().String())
		assert.IsType(t, new(ModelId), userId)
		assert.Nil(t, err)
	})

	t.Run("ParseModelId:Error", func(t *testing.T) {
		userId, err := ParseModelId("")
		assert.Nil(t, userId)
		assert.Error(t, err)
	})
}
