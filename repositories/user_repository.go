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

func (c *userRepositoryImpl) ListUsers() (userEntities []*entities.UserEntity, err common.Error) {
	if gormErr := c.db.Find(&userEntities).Error; nil != gormErr {
		err = common.NewSystemError(gormErr)
	}

	return
}

func (c *userRepositoryImpl) GetUser(userId *models.UserId) (userEntity *entities.UserEntity, err common.Error) {
	userEntity = new(entities.UserEntity)

	if gormErr := c.db.First(userEntity, "id = ?", userId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewUserByIdNotFoundError(userId)
		} else {
			err = common.NewSystemError(gormErr)
		}
	}

	return
}

func (c *userRepositoryImpl) LookupUser(userIdentity string) (userEntity *entities.UserEntity, err common.Error) {
	userEntity = new(entities.UserEntity)

	if gormErr := c.db.First(userEntity, "email = ?", userIdentity).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewUserNotFoundError(userIdentity)
		} else {
			err = common.NewSystemError(gormErr)
		}
	}

	return
}

func (c *userRepositoryImpl) SaveUser(userEntity *entities.UserEntity) (err common.Error) {
	if gormErr := c.db.Save(userEntity).Error; gormErr != nil {
		err = common.NewSystemError(gormErr)
	}

	return
}

func (c *userRepositoryImpl) RemoveUser(userEntity *entities.UserEntity) (err common.Error) {
	if gormErr := c.db.Delete(userEntity).Error; gormErr != nil {
		err = common.NewSystemError(gormErr)
	}

	return
}
