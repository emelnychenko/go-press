package aggregator

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostAggregator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostAggregator", func(t *testing.T) {
		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		subjectResolver := mocks.NewMockSubjectResolver(ctrl)
		fileApi := mocks.NewMockFileApi(ctrl)
		postAggregator, isPostAggregator := NewPostAggregator(postModelFactory, subjectResolver, fileApi).(*postAggregatorImpl)

		assert.True(t, isPostAggregator)
		assert.Equal(t, postModelFactory, postAggregator.postModelFactory)
		assert.Equal(t, subjectResolver, postAggregator.subjectResolver)
		assert.Equal(t, fileApi, postAggregator.fileApi)
	})

	t.Run("AggregatePost", func(t *testing.T) {
		fileApi := mocks.NewMockFileApi(ctrl)
		postPictureId := new(models.FileId)
		postPicture := new(models.File)
		fileApi.EXPECT().GetFile(postPictureId).Return(postPicture, nil)

		postVideoId := new(models.FileId)
		postVideo := new(models.File)
		fileApi.EXPECT().GetFile(postVideoId).Return(postVideo, nil)

		post := new(models.Post)
		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		postModelFactory.EXPECT().CreatePost().Return(post)

		systemUser := models.NewSystemUser()
		subjectResolver := mocks.NewMockSubjectResolver(ctrl)
		subjectResolver.EXPECT().ResolveSubject(gomock.Any(), gomock.Any()).Return(systemUser, nil)

		postEntity := &entities.PostEntity{PictureId: postPictureId, VideoId: postVideoId}
		postAggregator := &postAggregatorImpl{
			postModelFactory: postModelFactory,
			subjectResolver: subjectResolver,
			fileApi: fileApi,
		}
		response := postAggregator.AggregatePost(postEntity)
		assert.Equal(t, post, response)
		assert.Equal(t, postPicture, response.Picture)
		assert.Equal(t, postVideo, response.Video)
	})

	t.Run("AggregatePosts", func(t *testing.T) {
		post := new(models.Post)
		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		postModelFactory.EXPECT().CreatePost().Return(post)

		systemUser := models.NewSystemUser()
		subjectResolver := mocks.NewMockSubjectResolver(ctrl)
		subjectResolver.EXPECT().ResolveSubject(gomock.Any(), gomock.Any()).Return(systemUser, nil)

		postEntities := []*entities.PostEntity{entities.NewPostEntity()}
		postAggregator := &postAggregatorImpl{postModelFactory: postModelFactory, subjectResolver: subjectResolver}
		posts := postAggregator.AggregatePosts(postEntities)

		assert.IsType(t, []*models.Post{}, posts)
		assert.Equal(t, len(postEntities), len(posts))
	})
}
