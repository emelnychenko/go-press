package apis

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/events"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewTagApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		tagEventFactory := mocks.NewMockTagEventFactory(ctrl)
		tagService := mocks.NewMockTagService(ctrl)
		tagAggregator := mocks.NewMockTagAggregator(ctrl)

		tagApi, isTagApi := NewTagApi(
			eventDispatcher, tagEventFactory, tagService, tagAggregator,
		).(*tagApiImpl)

		assert.True(t, isTagApi)
		assert.Equal(t, eventDispatcher, tagApi.eventDispatcher)
		assert.Equal(t, tagEventFactory, tagApi.tagEventFactory)
		assert.Equal(t, tagService, tagApi.tagService)
		assert.Equal(t, tagAggregator, tagApi.tagAggregator)
	})

	t.Run("ListTags", func(t *testing.T) {
		paginationQuery := new(models.TagPaginationQuery)
		entityPaginationResult := new(models.PaginationResult)
		paginationResult := new(models.PaginationResult)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().ListTags(paginationQuery).Return(entityPaginationResult, nil)

		tagAggregator := mocks.NewMockTagAggregator(ctrl)
		tagAggregator.EXPECT().AggregatePaginationResult(entityPaginationResult).Return(paginationResult)

		tagApi := &tagApiImpl{tagService: tagService, tagAggregator: tagAggregator}
		response, err := tagApi.ListTags(paginationQuery)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListTags:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		paginationQuery := new(models.TagPaginationQuery)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().ListTags(paginationQuery).Return(nil, systemErr)

		tagApi := &tagApiImpl{tagService: tagService}
		response, err := tagApi.ListTags(paginationQuery)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetTag", func(t *testing.T) {
		tagId := new(models.TagId)
		tagEntity := new(entities.TagEntity)
		tag := new(models.Tag)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(tagEntity, nil)

		tagAggregator := mocks.NewMockTagAggregator(ctrl)
		tagAggregator.EXPECT().AggregateTag(tagEntity).Return(tag)

		tagApi := &tagApiImpl{tagService: tagService, tagAggregator: tagAggregator}
		response, err := tagApi.GetTag(tagId)

		assert.Equal(t, tag, response)
		assert.Nil(t, err)
	})

	t.Run("GetTag:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		tagId := new(models.TagId)
		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(nil, systemErr)

		tagApi := &tagApiImpl{tagService: tagService}
		response, err := tagApi.GetTag(tagId)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateTag", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		tag := new(models.Tag)
		data := new(models.TagCreate)

		tagEvent := new(events.TagEvent)
		tagEventFactory := mocks.NewMockTagEventFactory(ctrl)
		tagEventFactory.EXPECT().CreateTagCreatedEvent(tagEntity).Return(tagEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(tagEvent)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().CreateTag(data).Return(tagEntity, nil)

		tagAggregator := mocks.NewMockTagAggregator(ctrl)
		tagAggregator.EXPECT().AggregateTag(tagEntity).Return(tag)

		tagApi := &tagApiImpl{
			eventDispatcher: eventDispatcher,
			tagEventFactory: tagEventFactory,
			tagService:      tagService,
			tagAggregator:   tagAggregator,
		}
		response, err := tagApi.CreateTag(data)

		assert.Equal(t, tag, response)
		assert.Nil(t, err)
	})

	t.Run("CreateTag:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		data := new(models.TagCreate)
		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().CreateTag(data).Return(nil, systemErr)

		tagApi := &tagApiImpl{tagService: tagService}
		response, err := tagApi.CreateTag(data)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateTag", func(t *testing.T) {
		tagId := new(models.TagId)
		tagEntity := new(entities.TagEntity)
		data := new(models.TagUpdate)

		tagEvent := new(events.TagEvent)
		tagEventFactory := mocks.NewMockTagEventFactory(ctrl)
		tagEventFactory.EXPECT().CreateTagUpdatedEvent(tagEntity).Return(tagEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(tagEvent)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(tagEntity, nil)
		tagService.EXPECT().UpdateTag(tagEntity, data).Return(nil)

		tagApi := &tagApiImpl{
			eventDispatcher: eventDispatcher,
			tagEventFactory: tagEventFactory,
			tagService:      tagService,
		}
		assert.Nil(t, tagApi.UpdateTag(tagId, data))
	})

	t.Run("UpdateTag:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		tagId := new(models.TagId)
		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(nil, systemErr)

		data := new(models.TagUpdate)
		tagApi := &tagApiImpl{tagService: tagService}
		assert.Equal(t, systemErr, tagApi.UpdateTag(tagId, data))
	})

	t.Run("UpdateTag:UpdateTagError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		tagId := new(models.TagId)
		tagEntity := new(entities.TagEntity)
		data := new(models.TagUpdate)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(tagEntity, nil)
		tagService.EXPECT().UpdateTag(tagEntity, data).Return(systemErr)

		tagApi := &tagApiImpl{
			tagService: tagService,
		}

		err := tagApi.UpdateTag(tagId, data)
		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteTag", func(t *testing.T) {
		tagId := new(models.TagId)
		tagEntity := new(entities.TagEntity)

		tagEvent := new(events.TagEvent)
		tagEventFactory := mocks.NewMockTagEventFactory(ctrl)
		tagEventFactory.EXPECT().CreateTagDeletedEvent(tagEntity).Return(tagEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(tagEvent)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(tagEntity, nil)
		tagService.EXPECT().DeleteTag(tagEntity).Return(nil)

		tagApi := &tagApiImpl{
			eventDispatcher: eventDispatcher,
			tagEventFactory: tagEventFactory,
			tagService:      tagService,
		}
		assert.Nil(t, tagApi.DeleteTag(tagId))
	})

	t.Run("DeleteTag:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		tagId := new(models.TagId)
		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(nil, systemErr)

		tagApi := &tagApiImpl{tagService: tagService}
		assert.Equal(t, systemErr, tagApi.DeleteTag(tagId))
	})

	t.Run("DeleteTag:DeleteTagError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		tagId := new(models.TagId)
		tagEntity := new(entities.TagEntity)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(tagEntity, nil)
		tagService.EXPECT().DeleteTag(tagEntity).Return(systemErr)

		tagApi := &tagApiImpl{
			tagService: tagService,
		}

		err := tagApi.DeleteTag(tagId)
		assert.Equal(t, systemErr, err)
	})
}
