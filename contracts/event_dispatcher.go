package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	EventSubscriberFunc func(event Event) errors.Error

	EventDispatcher interface {
		Dispatch(event Event)
		Subscribe(eventName string, eventSubscriberFunc EventSubscriberFunc)
	}
)
