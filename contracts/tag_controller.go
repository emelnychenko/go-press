package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	TagController interface {
		ListTags(httpContext HttpContext) (response interface{}, err common.Error)
		GetTag(httpContext HttpContext) (response interface{}, err common.Error)
		CreateTag(httpContext HttpContext) (response interface{}, err common.Error)
		UpdateTag(httpContext HttpContext) (_ interface{}, err common.Error)
		DeleteTag(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
