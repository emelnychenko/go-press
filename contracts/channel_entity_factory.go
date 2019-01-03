package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	ChannelEntityFactory interface {
		CreateChannelEntity() *entities.ChannelEntity
	}
)
