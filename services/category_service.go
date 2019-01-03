package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"time"
)

type (
	categoryServiceImpl struct {
		categoryEntityFactory contracts.CategoryEntityFactory
		categoryRepository    contracts.CategoryRepository
	}
)

func NewCategoryService(
	categoryEntityFactory contracts.CategoryEntityFactory,
	categoryRepository contracts.CategoryRepository,
) (categoryService contracts.CategoryService) {
	return &categoryServiceImpl{
		categoryEntityFactory,
		categoryRepository,
	}
}

func (s *categoryServiceImpl) ListCategories(
	categoryPaginationQuery *models.CategoryPaginationQuery,
) (*models.PaginationResult, common.Error) {
	return s.categoryRepository.ListCategories(categoryPaginationQuery)
}

func (s *categoryServiceImpl) GetCategory(categoryId *models.CategoryId) (*entities.CategoryEntity, common.Error) {
	return s.categoryRepository.GetCategory(categoryId)
}

func (s *categoryServiceImpl) CreateCategory(data *models.CategoryCreate) (
	categoryEntity *entities.CategoryEntity, err common.Error,
) {
	categoryEntity = s.categoryEntityFactory.CreateCategoryEntity()
	categoryEntity.Name = data.Name

	err = s.categoryRepository.SaveCategory(categoryEntity)
	return
}

func (s *categoryServiceImpl) UpdateCategory(
	categoryEntity *entities.CategoryEntity, data *models.CategoryUpdate,
) common.Error {
	categoryEntity.Name = data.Name

	updated := time.Now().UTC()
	categoryEntity.Updated = &updated

	return s.categoryRepository.SaveCategory(categoryEntity)
}

func (s *categoryServiceImpl) DeleteCategory(categoryEntity *entities.CategoryEntity) common.Error {
	return s.categoryRepository.RemoveCategory(categoryEntity)
}
