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

func TestPostApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postId := models.NewModelId()
	testErr := common.ServerError("err0")

	t.Run("ListPosts", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postApi := NewPostApi(postService, postAggregator)
		var postEntities []*entities.PostEntity
		var commonReply []*models.Post

		postService.EXPECT().ListPosts().Return(postEntities, nil)
		postAggregator.EXPECT().AggregateCollection(postEntities).Return(commonReply)
		posts, err := postApi.ListPosts()

		assert.Equal(t, commonReply, posts)
		assert.Nil(t, err)
	})

	t.Run("ListPosts:Error", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postApi := NewPostApi(postService, nil)

		postService.EXPECT().ListPosts().Return(nil, testErr)
		posts, err := postApi.ListPosts()

		assert.Nil(t, posts)
		assert.Error(t, err)
	})

	t.Run("GetPost", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postApi := NewPostApi(postService, postAggregator)
		var postEntity *entities.PostEntity
		var reply *models.Post

		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postAggregator.EXPECT().AggregateObject(postEntity).Return(reply)
		post, err := postApi.GetPost(postId)

		assert.Equal(t, reply, post)
		assert.Nil(t, err)
	})

	t.Run("GetPost:Error", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postApi := NewPostApi(postService, nil)

		postService.EXPECT().GetPost(postId).Return(nil, testErr)
		post, err := postApi.GetPost(postId)

		assert.Nil(t, post)
		assert.Error(t, err)
	})

	t.Run("CreatePost", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postApi := NewPostApi(postService, postAggregator)
		postAuthor := models.NewSystemUser()
		var postEntity *entities.PostEntity
		var data *models.PostCreate
		var reply *models.Post

		postService.EXPECT().CreatePost(postAuthor, data).Return(postEntity, nil)
		postAggregator.EXPECT().AggregateObject(postEntity).Return(reply)
		post, err := postApi.CreatePost(postAuthor, data)

		assert.Equal(t, reply, post)
		assert.Nil(t, err)
	})

	t.Run("CreatePost:Error", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postApi := NewPostApi(postService, nil)
		postAuthor := models.NewSystemUser()
		var data *models.PostCreate

		postService.EXPECT().CreatePost(postAuthor, data).Return(nil, testErr)
		post, err := postApi.CreatePost(postAuthor, data)

		assert.Nil(t, post)
		assert.Error(t, err)
	})

	t.Run("UpdatePost", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postApi := NewPostApi(postService, nil)
		var postEntity *entities.PostEntity
		var data *models.PostUpdate

		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().UpdatePost(postEntity, data).Return(nil)
		assert.Nil(t, postApi.UpdatePost(postId, data))
	})

	t.Run("UpdatePost:Error", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postApi := NewPostApi(postService, nil)
		var data *models.PostUpdate

		postService.EXPECT().GetPost(postId).Return(nil, testErr)
		assert.Error(t, postApi.UpdatePost(postId, data))
	})

	t.Run("ChangePostAuthor", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postApi := NewPostApi(postService, nil)
		systemUser := models.NewSystemUser()
		var postEntity *entities.PostEntity

		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().ChangePostAuthor(postEntity, systemUser).Return(nil)
		assert.Nil(t, postApi.ChangePostAuthor(postId, systemUser))
	})

	t.Run("ChangePostAuthor:Error", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postApi := NewPostApi(postService, nil)
		systemUser := models.NewSystemUser()

		postService.EXPECT().GetPost(postId).Return(nil, testErr)
		assert.Error(t, postApi.ChangePostAuthor(postId, systemUser))
	})

	t.Run("DeletePost", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postApi := NewPostApi(postService, nil)
		var postEntity *entities.PostEntity

		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().DeletePost(postEntity).Return(nil)
		assert.Nil(t, postApi.DeletePost(postId))
	})

	t.Run("DeletePost:Error", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postApi := NewPostApi(postService, nil)

		postService.EXPECT().GetPost(postId).Return(nil, testErr)
		assert.Error(t, postApi.DeletePost(postId))
	})
}