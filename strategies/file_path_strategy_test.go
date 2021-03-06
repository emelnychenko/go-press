package strategies

import (
	"fmt"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"mime"
	"testing"
	"time"
)

func TestFilePathStrategy(t *testing.T) {
	t.Run("NewFilePathStrategy", func(t *testing.T) {
		_, isFilePathStrategy := NewFilePathStrategy().(*filePathStrategyImpl)

		assert.True(t, isFilePathStrategy)
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
		fileExtensions, _ := mime.ExtensionsByType(fileType)
		filePath := fmt.Sprintf("uploads/%d/%d/%s%s", fileCreated.Year(), int(fileCreated.Month()), fileId, fileExtensions[0])
		filePathStrategy := &filePathStrategyImpl{}

		response, err := filePathStrategy.BuildPath(fileEntity)

		assert.Equal(t, filePath, response)
		assert.Nil(t, err)
	})
}
