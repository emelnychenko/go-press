package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	FileUploadedEventName = "FileUploadedEvent"
	FileUpdatedEventName  = "FileUpdatedEvent"
	FileDeletedEventName  = "FileDeletedEvent"
)

type (
	FileEvent struct {
		*Event
		fileEntity *entities.FileEntity
	}
)

func (e *FileEvent) FileEntity() *entities.FileEntity {
	return e.fileEntity
}

func NewFileUploadedEvent(fileEntity *entities.FileEntity) contracts.FileEvent {
	event := &Event{name: FileUploadedEventName}
	return &FileEvent{fileEntity: fileEntity, Event: event}
}

func NewFileUpdatedEvent(fileEntity *entities.FileEntity) contracts.FileEvent {
	event := &Event{name: FileUpdatedEventName}
	return &FileEvent{fileEntity: fileEntity, Event: event}
}

func NewFileDeletedEvent(fileEntity *entities.FileEntity) contracts.FileEvent {
	event := &Event{name: FileDeletedEventName}
	return &FileEvent{fileEntity: fileEntity, Event: event}
}
