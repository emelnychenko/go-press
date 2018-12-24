package validators

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"strings"
)

type (
	contentTypeValidatorImpl struct {
	}
)

func NewContentTypeValidator() contracts.ContentTypeValidator {
	return &contentTypeValidatorImpl{}
}

func (*contentTypeValidatorImpl) ValidateImage(contentType string) (err common.Error) {
	if !strings.HasPrefix(contentType, "image/") {
		err = common.NewBadRequestError("ContentType is not Image")
	}

	return
}

func (*contentTypeValidatorImpl) ValidateVideo(contentType string) (err common.Error) {
	if !strings.HasPrefix(contentType, "video/") {
		err = common.NewBadRequestError("ContentType is not Video")
	}

	return
}

func (*contentTypeValidatorImpl) ValidateAudio(contentType string) (err common.Error) {
	if !strings.HasPrefix(contentType, "audio/") {
		err = common.NewBadRequestError("ContentType is not Audio")
	}

	return
}
