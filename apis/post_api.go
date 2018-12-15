package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	postApiImpl struct {
		postService    contracts.PostService
		postAggregator contracts.PostAggregator
	}
)

func NewPostApi(postService contracts.PostService, postAggregator contracts.PostAggregator) (postApi contracts.PostApi) {
	return &postApiImpl{postService, postAggregator}
}


func (c *postApiImpl) ListPosts() (posts []*models.Post, err common.Error) {
	postEntities, err := c.postService.ListPosts()

	if nil != err {
		return
	}

	posts = c.postAggregator.AggregateCollection(postEntities)
	return
}

func (c *postApiImpl) GetPost(postId *models.PostId) (post *models.Post, err common.Error) {
	postEntity, err := c.postService.GetPost(postId)

	if nil != err {
		return
	}

	post = c.postAggregator.AggregateObject(postEntity)
	return
}

func (c *postApiImpl) CreatePost(data *models.PostCreate) (post *models.Post, err common.Error) {
	postEntity, err := c.postService.CreatePost(data)

	if nil != err {
		return
	}

	post = c.postAggregator.AggregateObject(postEntity)
	return
}

func (c *postApiImpl) UpdatePost(postId *models.PostId, data *models.PostUpdate) (err common.Error) {
	postService := c.postService
	postEntity, err := postService.GetPost(postId)

	if nil != err {
		return
	}

	return postService.UpdatePost(postEntity, data)
}

func (c *postApiImpl) DeletePost(postId *models.PostId) (err common.Error) {
	postService := c.postService
	postEntity, err := postService.GetPost(postId)

	if nil != err {
		return
	}

	return postService.DeletePost(postEntity)
}


