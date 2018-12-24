package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileEventFactory(t *testing.T) {
	t.Run("NewFileEventFactory", func(t *testing.T) {
		_, isFileEventFactory := NewFileEventFactory().(*fileEventFactoryImpl)

		assert.True(t, isFileEventFactory)
	})

	t.Run("CreateFileUploadedEvent", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		fileEventFactory := new(fileEventFactoryImpl)
		fileEvent := fileEventFactory.CreateFileUploadedEvent(fileEntity)

		assert.Equal(t, events.FileUploadedEventName, fileEvent.Name())
	})

	t.Run("CreateFileUpdatedEvent", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		fileEventFactory := new(fileEventFactoryImpl)
		fileEvent := fileEventFactory.CreateFileUpdatedEvent(fileEntity)

		assert.Equal(t, events.FileUpdatedEventName, fileEvent.Name())
	})

	t.Run("CreateFileDeletedEvent", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		fileEventFactory := new(fileEventFactoryImpl)
		fileEvent := fileEventFactory.CreateFileDeletedEvent(fileEntity)

		assert.Equal(t, events.FileDeletedEventName, fileEvent.Name())
	})
}
