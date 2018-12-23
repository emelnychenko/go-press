package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	UserController interface {
		ListUsers(httpContext HttpContext) (response interface{}, err common.Error)
		GetUser(httpContext HttpContext) (response interface{}, err common.Error)
		CreateUser(httpContext HttpContext) (response interface{}, err common.Error)
		UpdateUser(httpContext HttpContext) (_ interface{}, err common.Error)
		ChangeUserIdentity(httpContext HttpContext) (_ interface{}, err common.Error)
		ChangeUserPassword(httpContext HttpContext) (_ interface{}, err common.Error)
		DeleteUser(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
