package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	postAuthorEventFactoryImpl struct {
	}
)

func NewPostAuthorEventFactory() contracts.PostAuthorEventFactory {
	return new(postAuthorEventFactoryImpl)
}

func (*postAuthorEventFactoryImpl) CreatePostAuthorChangedEvent(
	postEntity *entities.PostEntity,
	postAuthorEntity *entities.UserEntity,
) contracts.PostAuthorEvent {
	return events.NewPostAuthorChangedEvent(postEntity, postAuthorEntity)
}
