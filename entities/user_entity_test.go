package entities

import (
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserEntity(t *testing.T) {
	t.Run("NewUserEntity", func(t *testing.T) {
		e := NewUserEntity()
		assert.IsType(t, new(UserEntity), e)
		assert.IsType(t, new(models.UserId), e.Id)
	})

	t.Run("TableName", func(t *testing.T) {
		e := NewUserEntity()
		assert.Equal(t, UserTable, e.TableName())
	})

	t.Run("SubjectId", func(t *testing.T) {
		e := NewUserEntity()
		assert.Equal(t, e.Id, e.SubjectId())
	})

	t.Run("SubjectType", func(t *testing.T) {
		e := NewUserEntity()
		assert.Equal(t, enums.UserSubjectType, e.SubjectType())
	})
}
