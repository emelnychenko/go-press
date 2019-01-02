package main

import (
	"github.com/emelnychenko/go-press/aggregator"
	"github.com/emelnychenko/go-press/apis"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/controllers"
	"github.com/emelnychenko/go-press/dispatchers"
	echoImpl "github.com/emelnychenko/go-press/echo"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/factories"
	"github.com/emelnychenko/go-press/hasher"
	"github.com/emelnychenko/go-press/helpers"
	"github.com/emelnychenko/go-press/normalizers"
	"github.com/emelnychenko/go-press/paginators"
	"github.com/emelnychenko/go-press/parameters"
	"github.com/emelnychenko/go-press/providers"
	"github.com/emelnychenko/go-press/jobs"
	"github.com/emelnychenko/go-press/repositories"
	"github.com/emelnychenko/go-press/resolvers"
	"github.com/emelnychenko/go-press/services"
	"github.com/emelnychenko/go-press/strategies"
	"github.com/emelnychenko/go-press/validators"
	"github.com/emelnychenko/go-press/workers"
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
	db, err = gorm.Open("sqlite3", "./runtime.db")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(new(entities.UserEntity))
	db.AutoMigrate(new(entities.PostEntity))
	db.AutoMigrate(new(entities.FileEntity))
	return
}

func StartWorkers(postWorker contracts.PostPublisherWorker) {
	go func() {
		if err := postWorker.Start(); nil != err {
			panic(err)
		}
	}()
}

func BuildContainer() (container *dig.Container) {
	container = dig.New()

	_ = container.Provide(dispatchers.NewEventDispatcher)
	_ = container.Provide(ConnectDatabase)
	_ = container.Provide(NewServer)
	_ = container.Provide(echoImpl.NewRouter)
	_ = container.Provide(hasher.NewBCryptHasher)
	_ = container.Provide(helpers.NewUserEchoHelper)
	_ = container.Provide(helpers.NewPostHttpHelper)
	_ = container.Provide(helpers.NewFileHttpHelper)
	_ = container.Provide(paginators.NewDbPaginator)
	_ = container.Provide(jobs.NewPostPublisherJob)
	_ = container.Provide(workers.NewPostPublisherWorker)
	_ = container.Provide(validators.NewModelValidator)
	_ = container.Provide(validators.NewContentTypeValidator)
	_ = container.Provide(validators.NewPostStatusValidator)
	_ = container.Provide(parameters.NewAwsS3Parameters)
	_ = container.Provide(factories.NewAwsS3WriterProxyFactory)
	_ = container.Provide(factories.NewAwsS3Factory)
	_ = container.Provide(factories.NewAwsS3UploaderFactory)
	_ = container.Provide(factories.NewAwsS3DownloaderFactory)
	_ = container.Provide(factories.NewPaginationModelFactory)
	_ = container.Provide(factories.NewUserEntityFactory)
	_ = container.Provide(factories.NewUserEventFactory)
	_ = container.Provide(factories.NewUserPictureEventFactory)
	_ = container.Provide(factories.NewUserModelFactory)
	_ = container.Provide(factories.NewPostEntityFactory)
	_ = container.Provide(factories.NewPostEventFactory)
	_ = container.Provide(factories.NewPostPictureEventFactory)
	_ = container.Provide(factories.NewPostVideoEventFactory)
	_ = container.Provide(factories.NewPostAuthorEventFactory)
	_ = container.Provide(factories.NewPostModelFactory)
	_ = container.Provide(factories.NewFileEntityFactory)
	_ = container.Provide(factories.NewFileEventFactory)
	_ = container.Provide(factories.NewFileModelFactory)
	_ = container.Provide(providers.NewAwsS3StorageProvider)
	_ = container.Provide(strategies.NewFilePathStrategy)
	_ = container.Provide(normalizers.NewPostNormalizer)
	_ = container.Provide(resolvers.NewSubjectResolver)
	_ = container.Provide(repositories.NewUserRepository)
	_ = container.Provide(repositories.NewPostRepository)
	_ = container.Provide(repositories.NewFileRepository)
	_ = container.Provide(services.NewUserService)
	_ = container.Provide(services.NewUserPictureService)
	_ = container.Provide(services.NewFileService)
	_ = container.Provide(services.NewPostService)
	_ = container.Provide(services.NewPostAuthorService)
	_ = container.Provide(services.NewPostPictureService)
	_ = container.Provide(services.NewPostVideoService)
	_ = container.Provide(aggregator.NewUserAggregator)
	_ = container.Provide(aggregator.NewPostAggregator)
	_ = container.Provide(aggregator.NewFileAggregator)
	_ = container.Provide(apis.NewUserApi)
	_ = container.Provide(apis.NewUserPictureApi)
	_ = container.Provide(apis.NewPostApi)
	_ = container.Provide(apis.NewPostAuthorApi)
	_ = container.Provide(apis.NewPostPictureApi)
	_ = container.Provide(apis.NewPostVideoApi)
	_ = container.Provide(apis.NewFileApi)
	_ = container.Provide(controllers.NewUserController)
	_ = container.Provide(controllers.NewUserPictureController)
	_ = container.Provide(controllers.NewPostController)
	_ = container.Provide(controllers.NewPostAuthorController)
	_ = container.Provide(controllers.NewPostPictureController)
	_ = container.Provide(controllers.NewPostVideoController)
	_ = container.Provide(controllers.NewFileController)
	return
}

func BindRoutes(
	router contracts.Router,
	userController contracts.UserController,
	userPictureController contracts.UserPictureController,
	postController contracts.PostController,
	postAuthorController contracts.PostAuthorController,
	postPictureController contracts.PostPictureController,
	PostVideoController contracts.PostVideoController,
	fileController contracts.FileController,
) {
	controllers.BindUserRoutes(router, userController)
	controllers.BindUserPictureRoutes(router, userPictureController)
	controllers.BindPostRoutes(router, postController)
	controllers.BindPostAuthorRoutes(router, postAuthorController)
	controllers.BindPostPictureRoutes(router, postPictureController)
	controllers.BindPostVideoRoutes(router, PostVideoController)
	controllers.BindFileRoutes(router, fileController)
}

func main() {
	container := BuildContainer()

	if err := container.Invoke(BindRoutes); err != nil {
		panic(err)
	}

	if err := container.Invoke(StartWorkers); err != nil {
		panic(err)
	}

	err := container.Invoke(func(e *echo.Echo, db *gorm.DB, worker contracts.PostPublisherWorker) {
		defer db.Close()
		db.LogMode(true)
		e.Logger.Fatal(e.Start(":1323"))
	})

	if err != nil {
		panic(err)
	}
}
