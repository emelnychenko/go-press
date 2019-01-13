package apis

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	postTagApiImpl struct {
		eventDispatcher          contracts.EventDispatcher
		postTagEventFactory contracts.PostTagEventFactory
		postService              contracts.PostService
		tagService          contracts.TagService
		postTagService      contracts.PostTagService
		tagAggregator       contracts.TagAggregator
	}
)

//NewPostTagApi
func NewPostTagApi(
	eventDispatcher contracts.EventDispatcher,
	postTagEventFactory contracts.PostTagEventFactory,
	postService contracts.PostService,
	tagService contracts.TagService,
	postTagService contracts.PostTagService,
	tagAggregator contracts.TagAggregator,
) (postTagApi contracts.PostTagApi) {
	return &postTagApiImpl{
		eventDispatcher:          eventDispatcher,
		postTagEventFactory: postTagEventFactory,
		postService:              postService,
		tagService:          tagService,
		postTagService:      postTagService,
		tagAggregator:       tagAggregator,
	}
}

//ListPostTags
func (a *postTagApiImpl) ListPostTags(
	postId *models.PostId, paginationQuery *models.TagPaginationQuery,
) (
	paginationResult *models.PaginationResult, err errors.Error,
) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	entityPaginationResult, err := a.postTagService.ListPostTags(postEntity, paginationQuery)

	if nil != err {
		return
	}

	paginationResult = a.tagAggregator.AggregatePaginationResult(entityPaginationResult)
	return
}

//AddPostTag
func (a *postTagApiImpl) AddPostTag(postId *models.PostId, tagId *models.TagId) (err errors.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	tagEntity, err := a.tagService.GetTag(tagId)

	if nil != err {
		return
	}

	err = a.postTagService.AddPostTag(postEntity, tagEntity)

	if nil != err {
		return
	}

	postTagEvent := a.postTagEventFactory.CreatePostTagAddedEvent(postEntity, tagEntity)
	a.eventDispatcher.Dispatch(postTagEvent)

	return
}

//RemovePostTag
func (a *postTagApiImpl) RemovePostTag(postId *models.PostId, tagId *models.TagId) (err errors.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	tagEntity, err := a.tagService.GetTag(tagId)

	if nil != err {
		return
	}

	err = a.postTagService.RemovePostTag(postEntity, tagEntity)

	if nil != err {
		return
	}

	postTagEvent := a.postTagEventFactory.CreatePostTagRemovedEvent(postEntity, tagEntity)
	a.eventDispatcher.Dispatch(postTagEvent)

	return
}
