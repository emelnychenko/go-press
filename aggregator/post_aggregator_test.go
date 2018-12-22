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
		postAggregator, isPostAggregator := NewPostAggregator(postModelFactory, subjectResolver).(*postAggregatorImpl)

		assert.True(t, isPostAggregator)
		assert.Equal(t, postModelFactory, postAggregator.postModelFactory)
		assert.Equal(t, subjectResolver, postAggregator.subjectResolver)
	})

	t.Run("AggregatePost", func(t *testing.T) {
		post := new(models.Post)
		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		postModelFactory.EXPECT().CreatePost().Return(post)

		systemUser := models.NewSystemUser()
		subjectResolver := mocks.NewMockSubjectResolver(ctrl)
		subjectResolver.EXPECT().ResolveSubject(gomock.Any(), gomock.Any()).Return(systemUser, nil)

		postAggregator := &postAggregatorImpl{postModelFactory: postModelFactory, subjectResolver: subjectResolver}
		response := postAggregator.AggregatePost(new(entities.PostEntity))

		assert.Equal(t, post, response)
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
