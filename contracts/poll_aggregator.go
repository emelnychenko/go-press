package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	PollAggregator interface {
		AggregatePoll(pollEntity *entities.PollEntity) *models.Poll
		AggregatePolls(pollEntities []*entities.PollEntity) []*models.Poll
		AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult
	}
)
