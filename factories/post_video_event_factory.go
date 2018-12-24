package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	postVideoEventFactoryImpl struct {
	}
)

func NewPostVideoEventFactory() contracts.PostVideoEventFactory {
	return new(postVideoEventFactoryImpl)
}

func (*postVideoEventFactoryImpl) CreatePostVideoChangedEvent(
	postEntity *entities.PostEntity,
	postVideo *entities.FileEntity,
) contracts.PostVideoEvent {
	return events.NewPostVideoChangedEvent(postEntity, postVideo)
}

func (*postVideoEventFactoryImpl) CreatePostVideoRemovedEvent(postEntity *entities.PostEntity) contracts.PostVideoEvent {
	return events.NewPostVideoRemovedEvent(postEntity)
}
