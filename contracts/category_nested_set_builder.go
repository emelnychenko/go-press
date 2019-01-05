package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	CategoryNestedSetBuilder interface {
		BuildCategoryEntityNestedSet(categoryEntities []*entities.CategoryEntity) *entities.CategoryEntityNestedSet
		BuildCategoryEntityNestedSetFromTree(categoryEntityTree *entities.CategoryEntityTree) *entities.CategoryEntityNestedSet
	}
)
