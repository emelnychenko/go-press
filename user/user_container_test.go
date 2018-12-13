package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContainer(t *testing.T) {
	container := NewUserContainer(nil)
	assert.NotNil(t, container.UserRepository)
	assert.NotNil(t, container.UserService)
	assert.NotNil(t, container.UserApi)
}