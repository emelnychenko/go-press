package events

import "github.com/emelnychenko/go-press/contracts"

type (
	Event struct {
		name string
	}
)

func NewEvent(name string) contracts.Event {
	return &Event{name: name}
}

func (e *Event) Name() string {
	return e.name
}
