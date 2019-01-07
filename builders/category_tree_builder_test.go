package builders

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryTreeBuilder(t *testing.T) {
	t.Run("NewCategoryTreeBuilder", func(t *testing.T) {
		_, isCategoryTreeBuilder := NewCategoryTreeBuilder().(*categoryTreeBuilderImpl)
		assert.True(t, isCategoryTreeBuilder)
	})

	t.Run("BuildCategoryEntityTree", func(t *testing.T) {
		categoryEntityRoot0 := &entities.CategoryEntity{Id: common.NewModelId()}
		categoryEntityRoot0Node0 := &entities.CategoryEntity{
			Id: common.NewModelId(), ParentCategoryId: categoryEntityRoot0.Id}
		categoryEntityRoot0Node1 := &entities.CategoryEntity{
			Id: common.NewModelId(), ParentCategoryId: categoryEntityRoot0.Id}
		categoryEntityRoot0Node1Node0 := &entities.CategoryEntity{
			Id: common.NewModelId(), ParentCategoryId: categoryEntityRoot0Node1.Id}
		categoryEntityRoot0Node1Node1 := &entities.CategoryEntity{
			Id: common.NewModelId(), ParentCategoryId: categoryEntityRoot0Node1.Id}
		categoryEntityRoot1 := &entities.CategoryEntity{Id: common.NewModelId(), ParentCategoryId: common.NewModelId()}
		categoryEntityRoot1Node0 := &entities.CategoryEntity{
			Id: common.NewModelId(), ParentCategoryId: categoryEntityRoot1.Id}

		categoryEntities := []*entities.CategoryEntity{
			categoryEntityRoot0,
			categoryEntityRoot0Node0,
			categoryEntityRoot0Node1,
			categoryEntityRoot0Node1Node0,
			categoryEntityRoot0Node1Node1,
			categoryEntityRoot1,
			categoryEntityRoot1Node0,
		}

		categoryTreeBuilder := new(categoryTreeBuilderImpl)
		categoryEntityTree, err := categoryTreeBuilder.BuildCategoryEntityTree(categoryEntities)

		assert.Nil(t, err)
		// Test Roots
		assert.Len(t, categoryEntityTree.Roots, 2)
		// Test Root0
		assert.Equal(t, categoryEntityRoot0, categoryEntityTree.Roots[0].CategoryEntity)
		assert.Len(t, categoryEntityTree.Roots[0].Children, 2)
		// Test Root0Node0
		assert.Equal(t, categoryEntityRoot0Node0, categoryEntityTree.Roots[0].Children[0].CategoryEntity)
		assert.Len(t, categoryEntityTree.Roots[0].Children[0].Children, 0)
		// Test Root0Node1
		assert.Equal(t, categoryEntityRoot0Node1, categoryEntityTree.Roots[0].Children[1].CategoryEntity)
		assert.Len(t, categoryEntityTree.Roots[0].Children[1].Children, 2)
		// Test Root0Node1Node0
		assert.Equal(t, categoryEntityRoot0Node1Node0, categoryEntityTree.Roots[0].Children[1].Children[0].CategoryEntity)
		assert.Len(t, categoryEntityTree.Roots[0].Children[1].Children[0].Children, 0)
		// Test Root0Node1Node1
		assert.Equal(t, categoryEntityRoot0Node1Node1, categoryEntityTree.Roots[0].Children[1].Children[1].CategoryEntity)
		assert.Len(t, categoryEntityTree.Roots[0].Children[1].Children[1].Children, 0)
		// Test Root1
		assert.Equal(t, categoryEntityRoot1, categoryEntityTree.Roots[1].CategoryEntity)
		assert.Len(t, categoryEntityTree.Roots[1].Children, 1)
		// Test Root1Node0
		assert.Equal(t, categoryEntityRoot1Node0, categoryEntityTree.Roots[1].Children[0].CategoryEntity)
		assert.Len(t, categoryEntityTree.Roots[1].Children[0].Children, 0)
	})

	t.Run("BuildCategoryEntityTree:CircularDependencyError", func(t *testing.T) {
		chainConnectorId := common.NewModelId()

		categoryEntity0 := &entities.CategoryEntity{Id: common.NewModelId(), ParentCategoryId: chainConnectorId}
		categoryEntity1 := &entities.CategoryEntity{Id: common.NewModelId(), ParentCategoryId: categoryEntity0.Id}
		categoryEntity2 := &entities.CategoryEntity{Id: chainConnectorId, ParentCategoryId: categoryEntity1.Id}

		categoryEntities := []*entities.CategoryEntity{
			categoryEntity0,
			categoryEntity1,
			categoryEntity2,
		}

		categoryTreeBuilder := new(categoryTreeBuilderImpl)
		categoryEntityTree, err := categoryTreeBuilder.BuildCategoryEntityTree(categoryEntities)

		assert.Nil(t, categoryEntityTree)
		assert.Error(t, err)
	})

	t.Run("prepareCategoryEntities", func(t *testing.T) {
		parentCategoryId := new(models.CategoryId)
		rootCategoryEntity := &entities.CategoryEntity{Id: parentCategoryId}
		nodeCategoryEntity := &entities.CategoryEntity{Id: common.NewModelId(), ParentCategoryId: parentCategoryId}
		categoryEntities := []*entities.CategoryEntity{rootCategoryEntity, nodeCategoryEntity}

		categoryTreeBuilder := new(categoryTreeBuilderImpl)
		rootCategoryEntities, nodeCategoryEntities, err := categoryTreeBuilder.prepareCategoryEntities(categoryEntities)

		assert.Nil(t, err)
		assert.Len(t, rootCategoryEntities, 1)
		assert.Equal(t, rootCategoryEntity, rootCategoryEntities[0])

		assert.Len(t, nodeCategoryEntities, 1)
		assert.Equal(t, nodeCategoryEntity, nodeCategoryEntities[0])
	})
}
