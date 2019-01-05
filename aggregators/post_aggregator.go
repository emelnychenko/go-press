package aggregators

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type postAggregatorImpl struct {
	postModelFactory contracts.PostModelFactory
	subjectResolver  contracts.SubjectResolver
	fileApi          contracts.FileApi
}

func NewPostAggregator(
	postModelFactory contracts.PostModelFactory,
	subjectResolver contracts.SubjectResolver,
	fileApi contracts.FileApi,
) contracts.PostAggregator {
	return &postAggregatorImpl{postModelFactory, subjectResolver, fileApi}
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

	if nil != postEntity.PictureId {
		post.Picture, _ = a.fileApi.GetFile(postEntity.PictureId)
	}

	if nil != postEntity.VideoId {
		post.Video, _ = a.fileApi.GetFile(postEntity.VideoId)
	}

	return
}

func (a *postAggregatorImpl) AggregatePosts(postEntities []*entities.PostEntity) (posts []*models.Post) {
	posts = make([]*models.Post, len(postEntities))

	for k, postEntity := range postEntities {
		posts[k] = a.AggregatePost(postEntity)
	}

	return
}

func (a *postAggregatorImpl) AggregatePaginationResult(
	entityPaginationResult *models.PaginationResult,
) (
	paginationResult *models.PaginationResult,
) {
	postEntities := entityPaginationResult.Data.([]*entities.PostEntity)
	posts := a.AggregatePosts(postEntities)
	return &models.PaginationResult{Total: entityPaginationResult.Total, Data: posts}
}
