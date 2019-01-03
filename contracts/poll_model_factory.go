package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	PollModelFactory interface {
		CreatePollPaginationQuery() *models.PollPaginationQuery
		CreatePoll() *models.Poll
		CreatePollCreate() *models.PollCreate
		CreatePollUpdate() *models.PollUpdate
	}
)
