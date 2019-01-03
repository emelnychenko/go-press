package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	ChannelEvent interface {
		Event
		ChannelEntity() *entities.ChannelEntity
	}
)
