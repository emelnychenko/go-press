package user_contract

import (
	"../common"
	"../user_domain"
)

type (
	UserService interface {
		ListUsers() ([]*user_domain.UserEntity, common.Error)
		GetUser(userId *user_domain.UserId) (*user_domain.UserEntity, common.Error)
		LookupUser(userIdentity string) (*user_domain.UserEntity, common.Error)
		ChallengeUser(userEntity *user_domain.UserEntity, password string) common.Error
		CreateUser(data *user_domain.UserCreate) (*user_domain.UserEntity, common.Error)
		VerifyUser(userEntity *user_domain.UserEntity) common.Error
		ChangeUserIdentity(userEntity *user_domain.UserEntity, data *user_domain.UserChangeIdentity) common.Error
		ChangeUserPassword(userEntity *user_domain.UserEntity, data *user_domain.UserChangePassword) common.Error
		UpdateUser(userEntity *user_domain.UserEntity, data *user_domain.UserUpdate) common.Error
		DeleteUser(userEntity *user_domain.UserEntity) common.Error
	}
)
