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

func TestNewPostTagApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostTagApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		postTagEventFactory := mocks.NewMockPostTagEventFactory(ctrl)
		postService := mocks.NewMockPostService(ctrl)
		tagService := mocks.NewMockTagService(ctrl)
		postTagService := mocks.NewMockPostTagService(ctrl)
		tagAggregator := mocks.NewMockTagAggregator(ctrl)

		postTagApi, isPostTagApi := NewPostTagApi(
			eventDispatcher,
			postTagEventFactory,
			postService,
			tagService,
			postTagService,
			tagAggregator,
		).(*postTagApiImpl)

		assert.True(t, isPostTagApi)
		assert.Equal(t, eventDispatcher, postTagApi.eventDispatcher)
		assert.Equal(t, postTagEventFactory, postTagApi.postTagEventFactory)
		assert.Equal(t, postService, postTagApi.postService)
		assert.Equal(t, tagService, postTagApi.tagService)
		assert.Equal(t, postTagService, postTagApi.postTagService)
		assert.Equal(t, tagAggregator, postTagApi.tagAggregator)
	})

	t.Run("ListPostTags", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		tagPaginationQuery := new(models.TagPaginationQuery)
		entityPaginationQuery := new(models.PaginationResult)
		paginationResult := new(models.PaginationResult)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postTagService := mocks.NewMockPostTagService(ctrl)
		postTagService.EXPECT().ListPostTags(postEntity, tagPaginationQuery).
			Return(entityPaginationQuery, nil)

		tagAggregator := mocks.NewMockTagAggregator(ctrl)
		tagAggregator.EXPECT().AggregatePaginationResult(entityPaginationQuery).
			Return(paginationResult)

		postTagApi := &postTagApiImpl{
			postService:         postService,
			postTagService: postTagService,
			tagAggregator:  tagAggregator,
		}
		result, err := postTagApi.ListPostTags(postId, tagPaginationQuery)

		assert.Equal(t, result, paginationResult)
		assert.Nil(t, err)
	})

	t.Run("ListPostTags:GetPostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		tagPaginationQuery := new(models.TagPaginationQuery)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postTagApi := &postTagApiImpl{
			postService: postService,
		}
		_, err := postTagApi.ListPostTags(postId, tagPaginationQuery)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListPostTags:ListPostTagsError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		tagPaginationQuery := new(models.TagPaginationQuery)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		postTagService := mocks.NewMockPostTagService(ctrl)
		postTagService.EXPECT().ListPostTags(postEntity, tagPaginationQuery).
			Return(nil, systemErr)

		postTagApi := &postTagApiImpl{
			postService:         postService,
			postTagService: postTagService,
		}
		_, err := postTagApi.ListPostTags(postId, tagPaginationQuery)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostTag", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		tagId := new(models.TagId)
		tagEntity := new(entities.TagEntity)
		postTagEvent := new(events.PostTagEvent)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(tagEntity, nil)

		postTagService := mocks.NewMockPostTagService(ctrl)
		postTagService.EXPECT().AddPostTag(postEntity, tagEntity).
			Return(nil)

		postTagEventFactory := mocks.NewMockPostTagEventFactory(ctrl)
		postTagEventFactory.EXPECT().
			CreatePostTagAddedEvent(postEntity, tagEntity).
			Return(postTagEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postTagEvent)

		postTagApi := &postTagApiImpl{
			eventDispatcher:          eventDispatcher,
			postTagEventFactory: postTagEventFactory,
			postService:              postService,
			tagService:          tagService,
			postTagService:      postTagService,
		}
		err := postTagApi.AddPostTag(postId, tagId)
		assert.Nil(t, err)
	})

	t.Run("AddPostTag:GetPostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		tagId := new(models.TagId)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postTagApi := &postTagApiImpl{
			postService: postService,
		}
		err := postTagApi.AddPostTag(postId, tagId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostTag:GetTagError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		tagId := new(models.TagId)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(nil, systemErr)

		postTagApi := &postTagApiImpl{
			postService:     postService,
			tagService: tagService,
		}
		err := postTagApi.AddPostTag(postId, tagId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostTag:AddPostTagError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		tagId := new(models.TagId)
		tagEntity := new(entities.TagEntity)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(tagEntity, nil)

		postTagService := mocks.NewMockPostTagService(ctrl)
		postTagService.EXPECT().AddPostTag(postEntity, tagEntity).
			Return(systemErr)

		postTagApi := &postTagApiImpl{
			postService:         postService,
			tagService:     tagService,
			postTagService: postTagService,
		}
		err := postTagApi.AddPostTag(postId, tagId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostTag", func(t *testing.T) {
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		tagId := new(models.TagId)
		tagEntity := new(entities.TagEntity)
		postTagEvent := new(events.PostTagEvent)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(tagEntity, nil)

		postTagService := mocks.NewMockPostTagService(ctrl)
		postTagService.EXPECT().RemovePostTag(postEntity, tagEntity).
			Return(nil)

		postTagEventFactory := mocks.NewMockPostTagEventFactory(ctrl)
		postTagEventFactory.EXPECT().
			CreatePostTagRemovedEvent(postEntity, tagEntity).
			Return(postTagEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(postTagEvent)

		postTagApi := &postTagApiImpl{
			eventDispatcher:          eventDispatcher,
			postTagEventFactory: postTagEventFactory,
			postService:              postService,
			tagService:          tagService,
			postTagService:      postTagService,
		}
		err := postTagApi.RemovePostTag(postId, tagId)
		assert.Nil(t, err)
	})

	t.Run("RemovePostTag:GetPostError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		tagId := new(models.TagId)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(nil, systemErr)

		postTagApi := &postTagApiImpl{
			postService: postService,
		}
		err := postTagApi.RemovePostTag(postId, tagId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostTag:GetTagError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		tagId := new(models.TagId)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(nil, systemErr)

		postTagApi := &postTagApiImpl{
			postService:     postService,
			tagService: tagService,
		}
		err := postTagApi.RemovePostTag(postId, tagId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostTag:RemovePostTagError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postId := new(models.PostId)
		postEntity := new(entities.PostEntity)
		tagId := new(models.TagId)
		tagEntity := new(entities.TagEntity)

		postService := mocks.NewMockPostService(ctrl)
		postService.EXPECT().GetPost(postId).Return(postEntity, nil)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTag(tagId).Return(tagEntity, nil)

		postTagService := mocks.NewMockPostTagService(ctrl)
		postTagService.EXPECT().RemovePostTag(postEntity, tagEntity).
			Return(systemErr)

		postTagApi := &postTagApiImpl{
			postService:              postService,
			tagService:          tagService,
			postTagService:      postTagService,
		}
		err := postTagApi.RemovePostTag(postId, tagId)
		assert.Equal(t, systemErr, err)
	})
}
