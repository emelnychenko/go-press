package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserApi interface {
		ListUsers() (users []*models.User, err common.Error)
		GetUser(userId *models.UserId) (user *models.User, err common.Error)
		CreateUser(data *models.UserCreate) (user *models.User, err common.Error)
		UpdateUser(userId *models.UserId, data *models.UserUpdate) (err common.Error)
		VerifyUser(userId *models.UserId) (err common.Error)
		ChangeUserIdentity(userId *models.UserId, data *models.UserChangeIdentity) (err common.Error)
		ChangeUserPassword(userId *models.UserId, data *models.UserChangePassword) (err common.Error)
		DeleteUser(userId *models.UserId) (err common.Error)
	}
)
