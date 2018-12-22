package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileEntityFactory(t *testing.T) {
	t.Run("NewFileEntityFactory", func(t *testing.T) {
		_, isFileEntityFactory := NewFileEntityFactory().(*fileEntityFactoryImpl)

		assert.True(t, isFileEntityFactory)
	})

	t.Run("CreateFileEntity", func(t *testing.T) {
		fileEntityFactory := new(fileEntityFactoryImpl)
		assert.IsType(t, new(entities.FileEntity), fileEntityFactory.CreateFileEntity())
	})
}
