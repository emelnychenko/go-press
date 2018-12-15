package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
	"github.com/labstack/echo"
	"net/http"
)

const UserId = "userId"

type UserController struct {
	api contracts.UserApi
}

func NewUserController(api contracts.UserApi) (c *UserController) {
	return &UserController{api}
}

func parseUserId(context echo.Context) (*models.UserId, error) {
	return models.ParseModelId(context.Param(UserId))
}

func (c *UserController) ListUsers(context echo.Context) error {
	users, err := c.api.ListUsers()

	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, users)
}

func (c *UserController) CreateUser(context echo.Context) error {
	data := new(models.UserCreate)

	if err := context.Bind(data); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	user, err := c.api.CreateUser(data)

	if err != nil {
		return context.JSON(err.Code(), err)
	}

	return context.JSON(http.StatusOK, user)
}

func (c *UserController) GetUser(context echo.Context) error {
	userId, err := parseUserId(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	user, err2 := c.api.GetUser(userId)

	if err2 != nil {
		return context.JSON(err2.Code(), err2)
	}

	return context.JSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(context echo.Context) error {
	userId, err := parseUserId(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	data := new(models.UserUpdate)

	if err := context.Bind(data); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	if err2 := c.api.UpdateUser(userId, data); err2 != nil {
		return context.JSON(err2.Code(), err2)
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *UserController) ChangeUserIdentity(context echo.Context) error {
	userId, err := parseUserId(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	data := new(models.UserChangeIdentity)

	if err := context.Bind(data); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	if err := c.api.ChangeUserIdentity(userId, data); err != nil {
		return context.JSON(err.Code(), err)
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *UserController) ChangeUserPassword(context echo.Context) error {
	userId, err := parseUserId(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	data := new(models.UserChangePassword)

	if err := context.Bind(data); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	if err := c.api.ChangeUserPassword(userId, data); err != nil {
		return context.JSON(err.Code(), err)
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *UserController) DeleteUser(context echo.Context) error {
	userId, err := parseUserId(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.api.DeleteUser(userId); err != nil {
		return context.JSON(err.Code(), err.Error())
	}

	return context.JSON(http.StatusOK, nil)
}
