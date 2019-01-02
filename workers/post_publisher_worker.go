package workers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"time"
)

type (
	postPublisherWorkerImpl struct {
		sleepingTime     time.Duration
		working          bool
		postService      contracts.PostService
		postPublisherJob contracts.PostPublisherJob
	}
)

func NewPostPublisherWorker(
	postService contracts.PostService,
	postPublisher contracts.PostPublisherJob,
) contracts.PostPublisherWorker {
	return &postPublisherWorkerImpl{
		sleepingTime:     time.Minute,
		working:          false,
		postService:      postService,
		postPublisherJob: postPublisher,
	}
}

func (w *postPublisherWorkerImpl) Start() (err common.Error) {
	w.working = true
	for w.working {
		postEntities, serviceErr := w.postService.GetScheduledPosts()

		if nil != serviceErr {
			w.working = false
			err = serviceErr
			return
		}

		if 0 < len(postEntities) {
			queueErrs := make(chan common.Error, len(postEntities))

			for _, postEntity := range postEntities {
				go func(postEntity *entities.PostEntity) {
					queueErrs <- w.postPublisherJob.PublishPost(postEntity)
				}(postEntity)
			}

			for range postEntities {
				if queueErr := <-queueErrs; nil != queueErr {
					err = queueErr
					close(queueErrs)
					return
				}
			}

			close(queueErrs)
		}

		time.Sleep(w.sleepingTime)
	}
	return
}

func (w *postPublisherWorkerImpl) Stop() (err common.Error) {
	w.working = false
	return
}
