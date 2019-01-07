package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PollApi interface {
		ListPolls(pollPaginationQuery *models.PollPaginationQuery) (*models.PaginationResult, errors.Error)
		GetPoll(pollId *models.PollId) (poll *models.Poll, err errors.Error)
		CreatePoll(data *models.PollCreate) (poll *models.Poll, err errors.Error)
		UpdatePoll(pollId *models.PollId, data *models.PollUpdate) (err errors.Error)
		DeletePoll(pollId *models.PollId) (err errors.Error)
	}
)
