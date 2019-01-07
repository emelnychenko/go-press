package echo

import (
	builtinErr "errors"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
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
		context := mocks.NewMockContext(ctrl)
		modelValidator := mocks.NewMockModelValidator(ctrl)
		httpContext, isHttpContext := NewHttpContext(context, modelValidator).(*httpContextImpl)

		assert.True(t, isHttpContext)
		assert.Equal(t, context, httpContext.context)
		assert.Equal(t, modelValidator, httpContext.modelValidator)
	})

	t.Run("Request", func(t *testing.T) {
		request := new(http.Request)

		context := mocks.NewMockContext(ctrl)
		context.EXPECT().Request().Return(request)

		httpContext := &httpContextImpl{context: context}
		assert.Equal(t, request, httpContext.Request())
	})

	t.Run("Request", func(t *testing.T) {
		response := new(echo.Response)

		context := mocks.NewMockContext(ctrl)
		context.EXPECT().Response().Return(response)

		httpContext := &httpContextImpl{context: context}
		assert.Equal(t, response, httpContext.Response())
	})

	t.Run("Parameter", func(t *testing.T) {
		parameterName := ":test"
		parameter := "test"

		context := mocks.NewMockContext(ctrl)
		context.EXPECT().Param(parameterName).Return(parameter)

		httpContext := &httpContextImpl{context: context}
		assert.Equal(t, parameter, httpContext.Parameter(parameterName))
	})

	t.Run("BindModel", func(t *testing.T) {
		var data interface{}

		context := mocks.NewMockContext(ctrl)
		context.EXPECT().Bind(data).Return(nil)

		modelValidator := mocks.NewMockModelValidator(ctrl)
		modelValidator.EXPECT().ValidateModel(data).Return(nil)

		httpContext := &httpContextImpl{context: context, modelValidator: modelValidator}
		assert.Nil(t, httpContext.BindModel(data))
	})

	t.Run("BindModel:BindError", func(t *testing.T) {
		systemErr := builtinErr.New("")
		var data interface{}

		context := mocks.NewMockContext(ctrl)
		context.EXPECT().Bind(data).Return(systemErr)

		httpContext := &httpContextImpl{context: context}
		assert.Error(t, httpContext.BindModel(data))
	})

	t.Run("BindModel:ValidateModelError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		var data interface{}

		context := mocks.NewMockContext(ctrl)
		context.EXPECT().Bind(data).Return(nil)

		modelValidator := mocks.NewMockModelValidator(ctrl)
		modelValidator.EXPECT().ValidateModel(data).Return(systemErr)

		httpContext := &httpContextImpl{context: context, modelValidator: modelValidator}
		assert.Equal(t, systemErr, httpContext.BindModel(data))
	})

	t.Run("FormFile", func(t *testing.T) {
		formFileName := ":test"
		fileHeader := new(multipart.FileHeader)

		context := mocks.NewMockContext(ctrl)
		context.EXPECT().FormFile(formFileName).Return(fileHeader, nil)

		httpContext := &httpContextImpl{context: context}
		response, err := httpContext.FormFile(formFileName)

		assert.Equal(t, fileHeader, response)
		assert.Nil(t, err)
	})

	t.Run("FormFile:Error", func(t *testing.T) {
		systemErr := builtinErr.New("")
		formFileName := ":test"

		context := mocks.NewMockContext(ctrl)
		context.EXPECT().FormFile(formFileName).Return(nil, systemErr)

		httpContext := &httpContextImpl{context: context}
		response, err := httpContext.FormFile(formFileName)

		assert.Nil(t, response)
		assert.Error(t, err)
	})
}
