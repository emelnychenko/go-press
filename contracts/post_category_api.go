package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostCategoryApi interface {
		ListPostCategories(*models.PostId, *models.CategoryPaginationQuery) (*models.PaginationResult, errors.Error)
		AddPostCategory(*models.PostId, *models.CategoryId) errors.Error
		RemovePostCategory(*models.PostId, *models.CategoryId) errors.Error
	}
)
