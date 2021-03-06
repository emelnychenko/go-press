package builders

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
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
) (categoryEntityNestedSet *entities.CategoryEntityNestedSet, err errors.Error) {
	categoryEntityTree, err := b.categoryTreeBuilder.BuildCategoryEntityTree(categoryEntities)

	if nil != err {
		return
	}

	categoryEntityNestedSet = b.BuildCategoryEntityNestedSetFromTree(categoryEntityTree)
	return
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

		leftCursor = leftCursor + len(categoryEntityNestedSetNodesFromTreeBranch)*2
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

		leftCursor = leftCursor + len(categoryEntityNestedSetNodesFromTreeBranch)*2
		categoryEntityNestedSetNodes = append(categoryEntityNestedSetNodes, categoryEntityNestedSetNodesFromTreeBranch...)
	}

	categoryEntityNestedSetNode.Right = leftCursor
	return
}
