package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	postEntityFactoryImpl struct {
	}
)

func NewPostEntityFactory() contracts.PostEntityFactory {
	return new(postEntityFactoryImpl)
}

func (*postEntityFactoryImpl) CreatePostEntity() *entities.PostEntity {
	return entities.NewPostEntity()
}
