package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryService interface {
		ListCategories(categoryPaginationQuery *models.CategoryPaginationQuery) (*models.PaginationResult, common.Error)
		GetCategoriesTree() (*entities.CategoryEntityTree, common.Error)
		GetCategory(categoryId *models.CategoryId) (*entities.CategoryEntity, common.Error)
		GetCategoryTree(categoryId *models.CategoryId) (*entities.CategoryEntityTree, common.Error)
		CreateCategory(data *models.CategoryCreate) (*entities.CategoryEntity, common.Error)
		UpdateCategory(categoryEntity *entities.CategoryEntity, data *models.CategoryUpdate) common.Error
		ChangeCategoryParent(categoryEntity *entities.CategoryEntity, parentCategoryEntity *entities.CategoryEntity) common.Error
		RemoveCategoryParent(categoryEntity *entities.CategoryEntity) common.Error
		DeleteCategory(categoryEntity *entities.CategoryEntity) common.Error
	}
)
