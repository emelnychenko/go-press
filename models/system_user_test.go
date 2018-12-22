package models

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystemSubject(t *testing.T) {
	t.Run("NewSystemUser", func(t *testing.T) {
		systemUser := NewSystemUser()
		assert.Equal(t, uuid.Nil, systemUser.Id.UUID)
	})

	t.Run("SubjectId", func(t *testing.T) {
		modelId := new(common.ModelId)
		systemUser := &SystemUser{Id: modelId}
		assert.Equal(t, modelId, systemUser.SubjectId())
	})

	t.Run("SubjectType", func(t *testing.T) {
		systemUser := new(SystemUser)
		assert.Equal(t, enums.SystemSubjectType, systemUser.SubjectType())
	})
}
