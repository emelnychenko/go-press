package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	UserPictureController interface {
		ChangeUserPicture(httpContext HttpContext) (_ interface{}, err errors.Error)
		RemoveUserPicture(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
