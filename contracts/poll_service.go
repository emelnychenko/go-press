package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PollService interface {
		ListPolls(pollPaginationQuery *models.PollPaginationQuery) (*models.PaginationResult, errors.Error)
		GetPoll(pollId *models.PollId) (pollEntity *entities.PollEntity, err errors.Error)
		CreatePoll(data *models.PollCreate) (pollEntity *entities.PollEntity, err errors.Error)
		UpdatePoll(pollEntity *entities.PollEntity, data *models.PollUpdate) (err errors.Error)
		DeletePoll(pollEntity *entities.PollEntity) (err errors.Error)
	}
)
