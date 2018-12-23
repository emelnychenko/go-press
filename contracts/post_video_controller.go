package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	PostVideoController interface {
		ChangePostVideo(httpContext HttpContext) (_ interface{}, err common.Error)
		RemovePostVideo(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
