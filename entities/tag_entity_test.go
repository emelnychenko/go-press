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
}
