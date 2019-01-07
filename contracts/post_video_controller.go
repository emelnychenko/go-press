package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	PostVideoController interface {
		ChangePostVideo(httpContext HttpContext) (_ interface{}, err errors.Error)
		RemovePostVideo(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
