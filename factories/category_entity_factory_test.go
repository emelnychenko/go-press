package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryEntityFactory(t *testing.T) {
	t.Run("NewCategoryEntityFactory", func(t *testing.T) {
		_, isCategoryEntityFactory := NewCategoryEntityFactory().(*categoryEntityFactoryImpl)

		assert.True(t, isCategoryEntityFactory)
	})

	t.Run("CreateCategoryEntity", func(t *testing.T) {
		categoryEntityFactory := new(categoryEntityFactoryImpl)
		assert.IsType(t, new(entities.CategoryEntity), categoryEntityFactory.CreateCategoryEntity())
	})

	t.Run("CreateCategoryXrefEntity", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		postEntity := new(entities.PostEntity)

		categoryEntityFactory := new(categoryEntityFactoryImpl)
		categoryXrefEntity := categoryEntityFactory.CreateCategoryXrefEntity(categoryEntity, postEntity)

		assert.IsType(t, new(entities.CategoryXrefEntity), categoryXrefEntity)
	})
}
