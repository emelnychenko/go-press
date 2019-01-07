package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryRepository interface {
		ListCategories(categoryPaginationQuery *models.CategoryPaginationQuery) (*models.PaginationResult, errors.Error)
		GetCategories() ([]*entities.CategoryEntity, errors.Error)
		GetCategoriesExcept(categoryEntity *entities.CategoryEntity) ([]*entities.CategoryEntity, errors.Error)
		GetCategoriesTree() (*entities.CategoryEntityTree, errors.Error)
		GetCategory(categoryId *models.CategoryId) (*entities.CategoryEntity, errors.Error)
		GetCategoryTree(categoryId *models.CategoryId) (*entities.CategoryEntityTree, errors.Error)
		SaveCategory(categoryEntity *entities.CategoryEntity) (err errors.Error)
		RemoveCategory(categoryEntity *entities.CategoryEntity) (err errors.Error)

		GetCategoryXrefs(*entities.CategoryEntity) ([]*entities.CategoryXrefEntity, errors.Error)
		GetCategoryObjectXrefs(models.Object) ([]*entities.CategoryXrefEntity, errors.Error)
		GetCategoryXref(*entities.CategoryEntity, models.Object) (*entities.CategoryXrefEntity, errors.Error)
		SaveCategoryXref(*entities.CategoryXrefEntity) errors.Error
		RemoveCategoryXref(*entities.CategoryXrefEntity) errors.Error

		ListObjectCategories(models.Object, *models.CategoryPaginationQuery) (*models.PaginationResult, errors.Error)
	}
)
