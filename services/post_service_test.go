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

	postId := models.NewModelId()

	t.Run("ListPosts", func(t *testing.T) {
		postRepository := mocks.NewMockPostRepository(ctrl)
		postService := NewPostService(postRepository)
		var replies []*entities.PostEntity

		postRepository.EXPECT().ListPosts().Return(replies, nil)
		postEntities, err := postService.ListPosts()

		assert.Equal(t, replies, postEntities)
		assert.Nil(t, err)
	})

	t.Run("CreatePost", func(t *testing.T) {
		postRepository := mocks.NewMockPostRepository(ctrl)
		postService := NewPostService(postRepository)
		systemUser := models.NewSystemUser()
		data := &models.PostCreate{
			Title:       "0",
			Description: "1",
			Content:     "2",
			Status:      "3",
			Privacy:     "4",
			Views:       5,
			Published:   new(time.Time),
		}

		postRepository.EXPECT().SavePost(gomock.Any()).Return(nil)
		postEntity, err := postService.CreatePost(systemUser, data)

		assert.IsType(t, new(entities.PostEntity), postEntity)
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
		postRepository := mocks.NewMockPostRepository(ctrl)
		postService := NewPostService(postRepository)
		var reply *entities.PostEntity

		postRepository.EXPECT().GetPost(postId).Return(reply, nil)
		postEntity, err := postService.GetPost(postId)

		assert.Equal(t, reply, postEntity)
		assert.Nil(t, err)
	})

	t.Run("UpdatePost", func(t *testing.T) {
		postRepository := mocks.NewMockPostRepository(ctrl)
		postService := NewPostService(postRepository)
		data := &models.PostUpdate{
			Title:       "0",
			Description: "1",
			Content:     "2",
			Status:      "3",
			Privacy:     "4",
			Views:       5,
			Published:   new(time.Time),
		}
		postEntity := &entities.PostEntity{
			Id:          nil,
			Title:       "",
			Description: "",
			Content:     "",
			Status:      "",
			Privacy:     "",
			Views:       0,
			Published:   nil,
			Created:     nil,
			Updated:     nil,
		}

		postRepository.EXPECT().SavePost(postEntity).Return(nil)
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
		postRepository := mocks.NewMockPostRepository(ctrl)
		postService := NewPostService(postRepository)
		systemUser := models.NewSystemUser()
		postEntity := &entities.PostEntity{}

		postRepository.EXPECT().SavePost(postEntity).Return(nil)
		assert.Nil(t, postService.ChangePostAuthor(postEntity, systemUser))
	})

	t.Run("DeletePost", func(t *testing.T) {
		postRepository := mocks.NewMockPostRepository(ctrl)
		postService := NewPostService(postRepository)
		postEntity := entities.NewPostEntity()

		postRepository.EXPECT().RemovePost(postEntity).Return(nil)
		assert.Nil(t, postService.DeletePost(postEntity))
	})
}