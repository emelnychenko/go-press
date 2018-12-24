package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	userEventFactoryImpl struct {
	}
)

func NewUserEventFactory() contracts.UserEventFactory {
	return new(userEventFactoryImpl)
}

func (*userEventFactoryImpl) CreateUserCreatedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	return events.NewUserCreatedEvent(userEntity)
}

func (*userEventFactoryImpl) CreateUserUpdatedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	return events.NewUserUpdatedEvent(userEntity)
}

func (*userEventFactoryImpl) CreateUserVerifiedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	return events.NewUserVerifiedEvent(userEntity)
}

func (*userEventFactoryImpl) CreateUserIdentityChangedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	return events.NewUserIdentityChangedEvent(userEntity)
}

func (*userEventFactoryImpl) CreateUserPasswordChangedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	return events.NewUserPasswordChangedEvent(userEntity)
}

func (*userEventFactoryImpl) CreateUserDeletedEvent(userEntity *entities.UserEntity) contracts.UserEvent {
	return events.NewUserDeletedEvent(userEntity)
}
