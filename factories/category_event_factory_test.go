package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryEventFactory(t *testing.T) {
	t.Run("NewCategoryEventFactory", func(t *testing.T) {
		_, isCategoryEventFactory := NewCategoryEventFactory().(*categoryEventFactoryImpl)

		assert.True(t, isCategoryEventFactory)
	})

	t.Run("CreateCategoryCreatedEvent", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryEventFactory := new(categoryEventFactoryImpl)
		categoryEvent := categoryEventFactory.CreateCategoryCreatedEvent(categoryEntity)

		assert.Equal(t, events.CategoryCreatedEventName, categoryEvent.Name())
		assert.Equal(t, categoryEntity, categoryEvent.CategoryEntity())
	})

	t.Run("CreateCategoryUpdatedEvent", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryEventFactory := new(categoryEventFactoryImpl)
		categoryEvent := categoryEventFactory.CreateCategoryUpdatedEvent(categoryEntity)

		assert.Equal(t, events.CategoryUpdatedEventName, categoryEvent.Name())
		assert.Equal(t, categoryEntity, categoryEvent.CategoryEntity())
	})

	t.Run("CreateCategoryDeletedEvent", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryEventFactory := new(categoryEventFactoryImpl)
		categoryEvent := categoryEventFactory.CreateCategoryDeletedEvent(categoryEntity)

		assert.Equal(t, events.CategoryDeletedEventName, categoryEvent.Name())
		assert.Equal(t, categoryEntity, categoryEvent.CategoryEntity())
	})

	t.Run("CreateCategoryParentChangedEvent", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryEventFactory := new(categoryEventFactoryImpl)
		categoryEvent := categoryEventFactory.CreateCategoryParentChangedEvent(categoryEntity)

		assert.Equal(t, events.CategoryParentChangedEventName, categoryEvent.Name())
		assert.Equal(t, categoryEntity, categoryEvent.CategoryEntity())
	})

	t.Run("CreateCategoryParentRemovedEvent", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryEventFactory := new(categoryEventFactoryImpl)
		categoryEvent := categoryEventFactory.CreateCategoryParentRemovedEvent(categoryEntity)

		assert.Equal(t, events.CategoryParentRemovedEventName, categoryEvent.Name())
		assert.Equal(t, categoryEntity, categoryEvent.CategoryEntity())
	})
}
