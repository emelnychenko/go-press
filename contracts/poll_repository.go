package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	PollRepository interface {
		ListPolls(pollPaginationQuery *models.PollPaginationQuery) (*models.PaginationResult, common.Error)
		GetPoll(pollId *models.PollId) (pollEntity *entities.PollEntity, err common.Error)
		SavePoll(pollEntity *entities.PollEntity) (err common.Error)
		RemovePoll(pollEntity *entities.PollEntity) (err common.Error)
	}
)
