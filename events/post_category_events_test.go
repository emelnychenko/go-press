package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostCategoryEvents(t *testing.T) {
	t.Run("NewPostCategoryAddedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		categoryEntity := new(entities.CategoryEntity)

		postCategoryEvent, isPostCategoryEvent := NewPostCategoryAddedEvent(postEntity, categoryEntity).(*PostCategoryEvent)

		assert.True(t, isPostCategoryEvent)
		assert.Equal(t, postEntity, postCategoryEvent.postEntity)
		assert.Equal(t, categoryEntity, postCategoryEvent.categoryEntity)
		assert.Equal(t, PostCategoryAddedEventName, postCategoryEvent.name)
	})

	t.Run("NewPostCategoryRemovedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		categoryEntity := new(entities.CategoryEntity)

		postCategoryEvent, isPostCategoryEvent := NewPostCategoryRemovedEvent(postEntity, categoryEntity).(*PostCategoryEvent)

		assert.True(t, isPostCategoryEvent)
		assert.Equal(t, postEntity, postCategoryEvent.postEntity)
		assert.Equal(t, categoryEntity, postCategoryEvent.categoryEntity)
		assert.Equal(t, PostCategoryRemovedEventName, postCategoryEvent.name)
	})

	t.Run("PostCategoryEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		categoryEntity := new(entities.CategoryEntity)

		postCategoryEvent := &PostCategoryEvent{postEntity: postEntity, categoryEntity: categoryEntity}

		assert.Equal(t, postEntity, postCategoryEvent.PostEntity())
		assert.Equal(t, categoryEntity, postCategoryEvent.CategoryEntity())
	})
}
