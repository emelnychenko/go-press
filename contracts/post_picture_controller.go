package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	PostPictureController interface {
		ChangePostPicture(httpContext HttpContext) (_ interface{}, err errors.Error)
		RemovePostPicture(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
