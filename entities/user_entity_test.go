package entities

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserEntity(t *testing.T) {
	t.Run("NewUserEntity", func(t *testing.T) {
		userEntity := NewUserEntity()

		assert.IsType(t, new(models.UserId), userEntity.Id)
		assert.False(t, userEntity.Verified)
		assert.NotNil(t, userEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		userEntity := new(UserEntity)

		assert.Equal(t, UserTableName, userEntity.TableName())
	})

	t.Run("SubjectId", func(t *testing.T) {
		modelId := new(models.SubjectId)
		userEntity := &UserEntity{Id: modelId}
		assert.Equal(t, modelId, userEntity.SubjectId())
	})

	t.Run("SubjectType", func(t *testing.T) {
		userEntity := new(UserEntity)

		assert.Equal(t, UserEntitySubjectType, userEntity.SubjectType())
	})

	t.Run("SetPicture", func(t *testing.T) {
		modelId := new(models.ModelId)
		postPicture := &FileEntity{Id: modelId}
		userEntity := new(UserEntity)

		userEntity.SetPicture(postPicture)
		assert.Equal(t, modelId, userEntity.PictureId)
	})

	t.Run("RemovePicture", func(t *testing.T) {
		modelId := new(models.ModelId)
		postEntity := &UserEntity{PictureId: modelId}

		postEntity.RemovePicture()
		assert.Nil(t, postEntity.PictureId)
	})
}
