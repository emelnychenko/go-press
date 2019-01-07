package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
)

type (
	CategoryTreeBuilder interface {
		BuildCategoryEntityTree(categoryEntities []*entities.CategoryEntity) (*entities.CategoryEntityTree, errors.Error)
	}
)
