package services

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"time"
)

type (
	categoryServiceImpl struct {
		categoryEntityFactory   contracts.CategoryEntityFactory
		categoryRepository      contracts.CategoryRepository
		categoryTreeBuilder     contracts.CategoryTreeBuilder
		categoryEdgesBuilderJob contracts.CategoryEdgesBuilderJob
	}
)

//NewCategoryService
func NewCategoryService(
	categoryEntityFactory contracts.CategoryEntityFactory,
	categoryRepository contracts.CategoryRepository,
	categoryTreeBuilder contracts.CategoryTreeBuilder,
	categoryEdgesBuilderJob contracts.CategoryEdgesBuilderJob,
) (categoryService contracts.CategoryService) {
	return &categoryServiceImpl{
		categoryEntityFactory,
		categoryRepository,
		categoryTreeBuilder,
		categoryEdgesBuilderJob,
	}
}

//ListCategories
func (s *categoryServiceImpl) ListCategories(
	categoryPaginationQuery *models.CategoryPaginationQuery,
) (*models.PaginationResult, errors.Error) {
	return s.categoryRepository.ListCategories(categoryPaginationQuery)
}

//GetCategoriesTree
func (s *categoryServiceImpl) GetCategoriesTree() (*entities.CategoryEntityTree, errors.Error) {
	return s.categoryRepository.GetCategoriesTree()
}

//GetCategory
func (s *categoryServiceImpl) GetCategory(categoryId *models.CategoryId) (*entities.CategoryEntity, errors.Error) {
	return s.categoryRepository.GetCategory(categoryId)
}

//GetCategoryTree
func (s *categoryServiceImpl) GetCategoryTree(categoryId *models.CategoryId) (*entities.CategoryEntityTree, errors.Error) {
	return s.categoryRepository.GetCategoryTree(categoryId)
}

//CreateCategory
func (s *categoryServiceImpl) CreateCategory(data *models.CategoryCreate) (
	categoryEntity *entities.CategoryEntity, err errors.Error,
) {
	categoryEntity = s.categoryEntityFactory.CreateCategoryEntity()
	categoryEntity.Name = data.Name

	err = s.categoryRepository.SaveCategory(categoryEntity)
	return
}

//UpdateCategory
func (s *categoryServiceImpl) UpdateCategory(
	categoryEntity *entities.CategoryEntity, data *models.CategoryUpdate,
) errors.Error {
	categoryEntity.Name = data.Name

	updated := time.Now().UTC()
	categoryEntity.Updated = &updated

	return s.categoryRepository.SaveCategory(categoryEntity)
}

//ChangeCategoryParent
func (s *categoryServiceImpl) ChangeCategoryParent(
	categoryEntity *entities.CategoryEntity, parentCategoryEntity *entities.CategoryEntity,
) (err errors.Error) {
	categoryEntity.SetParentCategory(parentCategoryEntity)
	categoryEntities, err := s.categoryRepository.GetCategoriesExcept(categoryEntity)

	if nil != err {
		return
	}

	categoryEntities = append(categoryEntities, categoryEntity)
	// Assert circular dependency
	_, err = s.categoryTreeBuilder.BuildCategoryEntityTree(categoryEntities)

	if nil != err {
		return
	}

	err = s.categoryRepository.SaveCategory(categoryEntity)

	if nil != err {
		return
	}

	return s.categoryEdgesBuilderJob.BuildCategoriesEdges()
}

//RemoveCategoryParent
func (s *categoryServiceImpl) RemoveCategoryParent(categoryEntity *entities.CategoryEntity) (err errors.Error) {
	categoryEntity.RemoveParentCategory()
	err = s.categoryRepository.SaveCategory(categoryEntity)

	if nil != err {
		return
	}

	return s.categoryEdgesBuilderJob.BuildCategoriesEdges()
}

//DeleteCategory
func (s *categoryServiceImpl) DeleteCategory(categoryEntity *entities.CategoryEntity) errors.Error {
	return s.categoryRepository.RemoveCategory(categoryEntity)
}

//GetCategoryXrefs
func (s *categoryServiceImpl) GetCategoryXrefs(categoryEntity *entities.CategoryEntity) (
	[]*entities.CategoryXrefEntity, errors.Error,
) {
	return s.categoryRepository.GetCategoryXrefs(categoryEntity)
}

//GetCategoryObjectXrefs
func (s *categoryServiceImpl) GetCategoryObjectXrefs(categoryObject models.Object) (
	[]*entities.CategoryXrefEntity, errors.Error,
) {
	return s.categoryRepository.GetCategoryObjectXrefs(categoryObject)
}

//GetCategoryXref
func (s *categoryServiceImpl) GetCategoryXref(categoryEntity *entities.CategoryEntity, categoryObject models.Object) (
	*entities.CategoryXrefEntity, errors.Error,
) {
	return s.categoryRepository.GetCategoryXref(categoryEntity, categoryObject)
}

//CreateCategoryXref
func (s *categoryServiceImpl) CreateCategoryXref(categoryEntity *entities.CategoryEntity, categoryObject models.Object) (
	categoryXrefEntity *entities.CategoryXrefEntity, err errors.Error,
) {
	categoryXrefEntity = s.categoryEntityFactory.CreateCategoryXrefEntity(categoryEntity, categoryObject)
	err = s.categoryRepository.SaveCategoryXref(categoryXrefEntity)
	return
}

//DeleteCategoryXref
func (s *categoryServiceImpl) DeleteCategoryXref(categoryXrefEntity *entities.CategoryXrefEntity) errors.Error {
	return s.categoryRepository.RemoveCategoryXref(categoryXrefEntity)
}
