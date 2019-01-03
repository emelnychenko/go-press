package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	CategoryController interface {
		ListCategories(httpContext HttpContext) (response interface{}, err common.Error)
		GetCategory(httpContext HttpContext) (response interface{}, err common.Error)
		CreateCategory(httpContext HttpContext) (response interface{}, err common.Error)
		UpdateCategory(httpContext HttpContext) (_ interface{}, err common.Error)
		DeleteCategory(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
