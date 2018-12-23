package entities

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostEntity(t *testing.T) {
	t.Run("NewPostEntity", func(t *testing.T) {
		postEntity := NewPostEntity()

		assert.NotNil(t, postEntity.Id)
		assert.Equal(t, enums.PostDraftStatus, postEntity.Status)
		assert.Equal(t, enums.PostPublicPrivacy, postEntity.Privacy)
		assert.NotNil(t, postEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		postEntity := new(PostEntity)

		assert.Equal(t, PostTable, postEntity.TableName())
	})

	t.Run("SetPicture", func(t *testing.T) {
		modelId := new(common.ModelId)
		postPicture := &FileEntity{Id: modelId}
		postEntity := new(PostEntity)

		postEntity.SetPicture(postPicture)
		assert.Equal(t, modelId, postEntity.PictureId)
	})

	t.Run("RemovePicture", func(t *testing.T) {
		modelId := new(common.ModelId)
		postEntity := &PostEntity{PictureId: modelId}

		postEntity.RemovePicture()
		assert.Nil(t, postEntity.PictureId)
	})

	t.Run("SetVideo", func(t *testing.T) {
		modelId := new(common.ModelId)
		postVideo := &FileEntity{Id: modelId}
		postEntity := new(PostEntity)

		postEntity.SetVideo(postVideo)
		assert.Equal(t, modelId, postEntity.VideoId)
	})

	t.Run("RemoveVideo", func(t *testing.T) {
		modelId := new(common.ModelId)
		postEntity := &PostEntity{VideoId: modelId}

		postEntity.RemoveVideo()
		assert.Nil(t, postEntity.VideoId)
	})

	t.Run("SetAuthor", func(t *testing.T) {
		modelId := new(common.ModelId)
		postEntity := new(PostEntity)
		userEntity := &UserEntity{Id: modelId}
		postEntity.SetAuthor(userEntity)

		assert.Equal(t, modelId, postEntity.AuthorId)
		assert.Equal(t, userEntity.SubjectType(), postEntity.AuthorType)
	})
}
