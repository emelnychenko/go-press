package controllers

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewTagController", func(t *testing.T) {
		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagApi := mocks.NewMockTagApi(ctrl)
		tagController, isTagController := NewTagController(
			tagHttpHelper,
			tagModelFactory,
			tagApi,
		).(*tagControllerImpl)

		assert.True(t, isTagController)
		assert.Equal(t, tagHttpHelper, tagController.tagHttpHelper)
		assert.Equal(t, tagModelFactory, tagController.tagModelFactory)
		assert.Equal(t, tagApi, tagController.tagApi)
	})

	t.Run("ListTags", func(t *testing.T) {
		tagPaginationQuery := new(models.TagPaginationQuery)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagPaginationQuery().Return(tagPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(tagPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(tagPaginationQuery).Return(nil)

		var paginationResult *models.PaginationResult
		tagApi := mocks.NewMockTagApi(ctrl)
		tagApi.EXPECT().ListTags(tagPaginationQuery).Return(paginationResult, nil)

		tagController := &tagControllerImpl{
			tagModelFactory: tagModelFactory,
			tagApi:          tagApi,
		}
		response, err := tagController.ListTags(httpContext)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListTags:BindPaginationQueryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		tagPaginationQuery := new(models.TagPaginationQuery)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagPaginationQuery().Return(tagPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(tagPaginationQuery.PaginationQuery).Return(systemErr)

		tagController := &tagControllerImpl{
			tagModelFactory: tagModelFactory,
		}
		response, err := tagController.ListTags(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListTags:BindTagPaginationQueryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		tagPaginationQuery := new(models.TagPaginationQuery)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagPaginationQuery().Return(tagPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(tagPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(tagPaginationQuery).Return(systemErr)

		tagController := &tagControllerImpl{
			tagModelFactory: tagModelFactory,
		}
		response, err := tagController.ListTags(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetTag", func(t *testing.T) {
		tagId := new(models.TagId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		var tag *models.Tag
		tagApi := mocks.NewMockTagApi(ctrl)
		tagApi.EXPECT().GetTag(tagId).Return(tag, nil)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(tagId, nil)

		tagController := &tagControllerImpl{tagHttpHelper: tagHttpHelper, tagApi: tagApi}
		response, err := tagController.GetTag(httpContext)

		assert.Equal(t, tag, response)
		assert.Nil(t, err)
	})

	t.Run("GetTag:ParserError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(nil, systemErr)

		tagController := &tagControllerImpl{tagHttpHelper: tagHttpHelper}
		response, err := tagController.GetTag(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetTag:ApiError", func(t *testing.T) {
		tagId := new(models.TagId)
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		tagApi := mocks.NewMockTagApi(ctrl)
		tagApi.EXPECT().GetTag(tagId).Return(nil, systemErr)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(tagId, nil)

		tagController := &tagControllerImpl{tagHttpHelper: tagHttpHelper, tagApi: tagApi}
		response, err := tagController.GetTag(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateTag", func(t *testing.T) {
		tag := new(models.Tag)
		data := new(models.TagCreate)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagCreate().Return(data)

		tagApi := mocks.NewMockTagApi(ctrl)
		tagApi.EXPECT().CreateTag(data).Return(tag, nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		tagController := &tagControllerImpl{
			tagModelFactory: tagModelFactory,
			tagApi:          tagApi,
		}
		response, err := tagController.CreateTag(httpContext)

		assert.Equal(t, tag, response)
		assert.Nil(t, err)
	})

	t.Run("CreateTag:BindTagUpdateError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		data := new(models.TagCreate)

		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagCreate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		tagController := &tagControllerImpl{
			tagModelFactory: tagModelFactory,
		}
		_, err := tagController.CreateTag(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateTag:ApiError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		data := new(models.TagCreate)

		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagCreate().Return(data)

		tagApi := mocks.NewMockTagApi(ctrl)
		tagApi.EXPECT().CreateTag(data).Return(nil, systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		tagController := &tagControllerImpl{
			tagModelFactory: tagModelFactory,
			tagApi:          tagApi,
		}
		_, err := tagController.CreateTag(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateTag", func(t *testing.T) {
		tagId := new(models.TagId)
		data := new(models.TagUpdate)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagUpdate().Return(data)

		tagApi := mocks.NewMockTagApi(ctrl)
		tagApi.EXPECT().UpdateTag(tagId, data).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(tagId, nil)

		tagController := &tagControllerImpl{
			tagHttpHelper:   tagHttpHelper,
			tagModelFactory: tagModelFactory,
			tagApi:          tagApi,
		}
		_, err := tagController.UpdateTag(httpContext)

		assert.Nil(t, err)
	})

	t.Run("UpdateTag:ParseError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(nil, systemErr)

		tagController := &tagControllerImpl{tagHttpHelper: tagHttpHelper}
		_, err := tagController.UpdateTag(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateTag:BindTagUpdateError", func(t *testing.T) {
		tagId := new(models.TagId)
		systemErr := errors.NewUnknownError()
		data := new(models.TagUpdate)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagUpdate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(tagId, nil)

		tagController := &tagControllerImpl{
			tagHttpHelper:   tagHttpHelper,
			tagModelFactory: tagModelFactory,
		}
		_, err := tagController.UpdateTag(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateTag:ApiError", func(t *testing.T) {
		tagId := new(models.TagId)
		systemErr := errors.NewUnknownError()

		data := new(models.TagUpdate)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagUpdate().Return(data)

		tagApi := mocks.NewMockTagApi(ctrl)
		tagApi.EXPECT().UpdateTag(tagId, data).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(tagId, nil)

		tagController := &tagControllerImpl{
			tagHttpHelper:   tagHttpHelper,
			tagModelFactory: tagModelFactory,
			tagApi:          tagApi,
		}
		_, err := tagController.UpdateTag(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteTag", func(t *testing.T) {
		tagId := new(models.TagId)

		tagApi := mocks.NewMockTagApi(ctrl)
		tagApi.EXPECT().DeleteTag(tagId).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(tagId, nil)

		tagController := &tagControllerImpl{tagHttpHelper: tagHttpHelper, tagApi: tagApi}
		_, err := tagController.DeleteTag(httpContext)

		assert.Nil(t, err)
	})

	t.Run("DeleteTag:ParseError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		httpContext := mocks.NewMockHttpContext(ctrl)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(nil, systemErr)

		tagController := &tagControllerImpl{tagHttpHelper: tagHttpHelper}
		_, err := tagController.DeleteTag(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteTag:ApiError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		tagId := new(models.TagId)

		tagApi := mocks.NewMockTagApi(ctrl)
		tagApi.EXPECT().DeleteTag(tagId).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(tagId, nil)

		tagController := &tagControllerImpl{tagHttpHelper: tagHttpHelper, tagApi: tagApi}
		_, err := tagController.DeleteTag(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
