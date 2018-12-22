package entities

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserEntity(t *testing.T) {
	t.Run("NewUserEntity", func(t *testing.T) {
		userEntity := NewUserEntity()

		assert.NotNil(t, userEntity.Id)
		assert.False(t, userEntity.Verified)
		assert.NotNil(t, userEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		userEntity := new(UserEntity)

		assert.Equal(t, UserTable, userEntity.TableName())
	})

	t.Run("SubjectId", func(t *testing.T) {
		modelId := new(common.ModelId)
		userEntity := &UserEntity{Id: modelId}
		assert.Equal(t, modelId, userEntity.SubjectId())
	})

	t.Run("SubjectType", func(t *testing.T) {
		userEntity := new(UserEntity)

		assert.Equal(t, enums.UserSubjectType, userEntity.SubjectType())
	})
}
