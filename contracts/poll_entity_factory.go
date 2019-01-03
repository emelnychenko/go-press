package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	PollEntityFactory interface {
		CreatePollEntity() *entities.PollEntity
	}
)
