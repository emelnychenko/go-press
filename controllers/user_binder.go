package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

func BindUserController(e *echo.Echo, c *UserController) {
	e.GET("/users", c.ListUsers)
	e.GET(fmt.Sprintf("/user/:%s", UserId), c.GetUser)

	e.POST("/users", c.CreateUser)
	e.POST(fmt.Sprintf("/user/:%s", UserId), c.UpdateUser)
	e.POST(fmt.Sprintf("/user/:%s/identity", UserId), c.ChangeUserIdentity)
	e.POST(fmt.Sprintf("/user/:%s/password", UserId), c.ChangeUserPassword)

	e.DELETE(fmt.Sprintf("/user/:%s", UserId), c.DeleteUser)
}


