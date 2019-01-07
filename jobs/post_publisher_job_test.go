package jobs

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/events"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostPublisherJob(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostPublisherJob", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		postEventFactory := mocks.NewMockPostEventFactory(ctrl)
		postRepository := mocks.NewMockPostRepository(ctrl)

		postPublisherJob, isPostPublisherJob := NewPostPublisherJob(
			eventDispatcher,
			postEventFactory,
			postRepository,
		).(*postPublisherJobImpl)

		assert.True(t, isPostPublisherJob)
		assert.Equal(t, eventDispatcher, postPublisherJob.eventDispatcher)
		assert.Equal(t, postEventFactory, postPublisherJob.postEventFactory)
		assert.Equal(t, postRepository, postPublisherJob.postRepository)
	})

	t.Run("PublishPost", func(t *testing.T) {
		postEntity := new(entities.PostEntity)

		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().SavePost(postEntity).Return(nil)

		postEventFactory := mocks.NewMockPostEventFactory(ctrl)
		postUpdatedEvent := new(events.PostEvent)
		postEventFactory.EXPECT().CreatePostUpdatedEvent(postEntity).Return(postUpdatedEvent)
		postPublishedEvent := new(events.PostEvent)
		postEventFactory.EXPECT().CreatePostPublishedEvent(postEntity).Return(postPublishedEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postUpdatedEvent)
		eventDispatcher.EXPECT().Dispatch(postPublishedEvent)

		postPublisherJob := &postPublisherJobImpl{
			eventDispatcher:  eventDispatcher,
			postEventFactory: postEventFactory,
			postRepository:   postRepository,
		}

		err := postPublisherJob.PublishPost(postEntity)
		assert.Nil(t, err)
	})

	t.Run("PublishPost:SavePostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postEntity := new(entities.PostEntity)

		postRepository := mocks.NewMockPostRepository(ctrl)
		postRepository.EXPECT().SavePost(postEntity).Return(systemErr)

		postPublisherJob := &postPublisherJobImpl{
			postRepository: postRepository,
		}

		err := postPublisherJob.PublishPost(postEntity)
		assert.Equal(t, systemErr, err)
	})
}
