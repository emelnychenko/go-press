package builders

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	categoryNestedSetBuilderImpl struct {
		categoryTreeBuilder contracts.CategoryTreeBuilder
	}
)

//NewCategoryNestedSetBuilder
func NewCategoryNestedSetBuilder(categoryTreeBuilder contracts.CategoryTreeBuilder) contracts.CategoryNestedSetBuilder {
	return &categoryNestedSetBuilderImpl{categoryTreeBuilder}
}

//BuildCategoryEntityNestedSet
func (b *categoryNestedSetBuilderImpl) BuildCategoryEntityNestedSet(
	categoryEntities []*entities.CategoryEntity,
) *entities.CategoryEntityNestedSet {
	categoryEntityTree := b.categoryTreeBuilder.BuildCategoryEntityTree(categoryEntities)
	return b.BuildCategoryEntityNestedSetFromTree(categoryEntityTree)
}

//BuildCategoryEntityNestedSetFromTree
func (b *categoryNestedSetBuilderImpl) BuildCategoryEntityNestedSetFromTree(
	categoryEntityTree *entities.CategoryEntityTree,
) (categoryEntityNestedSet *entities.CategoryEntityNestedSet) {

	var categoryEntityNestedSetNodes []*entities.CategoryEntityNestedSetNode
	leftCursor := 1
	for _, categoryEntityTreeBranch := range categoryEntityTree.Roots {
		categoryEntityNestedSetNodesFromTreeBranch := b.buildCategoryEntityNestedSetNodesFromTreeBranch(
			categoryEntityTreeBranch, leftCursor)
		categoryEntityNestedSetNodes = append(
			categoryEntityNestedSetNodes, categoryEntityNestedSetNodesFromTreeBranch...)

		leftCursor = leftCursor + len(categoryEntityNestedSetNodesFromTreeBranch) * 2
	}

	categoryEntityNestedSet = &entities.CategoryEntityNestedSet{Nodes: categoryEntityNestedSetNodes}
	return
}

func (b *categoryNestedSetBuilderImpl) buildCategoryEntityNestedSetNodesFromTreeBranch(
	categoryEntityTreeBranch *entities.CategoryEntityTreeBranch,
	leftCursor int,
) (categoryEntityNestedSetNodes []*entities.CategoryEntityNestedSetNode) {
	categoryEntityNestedSetNode := &entities.CategoryEntityNestedSetNode{
		CategoryEntity: categoryEntityTreeBranch.CategoryEntity, Left: leftCursor}
	categoryEntityNestedSetNodes = append(categoryEntityNestedSetNodes, categoryEntityNestedSetNode)

	leftCursor++
	for _, categoryEntityTreeBranch := range categoryEntityTreeBranch.Children {
		categoryEntityNestedSetNodesFromTreeBranch := b.buildCategoryEntityNestedSetNodesFromTreeBranch(
			categoryEntityTreeBranch, leftCursor)

		leftCursor = leftCursor + len(categoryEntityNestedSetNodesFromTreeBranch) * 2
		categoryEntityNestedSetNodes = append(categoryEntityNestedSetNodes, categoryEntityNestedSetNodesFromTreeBranch...)
	}

	categoryEntityNestedSetNode.Right = leftCursor
	return
}
