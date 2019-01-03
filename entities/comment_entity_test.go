package entities

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentEntity(t *testing.T) {
	t.Run("NewCommentEntity", func(t *testing.T) {
		commentEntity := NewCommentEntity()

		assert.IsType(t, new(models.CommentId), commentEntity.Id)
		assert.NotNil(t, commentEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		commentEntity := new(CommentEntity)

		assert.Equal(t, CommentTableName, commentEntity.TableName())
	})
}
