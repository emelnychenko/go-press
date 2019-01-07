package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
)

type (
	CategoryNestedSetBuilder interface {
		BuildCategoryEntityNestedSet(categoryEntities []*entities.CategoryEntity) (*entities.CategoryEntityNestedSet, common.Error)
		BuildCategoryEntityNestedSetFromTree(categoryEntityTree *entities.CategoryEntityTree) *entities.CategoryEntityNestedSet
	}
)
