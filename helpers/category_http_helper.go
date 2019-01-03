package helpers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

const (
	CategoryIdParameterName = "categoryId"
)

type (
	categoryHttpHelperImpl struct {
	}
)

func NewCategoryHttpHelper() contracts.CategoryHttpHelper {
	return new(categoryHttpHelperImpl)
}

func (*categoryHttpHelperImpl) ParseCategoryId(httpContext contracts.HttpContext) (*models.CategoryId, common.Error) {
	return common.ParseModelId(httpContext.Parameter(CategoryIdParameterName))
}

