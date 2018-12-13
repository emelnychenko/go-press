package user_contract

import (
	"../common"
	"../user_domain"
)

type (
	UserApi interface {
		ListUsers() ([]*user_domain.User, common.Error)
		GetUser(userId *user_domain.UserId) (*user_domain.User, common.Error)
		CreateUser(data *user_domain.UserCreate) (*user_domain.User, common.Error)
		UpdateUser(userId *user_domain.UserId, data *user_domain.UserUpdate) common.Error
		VerifyUser(userId *user_domain.UserId) common.Error
		ChangeUserIdentity(userId *user_domain.UserId, data *user_domain.UserChangeIdentity) common.Error
		ChangeUserPassword(userId *user_domain.UserId, data *user_domain.UserChangePassword) common.Error
		DeleteUser(userId *user_domain.UserId) common.Error
	}
)
