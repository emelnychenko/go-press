package aggregator

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserAggregator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewUserAggregator", func(t *testing.T) {
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		fileApi := mocks.NewMockFileApi(ctrl)
		userAggregator, isUserAggregator := NewUserAggregator(userModelFactory, fileApi).(*userAggregatorImpl)

		assert.True(t, isUserAggregator)
		assert.Equal(t, userModelFactory, userAggregator.userModelFactory)
		assert.Equal(t, fileApi, userAggregator.fileApi)
	})

	t.Run("AggregateUser", func(t *testing.T) {
		fileApi := mocks.NewMockFileApi(ctrl)
		userPictureId := new(models.FileId)
		userPicture := new(models.File)
		fileApi.EXPECT().GetFile(userPictureId).Return(userPicture, nil)
		
		user := new(models.User)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUser().Return(user)

		userEntity := &entities.UserEntity{PictureId: userPictureId}
		userAggregator := &userAggregatorImpl{userModelFactory: userModelFactory, fileApi: fileApi}
		response := userAggregator.AggregateUser(userEntity)

		assert.Equal(t, user, response)
		assert.Equal(t, userPicture, response.Picture)
	})

	t.Run("AggregateUsers", func(t *testing.T) {
		user := new(models.User)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUser().Return(user)

		userAggregator := &userAggregatorImpl{userModelFactory: userModelFactory}
		userEntities := []*entities.UserEntity{new(entities.UserEntity)}
		response := userAggregator.AggregateUsers(userEntities)

		assert.IsType(t, []*models.User{}, response)
		assert.Equal(t, len(userEntities), len(response))
	})
}
