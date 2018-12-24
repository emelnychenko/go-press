package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserPictureService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewUserPictureService", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		userPictureService, isUserPictureService := NewUserPictureService(userRepository).(*userPictureServiceImpl)

		assert.True(t, isUserPictureService)
		assert.Equal(t, userRepository, userPictureService.userRepository)
	})

	t.Run("ChangeUserPicture", func(t *testing.T) {
		userPictureId := new(models.FileId)
		userPictureEntity := &entities.FileEntity{Id: userPictureId}
		userEntity := new(entities.UserEntity)

		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().SaveUser(userEntity).Return(nil)

		userPictureService := &userPictureServiceImpl{userRepository: userRepository}

		assert.Nil(t, userPictureService.ChangeUserPicture(userEntity, userPictureEntity))
		assert.Equal(t, userPictureId, userEntity.PictureId)
	})

	t.Run("RemoveUserPicture", func(t *testing.T) {
		userPictureId := new(models.FileId)
		userEntity := &entities.UserEntity{PictureId: userPictureId}

		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().SaveUser(userEntity).Return(nil)

		userPictureService := &userPictureServiceImpl{userRepository: userRepository}

		assert.Nil(t, userPictureService.RemoveUserPicture(userEntity))
		assert.Nil(t, userEntity.PictureId)
	})
}
