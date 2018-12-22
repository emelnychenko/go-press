package strategies

import (
	"fmt"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFilePathStrategy(t *testing.T) {
	t.Run("NewFilePathStrategy", func(t *testing.T) {
		_, isFilePathStrategy := NewFilePathStrategy().(*filePathStrategyImpl)

		assert.True(t, isFilePathStrategy)
	})

	t.Run("BuildPath:NoIdError", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		filePathStrategy := &filePathStrategyImpl{}

		response, err := filePathStrategy.BuildPath(fileEntity)

		assert.Empty(t, response)
		assert.Error(t, err)
	})

	t.Run("BuildPath:CreatedError", func(t *testing.T) {
		fileId := new(models.FileId)
		fileEntity := &entities.FileEntity{Id: fileId}
		filePathStrategy := &filePathStrategyImpl{}

		response, err := filePathStrategy.BuildPath(fileEntity)

		assert.Empty(t, response)
		assert.Error(t, err)
	})

	t.Run("BuildPath:MimeError", func(t *testing.T) {
		fileId := new(models.FileId)
		fileCreated := time.Now()
		fileEntity := &entities.FileEntity{Id: fileId, Created: &fileCreated}
		filePathStrategy := &filePathStrategyImpl{}

		response, err := filePathStrategy.BuildPath(fileEntity)

		assert.Empty(t, response)
		assert.Error(t, err)
	})

	t.Run("BuildPath", func(t *testing.T) {
		fileId := new(models.FileId)
		fileCreated := time.Now()
		fileType := "text/plain"
		fileEntity := &entities.FileEntity{Id: fileId, Type: fileType, Created: &fileCreated}
		filePath := fmt.Sprintf("uploads/%d/%d/%s.txt", fileCreated.Year(), int(fileCreated.Month()), fileId)
		filePathStrategy := &filePathStrategyImpl{}

		response, err := filePathStrategy.BuildPath(fileEntity)

		assert.Equal(t, filePath, response)
		assert.Nil(t, err)
	})
}
