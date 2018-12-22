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

	t.Run("SetAuthor", func(t *testing.T) {
		modelId := new(common.ModelId)
		postEntity := new(PostEntity)
		userEntity := &UserEntity{Id: modelId}
		postEntity.SetAuthor(userEntity)

		assert.Equal(t, modelId, postEntity.AuthorId)
		assert.Equal(t, userEntity.SubjectType(), postEntity.AuthorType)
	})
}
