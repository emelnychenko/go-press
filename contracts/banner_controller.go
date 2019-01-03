package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	BannerController interface {
		ListBanners(httpContext HttpContext) (response interface{}, err common.Error)
		GetBanner(httpContext HttpContext) (response interface{}, err common.Error)
		CreateBanner(httpContext HttpContext) (response interface{}, err common.Error)
		UpdateBanner(httpContext HttpContext) (_ interface{}, err common.Error)
		DeleteBanner(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
