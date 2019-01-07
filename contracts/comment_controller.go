package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	CommentController interface {
		ListComments(httpContext HttpContext) (response interface{}, err errors.Error)
		GetComment(httpContext HttpContext) (response interface{}, err errors.Error)
		CreateComment(httpContext HttpContext) (response interface{}, err errors.Error)
		UpdateComment(httpContext HttpContext) (_ interface{}, err errors.Error)
		DeleteComment(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
