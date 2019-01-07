package apis

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	tagApiImpl struct {
		eventDispatcher contracts.EventDispatcher
		tagEventFactory contracts.TagEventFactory
		tagService      contracts.TagService
		tagAggregator   contracts.TagAggregator
	}
)

func NewTagApi(
	eventDispatcher contracts.EventDispatcher,
	tagEventFactory contracts.TagEventFactory,
	tagService contracts.TagService,
	tagAggregator contracts.TagAggregator,
) (tagApi contracts.TagApi) {
	return &tagApiImpl{
		eventDispatcher,
		tagEventFactory,
		tagService,
		tagAggregator,
	}
}

func (a *tagApiImpl) ListTags(
	tagPaginationQuery *models.TagPaginationQuery,
) (paginationResult *models.PaginationResult, err errors.Error) {
	entityPaginationResult, err := a.tagService.ListTags(tagPaginationQuery)

	if nil != err {
		return
	}

	paginationResult = a.tagAggregator.AggregatePaginationResult(entityPaginationResult)
	return
}

func (a *tagApiImpl) GetTag(tagId *models.TagId) (tag *models.Tag, err errors.Error) {
	tagEntity, err := a.tagService.GetTag(tagId)

	if nil != err {
		return
	}

	tag = a.tagAggregator.AggregateTag(tagEntity)
	return
}

func (a *tagApiImpl) CreateTag(data *models.TagCreate) (tag *models.Tag, err errors.Error) {
	tagEntity, err := a.tagService.CreateTag(data)

	if nil != err {
		return
	}

	tagCreatedEvent := a.tagEventFactory.CreateTagCreatedEvent(tagEntity)
	a.eventDispatcher.Dispatch(tagCreatedEvent)

	tag = a.tagAggregator.AggregateTag(tagEntity)
	return
}

func (a *tagApiImpl) UpdateTag(tagId *models.TagId, data *models.TagUpdate) (err errors.Error) {
	tagService := a.tagService
	tagEntity, err := tagService.GetTag(tagId)

	if nil != err {
		return
	}

	err = tagService.UpdateTag(tagEntity, data)

	if nil != err {
		return
	}

	tagUpdatedEvent := a.tagEventFactory.CreateTagUpdatedEvent(tagEntity)
	a.eventDispatcher.Dispatch(tagUpdatedEvent)
	return
}

func (a *tagApiImpl) DeleteTag(tagId *models.TagId) (err errors.Error) {
	tagService := a.tagService
	tagEntity, err := tagService.GetTag(tagId)

	if nil != err {
		return
	}

	err = tagService.DeleteTag(tagEntity)

	if nil != err {
		return
	}

	tagDeletedEvent := a.tagEventFactory.CreateTagDeletedEvent(tagEntity)
	a.eventDispatcher.Dispatch(tagDeletedEvent)

	return
}

//ListObjectTags
func (a *tagApiImpl) ListObjectTags(
	tagObject models.Object, tagPaginationQuery *models.TagPaginationQuery,
) (
	paginationResult *models.PaginationResult, err errors.Error,
) {
	entityPaginationResult, err := a.tagService.ListObjectTags(tagObject, tagPaginationQuery)

	if nil != err {
		return
	}

	paginationResult = a.tagAggregator.AggregatePaginationResult(entityPaginationResult)
	return
}
