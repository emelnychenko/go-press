package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	PostPublisherWorker interface {
		Start() errors.Error
		Stop() errors.Error
	}
)
