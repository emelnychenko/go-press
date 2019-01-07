package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("SubjectId", func(t *testing.T) {
		modelId := new(ModelId)
		user := &User{Id: modelId}
		assert.Equal(t, modelId, user.SubjectId())
	})

	t.Run("SubjectType", func(t *testing.T) {
		user := new(User)
		assert.Equal(t, UserSubjectType, user.SubjectType())
	})
}
