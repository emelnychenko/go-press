package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostCategoryService interface {
		ListPostCategories(*entities.PostEntity, *models.CategoryPaginationQuery) (*models.PaginationResult, errors.Error)
		AddPostCategory(*entities.PostEntity, *entities.CategoryEntity) errors.Error
		RemovePostCategory(*entities.PostEntity, *entities.CategoryEntity) errors.Error
	}
)
