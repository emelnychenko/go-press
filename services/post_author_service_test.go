package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostAuthorService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostAuthorService", func(t *testing.T) {
		postRepository := mocks.NewMockPostRepository(ctrl)
		postAuthorService, isPostAuthorService := NewPostAuthorService(postRepository).(*postAuthorServiceImpl)

		assert.True(t, isPostAuthorService)
		assert.Equal(t, postRepository, postAuthorService.postRepository)
	})

	t.Run("ChangePostAuthor", func(t *testing.T) {
		postAuthorId := new(models.UserId)
		postAuthorEntity := &entities.UserEntity{Id: postAuthorId}
		postEntity := new(entities.PostEntity)

		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().SavePost(postEntity).Return(nil)

		postAuthorService := &postAuthorServiceImpl{postRepository: postRepository}

		assert.Nil(t, postAuthorService.ChangePostAuthor(postEntity, postAuthorEntity))
		assert.Equal(t, postAuthorId, postEntity.AuthorId)
	})
}
