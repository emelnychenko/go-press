package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	CategoryController interface {
		ListCategories(httpContext HttpContext) (response interface{}, err errors.Error)
		GetCategoriesTree(httpContext HttpContext) (response interface{}, err errors.Error)
		GetCategory(httpContext HttpContext) (response interface{}, err errors.Error)
		GetCategoryTree(httpContext HttpContext) (response interface{}, err errors.Error)
		CreateCategory(httpContext HttpContext) (response interface{}, err errors.Error)
		UpdateCategory(httpContext HttpContext) (_ interface{}, err errors.Error)
		ChangeCategoryParent(httpContext HttpContext) (_ interface{}, err errors.Error)
		RemoveCategoryParent(httpContext HttpContext) (_ interface{}, err errors.Error)
		DeleteCategory(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
