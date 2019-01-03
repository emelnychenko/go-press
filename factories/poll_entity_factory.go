package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	pollEntityFactoryImpl struct {
	}
)

func NewPollEntityFactory() contracts.PollEntityFactory {
	return new(pollEntityFactoryImpl)
}

func (*pollEntityFactoryImpl) CreatePollEntity() *entities.PollEntity {
	return entities.NewPollEntity()
}
