package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
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
		var postEntities []*entities.PostEntity
		var posts []*models.Post

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().ListPosts().Return(postEntities, nil)

		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postAggregator.EXPECT().AggregatePosts(postEntities).Return(posts)

		postApi := &postApiImpl{postService: postService, postAggregator: postAggregator}
		response, err := postApi.ListPosts()

		assert.Equal(t, posts, response)
		assert.Nil(t, err)
	})

	t.Run("ListPosts:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().ListPosts().Return(nil, systemErr)

		postApi := &postApiImpl{postService: postService}
		response, err := postApi.ListPosts()

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
		systemErr := common.NewUnknownError()

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
			eventDispatcher: eventDispatcher,
			postEventFactory: postEventFactory,
			postService: postService,
			postAggregator: postAggregator,
		}
		response, err := postApi.CreatePost(postAuthor, data)

		assert.Equal(t, post, response)
		assert.Nil(t, err)
	})

	t.Run("CreatePost:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

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
			eventDispatcher: eventDispatcher,
			postEventFactory: postEventFactory,
			postService: postService,
		}
		assert.Nil(t, postApi.UpdatePost(postId, data))
	})

	t.Run("UpdatePost:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		data := new(models.PostUpdate)
		postApi := &postApiImpl{postService: postService}
		assert.Equal(t, systemErr, postApi.UpdatePost(postId, data))
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
			eventDispatcher: eventDispatcher,
			postEventFactory: postEventFactory,
			postService: postService,
		}
		assert.Nil(t, postApi.DeletePost(postId))
	})

	t.Run("DeletePost:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postApi := &postApiImpl{postService: postService}
		assert.Equal(t, systemErr, postApi.DeletePost(postId))
	})
}
