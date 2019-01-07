package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	ContentTypeValidator interface {
		ValidateImage(contentType string) errors.Error
		ValidateVideo(contentType string) errors.Error
		ValidateAudio(contentType string) errors.Error
	}
)
