package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	PostAuthorController interface {
		ChangePostAuthor(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
