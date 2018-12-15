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

	t.Run("AggregateObject", func(t *testing.T) {
		subjectResolver := mocks.NewMockSubjectResolver(ctrl)
		postAggregator := NewPostAggregator(subjectResolver)
		systemUser := models.NewSystemUser()

		subjectResolver.EXPECT().ResolveSubject(gomock.Any(), gomock.Any()).Return(systemUser, nil)
		post := postAggregator.AggregateObject(new(entities.PostEntity))
		assert.IsType(t, new(models.Post), post)
	})

	t.Run("AggregateCollection", func(t *testing.T) {
		subjectResolver := mocks.NewMockSubjectResolver(ctrl)
		postAggregator := NewPostAggregator(subjectResolver)
		systemUser := models.NewSystemUser()
		postEntities := []*entities.PostEntity{entities.NewPostEntity()}

		subjectResolver.EXPECT().ResolveSubject(gomock.Any(), gomock.Any()).Return(systemUser, nil)
		posts := postAggregator.AggregateCollection(postEntities)
		assert.IsType(t, []*models.Post{}, posts)
		assert.Equal(t, len(postEntities), len(posts))
	})
}