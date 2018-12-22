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

func NewUserApi(userService contracts.UserService, userAggregator contracts.UserAggregator) (userApi contracts.UserApi) {
	return &userApiImpl{userService: userService, userAggregator: userAggregator}
}

func (c *userApiImpl) ListUsers() (users []*models.User, err common.Error) {
	userEntities, err := c.userService.ListUsers()

	if nil != err {
		return
	}

	users = c.userAggregator.AggregateUsers(userEntities)
	return
}

func (c *userApiImpl) CreateUser(data *models.UserCreate) (user *models.User, err common.Error) {
	userEntity, err := c.userService.CreateUser(data)

	if nil != err {
		return
	}

	user = c.userAggregator.AggregateUser(userEntity)
	return
}

func (c *userApiImpl) GetUser(userId *models.UserId) (user *models.User, err common.Error) {
	userEntity, err := c.userService.GetUser(userId)

	if nil != err {
		return
	}

	user = c.userAggregator.AggregateUser(userEntity)
	return
}

func (c *userApiImpl) UpdateUser(userId *models.UserId, data *models.UserUpdate) (err common.Error) {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return
	}

	return userService.UpdateUser(userEntity, data)
}

func (c *userApiImpl) VerifyUser(userId *models.UserId) (err common.Error) {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return
	}

	return userService.VerifyUser(userEntity)
}

func (c *userApiImpl) ChangeUserIdentity(userId *models.UserId, data *models.UserChangeIdentity) (err common.Error) {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return
	}

	return userService.ChangeUserIdentity(userEntity, data)
}

func (c *userApiImpl) ChangeUserPassword(userId *models.UserId, input *models.UserChangePassword) (err common.Error) {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return
	}

	if userEntity.Password != "" {
		if err = userService.ChallengeUser(userEntity, input.OldPassword); err != nil {
			return
		}
	}

	return userService.ChangeUserPassword(userEntity, input)
}

func (c *userApiImpl) DeleteUser(userId *models.UserId) (err common.Error) {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return
	}

	return userService.DeleteUser(userEntity)
}
