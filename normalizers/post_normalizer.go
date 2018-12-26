package normalizers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/enums"
	"time"
)

type (
	postNormalizerImpl struct {
	}
)

func NewPostNormalizer() (postNormalizer contracts.PostNormalizer) {
	return new(postNormalizerImpl)
}

func (*postNormalizerImpl) NormalizePostEntity(postEntity *entities.PostEntity) {
	switch postEntity.Status {
	case enums.PostPublishedStatus:
		postPublished := postEntity.Published
		currentTime := time.Now().UTC()

		if nil == postPublished {
			postEntity.Published = &currentTime
		} else if currentTime.Before(*postEntity.Published) {
			postEntity.Status = enums.PostScheduledStatus
		}
	case enums.PostScheduledStatus:
		postPublished := postEntity.Published
		currentTime := time.Now().UTC()

		if nil == postPublished {
			postEntity.Status = enums.PostPublishedStatus
			postEntity.Published = &currentTime
		} else if currentTime.After(*postEntity.Published) {
			postEntity.Status = enums.PostPublishedStatus
		}
	case enums.PostDraftStatus:
		postEntity.Published = nil
	}
}
