package apis

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/events"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostCategoryApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostCategoryApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		postCategoryEventFactory := mocks.NewMockPostCategoryEventFactory(ctrl)
		postService := mocks.NewMockPostService(ctrl)
		categoryService := mocks.NewMockCategoryService(ctrl)
		postCategoryService := mocks.NewMockPostCategoryService(ctrl)
		categoryAggregator := mocks.NewMockCategoryAggregator(ctrl)

		postCategoryApi, isPostCategoryApi := NewPostCategoryApi(
			eventDispatcher,
			postCategoryEventFactory,
			postService,
			categoryService,
			postCategoryService,
			categoryAggregator,
		).(*postCategoryApiImpl)

		assert.True(t, isPostCategoryApi)
		assert.Equal(t, eventDispatcher, postCategoryApi.eventDispatcher)
		assert.Equal(t, postCategoryEventFactory, postCategoryApi.postCategoryEventFactory)
		assert.Equal(t, postService, postCategoryApi.postService)
		assert.Equal(t, categoryService, postCategoryApi.categoryService)
		assert.Equal(t, postCategoryService, postCategoryApi.postCategoryService)
		assert.Equal(t, categoryAggregator, postCategoryApi.categoryAggregator)
	})

	t.Run("ListPostCategories", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		categoryPaginationQuery := new(models.CategoryPaginationQuery)
		entityPaginationQuery := new(models.PaginationResult)
		paginationResult := new(models.PaginationResult)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postCategoryService := mocks.NewMockPostCategoryService(ctrl)
		postCategoryService.EXPECT().ListPostCategories(postEntity, categoryPaginationQuery).
			Return(entityPaginationQuery, nil)

		categoryAggregator := mocks.NewMockCategoryAggregator(ctrl)
		categoryAggregator.EXPECT().AggregatePaginationResult(entityPaginationQuery).
			Return(paginationResult)

		postCategoryApi := &postCategoryApiImpl{
			postService:         postService,
			postCategoryService: postCategoryService,
			categoryAggregator:  categoryAggregator,
		}
		result, err := postCategoryApi.ListPostCategories(postId, categoryPaginationQuery)

		assert.Equal(t, result, paginationResult)
		assert.Nil(t, err)
	})

	t.Run("ListPostCategories:GetPostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		categoryPaginationQuery := new(models.CategoryPaginationQuery)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postCategoryApi := &postCategoryApiImpl{
			postService: postService,
		}
		_, err := postCategoryApi.ListPostCategories(postId, categoryPaginationQuery)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListPostCategories:ListPostCategoriesError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		categoryPaginationQuery := new(models.CategoryPaginationQuery)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postCategoryService := mocks.NewMockPostCategoryService(ctrl)
		postCategoryService.EXPECT().ListPostCategories(postEntity, categoryPaginationQuery).
			Return(nil, systemErr)

		postCategoryApi := &postCategoryApiImpl{
			postService:         postService,
			postCategoryService: postCategoryService,
		}
		_, err := postCategoryApi.ListPostCategories(postId, categoryPaginationQuery)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostCategory", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		categoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)
		postCategoryEvent := new(events.PostCategoryEvent)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(categoryEntity, nil)

		postCategoryService := mocks.NewMockPostCategoryService(ctrl)
		postCategoryService.EXPECT().AddPostCategory(postEntity, categoryEntity).
			Return(nil)

		postCategoryEventFactory := mocks.NewMockPostCategoryEventFactory(ctrl)
		postCategoryEventFactory.EXPECT().
			CreatePostCategoryAddedEvent(postEntity, categoryEntity).
			Return(postCategoryEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postCategoryEvent)

		postCategoryApi := &postCategoryApiImpl{
			eventDispatcher:          eventDispatcher,
			postCategoryEventFactory: postCategoryEventFactory,
			postService:              postService,
			categoryService:          categoryService,
			postCategoryService:      postCategoryService,
		}
		err := postCategoryApi.AddPostCategory(postId, categoryId)
		assert.Nil(t, err)
	})

	t.Run("AddPostCategory:GetPostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		categoryId := new(models.CategoryId)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postCategoryApi := &postCategoryApiImpl{
			postService: postService,
		}
		err := postCategoryApi.AddPostCategory(postId, categoryId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostCategory:GetCategoryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		categoryId := new(models.CategoryId)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(nil, systemErr)

		postCategoryApi := &postCategoryApiImpl{
			postService:     postService,
			categoryService: categoryService,
		}
		err := postCategoryApi.AddPostCategory(postId, categoryId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostCategory:AddPostCategoryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		categoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(categoryEntity, nil)

		postCategoryService := mocks.NewMockPostCategoryService(ctrl)
		postCategoryService.EXPECT().AddPostCategory(postEntity, categoryEntity).
			Return(systemErr)

		postCategoryApi := &postCategoryApiImpl{
			postService:         postService,
			categoryService:     categoryService,
			postCategoryService: postCategoryService,
		}
		err := postCategoryApi.AddPostCategory(postId, categoryId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostCategory", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		categoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)
		postCategoryEvent := new(events.PostCategoryEvent)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(categoryEntity, nil)

		postCategoryService := mocks.NewMockPostCategoryService(ctrl)
		postCategoryService.EXPECT().RemovePostCategory(postEntity, categoryEntity).
			Return(nil)

		postCategoryEventFactory := mocks.NewMockPostCategoryEventFactory(ctrl)
		postCategoryEventFactory.EXPECT().
			CreatePostCategoryRemovedEvent(postEntity, categoryEntity).
			Return(postCategoryEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postCategoryEvent)

		postCategoryApi := &postCategoryApiImpl{
			eventDispatcher:          eventDispatcher,
			postCategoryEventFactory: postCategoryEventFactory,
			postService:              postService,
			categoryService:          categoryService,
			postCategoryService:      postCategoryService,
		}
		err := postCategoryApi.RemovePostCategory(postId, categoryId)
		assert.Nil(t, err)
	})

	t.Run("RemovePostCategory:GetPostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		categoryId := new(models.CategoryId)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postCategoryApi := &postCategoryApiImpl{
			postService: postService,
		}
		err := postCategoryApi.RemovePostCategory(postId, categoryId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostCategory:GetCategoryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		categoryId := new(models.CategoryId)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(nil, systemErr)

		postCategoryApi := &postCategoryApiImpl{
			postService:     postService,
			categoryService: categoryService,
		}
		err := postCategoryApi.RemovePostCategory(postId, categoryId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostCategory:RemovePostCategoryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		categoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(categoryEntity, nil)

		postCategoryService := mocks.NewMockPostCategoryService(ctrl)
		postCategoryService.EXPECT().RemovePostCategory(postEntity, categoryEntity).
			Return(systemErr)

		postCategoryApi := &postCategoryApiImpl{
			postService:              postService,
			categoryService:          categoryService,
			postCategoryService:      postCategoryService,
		}
		err := postCategoryApi.RemovePostCategory(postId, categoryId)
		assert.Equal(t, systemErr, err)
	})
}
