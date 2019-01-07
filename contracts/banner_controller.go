package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	BannerController interface {
		ListBanners(httpContext HttpContext) (response interface{}, err errors.Error)
		GetBanner(httpContext HttpContext) (response interface{}, err errors.Error)
		CreateBanner(httpContext HttpContext) (response interface{}, err errors.Error)
		UpdateBanner(httpContext HttpContext) (_ interface{}, err errors.Error)
		DeleteBanner(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
