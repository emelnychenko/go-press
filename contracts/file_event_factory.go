package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	FileEventFactory interface {
		CreateFileUploadedEvent(fileEntity *entities.FileEntity) FileEvent
		CreateFileUpdatedEvent(fileEntity *entities.FileEntity) FileEvent
		CreateFileDeletedEvent(fileEntity *entities.FileEntity) FileEvent
	}
)
