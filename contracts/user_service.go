package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserService interface {
		ListUsers(userPaginationQuery *models.UserPaginationQuery) (*models.PaginationResult, errors.Error)
		GetUser(userId *models.UserId) (*entities.UserEntity, errors.Error)
		LookupUser(userIdentity string) (*entities.UserEntity, errors.Error)
		ChallengeUser(userEntity *entities.UserEntity, password string) errors.Error
		CreateUser(data *models.UserCreate) (*entities.UserEntity, errors.Error)
		VerifyUser(userEntity *entities.UserEntity) errors.Error
		ChangeUserIdentity(userEntity *entities.UserEntity, data *models.UserChangeIdentity) errors.Error
		ChangeUserPassword(userEntity *entities.UserEntity, data *models.UserChangePassword) errors.Error
		UpdateUser(userEntity *entities.UserEntity, data *models.UserUpdate) errors.Error
		DeleteUser(userEntity *entities.UserEntity) errors.Error
	}
)
