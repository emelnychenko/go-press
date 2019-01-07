package echo

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewRouter", func(t *testing.T) {
		instance := new(echo.Echo)
		modelValidator := mocks.NewMockModelValidator(ctrl)
		router, isRouter := NewRouter(instance, modelValidator).(*routerImpl)

		assert.True(t, isRouter)
		assert.Equal(t, instance, router.echo)
		assert.Equal(t, modelValidator, router.modelValidator)
	})

	t.Run("AddRoute", func(t *testing.T) {
		instance := echo.New()
		router := &routerImpl{echo: instance}

		httpMethod := http.MethodGet
		httpHandlerResponse := "test"
		routePath := "/test"
		router.AddRoute(httpMethod, routePath, func(httpContext contracts.HttpContext) (response interface{}, err errors.Error) {
			return httpHandlerResponse, nil
		})

		rootContext := echo.New().NewContext(nil, nil)
		instance.Router().Find(httpMethod, routePath, rootContext)

		responseWriter := httptest.NewRecorder()
		response := echo.NewResponse(responseWriter, instance)

		mockContext := mocks.NewMockContext(ctrl)
		mockContext.EXPECT().Response().Return(response)
		mockContext.EXPECT().JSON(http.StatusOK, httpHandlerResponse).Return(nil)

		err := rootContext.Handler()(mockContext)
		assert.Nil(t, err)
	})

	t.Run("AddRoute:CommittedError", func(t *testing.T) {
		instance := echo.New()
		router := &routerImpl{echo: instance}

		systemErr := errors.NewUnknownError()
		httpMethod := http.MethodGet
		routePath := "/test"
		router.AddRoute(httpMethod, routePath, func(httpContext contracts.HttpContext) (response interface{}, err errors.Error) {
			return nil, systemErr
		})

		rootContext := echo.New().NewContext(nil, nil)
		instance.Router().Find(httpMethod, routePath, rootContext)

		responseWriter := httptest.NewRecorder()
		response := echo.NewResponse(responseWriter, instance)
		response.Committed = true

		mockContext := mocks.NewMockContext(ctrl)
		mockContext.EXPECT().Response().Return(response)

		err := rootContext.Handler()(mockContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddRoute:JsonError", func(t *testing.T) {
		instance := echo.New()
		router := &routerImpl{echo: instance}

		systemErr := errors.NewUnknownError()
		httpMethod := http.MethodGet
		httpHandlerResponse := "test"
		routePath := "/test"
		router.AddRoute(httpMethod, routePath, func(httpContext contracts.HttpContext) (response interface{}, err errors.Error) {
			return httpHandlerResponse, systemErr
		})

		rootContext := echo.New().NewContext(nil, nil)
		instance.Router().Find(httpMethod, routePath, rootContext)

		responseWriter := httptest.NewRecorder()
		response := echo.NewResponse(responseWriter, instance)

		mockContext := mocks.NewMockContext(ctrl)
		mockContext.EXPECT().Response().Return(response)
		mockContext.EXPECT().JSON(systemErr.Code(), systemErr.Error()).Return(nil)

		err := rootContext.Handler()(mockContext)
		assert.Nil(t, err)
	})
}
