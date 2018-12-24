package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	UserEventFactory interface {
		CreateUserCreatedEvent(userEntity *entities.UserEntity) UserEvent
		CreateUserUpdatedEvent(userEntity *entities.UserEntity) UserEvent
		CreateUserVerifiedEvent(userEntity *entities.UserEntity) UserEvent
		CreateUserIdentityChangedEvent(userEntity *entities.UserEntity) UserEvent
		CreateUserPasswordChangedEvent(userEntity *entities.UserEntity) UserEvent
		CreateUserDeletedEvent(userEntity *entities.UserEntity) UserEvent
	}
)
