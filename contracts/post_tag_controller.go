package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	PostTagController interface {
		ListPostTags(HttpContext) (interface{}, errors.Error)
		AddPostTag(HttpContext) (interface{}, errors.Error)
		RemovePostTag(HttpContext) (interface{}, errors.Error)
	}
)
