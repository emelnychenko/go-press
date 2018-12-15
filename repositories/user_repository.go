package repositories

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"github.com/jinzhu/gorm"
)

type (
	userRepositoryImpl struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) contracts.UserRepository {
	return &userRepositoryImpl{db}
}

func (c *userRepositoryImpl) ListUsers() ([]*entities.UserEntity, common.Error) {
	var userEntities []*entities.UserEntity

	if err := c.db.Find(&userEntities).Error; nil != err {
		return nil, common.NewServerError(err)
	}

	return userEntities, nil
}

func (c *userRepositoryImpl) GetUser(userId *models.UserId) (*entities.UserEntity, common.Error) {
	var userEntity entities.UserEntity

	if err := c.db.First(&userEntity, "id = ?", userId).Error; err != nil {
		return nil, errors.NewUserNotFoundError(userId.String())
	}

	return &userEntity, nil
}

func (c *userRepositoryImpl) LookupUser(userIdentity string) (*entities.UserEntity, common.Error) {
	var userEntity entities.UserEntity

	if err := c.db.First(&userEntity, "email = ?", userIdentity).Error; err != nil {
		return nil, errors.NewUserNotFoundError(userIdentity)
	}

	return &userEntity, nil
}

func (c *userRepositoryImpl) SaveUser(userEntity *entities.UserEntity) common.Error {
	if err := c.db.Save(userEntity).Error; err != nil {
		return common.NewServerError(err)
	}

	return nil
}

func (c *userRepositoryImpl) RemoveUser(userEntity *entities.UserEntity) common.Error {
	if err := c.db.Delete(userEntity).Error; err != nil {
		return common.NewServerError(err)
	}

	return nil
}
