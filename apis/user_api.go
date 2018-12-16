package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	userApiImpl struct {
		userService    contracts.UserService
		userAggregator contracts.UserAggregator
	}
)

func NewUserApi(userService contracts.UserService, aggregator contracts.UserAggregator) contracts.UserApi {
	return &userApiImpl{userService: userService, userAggregator: aggregator}
}

func (c *userApiImpl) ListUsers() ([]*models.User, common.Error) {
	userEntities, err := c.userService.ListUsers()

	if nil != err {
		return nil, err
	}

	return c.userAggregator.AggregateCollection(userEntities), nil
}

func (c *userApiImpl) CreateUser(data *models.UserCreate) (*models.User, common.Error) {
	userEntity, err := c.userService.CreateUser(data)

	if nil != err {
		return nil, err
	}

	return c.userAggregator.AggregateObject(userEntity), nil
}

func (c *userApiImpl) GetUser(userId *models.UserId) (*models.User, common.Error) {
	userEntity, err := c.userService.GetUser(userId)

	if nil != err {
		return nil, err
	}

	return c.userAggregator.AggregateObject(userEntity), nil
}

func (c *userApiImpl) UpdateUser(userId *models.UserId, data *models.UserUpdate) common.Error {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return err
	}

	return userService.UpdateUser(userEntity, data)
}

func (c *userApiImpl) VerifyUser(userId *models.UserId) common.Error {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return err
	}

	return userService.VerifyUser(userEntity)
}

func (c *userApiImpl) ChangeUserIdentity(userId *models.UserId, data *models.UserChangeIdentity) common.Error {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return err
	}

	return userService.ChangeUserIdentity(userEntity, data)
}

func (c *userApiImpl) ChangeUserPassword(userId *models.UserId, input *models.UserChangePassword) common.Error {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return err
	}

	if userEntity.Password != "" {
		if err = userService.ChallengeUser(userEntity, input.OldPassword); err != nil {
			return err
		}
	}

	return userService.ChangeUserPassword(userEntity, input)
}

func (c *userApiImpl) DeleteUser(userId *models.UserId) common.Error {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return err
	}

	return userService.DeleteUser(userEntity)
}
