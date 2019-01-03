package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryModelFactory interface {
		CreateCategoryPaginationQuery() *models.CategoryPaginationQuery
		CreateCategory() *models.Category
		CreateCategoryCreate() *models.CategoryCreate
		CreateCategoryUpdate() *models.CategoryUpdate
	}
)
