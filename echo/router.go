package echo

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/labstack/echo"
	"net/http"
)

type (
	routerImpl struct {
		echo *echo.Echo
	}
)

func NewRouter(echo *echo.Echo) (router contracts.Router) {
	return &routerImpl{echo: echo}
}

func (r *routerImpl) AddRoute(httpMethod string, routePath string, httpHandlerFunc contracts.HttpHandlerFunc) {
	r.echo.Add(httpMethod, routePath, func(context echo.Context) error {
		httpContext := NewHttpContext(context)
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

