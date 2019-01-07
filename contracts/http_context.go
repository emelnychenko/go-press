package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"mime/multipart"
	"net/http"
)

type (
	HttpContext interface {
		Request() *http.Request
		Response() http.ResponseWriter
		Parameter(parameterName string) string
		FormFile(formFileName string) (*multipart.FileHeader, errors.Error)
		BindModel(data interface{}) errors.Error
	}
)
