package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
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

func (a *postApiImpl) ListPosts() (posts []*models.Post, err common.Error) {
	postEntities, err := a.postService.ListPosts()

	if nil != err {
		return
	}

	posts = a.postAggregator.AggregatePosts(postEntities)
	return
}

func (a *postApiImpl) GetPost(postId *models.PostId) (post *models.Post, err common.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	post = a.postAggregator.AggregatePost(postEntity)
	return
}

func (a *postApiImpl) CreatePost(postAuthor common.Subject, data *models.PostCreate) (post *models.Post, err common.Error) {
	postEntity, err := a.postService.CreatePost(postAuthor, data)

	if nil != err {
		return
	}

	postEvent := a.postEventFactory.CreatePostCreatedEvent(postEntity)
	a.eventDispatcher.Dispatch(postEvent)

	post = a.postAggregator.AggregatePost(postEntity)
	return
}

func (a *postApiImpl) UpdatePost(postId *models.PostId, data *models.PostUpdate) (err common.Error) {
	postService := a.postService
	postEntity, err := postService.GetPost(postId)

	if nil != err {
		return
	}

	err = postService.UpdatePost(postEntity, data)

	postEvent := a.postEventFactory.CreatePostUpdatedEvent(postEntity)
	a.eventDispatcher.Dispatch(postEvent)

	return
}

func (a *postApiImpl) DeletePost(postId *models.PostId) (err common.Error) {
	postService := a.postService
	postEntity, err := postService.GetPost(postId)

	if nil != err {
		return
	}

	err = postService.DeletePost(postEntity)

	postEvent := a.postEventFactory.CreatePostDeletedEvent(postEntity)
	a.eventDispatcher.Dispatch(postEvent)

	return
}
