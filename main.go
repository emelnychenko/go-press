package main

import (
	userDomain "./user_domain"
	userEcho "./user_echo"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	db, err := gorm.Open("sqlite3", ":memory:")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&userDomain.UserEntity{})

	userEcho.Bind(e, db)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":1323"))
}
