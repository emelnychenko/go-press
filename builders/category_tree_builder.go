package builders

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	categoryTreeBuilderImpl struct {
	}
)

//NewCategoryTreeBuilder
func NewCategoryTreeBuilder() contracts.CategoryTreeBuilder {
	return &categoryTreeBuilderImpl{}
}

//TODO: prevent circular dependency
//BuildCategoryEntityTree
func (b *categoryTreeBuilderImpl) BuildCategoryEntityTree(
	categoryEntities []*entities.CategoryEntity,
) (categoryEntityTree *entities.CategoryEntityTree, err common.Error) {
	rootCategoryEntities, nodeCategoryEntities, err := b.prepareCategoryEntities(categoryEntities)

	if nil != err {
		return
	}

	categoryEntityRoots := make([]*entities.CategoryEntityTreeBranch, len(rootCategoryEntities))

	for index, rootCategoryEntity := range rootCategoryEntities {
		categoryEntityRoot := b.buildCategoryEntityBranch(rootCategoryEntity, nodeCategoryEntities)
		categoryEntityRoots[index] = categoryEntityRoot
	}

	categoryEntityTree = &entities.CategoryEntityTree{Roots: categoryEntityRoots}
	return
}

func (b *categoryTreeBuilderImpl) buildCategoryEntityBranch(
	parentCategoryEntity *entities.CategoryEntity,
	categoryEntities []*entities.CategoryEntity,
) (categoryEntityTreeBranch *entities.CategoryEntityTreeBranch) {
	categoryEntityTreeBranch = new(entities.CategoryEntityTreeBranch)
	categoryEntityTreeBranch.CategoryEntity = parentCategoryEntity

	var categoryEntityChildren []*entities.CategoryEntityTreeBranch
	for _, categoryEntity := range categoryEntities {
		// TODO: Find better solution to compare uuid
		if parentCategoryEntity.Id.String() == categoryEntity.ParentCategoryId.String() {
			categoryEntityChild := b.buildCategoryEntityBranch(categoryEntity, categoryEntities)
			categoryEntityChildren = append(categoryEntityChildren, categoryEntityChild)
		}
	}
	categoryEntityTreeBranch.Children = categoryEntityChildren
	return
}

func (*categoryTreeBuilderImpl) prepareCategoryEntities(
	categoryEntities []*entities.CategoryEntity,
) (
	rootCategoryEntities []*entities.CategoryEntity,
	nodeCategoryEntities []*entities.CategoryEntity,
	err common.Error,
) {
	for _, categoryEntity := range categoryEntities {
		if nil == categoryEntity.ParentCategoryId {
			rootCategoryEntities = append(rootCategoryEntities, categoryEntity)
		} else {
			isParentExists := false
			for _, categoryEntityToCheck := range categoryEntities {
				if categoryEntity.ParentCategoryId.String() == categoryEntityToCheck.Id.String() {
					isParentExists = true
					break
				}
			}

			if !isParentExists {
				rootCategoryEntities = append(rootCategoryEntities, categoryEntity)
			} else {
				nodeCategoryEntities = append(nodeCategoryEntities, categoryEntity)
			}
		}
	}

	// Prevent circular dependencies
	if len(categoryEntities) > 0 && len(rootCategoryEntities) == 0 {
		err = common.NewBadRequestError("Circular dependencies detected")
	}

	return
}
