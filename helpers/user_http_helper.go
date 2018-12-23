package helpers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

const (
	UserIdParameterName = "userId"
)

type (
	userEchoHelperImpl struct {
	}
)

func NewUserEchoHelper() contracts.UserHttpHelper {
	return new(userEchoHelperImpl)
}

func (*userEchoHelperImpl) ParseUserId(httpContext contracts.HttpContext) (*models.UserId, common.Error) {
	return common.ParseModelId(httpContext.Parameter(UserIdParameterName))
}
