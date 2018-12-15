package models

import (
	"github.com/emelnychenko/go-press/enums"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("SubjectId", func(t *testing.T) {
		user := User{}
		assert.Equal(t, user.Id, user.SubjectId())
	})

	t.Run("SubjectType", func(t *testing.T) {
		user := User{}
		assert.Equal(t, enums.UserSubjectType, user.SubjectType())
	})
}