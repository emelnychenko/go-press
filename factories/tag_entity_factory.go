package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
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
