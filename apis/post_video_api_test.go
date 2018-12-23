package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
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
		postService := mocks.NewMockPostService(ctrl)
		fileService := mocks.NewMockFileService(ctrl)
		postVideoService := mocks.NewMockPostVideoService(ctrl)

		postVideoApi, isPostVideoApi := NewPostVideoApi(
			postService, fileService, postVideoService,
		).(*postVideoApiImpl)

		assert.True(t, isPostVideoApi)
		assert.Equal(t, postService, postVideoApi.postService)
		assert.Equal(t, fileService, postVideoApi.fileService)
		assert.Equal(t, postVideoService, postVideoApi.postVideoService)
	})

	t.Run("ChangePostVideo", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postVideoId := new(models.FileId)
		postVideo := new(entities.FileEntity)
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(postVideoId).Return(postVideo, nil)

		postVideoService := mocks.NewMockPostVideoService(ctrl)
		postVideoService.EXPECT().ChangePostVideo(postEntity, postVideo).Return(nil)

		postVideoApi := &postVideoApiImpl{
			postService: postService,
			fileService: fileService,
			postVideoService: postVideoService,
		}
		assert.Nil(t, postVideoApi.ChangePostVideo(postId, postVideoId))
	})

	t.Run("ChangePostVideo:GetPostError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postVideoId := new(models.FileId)
		postVideoApi := &postVideoApiImpl{postService: postService}
		err := postVideoApi.ChangePostVideo(postId, postVideoId)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostVideo:GetFileError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

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

	t.Run("RemovePostVideo", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postVideoService := mocks.NewMockPostVideoService(ctrl)
		postVideoService.EXPECT().RemovePostVideo(postEntity).Return(nil)

		postVideoApi := &postVideoApiImpl{
			postService: postService,
			postVideoService: postVideoService,
		}
		assert.Nil(t, postVideoApi.RemovePostVideo(postId))
	})

	t.Run("RemovePostVideo:GetPostError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postVideoApi := &postVideoApiImpl{postService: postService}
		err := postVideoApi.RemovePostVideo(postId)
		assert.Equal(t, systemErr, err)
	})
}
