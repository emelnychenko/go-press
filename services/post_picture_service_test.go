package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostPictureService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostPictureService", func(t *testing.T) {
		postRepository := mocks.NewMockPostRepository(ctrl)
		postPictureService, isPostPictureService := NewPostPictureService(postRepository).(*postPictureServiceImpl)

		assert.True(t, isPostPictureService)
		assert.Equal(t, postRepository, postPictureService.postRepository)
	})

	t.Run("ChangePostPicture", func(t *testing.T) {
		postPictureId := new(models.FileId)
		postPicture := &entities.FileEntity{Id: postPictureId}
		postEntity := new(entities.PostEntity)

		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().SavePost(postEntity).Return(nil)

		postPictureService := &postPictureServiceImpl{postRepository: postRepository}

		assert.Nil(t, postPictureService.ChangePostPicture(postEntity, postPicture))
		assert.Equal(t, postPictureId, postEntity.PictureId)
	})

	t.Run("RemovePostPicture", func(t *testing.T) {
		postPictureId := new(models.FileId)
		postEntity := &entities.PostEntity{PictureId: postPictureId}

		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().SavePost(postEntity).Return(nil)

		postPictureService := &postPictureServiceImpl{postRepository: postRepository}

		assert.Nil(t, postPictureService.RemovePostPicture(postEntity))
		assert.Nil(t, postEntity.PictureId)
	})
}
