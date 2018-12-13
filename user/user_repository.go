package user

import (
	"../common"
	"../user_domain"
	"github.com/jinzhu/gorm"
)

type (
	userRepositoryImpl struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db}
}

func (c *userRepositoryImpl) ListUsers() ([]*user_domain.UserEntity, common.Error) {
	var userEntities []*user_domain.UserEntity

	if err := c.db.Find(&userEntities).Error; nil != err {
		return nil, common.NewServerError(err)
	}

	return userEntities, nil
}

func (c *userRepositoryImpl) GetUser(userId *user_domain.UserId) (*user_domain.UserEntity, common.Error) {
	var userEntity user_domain.UserEntity

	if err := c.db.First(&userEntity, "id = ?", userId).Error; err != nil {
		return nil, user_domain.NewUserNotFoundError(userId.String())
	}

	return &userEntity, nil
}

func (c *userRepositoryImpl) LookupUser(userIdentity string) (*user_domain.UserEntity, common.Error) {
	var userEntity user_domain.UserEntity

	if err := c.db.First(&userEntity, "email = ?", userIdentity).Error; err != nil {
		return nil, user_domain.NewUserNotFoundError(userIdentity)
	}

	return &userEntity, nil
}

func (c *userRepositoryImpl) SaveUser(userEntity *user_domain.UserEntity) common.Error {
	if err := c.db.Save(userEntity).Error; err != nil {
		return common.NewServerError(err)
	}

	return nil
}

func (c *userRepositoryImpl) RemoveUser(userEntity *user_domain.UserEntity) common.Error {
	if err := c.db.Delete(userEntity).Error; err != nil {
		return common.NewServerError(err)
	}

	return nil
}
