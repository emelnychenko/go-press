package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	userModelFactoryImpl struct {
	}
)

func NewUserModelFactory() contracts.UserModelFactory {
	return new(userModelFactoryImpl)
}

func (*userModelFactoryImpl) CreateUser() *models.User {
	return new(models.User)
}

func (*userModelFactoryImpl) CreateUserCreate() *models.UserCreate {
	return new(models.UserCreate)
}

func (*userModelFactoryImpl) CreateUserUpdate() *models.UserUpdate {
	return new(models.UserUpdate)
}

func (*userModelFactoryImpl) CreateUserChangeIdentity() *models.UserChangeIdentity {
	return new(models.UserChangeIdentity)
}

func (*userModelFactoryImpl) CreateUserChangePassword() *models.UserChangePassword {
	return new(models.UserChangePassword)
}
