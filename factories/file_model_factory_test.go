package factories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileModelsFactory(t *testing.T) {
	t.Run("NewFileModelFactory", func(t *testing.T) {
		_, isFileModelsFactory := NewFileModelFactory().(*fileModelFactoryImpl)

		assert.True(t, isFileModelsFactory)
	})

	t.Run("CreateFile", func(t *testing.T) {
		fileModelFactory := &fileModelFactoryImpl{}
		assert.NotNil(t, fileModelFactory.CreateFile())
	})

	t.Run("CreateFileUpload", func(t *testing.T) {
		fileModelFactory := &fileModelFactoryImpl{}
		assert.NotNil(t, fileModelFactory.CreateFileUpload())
	})

	t.Run("CreateFileUpdate", func(t *testing.T) {
		fileModelFactory := &fileModelFactoryImpl{}
		assert.NotNil(t, fileModelFactory.CreateFileUpdate())
	})
}
