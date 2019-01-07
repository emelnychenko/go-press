package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryHttpHelper interface {
		ParseCategoryId(HttpContext) (*models.CategoryId, errors.Error)
		ParseParentCategoryId(HttpContext) (*models.CategoryId, errors.Error)
	}
)
