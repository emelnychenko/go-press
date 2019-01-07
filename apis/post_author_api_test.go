package apis

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/events"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostAuthorApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostAuthorApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		postAuthorEventFactory := mocks.NewMockPostAuthorEventFactory(ctrl)
		postService := mocks.NewMockPostService(ctrl)
		userService := mocks.NewMockUserService(ctrl)
		postAuthorService := mocks.NewMockPostAuthorService(ctrl)

		postAuthorApi, isPostAuthorApi := NewPostAuthorApi(
			eventDispatcher, postAuthorEventFactory, postService, userService, postAuthorService,
		).(*postAuthorApiImpl)

		assert.True(t, isPostAuthorApi)
		assert.Equal(t, eventDispatcher, postAuthorApi.eventDispatcher)
		assert.Equal(t, postAuthorEventFactory, postAuthorApi.postAuthorEventFactory)
		assert.Equal(t, postService, postAuthorApi.postService)
		assert.Equal(t, userService, postAuthorApi.userService)
		assert.Equal(t, postAuthorService, postAuthorApi.postAuthorService)
	})

	t.Run("ChangePostAuthor", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postAuthorId := new(models.UserId)
		postAuthorEntity := new(entities.UserEntity)
		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(postAuthorId).Return(postAuthorEntity, nil)

		postAuthorEvent := new(events.PostAuthorEvent)
		postAuthorEventFactory := mocks.NewMockPostAuthorEventFactory(ctrl)
		postAuthorEventFactory.EXPECT().CreatePostAuthorChangedEvent(postEntity, postAuthorEntity).Return(postAuthorEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postAuthorEvent)

		postAuthorService := mocks.NewMockPostAuthorService(ctrl)
		postAuthorService.EXPECT().ChangePostAuthor(postEntity, postAuthorEntity).Return(nil)

		postAuthorApi := &postAuthorApiImpl{
			eventDispatcher:        eventDispatcher,
			postAuthorEventFactory: postAuthorEventFactory,
			postService:            postService,
			userService:            userService,
			postAuthorService:      postAuthorService,
		}
		assert.Nil(t, postAuthorApi.ChangePostAuthor(postId, postAuthorId))
	})

	t.Run("ChangePostAuthor:GetPostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postAuthorId := new(models.UserId)
		postAuthorApi := &postAuthorApiImpl{postService: postService}
		err := postAuthorApi.ChangePostAuthor(postId, postAuthorId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostAuthor:GetUserError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postAuthorId := new(models.UserId)
		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(postAuthorId).Return(nil, systemErr)

		postAuthorApi := &postAuthorApiImpl{
			postService: postService,
			userService: userService,
		}
		err := postAuthorApi.ChangePostAuthor(postId, postAuthorId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostAuthor:ChangePostAuthorError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postAuthorId := new(models.UserId)
		postAuthorEntity := new(entities.UserEntity)
		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(postAuthorId).Return(postAuthorEntity, nil)

		postAuthorService := mocks.NewMockPostAuthorService(ctrl)
		postAuthorService.EXPECT().ChangePostAuthor(postEntity, postAuthorEntity).Return(systemErr)

		postAuthorApi := &postAuthorApiImpl{
			postService:       postService,
			userService:       userService,
			postAuthorService: postAuthorService,
		}

		err := postAuthorApi.ChangePostAuthor(postId, postAuthorId)
		assert.Equal(t, systemErr, err)
	})
}
