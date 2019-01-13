package main

import (
	"github.com/emelnychenko/go-press/aggregators"
	"github.com/emelnychenko/go-press/apis"
	"github.com/emelnychenko/go-press/builders"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/controllers"
	"github.com/emelnychenko/go-press/dispatchers"
	echoImpl "github.com/emelnychenko/go-press/echo"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/factories"
	"github.com/emelnychenko/go-press/hashers"
	"github.com/emelnychenko/go-press/helpers"
	"github.com/emelnychenko/go-press/jobs"
	"github.com/emelnychenko/go-press/normalizers"
	"github.com/emelnychenko/go-press/paginators"
	"github.com/emelnychenko/go-press/parameters"
	"github.com/emelnychenko/go-press/providers"
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
	db.AutoMigrate(new(entities.ChannelEntity))
	db.AutoMigrate(new(entities.CategoryEntity))
	db.AutoMigrate(new(entities.CategoryXrefEntity))
	db.AutoMigrate(new(entities.TagEntity))
	db.AutoMigrate(new(entities.TagXrefEntity))
	db.AutoMigrate(new(entities.CommentEntity))
	db.AutoMigrate(new(entities.BannerEntity))
	db.AutoMigrate(new(entities.PollEntity))
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
	_ = container.Provide(hashers.NewBCryptHasher)
	_ = container.Provide(helpers.NewUserEchoHelper)
	_ = container.Provide(helpers.NewPostHttpHelper)
	_ = container.Provide(helpers.NewFileHttpHelper)
	_ = container.Provide(helpers.NewCommentHttpHelper)
	_ = container.Provide(helpers.NewTagHttpHelper)
	_ = container.Provide(helpers.NewCategoryHttpHelper)
	_ = container.Provide(helpers.NewChannelHttpHelper)
	_ = container.Provide(helpers.NewBannerHttpHelper)
	_ = container.Provide(helpers.NewPollHttpHelper)
	_ = container.Provide(paginators.NewDbPaginator)
	_ = container.Provide(jobs.NewPostPublisherJob)
	_ = container.Provide(jobs.NewCategoryEdgesBuilderJob)
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
	_ = container.Provide(factories.NewPostCategoryEventFactory)
	_ = container.Provide(factories.NewPostTagEventFactory)
	_ = container.Provide(factories.NewPostModelFactory)
	_ = container.Provide(factories.NewFileEntityFactory)
	_ = container.Provide(factories.NewFileEventFactory)
	_ = container.Provide(factories.NewFileModelFactory)
	_ = container.Provide(factories.NewChannelEntityFactory)
	_ = container.Provide(factories.NewChannelModelFactory)
	_ = container.Provide(factories.NewChannelEventFactory)
	_ = container.Provide(factories.NewCategoryEntityFactory)
	_ = container.Provide(factories.NewCategoryModelFactory)
	_ = container.Provide(factories.NewCategoryEventFactory)
	_ = container.Provide(factories.NewTagEntityFactory)
	_ = container.Provide(factories.NewTagModelFactory)
	_ = container.Provide(factories.NewTagEventFactory)
	_ = container.Provide(factories.NewCommentEntityFactory)
	_ = container.Provide(factories.NewCommentModelFactory)
	_ = container.Provide(factories.NewCommentEventFactory)
	_ = container.Provide(factories.NewBannerEntityFactory)
	_ = container.Provide(factories.NewBannerModelFactory)
	_ = container.Provide(factories.NewBannerEventFactory)
	_ = container.Provide(factories.NewPollEntityFactory)
	_ = container.Provide(factories.NewPollModelFactory)
	_ = container.Provide(factories.NewPollEventFactory)
	_ = container.Provide(providers.NewAwsS3StorageProvider)
	_ = container.Provide(strategies.NewFilePathStrategy)
	_ = container.Provide(normalizers.NewPostNormalizer)
	_ = container.Provide(resolvers.NewSubjectResolver)
	_ = container.Provide(builders.NewCategoryTreeBuilder)
	_ = container.Provide(builders.NewCategoryNestedSetBuilder)
	_ = container.Provide(repositories.NewUserRepository)
	_ = container.Provide(repositories.NewPostRepository)
	_ = container.Provide(repositories.NewFileRepository)
	_ = container.Provide(repositories.NewChannelRepository)
	_ = container.Provide(repositories.NewCategoryRepository)
	_ = container.Provide(repositories.NewTagRepository)
	_ = container.Provide(repositories.NewCommentRepository)
	_ = container.Provide(repositories.NewBannerRepository)
	_ = container.Provide(repositories.NewPollRepository)
	_ = container.Provide(services.NewUserService)
	_ = container.Provide(services.NewUserPictureService)
	_ = container.Provide(services.NewFileService)
	_ = container.Provide(services.NewPostService)
	_ = container.Provide(services.NewPostAuthorService)
	_ = container.Provide(services.NewPostPictureService)
	_ = container.Provide(services.NewPostVideoService)
	_ = container.Provide(services.NewPostCategoryService)
	_ = container.Provide(services.NewPostTagService)
	_ = container.Provide(services.NewChannelService)
	_ = container.Provide(services.NewCategoryService)
	_ = container.Provide(services.NewTagService)
	_ = container.Provide(services.NewCommentService)
	_ = container.Provide(services.NewBannerService)
	_ = container.Provide(services.NewPollService)
	_ = container.Provide(aggregators.NewUserAggregator)
	_ = container.Provide(aggregators.NewPostAggregator)
	_ = container.Provide(aggregators.NewFileAggregator)
	_ = container.Provide(aggregators.NewCommentAggregator)
	_ = container.Provide(aggregators.NewCategoryAggregator)
	_ = container.Provide(aggregators.NewTagAggregator)
	_ = container.Provide(aggregators.NewChannelAggregator)
	_ = container.Provide(aggregators.NewBannerAggregator)
	_ = container.Provide(aggregators.NewPollAggregator)
	_ = container.Provide(apis.NewUserApi)
	_ = container.Provide(apis.NewUserPictureApi)
	_ = container.Provide(apis.NewPostApi)
	_ = container.Provide(apis.NewPostAuthorApi)
	_ = container.Provide(apis.NewPostPictureApi)
	_ = container.Provide(apis.NewPostVideoApi)
	_ = container.Provide(apis.NewPostCategoryApi)
	_ = container.Provide(apis.NewPostTagApi)
	_ = container.Provide(apis.NewFileApi)
	_ = container.Provide(apis.NewCommentApi)
	_ = container.Provide(apis.NewCategoryApi)
	_ = container.Provide(apis.NewTagApi)
	_ = container.Provide(apis.NewChannelApi)
	_ = container.Provide(apis.NewBannerApi)
	_ = container.Provide(apis.NewPollApi)
	_ = container.Provide(controllers.NewUserController)
	_ = container.Provide(controllers.NewUserPictureController)
	_ = container.Provide(controllers.NewPostController)
	_ = container.Provide(controllers.NewPostAuthorController)
	_ = container.Provide(controllers.NewPostPictureController)
	_ = container.Provide(controllers.NewPostVideoController)
	_ = container.Provide(controllers.NewPostCategoryController)
	_ = container.Provide(controllers.NewPostTagController)
	_ = container.Provide(controllers.NewFileController)
	_ = container.Provide(controllers.NewCommentController)
	_ = container.Provide(controllers.NewCategoryController)
	_ = container.Provide(controllers.NewTagController)
	_ = container.Provide(controllers.NewChannelController)
	_ = container.Provide(controllers.NewBannerController)
	_ = container.Provide(controllers.NewPollController)
	return
}

func BindRoutes(container *dig.Container) (err error) {
	binders := []interface{}{
		controllers.BindUserRoutes,
		controllers.BindUserPictureRoutes,
		controllers.BindPostRoutes,
		controllers.BindPostAuthorRoutes,
		controllers.BindPostPictureRoutes,
		controllers.BindPostVideoRoutes,
		controllers.BindPostCategoryRoutes,
		controllers.BindPostTagRoutes,
		controllers.BindFileRoutes,
		controllers.BindCommentRoutes,
		controllers.BindCategoryRoutes,
		controllers.BindTagRoutes,
		controllers.BindChannelRoutes,
		controllers.BindBannerRoutes,
		controllers.BindPollRoutes,
	}

	for _, binder := range binders {
		if err = container.Invoke(binder); nil != err {
			return
		}
	}

	return
}

func main() {
	container := BuildContainer()

	if err := BindRoutes(container); err != nil {
		panic(err)
	}

	if err := container.Invoke(StartWorkers); err != nil {
		panic(err)
	}

	err := container.Invoke(func(e *echo.Echo, db *gorm.DB) {
		defer db.Close()
		db.LogMode(true)
		e.Logger.Fatal(e.Start(":1323"))
	})

	if err != nil {
		panic(err)
	}
}
