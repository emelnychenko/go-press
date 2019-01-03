package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryHttpHelper interface {
		ParseCategoryId(httpContext HttpContext) (*models.CategoryId, common.Error)
	}
)
