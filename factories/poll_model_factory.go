package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	pollModelFactoryImpl struct {
		paginationModelFactory contracts.PaginationModelFactory
	}
)

func NewPollModelFactory(paginationModelFactory contracts.PaginationModelFactory) contracts.PollModelFactory {
	return &pollModelFactoryImpl{paginationModelFactory}
}

func (f *pollModelFactoryImpl) CreatePollPaginationQuery() *models.PollPaginationQuery {
	paginationQuery := f.paginationModelFactory.CreatePaginationQuery()
	return &models.PollPaginationQuery{PaginationQuery: paginationQuery}
}

func (*pollModelFactoryImpl) CreatePoll() *models.Poll {
	return new(models.Poll)
}

func (*pollModelFactoryImpl) CreatePollCreate() *models.PollCreate {
	return new(models.PollCreate)
}

func (*pollModelFactoryImpl) CreatePollUpdate() *models.PollUpdate {
	return new(models.PollUpdate)
}
