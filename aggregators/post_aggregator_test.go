package aggregators

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
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryApi := mocks.NewMockCategoryApi(ctrl)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagApi := mocks.NewMockTagApi(ctrl)

		postAggregator, isPostAggregator := NewPostAggregator(
			postModelFactory,
			subjectResolver,
			fileApi,
			categoryModelFactory,
			categoryApi,
			tagModelFactory,
			tagApi,
		).(*postAggregatorImpl)

		assert.True(t, isPostAggregator)
		assert.Equal(t, postModelFactory, postAggregator.postModelFactory)
		assert.Equal(t, subjectResolver, postAggregator.subjectResolver)
		assert.Equal(t, fileApi, postAggregator.fileApi)
		assert.Equal(t, categoryModelFactory, postAggregator.categoryModelFactory)
		assert.Equal(t, categoryApi, postAggregator.categoryApi)
		assert.Equal(t, tagModelFactory, postAggregator.tagModelFactory)
		assert.Equal(t, tagApi, postAggregator.tagApi)
	})

	t.Run("AggregatePost", func(t *testing.T) {
		fileApi := mocks.NewMockFileApi(ctrl)
		postPictureId := new(models.FileId)
		postPicture := new(models.File)
		fileApi.EXPECT().GetFile(postPictureId).Return(postPicture, nil)

		postVideoId := new(models.FileId)
		postVideo := new(models.File)
		fileApi.EXPECT().GetFile(postVideoId).Return(postVideo, nil)

		postEntity := &entities.PostEntity{PictureId: postPictureId, VideoId: postVideoId}
		post := new(models.Post)
		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		postModelFactory.EXPECT().CreatePost().Return(post)

		categoryPaginationQuery := new(models.CategoryPaginationQuery)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryPaginationQuery().Return(categoryPaginationQuery)

		categoryPaginationResult := new(models.PaginationResult)
		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().ListObjectCategories(postEntity, categoryPaginationQuery).Return(categoryPaginationResult, nil)

		tagPaginationQuery := new(models.TagPaginationQuery)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagPaginationQuery().Return(tagPaginationQuery)

		tagPaginationResult := new(models.PaginationResult)
		tagApi := mocks.NewMockTagApi(ctrl)
		tagApi.EXPECT().ListObjectTags(postEntity, tagPaginationQuery).Return(tagPaginationResult, nil)

		systemUser := models.NewSystemUser()
		subjectResolver := mocks.NewMockSubjectResolver(ctrl)
		subjectResolver.EXPECT().ResolveSubject(gomock.Any(), gomock.Any()).Return(systemUser, nil)

		postAggregator := &postAggregatorImpl{
			postModelFactory:     postModelFactory,
			subjectResolver:      subjectResolver,
			fileApi:              fileApi,
			categoryModelFactory: categoryModelFactory,
			categoryApi:          categoryApi,
			tagModelFactory:      tagModelFactory,
			tagApi:               tagApi,
		}
		response := postAggregator.AggregatePost(postEntity)
		assert.Equal(t, post, response)
		assert.Equal(t, postPicture, response.Picture)
		assert.Equal(t, postVideo, response.Video)
	})

	t.Run("AggregatePosts", func(t *testing.T) {
		postEntities := []*entities.PostEntity{
			entities.NewPostEntity(),
		}

		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		subjectResolver := mocks.NewMockSubjectResolver(ctrl)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryApi := mocks.NewMockCategoryApi(ctrl)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagApi := mocks.NewMockTagApi(ctrl)

		for _, postEntity := range postEntities {
			post := new(models.Post)
			postModelFactory.EXPECT().CreatePost().Return(post)

			systemUser := models.NewSystemUser()
			subjectResolver.EXPECT().ResolveSubject(gomock.Any(), gomock.Any()).Return(systemUser, nil)

			categoryPaginationQuery := new(models.CategoryPaginationQuery)
			categoryModelFactory.EXPECT().CreateCategoryPaginationQuery().Return(categoryPaginationQuery)

			categoryPaginationResult := new(models.PaginationResult)
			categoryApi.EXPECT().ListObjectCategories(postEntity, categoryPaginationQuery).Return(categoryPaginationResult, nil)

			tagPaginationQuery := new(models.TagPaginationQuery)
			tagModelFactory.EXPECT().CreateTagPaginationQuery().Return(tagPaginationQuery)

			tagPaginationResult := new(models.PaginationResult)
			tagApi.EXPECT().ListObjectTags(postEntity, tagPaginationQuery).Return(tagPaginationResult, nil)
		}

		postAggregator := &postAggregatorImpl{
			postModelFactory:     postModelFactory,
			subjectResolver:      subjectResolver,
			categoryModelFactory: categoryModelFactory,
			categoryApi:          categoryApi,
			tagModelFactory:      tagModelFactory,
			tagApi:               tagApi,
		}
		posts := postAggregator.AggregatePosts(postEntities)

		assert.IsType(t, []*models.Post{}, posts)
		assert.Equal(t, len(postEntities), len(posts))
	})

	t.Run("AggregatePaginationResult", func(t *testing.T) {
		postEntities := []*entities.PostEntity{entities.NewPostEntity()}

		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		subjectResolver := mocks.NewMockSubjectResolver(ctrl)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryApi := mocks.NewMockCategoryApi(ctrl)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagApi := mocks.NewMockTagApi(ctrl)

		for _, postEntity := range postEntities {
			post := new(models.Post)
			postModelFactory.EXPECT().CreatePost().Return(post)

			systemUser := models.NewSystemUser()
			subjectResolver.EXPECT().ResolveSubject(gomock.Any(), gomock.Any()).Return(systemUser, nil)

			categoryPaginationQuery := new(models.CategoryPaginationQuery)
			categoryModelFactory.EXPECT().CreateCategoryPaginationQuery().Return(categoryPaginationQuery)

			categoryPaginationResult := new(models.PaginationResult)
			categoryApi.EXPECT().ListObjectCategories(postEntity, categoryPaginationQuery).Return(categoryPaginationResult, nil)

			tagPaginationQuery := new(models.TagPaginationQuery)
			tagModelFactory.EXPECT().CreateTagPaginationQuery().Return(tagPaginationQuery)

			tagPaginationResult := new(models.PaginationResult)
			tagApi.EXPECT().ListObjectTags(postEntity, tagPaginationQuery).Return(tagPaginationResult, nil)
		}

		postAggregator := &postAggregatorImpl{
			postModelFactory:     postModelFactory,
			subjectResolver:      subjectResolver,
			categoryModelFactory: categoryModelFactory,
			categoryApi:          categoryApi,
			tagModelFactory:      tagModelFactory,
			tagApi:               tagApi,
		}

		entityPaginationResult := &models.PaginationResult{Data: postEntities}
		paginationResult := postAggregator.AggregatePaginationResult(entityPaginationResult)

		assert.IsType(t, []*models.Post{}, paginationResult.Data)
		assert.Equal(t, len(postEntities), len(paginationResult.Data.([]*models.Post)))
	})
}
