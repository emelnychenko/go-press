package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostVideoService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostVideoService", func(t *testing.T) {
		postRepository := mocks.NewMockPostRepository(ctrl)
		postVideoService, isPostVideoService := NewPostVideoService(postRepository).(*postVideoServiceImpl)

		assert.True(t, isPostVideoService)
		assert.Equal(t, postRepository, postVideoService.postRepository)
	})

	t.Run("ChangePostVideo", func(t *testing.T) {
		postVideoId := new(models.FileId)
		postVideo := &entities.FileEntity{Id: postVideoId}
		postEntity := new(entities.PostEntity)

		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().SavePost(postEntity).Return(nil)

		postVideoService := &postVideoServiceImpl{postRepository: postRepository}

		assert.Nil(t, postVideoService.ChangePostVideo(postEntity, postVideo))
		assert.Equal(t, postVideoId, postEntity.VideoId)
	})

	t.Run("RemovePostVideo", func(t *testing.T) {
		postVideoId := new(models.FileId)
		postEntity := &entities.PostEntity{VideoId: postVideoId}

		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().SavePost(postEntity).Return(nil)

		postVideoService := &postVideoServiceImpl{postRepository: postRepository}

		assert.Nil(t, postVideoService.RemovePostVideo(postEntity))
		assert.Nil(t, postEntity.VideoId)
	})
}
