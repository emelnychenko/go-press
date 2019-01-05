package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryApi interface {
		ListCategories(categoryPaginationQuery *models.CategoryPaginationQuery) (*models.PaginationResult, common.Error)
		GetCategoriesTree() ([]*models.CategoryTree, common.Error)
		GetCategory(categoryId *models.CategoryId) (category *models.Category, err common.Error)
		GetCategoryTree(categoryId *models.CategoryId) (*models.CategoryTree, common.Error)
		CreateCategory(data *models.CategoryCreate) (category *models.Category, err common.Error)
		UpdateCategory(categoryId *models.CategoryId, data *models.CategoryUpdate) (err common.Error)
		DeleteCategory(categoryId *models.CategoryId) (err common.Error)
	}
)
