package aggregators

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type postAggregatorImpl struct {
	postModelFactory     contracts.PostModelFactory
	subjectResolver      contracts.SubjectResolver
	fileApi              contracts.FileApi
	categoryModelFactory contracts.CategoryModelFactory
	categoryApi          contracts.CategoryApi
	tagModelFactory      contracts.TagModelFactory
	tagApi               contracts.TagApi
}

//NewPostAggregator
func NewPostAggregator(
	postModelFactory contracts.PostModelFactory,
	subjectResolver contracts.SubjectResolver,
	fileApi contracts.FileApi,
	categoryModelFactory contracts.CategoryModelFactory,
	categoryApi contracts.CategoryApi,
	tagModelFactory contracts.TagModelFactory,
	tagApi contracts.TagApi,
) contracts.PostAggregator {
	return &postAggregatorImpl{
		postModelFactory,
		subjectResolver,
		fileApi,
		categoryModelFactory,
		categoryApi,
		tagModelFactory,
		tagApi,
	}
}

//AggregatePost
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

	queue := 0
	done := make(chan bool)

	queue++
	// TODO: Add error log
	go func(done chan bool) {
		post.Author, _ = a.subjectResolver.ResolveSubject(postEntity.AuthorId, postEntity.AuthorType)
		done <- true
	}(done)

	queue++
	// TODO: Add error log
	go func(done chan bool) {
		categoryPaginationQuery := a.categoryModelFactory.CreateCategoryPaginationQuery()
		post.Categories, _ = a.categoryApi.ListObjectCategories(postEntity, categoryPaginationQuery)
		done <- true
	}(done)

	queue++
	// TODO: Add error log
	go func(done chan bool) {
		tagPaginationQuery := a.tagModelFactory.CreateTagPaginationQuery()
		post.Tags, _ = a.tagApi.ListObjectTags(postEntity, tagPaginationQuery)
		done <- true
	}(done)

	if nil != postEntity.PictureId {
		queue++
		// TODO: Add error log
		go func(done chan bool) {
			post.Picture, _ = a.fileApi.GetFile(postEntity.PictureId)
			done <- true
		}(done)
	}

	if nil != postEntity.VideoId {
		queue++
		// TODO: Add error log
		go func(done chan bool) {
			post.Video, _ = a.fileApi.GetFile(postEntity.VideoId)
			done <- true
		}(done)
	}

	for i := 0; i < queue; i++ {
		<-done
	}

	close(done)
	return
}

//AggregatePosts
func (a *postAggregatorImpl) AggregatePosts(postEntities []*entities.PostEntity) (posts []*models.Post) {
	posts = make([]*models.Post, len(postEntities))
	post := make(chan *models.Post, len(postEntities))

	for _, postEntity := range postEntities {
		go func(postEntity *entities.PostEntity, post chan *models.Post) {
			post <- a.AggregatePost(postEntity)
		}(postEntity, post)
	}

	for k := range postEntities {
		posts[k] = <-post
	}

	close(post)
	return
}

//AggregatePaginationResult
func (a *postAggregatorImpl) AggregatePaginationResult(
	entityPaginationResult *models.PaginationResult,
) (
	paginationResult *models.PaginationResult,
) {
	postEntities := entityPaginationResult.Data.([]*entities.PostEntity)
	posts := a.AggregatePosts(postEntities)
	return &models.PaginationResult{Total: entityPaginationResult.Total, Data: posts}
}
