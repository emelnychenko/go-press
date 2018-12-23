package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"mime/multipart"
	"net/http"
)

type (
	HttpContext interface {
		Request() *http.Request
		Response() http.ResponseWriter
		Parameter(parameterName string) string
		FormFile(formFileName string) (*multipart.FileHeader, common.Error)
		BindModel(data interface{}) common.Error
	}
)
