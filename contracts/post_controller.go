package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	PostController interface {
		ListPosts(httpContext HttpContext) (response interface{}, err errors.Error)
		GetPost(httpContext HttpContext) (response interface{}, err errors.Error)
		CreatePost(httpContext HttpContext) (response interface{}, err errors.Error)
		UpdatePost(httpContext HttpContext) (_ interface{}, err errors.Error)
		DeletePost(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
