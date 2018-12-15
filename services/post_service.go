package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"time"
)

type (
	postServiceImpl struct {
		postRepository contracts.PostRepository
	}
)

func NewPostService(postRepository contracts.PostRepository) (postService contracts.PostService) {
	return &postServiceImpl{postRepository}
}

func (c *postServiceImpl) ListPosts() ([]*entities.PostEntity, common.Error) {
	return c.postRepository.ListPosts()
}

func (c *postServiceImpl) GetPost(postId *models.PostId) (*entities.PostEntity, common.Error) {
	return c.postRepository.GetPost(postId)
}

func (c *postServiceImpl) CreatePost(data *models.PostCreate) (postEntity *entities.PostEntity, err common.Error) {
	postEntity = entities.NewPostEntity()
	postEntity.Title = data.Title
	postEntity.Description = data.Description
	postEntity.Content = data.Content
	postEntity.Status = data.Status
	postEntity.Privacy = data.Privacy
	postEntity.Published = data.Published
	postEntity.Views = data.Views

	err = c.postRepository.SavePost(postEntity)
	return
}

func (c *postServiceImpl) UpdatePost(postEntity *entities.PostEntity, data *models.PostUpdate) common.Error {
	postEntity.Title = data.Title
	postEntity.Description = data.Description
	postEntity.Content = data.Content
	postEntity.Status = data.Status
	postEntity.Privacy = data.Privacy
	postEntity.Published = data.Published
	postEntity.Views = data.Views

	updated := time.Now().UTC()
	postEntity.Updated = &updated

	return c.postRepository.SavePost(postEntity)
}

func (c *postServiceImpl) DeletePost(postEntity *entities.PostEntity) common.Error {
	return c.postRepository.RemovePost(postEntity)
}
