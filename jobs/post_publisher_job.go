package jobs

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/enums"
)

type (
	postPublisherJobImpl struct {
		eventDispatcher  contracts.EventDispatcher
		postEventFactory contracts.PostEventFactory
		postRepository   contracts.PostRepository
	}
)

func NewPostPublisherJob(
	eventDispatcher contracts.EventDispatcher,
	postEventFactory contracts.PostEventFactory,
	postRepository contracts.PostRepository,
) contracts.PostPublisherJob {
	return &postPublisherJobImpl{
		eventDispatcher,
		postEventFactory,
		postRepository,
	}
}

func (p *postPublisherJobImpl) PublishPost(postEntity *entities.PostEntity) (err common.Error) {
	postEntity.Status = enums.PostPublishedStatus
	err = p.postRepository.SavePost(postEntity)

	if nil != err {
		return
	}

	postUpdatedEvent := p.postEventFactory.CreatePostUpdatedEvent(postEntity)
	p.eventDispatcher.Dispatch(postUpdatedEvent)

	postPublishedEvent := p.postEventFactory.CreatePostPublishedEvent(postEntity)
	p.eventDispatcher.Dispatch(postPublishedEvent)

	return
}
