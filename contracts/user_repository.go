package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserRepository interface {
		ListUsers(userPaginationQuery *models.UserPaginationQuery) (*models.PaginationResult, errors.Error)
		GetUser(userId *models.UserId) (*entities.UserEntity, errors.Error)
		LookupUser(userIdentity string) (*entities.UserEntity, errors.Error)
		SaveUser(userEntity *entities.UserEntity) errors.Error
		RemoveUser(userEntity *entities.UserEntity) errors.Error
	}
)
