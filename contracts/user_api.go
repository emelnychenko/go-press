package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserApi interface {
		ListUsers() ([]*models.User, common.Error)
		GetUser(userId *models.UserId) (*models.User, common.Error)
		CreateUser(data *models.UserCreate) (*models.User, common.Error)
		UpdateUser(userId *models.UserId, data *models.UserUpdate) common.Error
		VerifyUser(userId *models.UserId) common.Error
		ChangeUserIdentity(userId *models.UserId, data *models.UserChangeIdentity) common.Error
		ChangeUserPassword(userId *models.UserId, data *models.UserChangePassword) common.Error
		DeleteUser(userId *models.UserId) common.Error
	}
)
