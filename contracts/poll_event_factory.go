package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	PollEventFactory interface {
		CreatePollCreatedEvent(pollEntity *entities.PollEntity) PollEvent
		CreatePollUpdatedEvent(pollEntity *entities.PollEntity) PollEvent
		CreatePollDeletedEvent(pollEntity *entities.PollEntity) PollEvent
	}
)
