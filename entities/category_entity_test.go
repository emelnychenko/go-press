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

	t.Run("NewCategoryXrefEntity", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		categoryEntity := &CategoryEntity{Id: categoryId}
		postId := new(models.PostId)
		postEntity := &PostEntity{Id: postId}

		categoryXrefEntity := NewCategoryXrefEntity(categoryEntity, postEntity)

		assert.Equal(t, categoryId, categoryXrefEntity.CategoryId)
		assert.Equal(t, postEntity.ObjectType(), categoryXrefEntity.ObjectType)
		assert.Equal(t, postId, categoryXrefEntity.ObjectId)
		assert.NotNil(t, categoryXrefEntity.Created)
	})

	t.Run("XrefTableName", func(t *testing.T) {
		categoryXrefEntity := new(CategoryXrefEntity)

		assert.Equal(t, CategoryXrefTableName, categoryXrefEntity.TableName())
	})

	t.Run("XrefSetCategory", func(t *testing.T) {
		categoryXrefEntity := new(CategoryXrefEntity)
		categoryId := new(models.CategoryId)
		categoryEntity := &CategoryEntity{Id: categoryId}

		categoryXrefEntity.SetCategory(categoryEntity)
		assert.Equal(t, categoryId, categoryXrefEntity.CategoryId)
	})

	t.Run("XrefSetObject", func(t *testing.T) {
		postEntity := &PostEntity{Id: new(models.PostId)}

		categoryXrefEntity := new(CategoryXrefEntity)
		categoryXrefEntity.SetObject(postEntity)

		assert.Equal(t, postEntity.ObjectType(), categoryXrefEntity.ObjectType)
		assert.Equal(t, postEntity.ObjectId(), categoryXrefEntity.ObjectId)
	})

	t.Run("SetEntityEdges", func(t *testing.T) {
		categoryEntity := new(CategoryEntity)
		categoryEntityNestedSetNode := CategoryEntityNestedSetNode{CategoryEntity: categoryEntity, Left: 1, Right: 2}

		categoryEntityNestedSetNode.SetEntityEdges()
		assert.Equal(t, 1, categoryEntity.Left)
		assert.Equal(t, 2, categoryEntity.Right)
	})
}
