package apis

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	postApiImpl struct {
		eventDispatcher  contracts.EventDispatcher
		postEventFactory contracts.PostEventFactory
		postService      contracts.PostService
		postAggregator   contracts.PostAggregator
	}
)

func NewPostApi(
	eventDispatcher contracts.EventDispatcher,
	postEventFactory contracts.PostEventFactory,
	postService contracts.PostService,
	postAggregator contracts.PostAggregator,
) (postApi contracts.PostApi) {
	return &postApiImpl{
		eventDispatcher,
		postEventFactory,
		postService,
		postAggregator,
	}
}

func (a *postApiImpl) ListPosts(
	postPaginationQuery *models.PostPaginationQuery,
) (paginationResult *models.PaginationResult, err errors.Error) {
	entityPaginationResult, err := a.postService.ListPosts(postPaginationQuery)

	if nil != err {
		return
	}

	paginationResult = a.postAggregator.AggregatePaginationResult(entityPaginationResult)
	return
}

func (a *postApiImpl) GetPost(postId *models.PostId) (post *models.Post, err errors.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	post = a.postAggregator.AggregatePost(postEntity)
	return
}

func (a *postApiImpl) CreatePost(postAuthor models.Subject, data *models.PostCreate) (post *models.Post, err errors.Error) {
	postEntity, err := a.postService.CreatePost(postAuthor, data)

	if nil != err {
		return
	}

	postCreatedEvent := a.postEventFactory.CreatePostCreatedEvent(postEntity)
	a.eventDispatcher.Dispatch(postCreatedEvent)

	if enums.PostPublishedStatus == postEntity.Status {
		postPublishedEvent := a.postEventFactory.CreatePostPublishedEvent(postEntity)
		a.eventDispatcher.Dispatch(postPublishedEvent)
	}

	post = a.postAggregator.AggregatePost(postEntity)
	return
}

func (a *postApiImpl) UpdatePost(postId *models.PostId, data *models.PostUpdate) (err errors.Error) {
	postService := a.postService
	postEntity, err := postService.GetPost(postId)

	if nil != err {
		return
	}

	postStatusBefore := postEntity.Status

	err = postService.UpdatePost(postEntity, data)

	if nil != err {
		return
	}

	postUpdatedEvent := a.postEventFactory.CreatePostUpdatedEvent(postEntity)
	a.eventDispatcher.Dispatch(postUpdatedEvent)

	postStatusAfter := postEntity.Status

	if enums.PostPublishedStatus != postStatusBefore &&
		enums.PostPublishedStatus == postStatusAfter {
		postPublishedEvent := a.postEventFactory.CreatePostPublishedEvent(postEntity)
		a.eventDispatcher.Dispatch(postPublishedEvent)
	}

	if enums.PostPublishedStatus == postStatusBefore &&
		enums.PostPublishedStatus != postStatusAfter {
		postConcealedEvent := a.postEventFactory.CreatePostConcealedEvent(postEntity)
		a.eventDispatcher.Dispatch(postConcealedEvent)
	}

	return
}

func (a *postApiImpl) DeletePost(postId *models.PostId) (err errors.Error) {
	postService := a.postService
	postEntity, err := postService.GetPost(postId)

	if nil != err {
		return
	}

	err = postService.DeletePost(postEntity)

	if nil != err {
		return
	}

	if enums.PostPublishedStatus == postEntity.Status {
		postConcealedEvent := a.postEventFactory.CreatePostConcealedEvent(postEntity)
		a.eventDispatcher.Dispatch(postConcealedEvent)
	}

	postDeletedEvent := a.postEventFactory.CreatePostDeletedEvent(postEntity)
	a.eventDispatcher.Dispatch(postDeletedEvent)

	return
}
