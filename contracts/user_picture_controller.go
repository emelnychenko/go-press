package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	UserPictureController interface {
		ChangeUserPicture(httpContext HttpContext) (_ interface{}, err common.Error)
		RemoveUserPicture(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
