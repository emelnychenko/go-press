package apis

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/events"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		postEventFactory := mocks.NewMockPostEventFactory(ctrl)
		postService := mocks.NewMockPostService(ctrl)
		postAggregator := mocks.NewMockPostAggregator(ctrl)

		postApi, isPostApi := NewPostApi(
			eventDispatcher, postEventFactory, postService, postAggregator,
		).(*postApiImpl)

		assert.True(t, isPostApi)
		assert.Equal(t, eventDispatcher, postApi.eventDispatcher)
		assert.Equal(t, postEventFactory, postApi.postEventFactory)
		assert.Equal(t, postService, postApi.postService)
		assert.Equal(t, postAggregator, postApi.postAggregator)
	})

	t.Run("ListPosts", func(t *testing.T) {
		paginationQuery := new(models.PostPaginationQuery)
		entityPaginationResult := new(models.PaginationResult)
		paginationResult := new(models.PaginationResult)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().ListPosts(paginationQuery).Return(entityPaginationResult, nil)

		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postAggregator.EXPECT().AggregatePaginationResult(entityPaginationResult).Return(paginationResult)

		postApi := &postApiImpl{postService: postService, postAggregator: postAggregator}
		response, err := postApi.ListPosts(paginationQuery)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListPosts:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		paginationQuery := new(models.PostPaginationQuery)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().ListPosts(paginationQuery).Return(nil, systemErr)

		postApi := &postApiImpl{postService: postService}
		response, err := postApi.ListPosts(paginationQuery)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetPost", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		post := new(models.Post)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postAggregator.EXPECT().AggregatePost(postEntity).Return(post)

		postApi := &postApiImpl{postService: postService, postAggregator: postAggregator}
		response, err := postApi.GetPost(postId)

		assert.Equal(t, post, response)
		assert.Nil(t, err)
	})

	t.Run("GetPost:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postApi := &postApiImpl{postService: postService}
		response, err := postApi.GetPost(postId)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreatePost", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postAuthor := models.NewSystemUser()
		post := new(models.Post)
		data := new(models.PostCreate)

		postEvent := new(events.PostEvent)
		postEventFactory := mocks.NewMockPostEventFactory(ctrl)
		postEventFactory.EXPECT().CreatePostCreatedEvent(postEntity).Return(postEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postEvent)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().CreatePost(postAuthor, data).Return(postEntity, nil)

		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postAggregator.EXPECT().AggregatePost(postEntity).Return(post)

		postApi := &postApiImpl{
			eventDispatcher:  eventDispatcher,
			postEventFactory: postEventFactory,
			postService:      postService,
			postAggregator:   postAggregator,
		}
		response, err := postApi.CreatePost(postAuthor, data)

		assert.Equal(t, post, response)
		assert.Nil(t, err)
	})

	t.Run("CreatePost:PostPublishedEvent", func(t *testing.T) {
		postEntity := &entities.PostEntity{Status: enums.PostPublishedStatus}
		postAuthor := models.NewSystemUser()
		post := new(models.Post)
		data := new(models.PostCreate)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		postEventFactory := mocks.NewMockPostEventFactory(ctrl)

		postCreatedEvent := new(events.PostEvent)
		postEventFactory.EXPECT().CreatePostCreatedEvent(postEntity).Return(postCreatedEvent)
		eventDispatcher.EXPECT().Dispatch(postCreatedEvent)

		postPublishedEvent := new(events.PostEvent)
		postEventFactory.EXPECT().CreatePostPublishedEvent(postEntity).Return(postPublishedEvent)
		eventDispatcher.EXPECT().Dispatch(postPublishedEvent)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().CreatePost(postAuthor, data).Return(postEntity, nil)

		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postAggregator.EXPECT().AggregatePost(postEntity).Return(post)

		postApi := &postApiImpl{
			eventDispatcher:  eventDispatcher,
			postEventFactory: postEventFactory,
			postService:      postService,
			postAggregator:   postAggregator,
		}
		response, err := postApi.CreatePost(postAuthor, data)

		assert.Equal(t, post, response)
		assert.Nil(t, err)
	})

	t.Run("CreatePost:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postAuthor := new(models.SystemUser)
		data := new(models.PostCreate)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().CreatePost(postAuthor, data).Return(nil, systemErr)

		postApi := &postApiImpl{postService: postService}
		response, err := postApi.CreatePost(postAuthor, data)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdatePost", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		data := new(models.PostUpdate)

		postEvent := new(events.PostEvent)
		postEventFactory := mocks.NewMockPostEventFactory(ctrl)
		postEventFactory.EXPECT().CreatePostUpdatedEvent(postEntity).Return(postEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postEvent)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().UpdatePost(postEntity, data).Return(nil)

		postApi := &postApiImpl{
			eventDispatcher:  eventDispatcher,
			postEventFactory: postEventFactory,
			postService:      postService,
		}
		assert.Nil(t, postApi.UpdatePost(postId, data))
	})

	t.Run("UpdatePost:PostPublishedEvent", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		data := new(models.PostUpdate)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		postEventFactory := mocks.NewMockPostEventFactory(ctrl)

		postUpdatedEvent := new(events.PostEvent)
		postEventFactory.EXPECT().CreatePostUpdatedEvent(postEntity).Return(postUpdatedEvent)
		eventDispatcher.EXPECT().Dispatch(postUpdatedEvent)

		postPublishedEvent := new(events.PostEvent)
		postEventFactory.EXPECT().CreatePostPublishedEvent(postEntity).Return(postPublishedEvent)
		eventDispatcher.EXPECT().Dispatch(postPublishedEvent)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().UpdatePost(postEntity, data).Do(func(postEntity *entities.PostEntity, data *models.PostUpdate) {
			postEntity.Status = enums.PostPublishedStatus
		}).Return(nil)

		postApi := &postApiImpl{
			eventDispatcher:  eventDispatcher,
			postEventFactory: postEventFactory,
			postService:      postService,
		}
		assert.Nil(t, postApi.UpdatePost(postId, data))
	})

	t.Run("UpdatePost:PostConcealedEvent", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := &entities.PostEntity{Status: enums.PostPublishedStatus}
		data := new(models.PostUpdate)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		postEventFactory := mocks.NewMockPostEventFactory(ctrl)

		postUpdatedEvent := new(events.PostEvent)
		postEventFactory.EXPECT().CreatePostUpdatedEvent(postEntity).Return(postUpdatedEvent)
		eventDispatcher.EXPECT().Dispatch(postUpdatedEvent)

		postConcealedEvent := new(events.PostEvent)
		postEventFactory.EXPECT().CreatePostConcealedEvent(postEntity).Return(postConcealedEvent)
		eventDispatcher.EXPECT().Dispatch(postConcealedEvent)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().UpdatePost(postEntity, data).Do(func(postEntity *entities.PostEntity, data *models.PostUpdate) {
			postEntity.Status = enums.PostScheduledStatus
		}).Return(nil)

		postApi := &postApiImpl{
			eventDispatcher:  eventDispatcher,
			postEventFactory: postEventFactory,
			postService:      postService,
		}
		assert.Nil(t, postApi.UpdatePost(postId, data))
	})

	t.Run("UpdatePost:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		data := new(models.PostUpdate)
		postApi := &postApiImpl{postService: postService}
		assert.Equal(t, systemErr, postApi.UpdatePost(postId, data))
	})

	t.Run("UpdatePost:UpdatePostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		data := new(models.PostUpdate)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().UpdatePost(postEntity, data).Return(systemErr)

		postApi := &postApiImpl{
			postService: postService,
		}

		err := postApi.UpdatePost(postId, data)
		assert.Equal(t, systemErr, err)
	})

	t.Run("DeletePost", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)

		postEvent := new(events.PostEvent)
		postEventFactory := mocks.NewMockPostEventFactory(ctrl)
		postEventFactory.EXPECT().CreatePostDeletedEvent(postEntity).Return(postEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postEvent)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().DeletePost(postEntity).Return(nil)

		postApi := &postApiImpl{
			eventDispatcher:  eventDispatcher,
			postEventFactory: postEventFactory,
			postService:      postService,
		}
		assert.Nil(t, postApi.DeletePost(postId))
	})

	t.Run("DeletePost:PostConcealedEvent", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := &entities.PostEntity{Status: enums.PostPublishedStatus}

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		postEventFactory := mocks.NewMockPostEventFactory(ctrl)

		postDeletedEvent := new(events.PostEvent)
		postEventFactory.EXPECT().CreatePostDeletedEvent(postEntity).Return(postDeletedEvent)
		eventDispatcher.EXPECT().Dispatch(postDeletedEvent)

		postConcealedEvent := new(events.PostEvent)
		postEventFactory.EXPECT().CreatePostConcealedEvent(postEntity).Return(postConcealedEvent)
		eventDispatcher.EXPECT().Dispatch(postConcealedEvent)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().DeletePost(postEntity).Return(nil)

		postApi := &postApiImpl{
			eventDispatcher:  eventDispatcher,
			postEventFactory: postEventFactory,
			postService:      postService,
		}
		assert.Nil(t, postApi.DeletePost(postId))
	})

	t.Run("DeletePost:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postApi := &postApiImpl{postService: postService}
		assert.Equal(t, systemErr, postApi.DeletePost(postId))
	})

	t.Run("DeletePost:DeletePostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().DeletePost(postEntity).Return(systemErr)

		postApi := &postApiImpl{
			postService: postService,
		}

		err := postApi.DeletePost(postId)
		assert.Equal(t, systemErr, err)
	})
}
