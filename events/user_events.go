package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	UserCreatedEventName = "UserCreatedEvent"
	UserUpdatedEventName = "UserUpdatedEvent"
	UserVerifiedEventName = "UserVerifiedEvent"
	UserIdentityChangedEventName = "UserIdentityChangedEvent"
	UserPasswordChangedEventName = "UserPasswordChangedEvent"
	UserDeletedEventName = "UserDeletedEvent"
)

type (
	UserEvent struct {
		*Event
		userEntity *entities.UserEntity
	}
)

func (e *UserEvent) UserEntity() *entities.UserEntity {
	return e.userEntity
}

func NewUserCreatedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	event := &Event{name: UserCreatedEventName}
	return &UserEvent{userEntity: userEntity, Event: event}
}

func NewUserUpdatedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	event := &Event{name: UserUpdatedEventName}
	return &UserEvent{userEntity: userEntity, Event: event}
}

func NewUserVerifiedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	event := &Event{name: UserVerifiedEventName}
	return &UserEvent{userEntity: userEntity, Event: event}
}

func NewUserIdentityChangedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	event := &Event{name: UserIdentityChangedEventName}
	return &UserEvent{userEntity: userEntity, Event: event}
}

func NewUserPasswordChangedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	event := &Event{name: UserPasswordChangedEventName}
	return &UserEvent{userEntity: userEntity, Event: event}
}

func NewUserDeletedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	event := &Event{name: UserDeletedEventName}
	return &UserEvent{userEntity: userEntity, Event: event}
}
