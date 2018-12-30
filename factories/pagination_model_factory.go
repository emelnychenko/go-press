package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

const (
	paginationQueryDefaultLimit = 10
)

type (
	paginationModelFactoryImpl struct {
	}
)

func NewPaginationModelFactory() contracts.PaginationModelFactory {
	return new(paginationModelFactoryImpl)
}

func (*paginationModelFactoryImpl) CreatePaginationQuery() *models.PaginationQuery {
	return &models.PaginationQuery{Limit: paginationQueryDefaultLimit, Start: 0, Page: 1}
}

func (*paginationModelFactoryImpl) CreatePaginationResult() *models.PaginationResult {
	return new(models.PaginationResult)
}
