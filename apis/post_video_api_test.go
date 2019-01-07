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

func TestNewPostVideoApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostVideoApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		postVideoEventFactory := mocks.NewMockPostVideoEventFactory(ctrl)
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		postService := mocks.NewMockPostService(ctrl)
		fileService := mocks.NewMockFileService(ctrl)
		postVideoService := mocks.NewMockPostVideoService(ctrl)

		postVideoApi, isPostVideoApi := NewPostVideoApi(
			eventDispatcher, postVideoEventFactory, contentTypeValidator, postService, fileService, postVideoService,
		).(*postVideoApiImpl)

		assert.True(t, isPostVideoApi)
		assert.Equal(t, eventDispatcher, postVideoApi.eventDispatcher)
		assert.Equal(t, postVideoEventFactory, postVideoApi.postVideoEventFactory)
		assert.Equal(t, contentTypeValidator, postVideoApi.contentTypeValidator)
		assert.Equal(t, postService, postVideoApi.postService)
		assert.Equal(t, fileService, postVideoApi.fileService)
		assert.Equal(t, postVideoService, postVideoApi.postVideoService)
	})

	t.Run("ChangePostVideo", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		contentType := "video/mp4"
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		contentTypeValidator.EXPECT().ValidateVideo(contentType).Return(nil)

		postVideoId := new(models.FileId)
		postVideoEntity := &entities.FileEntity{Type: contentType}
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(postVideoId).Return(postVideoEntity, nil)

		postVideoEvent := new(events.PostVideoEvent)
		postVideoEventFactory := mocks.NewMockPostVideoEventFactory(ctrl)
		postVideoEventFactory.EXPECT().CreatePostVideoChangedEvent(postEntity, postVideoEntity).Return(postVideoEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postVideoEvent)

		postVideoService := mocks.NewMockPostVideoService(ctrl)
		postVideoService.EXPECT().ChangePostVideo(postEntity, postVideoEntity).Return(nil)

		postVideoApi := &postVideoApiImpl{
			eventDispatcher:       eventDispatcher,
			postVideoEventFactory: postVideoEventFactory,
			contentTypeValidator:  contentTypeValidator,
			postService:           postService,
			fileService:           fileService,
			postVideoService:      postVideoService,
		}
		assert.Nil(t, postVideoApi.ChangePostVideo(postId, postVideoId))
	})

	t.Run("ChangePostVideo:GetPostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postVideoId := new(models.FileId)
		postVideoApi := &postVideoApiImpl{postService: postService}
		err := postVideoApi.ChangePostVideo(postId, postVideoId)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostVideo:GetFileError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postVideoId := new(models.FileId)
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(postVideoId).Return(nil, systemErr)

		postVideoApi := &postVideoApiImpl{
			postService: postService,
			fileService: fileService,
		}
		err := postVideoApi.ChangePostVideo(postId, postVideoId)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostVideo:ValidateVideoError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		contentType := "audio/mp3"
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		contentTypeValidator.EXPECT().ValidateVideo(contentType).Return(systemErr)

		postVideoId := new(models.FileId)
		postVideoEntity := &entities.FileEntity{Type: contentType}
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(postVideoId).Return(postVideoEntity, nil)

		postVideoApi := &postVideoApiImpl{
			contentTypeValidator: contentTypeValidator,
			postService:          postService,
			fileService:          fileService,
		}

		err := postVideoApi.ChangePostVideo(postId, postVideoId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostVideo:ChangePostVideoError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		contentType := "video/mp4"
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		contentTypeValidator.EXPECT().ValidateVideo(contentType).Return(nil)

		postVideoId := new(models.FileId)
		postVideoEntity := &entities.FileEntity{Type: contentType}
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(postVideoId).Return(postVideoEntity, nil)

		postVideoService := mocks.NewMockPostVideoService(ctrl)
		postVideoService.EXPECT().ChangePostVideo(postEntity, postVideoEntity).Return(systemErr)

		postVideoApi := &postVideoApiImpl{
			contentTypeValidator: contentTypeValidator,
			postService:          postService,
			fileService:          fileService,
			postVideoService:     postVideoService,
		}

		err := postVideoApi.ChangePostVideo(postId, postVideoId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostVideo", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postVideoEvent := new(events.PostVideoEvent)
		postVideoEventFactory := mocks.NewMockPostVideoEventFactory(ctrl)
		postVideoEventFactory.EXPECT().CreatePostVideoRemovedEvent(postEntity).Return(postVideoEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postVideoEvent)

		postVideoService := mocks.NewMockPostVideoService(ctrl)
		postVideoService.EXPECT().RemovePostVideo(postEntity).Return(nil)

		postVideoApi := &postVideoApiImpl{
			eventDispatcher:       eventDispatcher,
			postVideoEventFactory: postVideoEventFactory,
			postService:           postService,
			postVideoService:      postVideoService,
		}
		assert.Nil(t, postVideoApi.RemovePostVideo(postId))
	})

	t.Run("RemovePostVideo:GetPostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postVideoApi := &postVideoApiImpl{postService: postService}
		err := postVideoApi.RemovePostVideo(postId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostVideo:RemovePostVideoError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postVideoService := mocks.NewMockPostVideoService(ctrl)
		postVideoService.EXPECT().RemovePostVideo(postEntity).Return(systemErr)

		postVideoApi := &postVideoApiImpl{
			postService:      postService,
			postVideoService: postVideoService,
		}

		err := postVideoApi.RemovePostVideo(postId)
		assert.Equal(t, systemErr, err)
	})
}
