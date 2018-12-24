package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserEventFactory(t *testing.T) {
	t.Run("NewUserEventFactory", func(t *testing.T) {
		_, isUserEventFactory := NewUserEventFactory().(*userEventFactoryImpl)

		assert.True(t, isUserEventFactory)
	})

	t.Run("CreateUserCreatedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEventFactory := new(userEventFactoryImpl)
		userEvent := userEventFactory.CreateUserCreatedEvent(userEntity)

		assert.Equal(t, events.UserCreatedEventName, userEvent.Name())
	})

	t.Run("CreateUserUpdatedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEventFactory := new(userEventFactoryImpl)
		userEvent := userEventFactory.CreateUserUpdatedEvent(userEntity)

		assert.Equal(t, events.UserUpdatedEventName, userEvent.Name())
	})

	t.Run("CreateUserVerifiedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEventFactory := new(userEventFactoryImpl)
		userEvent := userEventFactory.CreateUserVerifiedEvent(userEntity)

		assert.Equal(t, events.UserVerifiedEventName, userEvent.Name())
	})

	t.Run("CreateUserIdentityChangedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEventFactory := new(userEventFactoryImpl)
		userEvent := userEventFactory.CreateUserIdentityChangedEvent(userEntity)

		assert.Equal(t, events.UserIdentityChangedEventName, userEvent.Name())
	})

	t.Run("CreateUserPasswordChangedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEventFactory := new(userEventFactoryImpl)
		userEvent := userEventFactory.CreateUserPasswordChangedEvent(userEntity)

		assert.Equal(t, events.UserPasswordChangedEventName, userEvent.Name())
	})

	t.Run("CreateUserDeletedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEventFactory := new(userEventFactoryImpl)
		userEvent := userEventFactory.CreateUserDeletedEvent(userEntity)

		assert.Equal(t, events.UserDeletedEventName, userEvent.Name())
	})
}
