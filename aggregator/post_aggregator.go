package aggregator

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type postAggregatorImpl struct {
	subjectResolver contracts.SubjectResolver
}

func NewPostAggregator(subjectResolver contracts.SubjectResolver) contracts.PostAggregator {
	return &postAggregatorImpl{subjectResolver}
}

func (c *postAggregatorImpl) AggregateObject(postEntity *entities.PostEntity) (post *models.Post) {
	post = new(models.Post)
	post.Id = postEntity.Id
	post.Title = postEntity.Title
	post.Description = postEntity.Description
	post.Content = postEntity.Content
	post.Status = postEntity.Status
	post.Privacy = postEntity.Privacy
	post.Published = postEntity.Published
	post.Views = postEntity.Views
	post.Created = postEntity.Created
	post.Updated = postEntity.Updated
	post.Author, _ = c.subjectResolver.ResolveSubject(postEntity.AuthorId, postEntity.AuthorType)

	return
}

func (c *postAggregatorImpl) AggregateCollection(postEntities []*entities.PostEntity) (posts []*models.Post) {
	posts = make([]*models.Post, len(postEntities))

	for k, postEntity := range postEntities {
		posts[k] = c.AggregateObject(postEntity)
	}

	return
}


