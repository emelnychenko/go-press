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

func TestNewPostPictureApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostPictureApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		postPictureEventFactory := mocks.NewMockPostPictureEventFactory(ctrl)
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		postService := mocks.NewMockPostService(ctrl)
		fileService := mocks.NewMockFileService(ctrl)
		postPictureService := mocks.NewMockPostPictureService(ctrl)

		postPictureApi, isPostPictureApi := NewPostPictureApi(
			eventDispatcher, postPictureEventFactory, contentTypeValidator, postService, fileService, postPictureService,
		).(*postPictureApiImpl)

		assert.True(t, isPostPictureApi)
		assert.Equal(t, eventDispatcher, postPictureApi.eventDispatcher)
		assert.Equal(t, postPictureEventFactory, postPictureApi.postPictureEventFactory)
		assert.Equal(t, contentTypeValidator, postPictureApi.contentTypeValidator)
		assert.Equal(t, postService, postPictureApi.postService)
		assert.Equal(t, fileService, postPictureApi.fileService)
		assert.Equal(t, postPictureService, postPictureApi.postPictureService)
	})

	t.Run("ChangePostPicture", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		contentType := "image/png"
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		contentTypeValidator.EXPECT().ValidateImage(contentType).Return(nil)

		postPictureId := new(models.FileId)
		postPictureEntity := &entities.FileEntity{Type: contentType}
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(postPictureId).Return(postPictureEntity, nil)

		postPictureEvent := new(events.PostPictureEvent)
		postPictureEventFactory := mocks.NewMockPostPictureEventFactory(ctrl)
		postPictureEventFactory.EXPECT().CreatePostPictureChangedEvent(postEntity, postPictureEntity).Return(postPictureEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postPictureEvent)

		postPictureService := mocks.NewMockPostPictureService(ctrl)
		postPictureService.EXPECT().ChangePostPicture(postEntity, postPictureEntity).Return(nil)

		postPictureApi := &postPictureApiImpl{
			eventDispatcher:         eventDispatcher,
			postPictureEventFactory: postPictureEventFactory,
			contentTypeValidator:    contentTypeValidator,
			postService:             postService,
			fileService:             fileService,
			postPictureService:      postPictureService,
		}
		assert.Nil(t, postPictureApi.ChangePostPicture(postId, postPictureId))
	})

	t.Run("ChangePostPicture:GetPostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postPictureId := new(models.FileId)
		postPictureApi := &postPictureApiImpl{postService: postService}
		err := postPictureApi.ChangePostPicture(postId, postPictureId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostPicture:GetFileError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postPictureId := new(models.FileId)
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(postPictureId).Return(nil, systemErr)

		postPictureApi := &postPictureApiImpl{
			postService: postService,
			fileService: fileService,
		}
		err := postPictureApi.ChangePostPicture(postId, postPictureId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostPicture:ValidateImageError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		contentType := "audio/mp3"
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		contentTypeValidator.EXPECT().ValidateImage(contentType).Return(systemErr)

		postPictureId := new(models.FileId)
		postPictureEntity := &entities.FileEntity{Type: contentType}
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(postPictureId).Return(postPictureEntity, nil)

		postPictureApi := &postPictureApiImpl{
			contentTypeValidator: contentTypeValidator,
			postService:          postService,
			fileService:          fileService,
		}

		err := postPictureApi.ChangePostPicture(postId, postPictureId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostPicture:ChangePostPictureError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		contentType := "image/png"
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		contentTypeValidator.EXPECT().ValidateImage(contentType).Return(nil)

		postPictureId := new(models.FileId)
		postPictureEntity := &entities.FileEntity{Type: contentType}
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(postPictureId).Return(postPictureEntity, nil)

		postPictureService := mocks.NewMockPostPictureService(ctrl)
		postPictureService.EXPECT().ChangePostPicture(postEntity, postPictureEntity).Return(systemErr)

		postPictureApi := &postPictureApiImpl{
			contentTypeValidator: contentTypeValidator,
			postService:          postService,
			fileService:          fileService,
			postPictureService:   postPictureService,
		}

		err := postPictureApi.ChangePostPicture(postId, postPictureId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostPicture", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postPictureEvent := new(events.PostPictureEvent)
		postPictureEventFactory := mocks.NewMockPostPictureEventFactory(ctrl)
		postPictureEventFactory.EXPECT().CreatePostPictureRemovedEvent(postEntity).Return(postPictureEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postPictureEvent)

		postPictureService := mocks.NewMockPostPictureService(ctrl)
		postPictureService.EXPECT().RemovePostPicture(postEntity).Return(nil)

		postPictureApi := &postPictureApiImpl{
			eventDispatcher:         eventDispatcher,
			postPictureEventFactory: postPictureEventFactory,
			postService:             postService,
			postPictureService:      postPictureService,
		}
		assert.Nil(t, postPictureApi.RemovePostPicture(postId))
	})

	t.Run("RemovePostPicture:GetPostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postPictureApi := &postPictureApiImpl{postService: postService}
		err := postPictureApi.RemovePostPicture(postId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostPicture:RemovePostPictureError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postPictureService := mocks.NewMockPostPictureService(ctrl)
		postPictureService.EXPECT().RemovePostPicture(postEntity).Return(systemErr)

		postPictureApi := &postPictureApiImpl{
			postService:        postService,
			postPictureService: postPictureService,
		}

		err := postPictureApi.RemovePostPicture(postId)
		assert.Equal(t, systemErr, err)
	})
}
