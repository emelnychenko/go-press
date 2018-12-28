package echo

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/labstack/echo"
	"mime/multipart"
	"net/http"
)

type (
	httpContextImpl struct {
		context echo.Context
	}
)

func NewHttpContext(context echo.Context) (httpContext contracts.HttpContext) {
	return &httpContextImpl{context: context}
}

func (c *httpContextImpl) Request() *http.Request {
	return c.context.Request()
}

func (c *httpContextImpl) Response() http.ResponseWriter {
	return c.context.Response()
}

func (c *httpContextImpl) Parameter(parameterName string) string {
	return c.context.Param(parameterName)
}

func (c *httpContextImpl) BindModel(data interface{}) (err common.Error) {
	echoErr := c.context.Bind(data)

	if nil != echoErr {
		err = common.NewBadRequestErrorFromBuiltin(echoErr)
	}

	return
}

func (c *httpContextImpl) FormFile(formFileName string) (fileHeader *multipart.FileHeader, err common.Error) {
	fileHeader, echoErr := c.context.FormFile(formFileName)

	if nil != echoErr {
		err = common.NewBadRequestErrorFromBuiltin(echoErr)
	}

	return
}