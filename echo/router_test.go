package echo

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/echo_mocks"
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
		router, isRouter := NewRouter(instance).(*routerImpl)

		assert.True(t, isRouter)
		assert.Equal(t, instance, router.echo)
	})

	t.Run("AddRoute", func(t *testing.T) {
		instance := echo.New()
		router := &routerImpl{echo: instance}

		httpMethod := http.MethodGet
		httpHandlerResponse := "test"
		routePath := "/test"
		router.AddRoute(httpMethod, routePath, func(httpContext contracts.HttpContext) (response interface{}, err common.Error) {
			return httpHandlerResponse, nil
		})

		rootContext := echo.New().NewContext(nil, nil)
		instance.Router().Find(httpMethod, routePath, rootContext)

		responseWriter := httptest.NewRecorder()
		response := echo.NewResponse(responseWriter, instance)

		mockContext := echo_mocks.NewMockContext(ctrl)
		mockContext.EXPECT().Response().Return(response)
		mockContext.EXPECT().JSON(http.StatusOK, httpHandlerResponse).Return(nil)

		err := rootContext.Handler()(mockContext)
		assert.Nil(t, err)
	})

	t.Run("AddRoute:CommittedError", func(t *testing.T) {
		instance := echo.New()
		router := &routerImpl{echo: instance}

		systemErr := common.NewUnknownError()
		httpMethod := http.MethodGet
		routePath := "/test"
		router.AddRoute(httpMethod, routePath, func(httpContext contracts.HttpContext) (response interface{}, err common.Error) {
			return nil, systemErr
		})

		rootContext := echo.New().NewContext(nil, nil)
		instance.Router().Find(httpMethod, routePath, rootContext)

		responseWriter := httptest.NewRecorder()
		response := echo.NewResponse(responseWriter, instance)
		response.Committed = true

		mockContext := echo_mocks.NewMockContext(ctrl)
		mockContext.EXPECT().Response().Return(response)

		err := rootContext.Handler()(mockContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddRoute:JsonError", func(t *testing.T) {
		instance := echo.New()
		router := &routerImpl{echo: instance}

		systemErr := common.NewUnknownError()
		httpMethod := http.MethodGet
		httpHandlerResponse := "test"
		routePath := "/test"
		router.AddRoute(httpMethod, routePath, func(httpContext contracts.HttpContext) (response interface{}, err common.Error) {
			return httpHandlerResponse, systemErr
		})

		rootContext := echo.New().NewContext(nil, nil)
		instance.Router().Find(httpMethod, routePath, rootContext)

		responseWriter := httptest.NewRecorder()
		response := echo.NewResponse(responseWriter, instance)

		mockContext := echo_mocks.NewMockContext(ctrl)
		mockContext.EXPECT().Response().Return(response)
		mockContext.EXPECT().JSON(systemErr.Code(), systemErr.Error()).Return(nil)

		err := rootContext.Handler()(mockContext)
		assert.Nil(t, err)
	})
}
