package user_echo

import (
	"../user"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Bind(echo *echo.Echo, db *gorm.DB) {
	userContainer := user.NewUserContainer(db)
	userController := NewUserController(userContainer.UserApi)

	echo.GET("/users", userController.ListUsers)
	echo.GET(fmt.Sprintf("/user/:%s", UserId), userController.GetUser)

	echo.POST("/users", userController.CreateUser)
	echo.POST(fmt.Sprintf("/user/:%s", UserId), userController.UpdateUser)
	echo.POST(fmt.Sprintf("/user/:%s/identity", UserId), userController.ChangeUserIdentity)
	echo.POST(fmt.Sprintf("/user/:%s/password", UserId), userController.ChangeUserPassword)

	echo.DELETE(fmt.Sprintf("/user/:%s", UserId), userController.DeleteUser)
}


