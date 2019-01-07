package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	categoryEntityFactoryImpl struct {
	}
)

//NewCategoryEntityFactory
func NewCategoryEntityFactory() contracts.CategoryEntityFactory {
	return new(categoryEntityFactoryImpl)
}

//CreateCategoryEntity
func (*categoryEntityFactoryImpl) CreateCategoryEntity() *entities.CategoryEntity {
	return entities.NewCategoryEntity()
}

//CreateCategoryXrefEntity
func (*categoryEntityFactoryImpl) CreateCategoryXrefEntity(
	categoryEntity *entities.CategoryEntity, categoryObject models.Object,
) *entities.CategoryXrefEntity {
	return entities.NewCategoryXrefEntity(categoryEntity, categoryObject)
}
