package controllers

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewChannelController", func(t *testing.T) {
		channelHttpHelper := mocks.NewMockChannelHttpHelper(ctrl)
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelApi := mocks.NewMockChannelApi(ctrl)
		channelController, isChannelController := NewChannelController(
			channelHttpHelper,
			channelModelFactory,
			channelApi,
		).(*channelControllerImpl)

		assert.True(t, isChannelController)
		assert.Equal(t, channelHttpHelper, channelController.channelHttpHelper)
		assert.Equal(t, channelModelFactory, channelController.channelModelFactory)
		assert.Equal(t, channelApi, channelController.channelApi)
	})

	t.Run("ListChannels", func(t *testing.T) {
		channelPaginationQuery := new(models.ChannelPaginationQuery)
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannelPaginationQuery().Return(channelPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(channelPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(channelPaginationQuery).Return(nil)

		var paginationResult *models.PaginationResult
		channelApi := mocks.NewMockChannelApi(ctrl)
		channelApi.EXPECT().ListChannels(channelPaginationQuery).Return(paginationResult, nil)

		channelController := &channelControllerImpl{
			channelModelFactory: channelModelFactory,
			channelApi:          channelApi,
		}
		response, err := channelController.ListChannels(httpContext)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListChannels:BindPaginationQueryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		channelPaginationQuery := new(models.ChannelPaginationQuery)
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannelPaginationQuery().Return(channelPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(channelPaginationQuery.PaginationQuery).Return(systemErr)

		channelController := &channelControllerImpl{
			channelModelFactory: channelModelFactory,
		}
		response, err := channelController.ListChannels(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListChannels:BindChannelPaginationQueryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		channelPaginationQuery := new(models.ChannelPaginationQuery)
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannelPaginationQuery().Return(channelPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(channelPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(channelPaginationQuery).Return(systemErr)

		channelController := &channelControllerImpl{
			channelModelFactory: channelModelFactory,
		}
		response, err := channelController.ListChannels(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetChannel", func(t *testing.T) {
		channelId := new(models.ChannelId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		var channel *models.Channel
		channelApi := mocks.NewMockChannelApi(ctrl)
		channelApi.EXPECT().GetChannel(channelId).Return(channel, nil)

		channelHttpHelper := mocks.NewMockChannelHttpHelper(ctrl)
		channelHttpHelper.EXPECT().ParseChannelId(httpContext).Return(channelId, nil)

		channelController := &channelControllerImpl{channelHttpHelper: channelHttpHelper, channelApi: channelApi}
		response, err := channelController.GetChannel(httpContext)

		assert.Equal(t, channel, response)
		assert.Nil(t, err)
	})

	t.Run("GetChannel:ParserError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		channelHttpHelper := mocks.NewMockChannelHttpHelper(ctrl)
		channelHttpHelper.EXPECT().ParseChannelId(httpContext).Return(nil, systemErr)

		channelController := &channelControllerImpl{channelHttpHelper: channelHttpHelper}
		response, err := channelController.GetChannel(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetChannel:ApiError", func(t *testing.T) {
		channelId := new(models.ChannelId)
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		channelApi := mocks.NewMockChannelApi(ctrl)
		channelApi.EXPECT().GetChannel(channelId).Return(nil, systemErr)

		channelHttpHelper := mocks.NewMockChannelHttpHelper(ctrl)
		channelHttpHelper.EXPECT().ParseChannelId(httpContext).Return(channelId, nil)

		channelController := &channelControllerImpl{channelHttpHelper: channelHttpHelper, channelApi: channelApi}
		response, err := channelController.GetChannel(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateChannel", func(t *testing.T) {
		channel := new(models.Channel)
		data := new(models.ChannelCreate)
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannelCreate().Return(data)

		channelApi := mocks.NewMockChannelApi(ctrl)
		channelApi.EXPECT().CreateChannel(data).Return(channel, nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		channelController := &channelControllerImpl{
			channelModelFactory: channelModelFactory,
			channelApi:          channelApi,
		}
		response, err := channelController.CreateChannel(httpContext)

		assert.Equal(t, channel, response)
		assert.Nil(t, err)
	})

	t.Run("CreateChannel:BindChannelUpdateError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		data := new(models.ChannelCreate)

		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannelCreate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		channelController := &channelControllerImpl{
			channelModelFactory: channelModelFactory,
		}
		_, err := channelController.CreateChannel(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateChannel:ApiError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		data := new(models.ChannelCreate)

		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannelCreate().Return(data)

		channelApi := mocks.NewMockChannelApi(ctrl)
		channelApi.EXPECT().CreateChannel(data).Return(nil, systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		channelController := &channelControllerImpl{
			channelModelFactory: channelModelFactory,
			channelApi:          channelApi,
		}
		_, err := channelController.CreateChannel(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateChannel", func(t *testing.T) {
		channelId := new(models.ChannelId)
		data := new(models.ChannelUpdate)
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannelUpdate().Return(data)

		channelApi := mocks.NewMockChannelApi(ctrl)
		channelApi.EXPECT().UpdateChannel(channelId, data).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		channelHttpHelper := mocks.NewMockChannelHttpHelper(ctrl)
		channelHttpHelper.EXPECT().ParseChannelId(httpContext).Return(channelId, nil)

		channelController := &channelControllerImpl{
			channelHttpHelper:   channelHttpHelper,
			channelModelFactory: channelModelFactory,
			channelApi:          channelApi,
		}
		_, err := channelController.UpdateChannel(httpContext)

		assert.Nil(t, err)
	})

	t.Run("UpdateChannel:ParseError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		channelHttpHelper := mocks.NewMockChannelHttpHelper(ctrl)
		channelHttpHelper.EXPECT().ParseChannelId(httpContext).Return(nil, systemErr)

		channelController := &channelControllerImpl{channelHttpHelper: channelHttpHelper}
		_, err := channelController.UpdateChannel(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateChannel:BindChannelUpdateError", func(t *testing.T) {
		channelId := new(models.ChannelId)
		systemErr := errors.NewUnknownError()
		data := new(models.ChannelUpdate)
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannelUpdate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		channelHttpHelper := mocks.NewMockChannelHttpHelper(ctrl)
		channelHttpHelper.EXPECT().ParseChannelId(httpContext).Return(channelId, nil)

		channelController := &channelControllerImpl{
			channelHttpHelper:   channelHttpHelper,
			channelModelFactory: channelModelFactory,
		}
		_, err := channelController.UpdateChannel(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateChannel:ApiError", func(t *testing.T) {
		channelId := new(models.ChannelId)
		systemErr := errors.NewUnknownError()

		data := new(models.ChannelUpdate)
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannelUpdate().Return(data)

		channelApi := mocks.NewMockChannelApi(ctrl)
		channelApi.EXPECT().UpdateChannel(channelId, data).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		channelHttpHelper := mocks.NewMockChannelHttpHelper(ctrl)
		channelHttpHelper.EXPECT().ParseChannelId(httpContext).Return(channelId, nil)

		channelController := &channelControllerImpl{
			channelHttpHelper:   channelHttpHelper,
			channelModelFactory: channelModelFactory,
			channelApi:          channelApi,
		}
		_, err := channelController.UpdateChannel(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteChannel", func(t *testing.T) {
		channelId := new(models.ChannelId)

		channelApi := mocks.NewMockChannelApi(ctrl)
		channelApi.EXPECT().DeleteChannel(channelId).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)

		channelHttpHelper := mocks.NewMockChannelHttpHelper(ctrl)
		channelHttpHelper.EXPECT().ParseChannelId(httpContext).Return(channelId, nil)

		channelController := &channelControllerImpl{channelHttpHelper: channelHttpHelper, channelApi: channelApi}
		_, err := channelController.DeleteChannel(httpContext)

		assert.Nil(t, err)
	})

	t.Run("DeleteChannel:ParseError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		httpContext := mocks.NewMockHttpContext(ctrl)

		channelHttpHelper := mocks.NewMockChannelHttpHelper(ctrl)
		channelHttpHelper.EXPECT().ParseChannelId(httpContext).Return(nil, systemErr)

		channelController := &channelControllerImpl{channelHttpHelper: channelHttpHelper}
		_, err := channelController.DeleteChannel(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteChannel:ApiError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		channelId := new(models.ChannelId)

		channelApi := mocks.NewMockChannelApi(ctrl)
		channelApi.EXPECT().DeleteChannel(channelId).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)

		channelHttpHelper := mocks.NewMockChannelHttpHelper(ctrl)
		channelHttpHelper.EXPECT().ParseChannelId(httpContext).Return(channelId, nil)

		channelController := &channelControllerImpl{channelHttpHelper: channelHttpHelper, channelApi: channelApi}
		_, err := channelController.DeleteChannel(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
