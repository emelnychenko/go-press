package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
)

type (
	CategoryNestedSetBuilder interface {
		BuildCategoryEntityNestedSet(categoryEntities []*entities.CategoryEntity) (*entities.CategoryEntityNestedSet, errors.Error)
		BuildCategoryEntityNestedSetFromTree(categoryEntityTree *entities.CategoryEntityTree) *entities.CategoryEntityNestedSet
	}
)
