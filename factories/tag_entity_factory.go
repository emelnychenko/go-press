package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	tagEntityFactoryImpl struct {
	}
)

func NewTagEntityFactory() contracts.TagEntityFactory {
	return new(tagEntityFactoryImpl)
}

func (*tagEntityFactoryImpl) CreateTagEntity() *entities.TagEntity {
	return entities.NewTagEntity()
}

//CreateTagXrefEntity
func (*tagEntityFactoryImpl) CreateTagXrefEntity(
	tagEntity *entities.TagEntity, tagObject models.Object,
) *entities.TagXrefEntity {
	return entities.NewTagXrefEntity(tagEntity, tagObject)
}