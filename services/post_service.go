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
		postEntityFactory contracts.PostEntityFactory
		postNormalizer    contracts.PostNormalizer
		postRepository    contracts.PostRepository
	}
)

func NewPostService(
	postEntityFactory contracts.PostEntityFactory,
	postNormalizer contracts.PostNormalizer,
	postRepository contracts.PostRepository,
) (postService contracts.PostService) {
	return &postServiceImpl{
		postEntityFactory,
		postNormalizer,
		postRepository,
	}
}

func (s *postServiceImpl) ListPosts(
	postPaginationQuery *models.PostPaginationQuery,
) (*models.PaginationResult, common.Error) {
	return s.postRepository.ListPosts(postPaginationQuery)
}

func (s *postServiceImpl) GetScheduledPosts() (postEntities []*entities.PostEntity, err common.Error) {
	return s.postRepository.GetScheduledPosts()
}

func (s *postServiceImpl) GetPost(postId *models.PostId) (*entities.PostEntity, common.Error) {
	return s.postRepository.GetPost(postId)
}

func (s *postServiceImpl) CreatePost(postAuthor common.Subject, data *models.PostCreate) (postEntity *entities.PostEntity, err common.Error) {
	postEntity = s.postEntityFactory.CreatePostEntity()
	postEntity.Title = data.Title
	postEntity.Description = data.Description
	postEntity.Content = data.Content
	postEntity.Status = data.Status
	postEntity.Privacy = data.Privacy
	postEntity.Published = data.Published
	postEntity.Views = data.Views
	postEntity.SetAuthor(postAuthor)

	s.postNormalizer.NormalizePostEntity(postEntity)
	err = s.postRepository.SavePost(postEntity)
	return
}

func (s *postServiceImpl) UpdatePost(postEntity *entities.PostEntity, data *models.PostUpdate) common.Error {
	postEntity.Title = data.Title
	postEntity.Description = data.Description
	postEntity.Content = data.Content
	postEntity.Status = data.Status
	postEntity.Privacy = data.Privacy
	postEntity.Published = data.Published
	postEntity.Views = data.Views

	updated := time.Now().UTC()
	postEntity.Updated = &updated

	s.postNormalizer.NormalizePostEntity(postEntity)
	return s.postRepository.SavePost(postEntity)
}

func (s *postServiceImpl) DeletePost(postEntity *entities.PostEntity) common.Error {
	return s.postRepository.RemovePost(postEntity)
}
