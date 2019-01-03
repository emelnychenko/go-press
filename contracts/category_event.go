package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	CategoryEvent interface {
		Event
		CategoryEntity() *entities.CategoryEntity
	}
)