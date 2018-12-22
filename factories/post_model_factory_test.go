package factories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostModelFactory(t *testing.T) {
	t.Run("NewPostModelFactory", func(t *testing.T) {
		_, isPostModelFactory := NewPostModelFactory().(*postModelFactoryImpl)

		assert.True(t, isPostModelFactory)
	})

	t.Run("CreatePost", func(t *testing.T) {
		postModelFactory := new(postModelFactoryImpl)
		assert.NotNil(t, postModelFactory.CreatePost())
	})

	t.Run("CreatePostCreate", func(t *testing.T) {
		postModelFactory := new(postModelFactoryImpl)
		assert.NotNil(t, postModelFactory.CreatePostCreate())
	})

	t.Run("CreatePostUpdate", func(t *testing.T) {
		postModelFactory := new(postModelFactoryImpl)
		assert.NotNil(t, postModelFactory.CreatePostUpdate())
	})
}
