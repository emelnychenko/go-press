package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentEntityFactory(t *testing.T) {
	t.Run("NewCommentEntityFactory", func(t *testing.T) {
		_, isCommentEntityFactory := NewCommentEntityFactory().(*commentEntityFactoryImpl)

		assert.True(t, isCommentEntityFactory)
	})

	t.Run("CreateCommentEntity", func(t *testing.T) {
		commentEntityFactory := new(commentEntityFactoryImpl)
		assert.IsType(t, new(entities.CommentEntity), commentEntityFactory.CreateCommentEntity())
	})
}
