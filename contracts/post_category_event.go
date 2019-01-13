package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostCategoryEvent interface {
		Event
		PostEntity() *entities.PostEntity
		CategoryEntity() *entities.CategoryEntity
	}
)
