package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"time"
)

type (
	userServiceImpl struct {
		hasher     contracts.Hasher
		repository contracts.UserRepository
	}
)

func NewUserService(hasher contracts.Hasher, userRepository contracts.UserRepository) contracts.UserService {
	return &userServiceImpl{hasher: hasher, repository: userRepository}
}

func (c *userServiceImpl) ListUsers() ([]*entities.UserEntity, common.Error) {
	return c.repository.ListUsers()
}

func (c *userServiceImpl) CreateUser(data *models.UserCreate) (*entities.UserEntity, common.Error) {
	userEntity := entities.NewUserEntity()
	userEntity.FirstName = data.FirstName
	userEntity.LastName = data.LastName
	userEntity.Email = data.Email

	if data.Password != "" {
		hash, err := c.hasher.Make(data.Password)

		if nil != err {
			return nil, err
		}

		userEntity.Password = hash
	}

	if err := c.repository.SaveUser(userEntity); err != nil {
		return nil, err
	}

	return userEntity, nil
}

func (c *userServiceImpl) GetUser(userId *models.UserId) (*entities.UserEntity, common.Error) {
	return c.repository.GetUser(userId)
}

func (c *userServiceImpl) LookupUser(entityIdentity string) (*entities.UserEntity, common.Error) {
	return c.repository.LookupUser(entityIdentity)
}

func (c *userServiceImpl) ChallengeUser(userEntity *entities.UserEntity, password string) common.Error {
	if userEntity.Password == "" {
		return common.ServerError("password is not registered")
	}

	return c.hasher.Check(userEntity.Password, password)
}

func (c *userServiceImpl) UpdateUser(userEntity *entities.UserEntity, data *models.UserUpdate) common.Error {
	userEntity.FirstName = data.FirstName
	userEntity.LastName = data.LastName

	updated := time.Now().UTC()
	userEntity.Updated = &updated

	return c.repository.SaveUser(userEntity)
}

func (c *userServiceImpl) VerifyUser(entity *entities.UserEntity) common.Error {
	entity.Verified = true

	return c.repository.SaveUser(entity)
}

func (c *userServiceImpl) ChangeUserIdentity(userEntity *entities.UserEntity, data *models.UserChangeIdentity) common.Error {
	userEntity.Email = data.Email

	return c.repository.SaveUser(userEntity)
}

func (c *userServiceImpl) ChangeUserPassword(userEntity *entities.UserEntity, data *models.UserChangePassword) common.Error {
	if data.NewPassword != "" {
		hash, err := c.hasher.Make(data.NewPassword)

		if nil != err {
			return err
		}

		userEntity.Password = string(hash)
	} else {
		userEntity.Password = ""
	}

	return c.repository.SaveUser(userEntity)
}

func (c *userServiceImpl) DeleteUser(userEntity *entities.UserEntity) common.Error {
	return c.repository.RemoveUser(userEntity)
}
