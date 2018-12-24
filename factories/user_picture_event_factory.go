package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	userPictureEventFactoryImpl struct {
	}
)

func NewUserPictureEventFactory() contracts.UserPictureEventFactory {
	return new(userPictureEventFactoryImpl)
}

func (*userPictureEventFactoryImpl) CreateUserPictureChangedEvent(
	userEntity *entities.UserEntity,
	userPicture *entities.FileEntity,
) contracts.UserPictureEvent {
	return events.NewUserPictureChangedEvent(userEntity, userPicture)
}

func (*userPictureEventFactoryImpl) CreateUserPictureRemovedEvent(userEntity *entities.UserEntity) contracts.UserPictureEvent {
	return events.NewUserPictureRemovedEvent(userEntity)
}
