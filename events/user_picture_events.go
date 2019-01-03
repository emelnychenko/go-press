package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	UserPictureChangedEventName = "UserPictureChangedEvent"
	UserPictureRemovedEventName = "UserPictureRemovedEvent"
)

type (
	UserPictureEvent struct {
		*Event
		userEntity  *entities.UserEntity
		userPicture *entities.FileEntity
	}
)

func (e *UserPictureEvent) UserEntity() *entities.UserEntity {
	return e.userEntity
}

func (e *UserPictureEvent) UserPictureEntity() *entities.FileEntity {
	return e.userPicture
}

func NewUserPictureChangedEvent(
	userEntity *entities.UserEntity,
	userPicture *entities.FileEntity,
) contracts.UserPictureEvent {
	event := &Event{name: UserPictureChangedEventName}
	return &UserPictureEvent{userEntity: userEntity, userPicture: userPicture, Event: event}
}

func NewUserPictureRemovedEvent(userEntity *entities.UserEntity) contracts.UserPictureEvent {
	event := &Event{name: UserPictureRemovedEventName}
	return &UserPictureEvent{userEntity: userEntity, userPicture: nil, Event: event}
}
