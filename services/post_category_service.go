package services

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	postCategoryServiceImpl struct {
		categoryService contracts.CategoryService
	}
)

//NewPostCategoryService
func NewPostCategoryService(categoryService contracts.CategoryService) contracts.PostCategoryService {
	return &postCategoryServiceImpl{categoryService: categoryService}
}

//ListPostCategories
func (s *postCategoryServiceImpl) ListPostCategories(
	postEntity *entities.PostEntity, categoryPaginationQuery *models.CategoryPaginationQuery,
) (*models.PaginationResult, errors.Error) {
	return s.categoryService.ListObjectCategories(postEntity, categoryPaginationQuery)
}

//AddPostCategory
func (s *postCategoryServiceImpl) AddPostCategory(
	postEntity *entities.PostEntity, categoryEntity *entities.CategoryEntity,
) (err errors.Error) {
	_, err = s.categoryService.CreateCategoryXref(categoryEntity, postEntity)
	return
}

//RemovePostCategory
func (s *postCategoryServiceImpl) RemovePostCategory(
	postEntity *entities.PostEntity, categoryEntity *entities.CategoryEntity,
) (err errors.Error) {
	categoryXrefEntity, err := s.categoryService.GetCategoryXref(categoryEntity, postEntity)

	if nil != err {
		return
	}

	return s.categoryService.DeleteCategoryXref(categoryXrefEntity)
}
