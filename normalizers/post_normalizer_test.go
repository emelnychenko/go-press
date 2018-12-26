package normalizers

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/enums"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPostNormalizer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostNormalizer", func(t *testing.T) {
		_, isPostNormalizer := NewPostNormalizer().(*postNormalizerImpl)

		assert.True(t, isPostNormalizer)
	})

	t.Run("NormalizePostEntity:PostPublishedStatus:EmptyPostPublished", func(t *testing.T) {
		postNormalizer := &postNormalizerImpl{}
		postEntity := &entities.PostEntity{Status: enums.PostPublishedStatus}

		postNormalizer.NormalizePostEntity(postEntity)

		assert.Equal(t, enums.PostPublishedStatus, postEntity.Status)
		assert.NotNil(t, postEntity.Published)
	})

	t.Run("NormalizePostEntity:PostPublishedStatus:PostPublishedInFuture", func(t *testing.T) {
		postNormalizer := &postNormalizerImpl{}
		postPublished := time.Now().UTC().Add(time.Hour)
		postEntity := &entities.PostEntity{Status: enums.PostPublishedStatus, Published: &postPublished}

		postNormalizer.NormalizePostEntity(postEntity)

		assert.Equal(t, enums.PostScheduledStatus, postEntity.Status)
		assert.True(t, postEntity.Published.Equal(postPublished))
	})

	t.Run("NormalizePostEntity:PostScheduledStatus:EmptyPostPublished", func(t *testing.T) {
		postNormalizer := &postNormalizerImpl{}
		postEntity := &entities.PostEntity{Status: enums.PostScheduledStatus}

		postNormalizer.NormalizePostEntity(postEntity)

		assert.Equal(t, enums.PostPublishedStatus, postEntity.Status)
		assert.NotNil(t, postEntity.Published)
	})

	t.Run("NormalizePostEntity:PostScheduledStatus:PostPublishedInPast", func(t *testing.T) {
		postNormalizer := &postNormalizerImpl{}
		postPublished := time.Now().UTC().Add(-time.Hour)
		postEntity := &entities.PostEntity{Status: enums.PostScheduledStatus, Published: &postPublished}

		postNormalizer.NormalizePostEntity(postEntity)

		assert.Equal(t, enums.PostPublishedStatus, postEntity.Status)
		assert.True(t, postEntity.Published.Equal(postPublished))
	})

	t.Run("NormalizePostEntity:PostDraftStatus", func(t *testing.T) {
		postNormalizer := &postNormalizerImpl{}
		postPublished := time.Now().UTC()
		postEntity := &entities.PostEntity{Status: enums.PostDraftStatus, Published: &postPublished}

		postNormalizer.NormalizePostEntity(postEntity)

		assert.Equal(t, enums.PostDraftStatus, postEntity.Status)
		assert.Nil(t, postEntity.Published)
	})
}
