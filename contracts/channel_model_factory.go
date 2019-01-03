package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	ChannelModelFactory interface {
		CreateChannelPaginationQuery() *models.ChannelPaginationQuery
		CreateChannel() *models.Channel
		CreateChannelCreate() *models.ChannelCreate
		CreateChannelUpdate() *models.ChannelUpdate
	}
)
