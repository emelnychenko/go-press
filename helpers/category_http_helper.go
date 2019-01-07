package helpers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

const (
	CategoryIdParameterName       = "categoryId"
	ParentCategoryIdParameterName = "parentCategoryId"
)

type (
	categoryHttpHelperImpl struct {
	}
)

//NewCategoryHttpHelper
func NewCategoryHttpHelper() contracts.CategoryHttpHelper {
	return new(categoryHttpHelperImpl)
}

//ParseCategoryId
func (*categoryHttpHelperImpl) ParseCategoryId(httpContext contracts.HttpContext) (*models.CategoryId, errors.Error) {
	return models.ParseModelId(httpContext.Parameter(CategoryIdParameterName))
}

//ParseParentCategoryId
func (*categoryHttpHelperImpl) ParseParentCategoryId(httpContext contracts.HttpContext) (*models.CategoryId, errors.Error) {
	return models.ParseModelId(httpContext.Parameter(ParentCategoryIdParameterName))
}
