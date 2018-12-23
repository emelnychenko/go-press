package echo

import (
	"errors"
	"github.com/emelnychenko/go-press/echo_mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"testing"
)

func TestHttpContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewHttpContext", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)
		httpContext, isHttpContext := NewHttpContext(context).(*httpContextImpl)

		assert.True(t, isHttpContext)
		assert.Equal(t, context, httpContext.context)
	})

	t.Run("Request", func(t *testing.T) {
		request := new(http.Request)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().Request().Return(request)

		httpContext := &httpContextImpl{context: context}
		assert.Equal(t, request, httpContext.Request())
	})

	t.Run("Request", func(t *testing.T) {
		response := new(echo.Response)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().Response().Return(response)

		httpContext := &httpContextImpl{context: context}
		assert.Equal(t, response, httpContext.Response())
	})

	t.Run("Parameter", func(t *testing.T) {
		parameterName := ":test"
		parameter := "test"

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().Param(parameterName).Return(parameter)

		httpContext := &httpContextImpl{context: context}
		assert.Equal(t, parameter, httpContext.Parameter(parameterName))
	})

	t.Run("BindModel", func(t *testing.T) {
		var data interface{}

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().Bind(data).Return(nil)

		httpContext := &httpContextImpl{context: context}
		assert.Nil(t, httpContext.BindModel(data))
	})

	t.Run("BindModel:Error", func(t *testing.T) {
		systemErr := errors.New("")
		var data interface{}

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().Bind(data).Return(systemErr)

		httpContext := &httpContextImpl{context: context}
		assert.Error(t, httpContext.BindModel(data))
	})

	t.Run("FormFile", func(t *testing.T) {
		formFileName := ":test"
		fileHeader := new(multipart.FileHeader)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().FormFile(formFileName).Return(fileHeader, nil)

		httpContext := &httpContextImpl{context: context}
		response, err := httpContext.FormFile(formFileName)

		assert.Equal(t, fileHeader, response)
		assert.Nil(t, err)
	})

	t.Run("FormFile:Error", func(t *testing.T) {
		systemErr := errors.New("")
		formFileName := ":test"

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().FormFile(formFileName).Return(nil, systemErr)

		httpContext := &httpContextImpl{context: context}
		response, err := httpContext.FormFile(formFileName)

		assert.Nil(t, response)
		assert.Error(t, err)
	})
}
