package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	TagController interface {
		ListTags(httpContext HttpContext) (response interface{}, err errors.Error)
		GetTag(httpContext HttpContext) (response interface{}, err errors.Error)
		CreateTag(httpContext HttpContext) (response interface{}, err errors.Error)
		UpdateTag(httpContext HttpContext) (_ interface{}, err errors.Error)
		DeleteTag(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
