package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
)

type (
	PostPublisherJob interface {
		PublishPost(postEntity *entities.PostEntity) errors.Error
	}
)
