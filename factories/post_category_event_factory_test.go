package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostCategoryEventFactory(t *testing.T) {
	t.Run("NewPostCategoryEventFactory", func(t *testing.T) {
		_, isPostCategoryEventFactory := NewPostCategoryEventFactory().(*postCategoryEventFactoryImpl)

		assert.True(t, isPostCategoryEventFactory)
	})

	t.Run("CreatePostCategoryAddedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		categoryEntity := new(entities.CategoryEntity)

		postCategoryEventFactory := new(postCategoryEventFactoryImpl)
		postCategoryEvent := postCategoryEventFactory.CreatePostCategoryAddedEvent(postEntity, categoryEntity)

		assert.Equal(t, events.PostCategoryAddedEventName, postCategoryEvent.Name())
		assert.Equal(t, postEntity, postCategoryEvent.PostEntity())
		assert.Equal(t, categoryEntity, postCategoryEvent.CategoryEntity())
	})

	t.Run("CreatePostCategoryRemovedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		categoryEntity := new(entities.CategoryEntity)

		postCategoryEventFactory := new(postCategoryEventFactoryImpl)
		postCategoryEvent := postCategoryEventFactory.CreatePostCategoryRemovedEvent(postEntity, categoryEntity)

		assert.Equal(t, events.PostCategoryRemovedEventName, postCategoryEvent.Name())
		assert.Equal(t, postEntity, postCategoryEvent.PostEntity())
		assert.Equal(t, categoryEntity, postCategoryEvent.CategoryEntity())
	})
}
