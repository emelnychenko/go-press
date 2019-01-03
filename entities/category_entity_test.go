package entities

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryEntity(t *testing.T) {
	t.Run("NewCategoryEntity", func(t *testing.T) {
		categoryEntity := NewCategoryEntity()

		assert.IsType(t, new(models.CategoryId), categoryEntity.Id)
		assert.NotNil(t, categoryEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		categoryEntity := new(CategoryEntity)

		assert.Equal(t, CategoryTableName, categoryEntity.TableName())
	})
}
