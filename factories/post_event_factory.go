package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	postEventFactoryImpl struct {
	}
)

func NewPostEventFactory() contracts.PostEventFactory {
	return new(postEventFactoryImpl)
}

func (*postEventFactoryImpl) CreatePostCreatedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	return events.NewPostCreatedEvent(postEntity)
}

func (*postEventFactoryImpl) CreatePostUpdatedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	return events.NewPostUpdatedEvent(postEntity)
}

func (*postEventFactoryImpl) CreatePostAuthorChangedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	return events.NewPostAuthorChangedEvent(postEntity)
}

func (*postEventFactoryImpl) CreatePostDeletedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	return events.NewPostDeletedEvent(postEntity)
}
