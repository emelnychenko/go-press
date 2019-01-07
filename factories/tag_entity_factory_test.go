package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagEntityFactory(t *testing.T) {
	t.Run("NewTagEntityFactory", func(t *testing.T) {
		_, isTagEntityFactory := NewTagEntityFactory().(*tagEntityFactoryImpl)

		assert.True(t, isTagEntityFactory)
	})

	t.Run("CreateTagEntity", func(t *testing.T) {
		tagEntityFactory := new(tagEntityFactoryImpl)
		assert.IsType(t, new(entities.TagEntity), tagEntityFactory.CreateTagEntity())
	})

	t.Run("CreateTagXrefEntity", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		postEntity := new(entities.PostEntity)

		tagEntityFactory := new(tagEntityFactoryImpl)
		tagXrefEntity := tagEntityFactory.CreateTagXrefEntity(tagEntity, postEntity)

		assert.IsType(t, new(entities.TagXrefEntity), tagXrefEntity)
	})
}
