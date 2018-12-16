package entities

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFileEntity(t *testing.T) {
	t.Run("NewFileEntity", func(t *testing.T) {
		fileEntity := NewFileEntity()
		assert.IsType(t, new(models.FileId), fileEntity.Id)
		assert.IsType(t, new(time.Time), fileEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		e := &FileEntity{}
		assert.Equal(t, FileTable, e.TableName())
	})
}
