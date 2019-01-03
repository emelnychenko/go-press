package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	PollService interface {
		ListPolls(pollPaginationQuery *models.PollPaginationQuery) (*models.PaginationResult, common.Error)
		GetPoll(pollId *models.PollId) (pollEntity *entities.PollEntity, err common.Error)
		CreatePoll(data *models.PollCreate) (pollEntity *entities.PollEntity, err common.Error)
		UpdatePoll(pollEntity *entities.PollEntity, data *models.PollUpdate) (err common.Error)
		DeletePoll(pollEntity *entities.PollEntity) (err common.Error)
	}
)
