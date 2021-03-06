package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	UserModelFactory interface {
		CreateUserPaginationQuery() *models.UserPaginationQuery
		CreateUser() *models.User
		CreateUserCreate() *models.UserCreate
		CreateUserUpdate() *models.UserUpdate
		CreateUserChangeIdentity() *models.UserChangeIdentity
		CreateUserChangePassword() *models.UserChangePassword
	}
)
