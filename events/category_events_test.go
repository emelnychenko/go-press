package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryEvents(t *testing.T) {
	t.Run("NewCategoryCreatedEvent", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryEvent, isCategoryEvent := NewCategoryCreatedEvent(categoryEntity).(*CategoryEvent)

		assert.True(t, isCategoryEvent)
		assert.Equal(t, categoryEntity, categoryEvent.categoryEntity)
		assert.Equal(t, CategoryCreatedEventName, categoryEvent.name)
	})

	t.Run("NewCategoryUpdatedEvent", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryEvent, isCategoryEvent := NewCategoryUpdatedEvent(categoryEntity).(*CategoryEvent)

		assert.True(t, isCategoryEvent)
		assert.Equal(t, categoryEntity, categoryEvent.categoryEntity)
		assert.Equal(t, CategoryUpdatedEventName, categoryEvent.name)
	})

	t.Run("NewCategoryDeletedEvent", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryEvent, isCategoryEvent := NewCategoryDeletedEvent(categoryEntity).(*CategoryEvent)

		assert.True(t, isCategoryEvent)
		assert.Equal(t, categoryEntity, categoryEvent.categoryEntity)
		assert.Equal(t, CategoryDeletedEventName, categoryEvent.name)
	})

	t.Run("CategoryEntity", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryEvent := &CategoryEvent{categoryEntity: categoryEntity}

		assert.Equal(t, categoryEntity, categoryEvent.CategoryEntity())
	})
}
