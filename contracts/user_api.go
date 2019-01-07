package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserApi interface {
		ListUsers(userPaginationQuery *models.UserPaginationQuery) (*models.PaginationResult, errors.Error)
		GetUser(userId *models.UserId) (user *models.User, err errors.Error)
		CreateUser(data *models.UserCreate) (user *models.User, err errors.Error)
		UpdateUser(userId *models.UserId, data *models.UserUpdate) (err errors.Error)
		VerifyUser(userId *models.UserId) (err errors.Error)
		ChangeUserIdentity(userId *models.UserId, data *models.UserChangeIdentity) (err errors.Error)
		ChangeUserPassword(userId *models.UserId, data *models.UserChangePassword) (err errors.Error)
		DeleteUser(userId *models.UserId) (err errors.Error)
	}
)
