package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPostService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostService", func(t *testing.T) {
		postEntityFactory := mocks.NewMockPostEntityFactory(ctrl)
		postRepository := mocks.NewMockPostRepository(ctrl)

		postService, isPostService := NewPostService(postEntityFactory, postRepository).(*postServiceImpl)

		assert.True(t, isPostService)
		assert.Equal(t, postEntityFactory, postService.postEntityFactory)
		assert.Equal(t, postRepository, postService.postRepository)
	})

	t.Run("ListPosts", func(t *testing.T) {
		var postEntities []*entities.PostEntity
		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().ListPosts().Return(postEntities, nil)

		postService := &postServiceImpl{postRepository: postRepository}
		response, err := postService.ListPosts()

		assert.Equal(t, postEntities, response)
		assert.Nil(t, err)
	})

	t.Run("CreatePost", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEntityFactory := mocks.NewMockPostEntityFactory(ctrl)
		postEntityFactory.EXPECT().CreatePostEntity().Return(postEntity)

		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().SavePost(postEntity).Return(nil)

		postAuthor := models.NewSystemUser()
		data := &models.PostCreate{
			Title:       "0",
			Description: "1",
			Content:     "2",
			Status:      "3",
			Privacy:     "4",
			Views:       5,
			Published:   new(time.Time),
		}
		postService := &postServiceImpl{postEntityFactory: postEntityFactory, postRepository: postRepository}
		response, err := postService.CreatePost(postAuthor, data)

		assert.IsType(t, postEntity, response)
		assert.Nil(t, err)
		assert.Equal(t, data.Title, postEntity.Title)
		assert.Equal(t, data.Description, postEntity.Description)
		assert.Equal(t, data.Content, postEntity.Content)
		assert.Equal(t, data.Status, postEntity.Status)
		assert.Equal(t, data.Privacy, postEntity.Privacy)
		assert.Equal(t, data.Views, postEntity.Views)
		assert.Equal(t, data.Published, postEntity.Published)
	})

	t.Run("GetPost", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().GetPost(postId).Return(postEntity, nil)

		postService := &postServiceImpl{postRepository: postRepository}
		response, err := postService.GetPost(postId)

		assert.Equal(t, postEntity, response)
		assert.Nil(t, err)
	})

	t.Run("UpdatePost", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().SavePost(postEntity).Return(nil)

		data := &models.PostUpdate{
			Title:       "0",
			Description: "1",
			Content:     "2",
			Status:      "3",
			Privacy:     "4",
			Views:       5,
			Published:   new(time.Time),
		}
		postService := &postServiceImpl{postRepository: postRepository}
		assert.Nil(t, postService.UpdatePost(postEntity, data))

		assert.Equal(t, data.Title, postEntity.Title)
		assert.Equal(t, data.Description, postEntity.Description)
		assert.Equal(t, data.Content, postEntity.Content)
		assert.Equal(t, data.Status, postEntity.Status)
		assert.Equal(t, data.Privacy, postEntity.Privacy)
		assert.Equal(t, data.Views, postEntity.Views)
		assert.Equal(t, data.Published, postEntity.Published)
		assert.NotNil(t, postEntity.Updated)
	})

	t.Run("ChangePostAuthor", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().SavePost(postEntity).Return(nil)

		postAuthor := models.NewSystemUser()
		postService := &postServiceImpl{postRepository: postRepository}

		assert.Nil(t, postService.ChangePostAuthor(postEntity, postAuthor))
	})

	t.Run("DeletePost", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().RemovePost(postEntity).Return(nil)

		postService := &postServiceImpl{postRepository: postRepository}
		assert.Nil(t, postService.DeletePost(postEntity))
	})
}
