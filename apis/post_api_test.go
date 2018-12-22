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

	testErr := common.ServerError("err0")

	t.Run("NewPostApi", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postApi, isPostApi := NewPostApi(postService, postAggregator).(*postApiImpl)

		assert.True(t, isPostApi)
		assert.Equal(t, postService, postApi.postService)
		assert.Equal(t, postAggregator, postApi.postAggregator)
	})

	t.Run("ListPosts", func(t *testing.T) {
		var postEntities []*entities.PostEntity
		var posts []*models.Post

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().ListPosts().Return(postEntities, nil)

		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postAggregator.EXPECT().AggregatePosts(postEntities).Return(posts)

		postApi := &postApiImpl{postService: postService, postAggregator: postAggregator}
		response, err := postApi.ListPosts()

		assert.Equal(t, posts, response)
		assert.Nil(t, err)
	})

	t.Run("ListPosts:Error", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().ListPosts().Return(nil, testErr)

		postApi := &postApiImpl{postService: postService}
		response, err := postApi.ListPosts()

		assert.Nil(t, response)
		assert.Error(t, err)
	})

	t.Run("GetPost", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		post := new(models.Post)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postAggregator.EXPECT().AggregatePost(postEntity).Return(post)

		postApi := &postApiImpl{postService: postService, postAggregator: postAggregator}
		response, err := postApi.GetPost(postId)

		assert.Equal(t, post, response)
		assert.Nil(t, err)
	})

	t.Run("GetPost:Error", func(t *testing.T) {
		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, testErr)

		postApi := &postApiImpl{postService: postService}
		response, err := postApi.GetPost(postId)

		assert.Nil(t, response)
		assert.Error(t, err)
	})

	t.Run("CreatePost", func(t *testing.T) {
		postAuthor := models.NewSystemUser()
		postEntity := new(entities.PostEntity)
		post := new(models.Post)
		data := new(models.PostCreate)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().CreatePost(postAuthor, data).Return(postEntity, nil)

		postAggregator := mocks.NewMockPostAggregator(ctrl)
		postAggregator.EXPECT().AggregatePost(postEntity).Return(post)

		postApi := &postApiImpl{postService: postService, postAggregator: postAggregator}
		response, err := postApi.CreatePost(postAuthor, data)

		assert.Equal(t, post, response)
		assert.Nil(t, err)
	})

	t.Run("CreatePost:Error", func(t *testing.T) {
		postAuthor := new(models.SystemUser)
		data := new(models.PostCreate)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().CreatePost(postAuthor, data).Return(nil, testErr)

		postApi := &postApiImpl{postService: postService}
		response, err := postApi.CreatePost(postAuthor, data)

		assert.Nil(t, response)
		assert.Error(t, err)
	})

	t.Run("UpdatePost", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		data := new(models.PostUpdate)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().UpdatePost(postEntity, data).Return(nil)

		postApi := &postApiImpl{postService: postService}
		assert.Nil(t, postApi.UpdatePost(postId, data))
	})

	t.Run("UpdatePost:Error", func(t *testing.T) {
		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, testErr)

		data := new(models.PostUpdate)
		postApi := &postApiImpl{postService: postService}
		assert.Error(t, postApi.UpdatePost(postId, data))
	})

	t.Run("ChangePostAuthor", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		postAuthor := new(models.SystemUser)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().ChangePostAuthor(postEntity, postAuthor).Return(nil)

		postApi := &postApiImpl{postService: postService}
		assert.Nil(t, postApi.ChangePostAuthor(postId, postAuthor))
	})

	t.Run("ChangePostAuthor:Error", func(t *testing.T) {
		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, testErr)

		postAuthor := new(models.SystemUser)
		postApi := &postApiImpl{postService: postService}
		assert.Error(t, postApi.ChangePostAuthor(postId, postAuthor))
	})

	t.Run("DeletePost", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)
		postService.EXPECT().DeletePost(postEntity).Return(nil)

		postApi := &postApiImpl{postService: postService}
		assert.Nil(t, postApi.DeletePost(postId))
	})

	t.Run("DeletePost:Error", func(t *testing.T) {
		postId := new(models.PostId)
		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, testErr)

		postApi := &postApiImpl{postService: postService}
		assert.Error(t, postApi.DeletePost(postId))
	})
}
