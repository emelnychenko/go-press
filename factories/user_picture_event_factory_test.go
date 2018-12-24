package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserPictureEventFactory(t *testing.T) {
	t.Run("NewUserPictureEventFactory", func(t *testing.T) {
		_, isUserPictureEventFactory := NewUserPictureEventFactory().(*userPictureEventFactoryImpl)

		assert.True(t, isUserPictureEventFactory)
	})

	t.Run("CreateUserPictureChangedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userPictureEntity := new(entities.FileEntity)

		userPictureEventFactory := new(userPictureEventFactoryImpl)
		userPictureEvent := userPictureEventFactory.CreateUserPictureChangedEvent(userEntity, userPictureEntity)

		assert.Equal(t, events.UserPictureChangedEventName, userPictureEvent.Name())
		assert.Equal(t, userEntity, userPictureEvent.UserEntity())
		assert.Equal(t, userPictureEntity, userPictureEvent.UserPictureEntity())
	})

	t.Run("CreateUserCreatedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userPictureEventFactory := new(userPictureEventFactoryImpl)
		userPictureEvent := userPictureEventFactory.CreateUserPictureRemovedEvent(userEntity)

		assert.Equal(t, events.UserPictureRemovedEventName, userPictureEvent.Name())
		assert.Equal(t, userEntity, userPictureEvent.UserEntity())
	})
}
