package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryRepository interface {
		ListCategories(categoryPaginationQuery *models.CategoryPaginationQuery) (*models.PaginationResult, common.Error)
		GetCategories() ([]*entities.CategoryEntity, common.Error)
		GetCategoriesExcept(categoryEntity *entities.CategoryEntity) ([]*entities.CategoryEntity, common.Error)
		GetCategoriesTree() (*entities.CategoryEntityTree, common.Error)
		GetCategory(categoryId *models.CategoryId) (categoryEntity *entities.CategoryEntity, err common.Error)
		GetCategoryTree(categoryId *models.CategoryId) (*entities.CategoryEntityTree, common.Error)
		SaveCategory(categoryEntity *entities.CategoryEntity) (err common.Error)
		RemoveCategory(categoryEntity *entities.CategoryEntity) (err common.Error)
	}
)
