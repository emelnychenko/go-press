package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	UserController interface {
		ListUsers(httpContext HttpContext) (response interface{}, err errors.Error)
		GetUser(httpContext HttpContext) (response interface{}, err errors.Error)
		CreateUser(httpContext HttpContext) (response interface{}, err errors.Error)
		UpdateUser(httpContext HttpContext) (_ interface{}, err errors.Error)
		ChangeUserIdentity(httpContext HttpContext) (_ interface{}, err errors.Error)
		ChangeUserPassword(httpContext HttpContext) (_ interface{}, err errors.Error)
		DeleteUser(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
