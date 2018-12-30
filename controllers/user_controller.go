package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
)

type userControllerImpl struct {
	userHttpHelper   contracts.UserHttpHelper
	userModelFactory contracts.UserModelFactory
	userApi          contracts.UserApi
}

func NewUserController(
	userHttpHelper contracts.UserHttpHelper,
	userModelFactory contracts.UserModelFactory,
	userApi contracts.UserApi,
) (userController contracts.UserController) {
	return &userControllerImpl{userHttpHelper, userModelFactory, userApi}
}

func (c *userControllerImpl) ListUsers(httpContext contracts.HttpContext) (paginationResult interface{}, err common.Error) {
	userPaginationQuery := c.userModelFactory.CreateUserPaginationQuery()

	if err = httpContext.BindModel(userPaginationQuery.PaginationQuery); err != nil {
		return
	}

	if err = httpContext.BindModel(userPaginationQuery); err != nil {
		return
	}

	paginationResult, err = c.userApi.ListUsers(userPaginationQuery)
	return
}

func (c *userControllerImpl) GetUser(httpContext contracts.HttpContext) (user interface{}, err common.Error) {
	userId, err := c.userHttpHelper.ParseUserId(httpContext)

	if err != nil {
		return
	}

	user, err = c.userApi.GetUser(userId)
	return
}

func (c *userControllerImpl) CreateUser(httpContext contracts.HttpContext) (user interface{}, err common.Error) {
	data := c.userModelFactory.CreateUserCreate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	user, err = c.userApi.CreateUser(data)
	return
}

func (c *userControllerImpl) UpdateUser(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	userId, err := c.userHttpHelper.ParseUserId(httpContext)

	if err != nil {
		return
	}

	data := c.userModelFactory.CreateUserUpdate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	err = c.userApi.UpdateUser(userId, data)
	return
}

func (c *userControllerImpl) ChangeUserIdentity(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	userId, err := c.userHttpHelper.ParseUserId(httpContext)

	if err != nil {
		return
	}

	data := c.userModelFactory.CreateUserChangeIdentity()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	err = c.userApi.ChangeUserIdentity(userId, data)
	return
}

func (c *userControllerImpl) ChangeUserPassword(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	userId, err := c.userHttpHelper.ParseUserId(httpContext)

	if err != nil {
		return
	}

	data := c.userModelFactory.CreateUserChangePassword()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	err = c.userApi.ChangeUserPassword(userId, data)
	return
}

func (c *userControllerImpl) DeleteUser(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	userId, err := c.userHttpHelper.ParseUserId(httpContext)

	if err != nil {
		return
	}

	err = c.userApi.DeleteUser(userId)
	return
}

