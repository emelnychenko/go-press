package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	fileEventFactoryImpl struct {
	}
)

func NewFileEventFactory() contracts.FileEventFactory {
	return new(fileEventFactoryImpl)
}

func (*fileEventFactoryImpl) CreateFileUploadedEvent(fileEntity *entities.FileEntity) contracts.FileEvent {
	return events.NewFileUploadedEvent(fileEntity)
}

func (*fileEventFactoryImpl) CreateFileUpdatedEvent(fileEntity *entities.FileEntity) contracts.FileEvent {
	return events.NewFileUpdatedEvent(fileEntity)
}

func (*fileEventFactoryImpl) CreateFileDeletedEvent(fileEntity *entities.FileEntity) contracts.FileEvent {
	return events.NewFileDeletedEvent(fileEntity)
}
