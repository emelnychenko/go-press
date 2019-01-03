package entities

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileEntity(t *testing.T) {
	t.Run("NewFileEntity", func(t *testing.T) {
		fileEntity := NewFileEntity()

		assert.IsType(t, new(models.FileId), fileEntity.Id)
		assert.NotNil(t, fileEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		fileEntity := new(FileEntity)

		assert.Equal(t, FileTableName, fileEntity.TableName())
	})
}
