package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryRepository interface {
		ListCategories(categoryPaginationQuery *models.CategoryPaginationQuery) (*models.PaginationResult, common.Error)
		GetCategory(categoryId *models.CategoryId) (categoryEntity *entities.CategoryEntity, err common.Error)
		SaveCategory(categoryEntity *entities.CategoryEntity) (err common.Error)
		RemoveCategory(categoryEntity *entities.CategoryEntity) (err common.Error)
	}
)
