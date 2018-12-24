package contracts

import "github.com/emelnychenko/go-press/common"

type (
	ContentTypeValidator interface {
		ValidateImage(contentType string) common.Error
		ValidateVideo(contentType string) common.Error
		ValidateAudio(contentType string) common.Error
	}
)
