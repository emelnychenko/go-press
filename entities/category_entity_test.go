package entities

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryEntity(t *testing.T) {
	t.Run("NewCategoryEntity", func(t *testing.T) {
		categoryEntity := NewCategoryEntity()

		assert.IsType(t, new(models.CategoryId), categoryEntity.Id)
		assert.NotNil(t, categoryEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		categoryEntity := new(CategoryEntity)

		assert.Equal(t, CategoryTableName, categoryEntity.TableName())
	})

	t.Run("SetParentCategory", func(t *testing.T) {
		parentCategoryId := new(models.CategoryId)
		categoryEntity := new(CategoryEntity)
		parentCategoryEntity := &CategoryEntity{Id: parentCategoryId}

		categoryEntity.SetParentCategory(parentCategoryEntity)
		assert.Equal(t, parentCategoryId, categoryEntity.ParentCategoryId)
	})

	t.Run("RemoveParentCategory", func(t *testing.T) {
		categoryEntity := &CategoryEntity{ParentCategoryId: new(models.CategoryId)}
		categoryEntity.RemoveParentCategory()

		assert.Nil(t, categoryEntity.ParentCategoryId)
	})

	t.Run("EdgesDifferent", func(t *testing.T) {
		categoryEntity := &CategoryEntity{Left: 1, Right: 2}
		categoryEntityNestedSetNode := CategoryEntityNestedSetNode{CategoryEntity: categoryEntity, Left: 1, Right: 4}
		assert.True(t, categoryEntityNestedSetNode.EdgesDifferent())
	})
}
