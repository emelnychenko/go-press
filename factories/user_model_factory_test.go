package factories

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserModelFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewUserModelFactory", func(t *testing.T) {
		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		userModelFactory, isUserModelFactory := NewUserModelFactory(paginationModelFactory).(*userModelFactoryImpl)

		assert.True(t, isUserModelFactory)
		assert.Equal(t, paginationModelFactory, userModelFactory.paginationModelFactory)
	})

	t.Run("CreateUserPaginationQuery", func(t *testing.T) {
		paginationQuery := new(models.PaginationQuery)

		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		paginationModelFactory.EXPECT().CreatePaginationQuery().Return(paginationQuery)

		userModelFactory := &userModelFactoryImpl{paginationModelFactory: paginationModelFactory}
		userPaginationQuery := userModelFactory.CreateUserPaginationQuery()

		assert.Equal(t, paginationQuery, userPaginationQuery.PaginationQuery)
	})

	t.Run("CreateUser", func(t *testing.T) {
		userModelFactory := new(userModelFactoryImpl)
		assert.NotNil(t, userModelFactory.CreateUser())
	})

	t.Run("CreateUserCreate", func(t *testing.T) {
		userModelFactory := new(userModelFactoryImpl)
		assert.NotNil(t, userModelFactory.CreateUserCreate())
	})

	t.Run("CreateUserUpdate", func(t *testing.T) {
		userModelFactory := new(userModelFactoryImpl)
		assert.NotNil(t, userModelFactory.CreateUserUpdate())
	})

	t.Run("CreateUserChangeIdentity", func(t *testing.T) {
		userModelFactory := new(userModelFactoryImpl)
		assert.NotNil(t, userModelFactory.CreateUserChangeIdentity())
	})

	t.Run("CreateUserChangePassword", func(t *testing.T) {
		userModelFactory := new(userModelFactoryImpl)
		assert.NotNil(t, userModelFactory.CreateUserChangePassword())
	})
}
