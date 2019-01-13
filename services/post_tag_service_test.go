package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostTagService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostTagService", func(t *testing.T) {
		tagService := mocks.NewMockTagService(ctrl)
		postTagService, isPostTagService := NewPostTagService(tagService).(*postTagServiceImpl)

		assert.True(t, isPostTagService)
		assert.Equal(t, tagService, postTagService.tagService)
	})

	t.Run("ListPostTags", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		tagPaginationQuery := new(models.TagPaginationQuery)
		paginationResult := new(models.PaginationResult)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().ListObjectTags(postEntity, tagPaginationQuery).Return(paginationResult, nil)

		postTagService := &postTagServiceImpl{tagService: tagService}
		result, err := postTagService.ListPostTags(postEntity, tagPaginationQuery)

		assert.Nil(t, err)
		assert.Equal(t, result, paginationResult)
	})

	t.Run("AddPostTag", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		tagEntity := new(entities.TagEntity)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().CreateTagXref(tagEntity, postEntity).Return(nil, nil)

		postTagService := &postTagServiceImpl{tagService: tagService}
		err := postTagService.AddPostTag(postEntity, tagEntity)

		assert.Nil(t, err)
	})

	t.Run("RemovePostTag", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		tagEntity := new(entities.TagEntity)
		tagXrefEntity := new(entities.TagXrefEntity)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTagXref(tagEntity, postEntity).Return(tagXrefEntity, nil)
		tagService.EXPECT().DeleteTagXref(tagXrefEntity).Return(nil)

		postTagService := &postTagServiceImpl{tagService: tagService}
		err := postTagService.RemovePostTag(postEntity, tagEntity)

		assert.Nil(t, err)
	})

	t.Run("RemovePostTag:GetTagXrefError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postEntity := new(entities.PostEntity)
		tagEntity := new(entities.TagEntity)

		tagService := mocks.NewMockTagService(ctrl)
		tagService.EXPECT().GetTagXref(tagEntity, postEntity).Return(nil, systemErr)

		postTagService := &postTagServiceImpl{tagService: tagService}
		err := postTagService.RemovePostTag(postEntity, tagEntity)

		assert.Equal(t, systemErr, err)
	})
}
