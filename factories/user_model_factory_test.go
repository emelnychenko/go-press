package factories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserModelFactory(t *testing.T) {
	t.Run("NewUserModelFactory", func(t *testing.T) {
		_, isUserModelFactory := NewUserModelFactory().(*userModelFactoryImpl)

		assert.True(t, isUserModelFactory)
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
