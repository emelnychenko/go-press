package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryService interface {
		ListCategories(categoryPaginationQuery *models.CategoryPaginationQuery) (*models.PaginationResult, errors.Error)
		GetCategoriesTree() (*entities.CategoryEntityTree, errors.Error)
		GetCategory(categoryId *models.CategoryId) (*entities.CategoryEntity, errors.Error)
		GetCategoryTree(categoryId *models.CategoryId) (*entities.CategoryEntityTree, errors.Error)
		CreateCategory(data *models.CategoryCreate) (*entities.CategoryEntity, errors.Error)
		UpdateCategory(categoryEntity *entities.CategoryEntity, data *models.CategoryUpdate) errors.Error
		ChangeCategoryParent(categoryEntity *entities.CategoryEntity, parentCategoryEntity *entities.CategoryEntity) errors.Error
		RemoveCategoryParent(categoryEntity *entities.CategoryEntity) errors.Error
		DeleteCategory(categoryEntity *entities.CategoryEntity) errors.Error

		GetCategoryXrefs(*entities.CategoryEntity) ([]*entities.CategoryXrefEntity, errors.Error)
		GetCategoryObjectXrefs(models.Object) ([]*entities.CategoryXrefEntity, errors.Error)
		GetCategoryXref(*entities.CategoryEntity, models.Object) (*entities.CategoryXrefEntity, errors.Error)
		CreateCategoryXref(*entities.CategoryEntity, models.Object) (*entities.CategoryXrefEntity, errors.Error)
		DeleteCategoryXref(*entities.CategoryXrefEntity) errors.Error
	}
)
