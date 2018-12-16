package entities

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPostEntity(t *testing.T) {
	t.Run("NewPostEntity", func(t *testing.T) {
		e := NewPostEntity()
		assert.IsType(t, new(models.PostId), e.Id)
		assert.IsType(t, new(time.Time), e.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		e := NewPostEntity()
		assert.Equal(t, PostTable, e.TableName())
	})

	t.Run("SetAuthor", func(t *testing.T) {
		postEntity := NewPostEntity()
		userEntity := NewUserEntity()
		postEntity.SetAuthor(userEntity)
		assert.Equal(t, userEntity.SubjectId(), postEntity.AuthorId)
		assert.Equal(t, userEntity.SubjectType(), postEntity.AuthorType)
	})
}
