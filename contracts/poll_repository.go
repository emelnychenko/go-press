package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PollRepository interface {
		ListPolls(pollPaginationQuery *models.PollPaginationQuery) (*models.PaginationResult, errors.Error)
		GetPoll(pollId *models.PollId) (pollEntity *entities.PollEntity, err errors.Error)
		SavePoll(pollEntity *entities.PollEntity) (err errors.Error)
		RemovePoll(pollEntity *entities.PollEntity) (err errors.Error)
	}
)
