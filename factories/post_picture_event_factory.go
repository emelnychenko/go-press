package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	postPictureEventFactoryImpl struct {
	}
)

func NewPostPictureEventFactory() contracts.PostPictureEventFactory {
	return new(postPictureEventFactoryImpl)
}

func (*postPictureEventFactoryImpl) CreatePostPictureChangedEvent(
	postEntity *entities.PostEntity,
	postPicture *entities.FileEntity,
) contracts.PostPictureEvent {
	return events.NewPostPictureChangedEvent(postEntity, postPicture)
}

func (*postPictureEventFactoryImpl) CreatePostPictureRemovedEvent(postEntity *entities.PostEntity) contracts.PostPictureEvent {
	return events.NewPostPictureRemovedEvent(postEntity)
}
