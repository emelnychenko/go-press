package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	categoryModelFactoryImpl struct {
		paginationModelFactory contracts.PaginationModelFactory
	}
)

func NewCategoryModelFactory(paginationModelFactory contracts.PaginationModelFactory) contracts.CategoryModelFactory {
	return &categoryModelFactoryImpl{paginationModelFactory}
}

func (f *categoryModelFactoryImpl) CreateCategoryPaginationQuery() *models.CategoryPaginationQuery {
	paginationQuery := f.paginationModelFactory.CreatePaginationQuery()
	return &models.CategoryPaginationQuery{PaginationQuery: paginationQuery}
}

func (*categoryModelFactoryImpl) CreateCategory() *models.Category {
	return new(models.Category)
}

func (*categoryModelFactoryImpl) CreateCategoryCreate() *models.CategoryCreate {
	return new(models.CategoryCreate)
}

func (*categoryModelFactoryImpl) CreateCategoryUpdate() *models.CategoryUpdate {
	return new(models.CategoryUpdate)
}
