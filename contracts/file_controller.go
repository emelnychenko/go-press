package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	FileController interface {
		ListFiles(httpContext HttpContext) (response interface{}, err common.Error)
		GetFile(httpContext HttpContext) (response interface{}, err common.Error)
		UploadFile(httpContext HttpContext) (response interface{}, err common.Error)
		DownloadFile(httpContext HttpContext) (_ interface{}, err common.Error)
		UpdateFile(httpContext HttpContext) (_ interface{}, err common.Error)
		DeleteFile(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
