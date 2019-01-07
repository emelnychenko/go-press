package entities

import (
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostEntity(t *testing.T) {
	t.Run("NewPostEntity", func(t *testing.T) {
		postEntity := NewPostEntity()

		assert.IsType(t, new(models.PostId), postEntity.Id)
		assert.Equal(t, enums.PostDraftStatus, postEntity.Status)
		assert.Equal(t, enums.PostPublicPrivacy, postEntity.Privacy)
		assert.NotNil(t, postEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		postEntity := new(PostEntity)

		assert.Equal(t, PostTableName, postEntity.TableName())
	})

	t.Run("ObjectId", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := &PostEntity{Id: postId}
		assert.Equal(t, postId, postEntity.ObjectId())
	})

	t.Run("ObjectType", func(t *testing.T) {
		postEntity := new(PostEntity)
		assert.Equal(t, PostEntityObjectType, postEntity.ObjectType())
	})

	t.Run("SetPicture", func(t *testing.T) {
		modelId := new(models.ModelId)
		postPicture := &FileEntity{Id: modelId}
		postEntity := new(PostEntity)

		postEntity.SetPicture(postPicture)
		assert.Equal(t, modelId, postEntity.PictureId)
	})

	t.Run("RemovePicture", func(t *testing.T) {
		modelId := new(models.ModelId)
		postEntity := &PostEntity{PictureId: modelId}

		postEntity.RemovePicture()
		assert.Nil(t, postEntity.PictureId)
	})

	t.Run("SetVideo", func(t *testing.T) {
		modelId := new(models.ModelId)
		postVideo := &FileEntity{Id: modelId}
		postEntity := new(PostEntity)

		postEntity.SetVideo(postVideo)
		assert.Equal(t, modelId, postEntity.VideoId)
	})

	t.Run("RemoveVideo", func(t *testing.T) {
		modelId := new(models.ModelId)
		postEntity := &PostEntity{VideoId: modelId}

		postEntity.RemoveVideo()
		assert.Nil(t, postEntity.VideoId)
	})

	t.Run("SetAuthor", func(t *testing.T) {
		userId := new(models.UserId)
		postEntity := new(PostEntity)
		userEntity := &UserEntity{Id: userId}
		postEntity.SetAuthor(userEntity)

		assert.Equal(t, userId, postEntity.AuthorId)
		assert.Equal(t, userEntity.SubjectType(), postEntity.AuthorType)
	})
}
