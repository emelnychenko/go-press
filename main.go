package main

import (
	"github.com/emelnychenko/go-press/aggregator"
	"github.com/emelnychenko/go-press/apis"
	"github.com/emelnychenko/go-press/controllers"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/hasher"
	"github.com/emelnychenko/go-press/repositories"
	"github.com/emelnychenko/go-press/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"go.uber.org/dig"
)

func NewServer() (e *echo.Echo, err error) {
	e = echo.New()
	e.HideBanner = true
	return
}

func ConnectDatabase() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", ":memory:")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(new(entities.UserEntity))
	db.AutoMigrate(new(entities.PostEntity))
	return
}

func BuildContainer() (container *dig.Container) {
	container = dig.New()

	_ = container.Provide(ConnectDatabase)
	_ = container.Provide(NewServer)
	_ = container.Provide(hasher.NewBCryptHasher)
	_ = container.Provide(repositories.NewUserRepository)
	_ = container.Provide(repositories.NewPostRepository)
	_ = container.Provide(services.NewUserService)
	_ = container.Provide(services.NewPostService)
	_ = container.Provide(aggregator.NewUserAggregator)
	_ = container.Provide(aggregator.NewPostAggregator)
	_ = container.Provide(apis.NewUserApi)
	_ = container.Provide(apis.NewPostApi)
	_ = container.Provide(controllers.NewUserController)
	_ = container.Provide(controllers.NewPostController)
	return
}

func BindControllers(
	echo *echo.Echo,
	userController *controllers.UserController,
	postController *controllers.PostController,
) {
	controllers.BindUserController(echo, userController)
	controllers.BindPostController(echo, postController)
}

func main() {
	container := BuildContainer()

	if err := container.Invoke(BindControllers); err != nil {
		panic(err)
	}

	err := container.Invoke(func(e *echo.Echo, db *gorm.DB) {
		defer db.Close()
		e.Logger.Fatal(e.Start(":1323"))
	})

	if err != nil {
		panic(err)
	}
}
