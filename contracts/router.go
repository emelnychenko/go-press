package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	HttpHandlerFunc func(httpContext HttpContext) (response interface{}, err errors.Error)

	Router interface {
		AddRoute(httpMethod string, routePath string, httpHandlerFunc HttpHandlerFunc)
	}
)
