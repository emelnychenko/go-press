package models

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("SubjectId", func(t *testing.T) {
		modelId := new(common.ModelId)
		user := &User{Id: modelId}
		assert.Equal(t, modelId, user.SubjectId())
	})

	t.Run("SubjectType", func(t *testing.T) {
		user := new(User)
		assert.Equal(t, enums.UserSubjectType, user.SubjectType())
	})
}
