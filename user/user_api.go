package user

import (
	"../common"
	"../user_contract"
	"../user_domain"
)

type (
	userApiImpl struct {
		userService    user_contract.UserService
		userAggregator user_contract.UserAggregator
	}
)

func NewUserApi(userService user_contract.UserService, aggregator user_contract.UserAggregator) *userApiImpl {
	return &userApiImpl{userService: userService, userAggregator: aggregator}
}

func (c *userApiImpl) ListUsers() ([]*user_domain.User, common.Error) {
	userEntities, err := c.userService.ListUsers()

	if nil != err {
		return nil, err
	}

	return c.userAggregator.AggregateCollection(userEntities), nil
}

func (c *userApiImpl) CreateUser(data *user_domain.UserCreate) (*user_domain.User, common.Error) {
	userEntity, err := c.userService.CreateUser(data)

	if nil != err {
		return nil, err
	}

	return c.userAggregator.AggregateObject(userEntity), nil
}

func (c *userApiImpl) GetUser(userId *user_domain.UserId) (*user_domain.User, common.Error) {
	userEntity, err := c.userService.GetUser(userId)

	if nil != err {
		return nil, err
	}

	return c.userAggregator.AggregateObject(userEntity), nil
}

func (c *userApiImpl) UpdateUser(userId *user_domain.UserId, data *user_domain.UserUpdate) common.Error {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return err
	}

	return userService.UpdateUser(userEntity, data)
}

func (c *userApiImpl) VerifyUser(userId *user_domain.UserId) common.Error {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return err
	}

	return userService.VerifyUser(userEntity)
}

func (c *userApiImpl) ChangeUserIdentity(userId *user_domain.UserId, data *user_domain.UserChangeIdentity) common.Error  {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return err
	}

	return userService.ChangeUserIdentity(userEntity, data)
}

func (c *userApiImpl) ChangeUserPassword(userId *user_domain.UserId, input *user_domain.UserChangePassword) common.Error  {
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

func (c *userApiImpl) DeleteUser(userId *user_domain.UserId) common.Error {
	userService := c.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return err
	}

	return userService.DeleteUser(userEntity)
}