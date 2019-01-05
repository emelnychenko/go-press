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

//NewCategoryModelFactory
func NewCategoryModelFactory(paginationModelFactory contracts.PaginationModelFactory) contracts.CategoryModelFactory {
	return &categoryModelFactoryImpl{paginationModelFactory}
}

//CreateCategoryPaginationQuery
func (f *categoryModelFactoryImpl) CreateCategoryPaginationQuery() *models.CategoryPaginationQuery {
	paginationQuery := f.paginationModelFactory.CreatePaginationQuery()
	return &models.CategoryPaginationQuery{PaginationQuery: paginationQuery}
}

//CreateCategory
func (*categoryModelFactoryImpl) CreateCategory() *models.Category {
	return new(models.Category)
}

//CreateCategoryTree
func (f *categoryModelFactoryImpl) CreateCategoryTree() *models.CategoryTree {
	return new(models.CategoryTree)
}

//CreateCategoryCreate
func (*categoryModelFactoryImpl) CreateCategoryCreate() *models.CategoryCreate {
	return new(models.CategoryCreate)
}

//CreateCategoryUpdate
func (*categoryModelFactoryImpl) CreateCategoryUpdate() *models.CategoryUpdate {
	return new(models.CategoryUpdate)
}
