package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	HttpHandlerFunc func (httpContext HttpContext) (response interface{}, err common.Error)

	Router interface {
		AddRoute(httpMethod string, routePath string, httpHandlerFunc HttpHandlerFunc)
	}
)
