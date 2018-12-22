package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostEntityFactory(t *testing.T) {
	t.Run("NewPostEntityFactory", func(t *testing.T) {
		_, isPostEntityFactory := NewPostEntityFactory().(*postEntityFactoryImpl)

		assert.True(t, isPostEntityFactory)
	})

	t.Run("CreatePostEntity", func(t *testing.T) {
		postEntityFactory := new(postEntityFactoryImpl)
		assert.IsType(t, new(entities.PostEntity), postEntityFactory.CreatePostEntity())
	})
}
