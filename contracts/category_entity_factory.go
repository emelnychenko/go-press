package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	CategoryEntityFactory interface {
		CreateCategoryEntity() *entities.CategoryEntity
	}
)
