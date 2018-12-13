package user_echo

import (
	"../user_contract"
	"../common"
	"../user_domain"
	"github.com/labstack/echo"
	"net/http"
)

const UserId = "userId"

type UserController struct {
	api user_contract.UserApi
}

func parseUserId(context echo.Context) (*user_domain.UserId, error) {
	return common.ParseModelId(context.Param(UserId))
}

func NewUserController(api user_contract.UserApi) (c *UserController) {
	c = &UserController{api}
	return c
}

func (c *UserController) ListUsers(context echo.Context) error {
	users, err := c.api.ListUsers()

	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, users)
}

func (c *UserController) CreateUser(context echo.Context) error {
	data := new(user_domain.UserCreate)

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

	data := new(user_domain.UserUpdate)

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

	data := new(user_domain.UserChangeIdentity)

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

	data := new(user_domain.UserChangePassword)

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
