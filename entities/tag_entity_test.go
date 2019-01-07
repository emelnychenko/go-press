package entities

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagEntity(t *testing.T) {
	t.Run("NewTagEntity", func(t *testing.T) {
		tagEntity := NewTagEntity()

		assert.IsType(t, new(models.TagId), tagEntity.Id)
		assert.NotNil(t, tagEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		tagEntity := new(TagEntity)

		assert.Equal(t, TagTableName, tagEntity.TableName())
	})

	t.Run("NewTagXrefEntity", func(t *testing.T) {
		tagId := new(models.TagId)
		tagEntity := &TagEntity{Id: tagId}
		postId := new(models.PostId)
		postEntity := &PostEntity{Id: postId}

		tagXrefEntity := NewTagXrefEntity(tagEntity, postEntity)

		assert.Equal(t, tagId, tagXrefEntity.TagId)
		assert.Equal(t, postEntity.ObjectType(), tagXrefEntity.ObjectType)
		assert.Equal(t, postId, tagXrefEntity.ObjectId)
		assert.NotNil(t, tagXrefEntity.Created)
	})

	t.Run("XrefTableName", func(t *testing.T) {
		tagXrefEntity := new(TagXrefEntity)

		assert.Equal(t, TagXrefTableName, tagXrefEntity.TableName())
	})

	t.Run("XrefSetTag", func(t *testing.T) {
		tagXrefEntity := new(TagXrefEntity)
		tagId := new(models.TagId)
		tagEntity := &TagEntity{Id: tagId}

		tagXrefEntity.SetTag(tagEntity)
		assert.Equal(t, tagId, tagXrefEntity.TagId)
	})

	t.Run("XrefSetObject", func(t *testing.T) {
		postEntity := &PostEntity{Id: new(models.PostId)}

		tagXrefEntity := new(TagXrefEntity)
		tagXrefEntity.SetObject(postEntity)

		assert.Equal(t, postEntity.ObjectType(), tagXrefEntity.ObjectType)
		assert.Equal(t, postEntity.ObjectId(), tagXrefEntity.ObjectId)
	})
}
