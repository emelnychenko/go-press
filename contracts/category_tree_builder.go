package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	CategoryTreeBuilder interface {
		BuildCategoryEntityTree(categoryEntities []*entities.CategoryEntity) *entities.CategoryEntityTree
	}
)