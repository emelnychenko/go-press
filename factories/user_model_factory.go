package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	userModelFactoryImpl struct {
		paginationModelFactory contracts.PaginationModelFactory
	}
)

func NewUserModelFactory(paginationModelFactory contracts.PaginationModelFactory) contracts.UserModelFactory {
	return &userModelFactoryImpl{paginationModelFactory}
}

func (f *userModelFactoryImpl) CreateUserPaginationQuery() *models.UserPaginationQuery {
	paginationQuery := f.paginationModelFactory.CreatePaginationQuery()
	return &models.UserPaginationQuery{PaginationQuery: paginationQuery}
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
