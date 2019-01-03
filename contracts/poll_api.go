package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	PollApi interface {
		ListPolls(pollPaginationQuery *models.PollPaginationQuery) (*models.PaginationResult, common.Error)
		GetPoll(pollId *models.PollId) (poll *models.Poll, err common.Error)
		CreatePoll(data *models.PollCreate) (poll *models.Poll, err common.Error)
		UpdatePoll(pollId *models.PollId, data *models.PollUpdate) (err common.Error)
		DeletePoll(pollId *models.PollId) (err common.Error)
	}
)
