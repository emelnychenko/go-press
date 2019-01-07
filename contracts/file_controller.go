package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	FileController interface {
		ListFiles(httpContext HttpContext) (response interface{}, err errors.Error)
		GetFile(httpContext HttpContext) (response interface{}, err errors.Error)
		UploadFile(httpContext HttpContext) (response interface{}, err errors.Error)
		DownloadFile(httpContext HttpContext) (_ interface{}, err errors.Error)
		UpdateFile(httpContext HttpContext) (_ interface{}, err errors.Error)
		DeleteFile(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
