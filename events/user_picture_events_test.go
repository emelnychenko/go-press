package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserPictureEvents(t *testing.T) {
	t.Run("NewUserPictureChangedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userPicture := new(entities.FileEntity)

		userPictureEvent, isUserPictureEvent := NewUserPictureChangedEvent(userEntity, userPicture).(*UserPictureEvent)

		assert.True(t, isUserPictureEvent)
		assert.Equal(t, userEntity, userPictureEvent.userEntity)
		assert.Equal(t, userPicture, userPictureEvent.userPicture)
		assert.Equal(t, UserPictureChangedEventName, userPictureEvent.name)
	})

	t.Run("NewUserPictureRemovedEvent", func(t *testing.T) {
		userEntity := new(entities.UserEntity)

		userPictureEvent, isUserPictureEvent := NewUserPictureRemovedEvent(userEntity).(*UserPictureEvent)

		assert.True(t, isUserPictureEvent)
		assert.Equal(t, userEntity, userPictureEvent.userEntity)
		assert.Nil(t, userPictureEvent.userPicture)
		assert.Equal(t, UserPictureRemovedEventName, userPictureEvent.name)
	})

	t.Run("UserPictureEntity", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userPicture := new(entities.FileEntity)

		userPictureEvent := &UserPictureEvent{userEntity: userEntity, userPicture: userPicture}

		assert.Equal(t, userEntity, userPictureEvent.UserEntity())
		assert.Equal(t, userPicture, userPictureEvent.UserPictureEntity())
	})
}
