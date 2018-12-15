package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserService interface {
		ListUsers() ([]*entities.UserEntity, common.Error)
		GetUser(userId *models.UserId) (*entities.UserEntity, common.Error)
		LookupUser(userIdentity string) (*entities.UserEntity, common.Error)
		ChallengeUser(userEntity *entities.UserEntity, password string) common.Error
		CreateUser(data *models.UserCreate) (*entities.UserEntity, common.Error)
		VerifyUser(userEntity *entities.UserEntity) common.Error
		ChangeUserIdentity(userEntity *entities.UserEntity, data *models.UserChangeIdentity) common.Error
		ChangeUserPassword(userEntity *entities.UserEntity, data *models.UserChangePassword) common.Error
		UpdateUser(userEntity *entities.UserEntity, data *models.UserUpdate) common.Error
		DeleteUser(userEntity *entities.UserEntity) common.Error
	}
)
