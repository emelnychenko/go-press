package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	tagModelFactoryImpl struct {
		paginationModelFactory contracts.PaginationModelFactory
	}
)

func NewTagModelFactory(paginationModelFactory contracts.PaginationModelFactory) contracts.TagModelFactory {
	return &tagModelFactoryImpl{paginationModelFactory}
}

func (f *tagModelFactoryImpl) CreateTagPaginationQuery() *models.TagPaginationQuery {
	paginationQuery := f.paginationModelFactory.CreatePaginationQuery()
	return &models.TagPaginationQuery{PaginationQuery: paginationQuery}
}

func (*tagModelFactoryImpl) CreateTag() *models.Tag {
	return new(models.Tag)
}

func (*tagModelFactoryImpl) CreateTagCreate() *models.TagCreate {
	return new(models.TagCreate)
}

func (*tagModelFactoryImpl) CreateTagUpdate() *models.TagUpdate {
	return new(models.TagUpdate)
}
