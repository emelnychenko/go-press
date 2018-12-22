package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	userEntityFactoryImpl struct {
	}
)

func NewUserEntityFactory() contracts.UserEntityFactory {
	return new(userEntityFactoryImpl)
}

func (*userEntityFactoryImpl) CreateUserEntity() *entities.UserEntity {
	return entities.NewUserEntity()
}
