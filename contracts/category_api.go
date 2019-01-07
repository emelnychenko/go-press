package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryApi interface {
		ListCategories(categoryPaginationQuery *models.CategoryPaginationQuery) (*models.PaginationResult, errors.Error)
		GetCategoriesTree() ([]*models.CategoryTree, errors.Error)
		GetCategory(categoryId *models.CategoryId) (category *models.Category, err errors.Error)
		GetCategoryTree(categoryId *models.CategoryId) (*models.CategoryTree, errors.Error)
		CreateCategory(data *models.CategoryCreate) (category *models.Category, err errors.Error)
		UpdateCategory(categoryId *models.CategoryId, data *models.CategoryUpdate) (err errors.Error)
		ChangeCategoryParent(categoryId *models.CategoryId, parentCategoryId *models.CategoryId) errors.Error
		RemoveCategoryParent(categoryId *models.CategoryId) errors.Error
		DeleteCategory(categoryId *models.CategoryId) (err errors.Error)

		ListObjectCategories(models.Object, *models.CategoryPaginationQuery) (*models.PaginationResult, errors.Error)
	}
)
