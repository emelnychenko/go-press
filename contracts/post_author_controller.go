package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	PostAuthorController interface {
		ChangePostAuthor(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
