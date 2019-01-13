package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	postTagEventFactoryImpl struct {
	}
)

//NewPostTagEventFactory
func NewPostTagEventFactory() contracts.PostTagEventFactory {
	return new(postTagEventFactoryImpl)
}

//CreatePostTagAddedEvent
func (*postTagEventFactoryImpl) CreatePostTagAddedEvent(
	postEntity *entities.PostEntity, tagEntity *entities.TagEntity,
) contracts.PostTagEvent {
	return events.NewPostTagAddedEvent(postEntity, tagEntity)
}

//CreatePostTagRemovedEvent
func (*postTagEventFactoryImpl) CreatePostTagRemovedEvent(
	postEntity *entities.PostEntity, tagEntity *entities.TagEntity,
) contracts.PostTagEvent {
	return events.NewPostTagRemovedEvent(postEntity, tagEntity)
}
