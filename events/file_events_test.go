package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileEvents(t *testing.T) {
	t.Run("NewFileUploadedEvent", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		fileEvent, isFileEvent := NewFileUploadedEvent(fileEntity).(*FileEvent)

		assert.True(t, isFileEvent)
		assert.Equal(t, fileEntity, fileEvent.fileEntity)
		assert.Equal(t, FileUploadedEventName, fileEvent.name)
	})

	t.Run("NewFileUpdatedEvent", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		fileEvent, isFileEvent := NewFileUpdatedEvent(fileEntity).(*FileEvent)

		assert.True(t, isFileEvent)
		assert.Equal(t, fileEntity, fileEvent.fileEntity)
		assert.Equal(t, FileUpdatedEventName, fileEvent.name)
	})

	t.Run("NewFileDeletedEvent", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		fileEvent, isFileEvent := NewFileDeletedEvent(fileEntity).(*FileEvent)

		assert.True(t, isFileEvent)
		assert.Equal(t, fileEntity, fileEvent.fileEntity)
		assert.Equal(t, FileDeletedEventName, fileEvent.name)
	})

	t.Run("FileEntity", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		fileEvent := &FileEvent{fileEntity: fileEntity}

		assert.Equal(t, fileEntity, fileEvent.FileEntity())
	})
}
