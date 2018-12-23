package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	PostPictureController interface {
		ChangePostPicture(httpContext HttpContext) (_ interface{}, err common.Error)
		RemovePostPicture(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
