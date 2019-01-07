package errors

import (
	"fmt"
)

func NewBannerNotFoundError(request string) Error {
	message := fmt.Sprintf("The banner was not found on request: %s", request)
	return NewNotFoundError(message)
}

func NewBannerByIdNotFoundError(bannerId fmt.Stringer) Error {
	request := fmt.Sprintf("id = %s", bannerId)
	return NewBannerNotFoundError(request)
}
