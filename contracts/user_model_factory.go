package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	UserModelFactory interface {
		CreateUser() *models.User
		CreateUserCreate() *models.UserCreate
		CreateUserUpdate() *models.UserUpdate
		CreateUserChangeIdentity() *models.UserChangeIdentity
		CreateUserChangePassword() *models.UserChangePassword
	}
)
