package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

func NewBannerNotFoundError(request string) common.Error {
	message := fmt.Sprintf("The banner was not found on request: %s", request)
	return common.NewNotFoundError(message)
}

func NewBannerByIdNotFoundError(bannerId *models.BannerId) common.Error {
	request := fmt.Sprintf("id = %s", bannerId)
	return NewBannerNotFoundError(request)
}
