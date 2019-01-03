package workers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPostPublisherWorker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostPublisherWorker", func(t *testing.T) {
		postService := mocks.NewMockPostService(ctrl)
		postPublisherJob := mocks.NewMockPostPublisherJob(ctrl)

		postPublisherWorker, isPostPublisherWorker := NewPostPublisherWorker(
			postService,
			postPublisherJob,
		).(*postPublisherWorkerImpl)

		assert.True(t, isPostPublisherWorker)
		assert.Equal(t, time.Minute, postPublisherWorker.sleepingTime)
		assert.Equal(t, false, postPublisherWorker.working)
		assert.Equal(t, postService, postPublisherWorker.postService)
		assert.Equal(t, postPublisherJob, postPublisherWorker.postPublisherJob)
	})

	t.Run("Start", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEntities := []*entities.PostEntity{postEntity}

		postService := mocks.NewMockPostService(ctrl)

		postPublisherJob := mocks.NewMockPostPublisherJob(ctrl)
		postPublisherJob.EXPECT().PublishPost(postEntity).Return(nil)

		postPublisherWorker := &postPublisherWorkerImpl{
			sleepingTime:     time.Second,
			postService:      postService,
			postPublisherJob: postPublisherJob,
		}

		postService.EXPECT().GetScheduledPosts().Do(func() {
			// Assert enabled loop
			assert.True(t, postPublisherWorker.working)
			// Stop worker loop
			postPublisherWorker.working = false
		}).Return(postEntities, nil)

		err := postPublisherWorker.Start()
		assert.Nil(t, err)
		assert.False(t, postPublisherWorker.working)
	})

	t.Run("Start:GetScheduledPostsError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetScheduledPosts().Return(nil, systemErr)

		postPublisherWorker := &postPublisherWorkerImpl{
			sleepingTime: time.Second,
			postService:  postService,
		}

		err := postPublisherWorker.Start()
		assert.Equal(t, systemErr, err)
	})

	t.Run("Start:PublishPostError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		postEntity := new(entities.PostEntity)
		postEntities := []*entities.PostEntity{postEntity}

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetScheduledPosts().Return(postEntities, nil)

		postPublisherJob := mocks.NewMockPostPublisherJob(ctrl)
		postPublisherJob.EXPECT().PublishPost(postEntity).Return(systemErr)

		postPublisherWorker := &postPublisherWorkerImpl{
			sleepingTime:     time.Second,
			postService:      postService,
			postPublisherJob: postPublisherJob,
		}

		err := postPublisherWorker.Start()
		assert.Equal(t, systemErr, err)
	})

	t.Run("Stop", func(t *testing.T) {
		postPublisherWorker := &postPublisherWorkerImpl{working: true}
		err := postPublisherWorker.Stop()

		assert.Nil(t, err)
		assert.False(t, postPublisherWorker.working)
	})
}
