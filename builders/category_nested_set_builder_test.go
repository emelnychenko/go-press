package builders

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryNestedSetBuilder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCategoryTreeBuilder", func(t *testing.T) {
		categoryTreeBuilder := mocks.NewMockCategoryTreeBuilder(ctrl)
		categoryNestedSetBuilder, isCategoryNestedSetBuilder := NewCategoryNestedSetBuilder(
			categoryTreeBuilder).(*categoryNestedSetBuilderImpl)

		assert.True(t, isCategoryNestedSetBuilder)
		assert.Equal(t, categoryTreeBuilder, categoryNestedSetBuilder.categoryTreeBuilder)
	})

	t.Run("BuildCategoryEntityNestedSet", func(t *testing.T) {
		categoryEntityRoot0 := &entities.CategoryEntity{Id: common.NewModelId()}
		categoryEntityRoot0Node0 := &entities.CategoryEntity{
			Id: common.NewModelId(), ParentCategoryId: categoryEntityRoot0.Id}
		categoryEntityRoot0Node1 := &entities.CategoryEntity{
			Id: common.NewModelId(), ParentCategoryId: categoryEntityRoot0.Id}
		categoryEntityRoot0Node1Node0 := &entities.CategoryEntity{
			Id: common.NewModelId(), ParentCategoryId: categoryEntityRoot0Node1.Id}
		categoryEntityRoot0Node1Node1 := &entities.CategoryEntity{
			Id: common.NewModelId(), ParentCategoryId: categoryEntityRoot0Node1.Id}
		categoryEntityRoot1 := &entities.CategoryEntity{Id: common.NewModelId()}
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

		categoryTreeBuilder := mocks.NewMockCategoryTreeBuilder(ctrl)
		categoryTreeBuilder.EXPECT().BuildCategoryEntityTree(categoryEntities).DoAndReturn(func(
			categoryEntities []*entities.CategoryEntity,
		) (categoryEntityTree *entities.CategoryEntityTree) {
			return new(categoryTreeBuilderImpl).BuildCategoryEntityTree(categoryEntities)
		})

		categoryNestedSetBuilder := &categoryNestedSetBuilderImpl{categoryTreeBuilder: categoryTreeBuilder}
		categoryEntityNestedSet := categoryNestedSetBuilder.BuildCategoryEntityNestedSet(categoryEntities)

		// Test Nodes
		assert.Len(t, categoryEntityNestedSet.Nodes, 7)
		// Test Node0
		assert.Equal(t, categoryEntityNestedSet.Nodes[0].Left, 1)
		assert.Equal(t, categoryEntityNestedSet.Nodes[0].Right, 10)
		// Test Node1
		assert.Equal(t, categoryEntityNestedSet.Nodes[1].Left, 2)
		assert.Equal(t, categoryEntityNestedSet.Nodes[1].Right, 3)
		// Test Node2
		assert.Equal(t, categoryEntityNestedSet.Nodes[2].Left, 4)
		assert.Equal(t, categoryEntityNestedSet.Nodes[2].Right, 9)
		// Test Node3
		assert.Equal(t, categoryEntityNestedSet.Nodes[3].Left, 5)
		assert.Equal(t, categoryEntityNestedSet.Nodes[3].Right, 6)
		// Test Node4
		assert.Equal(t, categoryEntityNestedSet.Nodes[4].Left, 7)
		assert.Equal(t, categoryEntityNestedSet.Nodes[4].Right, 8)
		// Test Node5
		assert.Equal(t, categoryEntityNestedSet.Nodes[5].Left, 11)
		assert.Equal(t, categoryEntityNestedSet.Nodes[5].Right, 14)
		// Test Node6
		assert.Equal(t, categoryEntityNestedSet.Nodes[6].Left, 12)
		assert.Equal(t, categoryEntityNestedSet.Nodes[6].Right, 13)
	})
}
