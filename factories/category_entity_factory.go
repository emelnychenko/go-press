package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	categoryEntityFactoryImpl struct {
	}
)

func NewCategoryEntityFactory() contracts.CategoryEntityFactory {
	return new(categoryEntityFactoryImpl)
}

func (*categoryEntityFactoryImpl) CreateCategoryEntity() *entities.CategoryEntity {
	return entities.NewCategoryEntity()
}
