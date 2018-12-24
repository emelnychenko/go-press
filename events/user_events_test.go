package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserEvents(t *testing.T) {
	t.Run("NewUserCreatedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEvent, isUserEvent := NewUserCreatedEvent(userEntity).(*UserEvent)

		assert.True(t, isUserEvent)
		assert.Equal(t, userEntity, userEvent.userEntity)
		assert.Equal(t, UserCreatedEventName, userEvent.name)
	})

	t.Run("NewUserUpdatedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEvent, isUserEvent := NewUserUpdatedEvent(userEntity).(*UserEvent)

		assert.True(t, isUserEvent)
		assert.Equal(t, userEntity, userEvent.userEntity)
		assert.Equal(t, UserUpdatedEventName, userEvent.name)
	})

	t.Run("NewUserVerifiedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEvent, isUserEvent := NewUserVerifiedEvent(userEntity).(*UserEvent)

		assert.True(t, isUserEvent)
		assert.Equal(t, userEntity, userEvent.userEntity)
		assert.Equal(t, UserVerifiedEventName, userEvent.name)
	})

	t.Run("NewUserIdentityChangedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEvent, isUserEvent := NewUserIdentityChangedEvent(userEntity).(*UserEvent)

		assert.True(t, isUserEvent)
		assert.Equal(t, userEntity, userEvent.userEntity)
		assert.Equal(t, UserIdentityChangedEventName, userEvent.name)
	})

	t.Run("NewUserPasswordChangedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEvent, isUserEvent := NewUserPasswordChangedEvent(userEntity).(*UserEvent)

		assert.True(t, isUserEvent)
		assert.Equal(t, userEntity, userEvent.userEntity)
		assert.Equal(t, UserPasswordChangedEventName, userEvent.name)
	})

	t.Run("NewUserDeletedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEvent, isUserEvent := NewUserDeletedEvent(userEntity).(*UserEvent)

		assert.True(t, isUserEvent)
		assert.Equal(t, userEntity, userEvent.userEntity)
		assert.Equal(t, UserDeletedEventName, userEvent.name)
	})

	t.Run("UserEntity", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEvent := &UserEvent{userEntity: userEntity}

		assert.Equal(t, userEntity, userEvent.UserEntity())
	})
}
