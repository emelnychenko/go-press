package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserEntityFactory(t *testing.T) {
	t.Run("NewUserEntityFactory", func(t *testing.T) {
		_, isUserEntityFactory := NewUserEntityFactory().(*userEntityFactoryImpl)

		assert.True(t, isUserEntityFactory)
	})

	t.Run("CreateUserEntity", func(t *testing.T) {
		userEntityFactory := new(userEntityFactoryImpl)
		assert.IsType(t, new(entities.UserEntity), userEntityFactory.CreateUserEntity())
	})
}
