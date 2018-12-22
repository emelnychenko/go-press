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
		userAggregator, isUserAggregator := NewUserAggregator(userModelFactory).(*userAggregatorImpl)

		assert.True(t, isUserAggregator)
		assert.Equal(t, userModelFactory, userAggregator.userModelFactory)
	})

	t.Run("AggregateUser", func(t *testing.T) {
		user := new(models.User)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUser().Return(user)

		userAggregator := &userAggregatorImpl{userModelFactory: userModelFactory}
		response := userAggregator.AggregateUser(new(entities.UserEntity))

		assert.Equal(t, user, response)
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
