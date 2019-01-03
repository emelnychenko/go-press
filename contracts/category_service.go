package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryService interface {
		ListCategories(categoryPaginationQuery *models.CategoryPaginationQuery) (*models.PaginationResult, common.Error)
		GetCategory(categoryId *models.CategoryId) (categoryEntity *entities.CategoryEntity, err common.Error)
		CreateCategory(data *models.CategoryCreate) (categoryEntity *entities.CategoryEntity, err common.Error)
		UpdateCategory(categoryEntity *entities.CategoryEntity, data *models.CategoryUpdate) (err common.Error)
		DeleteCategory(categoryEntity *entities.CategoryEntity) (err common.Error)
	}
)
