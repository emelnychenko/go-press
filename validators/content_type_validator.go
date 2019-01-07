package validators

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"strings"
)

type (
	contentTypeValidatorImpl struct {
	}
)

func NewContentTypeValidator() contracts.ContentTypeValidator {
	return &contentTypeValidatorImpl{}
}

func (*contentTypeValidatorImpl) ValidateImage(contentType string) (err errors.Error) {
	if !strings.HasPrefix(contentType, "image/") {
		err = errors.NewBadRequestError("ContentType is not Image")
	}

	return
}

func (*contentTypeValidatorImpl) ValidateVideo(contentType string) (err errors.Error) {
	if !strings.HasPrefix(contentType, "video/") {
		err = errors.NewBadRequestError("ContentType is not Video")
	}

	return
}

func (*contentTypeValidatorImpl) ValidateAudio(contentType string) (err errors.Error) {
	if !strings.HasPrefix(contentType, "audio/") {
		err = errors.NewBadRequestError("ContentType is not Audio")
	}

	return
}
