package echo

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/labstack/echo"
	"net/http"
)

type (
	routerImpl struct {
		echo           *echo.Echo
		modelValidator contracts.ModelValidator
	}
)

func NewRouter(echo *echo.Echo, modelValidator contracts.ModelValidator) (router contracts.Router) {
	return &routerImpl{echo: echo, modelValidator: modelValidator}
}

func (r *routerImpl) AddRoute(httpMethod string, routePath string, httpHandlerFunc contracts.HttpHandlerFunc) {
	r.echo.Add(httpMethod, routePath, func(context echo.Context) error {
		httpContext := NewHttpContext(context, r.modelValidator)
		response, err := httpHandlerFunc(httpContext)

		if context.Response().Committed {
			return err
		}

		if nil != err {
			return context.JSON(err.Code(), err.Error())
		}

		return context.JSON(http.StatusOK, response)
	})
}
