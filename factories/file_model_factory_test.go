package factories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileModelFactory(t *testing.T) {
	t.Run("NewFileModelFactory", func(t *testing.T) {
		_, isFileModelFactory := NewFileModelFactory().(*fileModelFactoryImpl)

		assert.True(t, isFileModelFactory)
	})

	t.Run("CreateFile", func(t *testing.T) {
		fileModelFactory := new(fileModelFactoryImpl)
		assert.NotNil(t, fileModelFactory.CreateFile())
	})

	t.Run("CreateFileUpload", func(t *testing.T) {
		fileModelFactory := new(fileModelFactoryImpl)
		assert.NotNil(t, fileModelFactory.CreateFileUpload())
	})

	t.Run("CreateFileUpdate", func(t *testing.T) {
		fileModelFactory := new(fileModelFactoryImpl)
		assert.NotNil(t, fileModelFactory.CreateFileUpdate())
	})
}
