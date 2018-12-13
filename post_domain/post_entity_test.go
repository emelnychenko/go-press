package post_domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostEntity(t *testing.T) {
	t.Run("NewPostEntity", func(t *testing.T) {
		e := NewPostEntity()
		assert.IsType(t, new(PostId), e.Id)
	})

	t.Run("TableName", func(t *testing.T) {
		e := NewPostEntity()
		assert.Equal(t, "posts", e.TableName())
	})
}