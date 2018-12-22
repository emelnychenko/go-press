package aggregator

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type postAggregatorImpl struct {
	postModelFactory contracts.PostModelFactory
	subjectResolver  contracts.SubjectResolver
}

func NewPostAggregator(
	postModelFactory contracts.PostModelFactory,
	subjectResolver contracts.SubjectResolver,
) contracts.PostAggregator {
	return &postAggregatorImpl{postModelFactory, subjectResolver}
}

func (a *postAggregatorImpl) AggregatePost(postEntity *entities.PostEntity) (post *models.Post) {
	post = a.postModelFactory.CreatePost()
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
	post.Author, _ = a.subjectResolver.ResolveSubject(postEntity.AuthorId, postEntity.AuthorType)

	return
}

func (a *postAggregatorImpl) AggregatePosts(postEntities []*entities.PostEntity) (posts []*models.Post) {
	posts = make([]*models.Post, len(postEntities))

	for k, postEntity := range postEntities {
		posts[k] = a.AggregatePost(postEntity)
	}

	return
}
