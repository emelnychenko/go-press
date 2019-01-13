package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	PostCategoryController interface {
		ListPostCategories(HttpContext) (interface{}, errors.Error)
		AddPostCategory(HttpContext) (interface{}, errors.Error)
		RemovePostCategory(HttpContext) (interface{}, errors.Error)
	}
)
