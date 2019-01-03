package helpers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

const (
	TagIdParameterName = "tagId"
)

type (
	tagHttpHelperImpl struct {
	}
)

func NewTagHttpHelper() contracts.TagHttpHelper {
	return new(tagHttpHelperImpl)
}

func (*tagHttpHelperImpl) ParseTagId(httpContext contracts.HttpContext) (*models.TagId, common.Error) {
	return common.ParseModelId(httpContext.Parameter(TagIdParameterName))
}

