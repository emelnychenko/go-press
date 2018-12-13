package user_contract

import (
	"../common"
	"../user_domain"
)

type (
	UserRepository interface {
		ListUsers() ([]*user_domain.UserEntity, common.Error)
		GetUser(userId *user_domain.UserId) (*user_domain.UserEntity, common.Error)
		LookupUser(userIdentity string) (*user_domain.UserEntity, common.Error)
		SaveUser(userEntity *user_domain.UserEntity) common.Error
		RemoveUser(userEntity *user_domain.UserEntity) common.Error
	}
)
