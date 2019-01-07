package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryEntityFactory interface {
		CreateCategoryEntity() *entities.CategoryEntity
		CreateCategoryXrefEntity(*entities.CategoryEntity, models.Object) *entities.CategoryXrefEntity
	}
)
