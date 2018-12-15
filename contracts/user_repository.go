package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserRepository interface {
		ListUsers() ([]*entities.UserEntity, common.Error)
		GetUser(userId *models.UserId) (*entities.UserEntity, common.Error)
		LookupUser(userIdentity string) (*entities.UserEntity, common.Error)
		SaveUser(userEntity *entities.UserEntity) common.Error
		RemoveUser(userEntity *entities.UserEntity) common.Error
	}
)
