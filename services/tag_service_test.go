package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewTagService", func(t *testing.T) {
		tagEntityFactory := mocks.NewMockTagEntityFactory(ctrl)
		tagRepository := mocks.NewMockTagRepository(ctrl)

		tagService, isTagService := NewTagService(tagEntityFactory, tagRepository).(*tagServiceImpl)

		assert.True(t, isTagService)
		assert.Equal(t, tagEntityFactory, tagService.tagEntityFactory)
		assert.Equal(t, tagRepository, tagService.tagRepository)
	})

	t.Run("ListTags", func(t *testing.T) {
		tagPaginationQuery := new(models.TagPaginationQuery)

		var tagEntities *models.PaginationResult
		tagRepository := mocks.NewMockTagRepository(ctrl)
		tagRepository.EXPECT().ListTags(tagPaginationQuery).Return(tagEntities, nil)

		tagService := &tagServiceImpl{tagRepository: tagRepository}
		response, err := tagService.ListTags(tagPaginationQuery)

		assert.Equal(t, tagEntities, response)
		assert.Nil(t, err)
	})

	t.Run("CreateTag", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		tagEntityFactory := mocks.NewMockTagEntityFactory(ctrl)
		tagEntityFactory.EXPECT().CreateTagEntity().Return(tagEntity)

		tagRepository := mocks.NewMockTagRepository(ctrl)
		tagRepository.EXPECT().SaveTag(tagEntity).Return(nil)

		data := &models.TagCreate{
			Name: "0",
		}
		tagService := &tagServiceImpl{
			tagEntityFactory: tagEntityFactory,
			tagRepository:    tagRepository,
		}
		response, err := tagService.CreateTag(data)

		assert.IsType(t, tagEntity, response)
		assert.Nil(t, err)
		assert.Equal(t, data.Name, tagEntity.Name)
	})

	t.Run("GetTag", func(t *testing.T) {
		tagId := new(models.TagId)
		tagEntity := new(entities.TagEntity)
		tagRepository := mocks.NewMockTagRepository(ctrl)
		tagRepository.EXPECT().GetTag(tagId).Return(tagEntity, nil)

		tagService := &tagServiceImpl{tagRepository: tagRepository}
		response, err := tagService.GetTag(tagId)

		assert.Equal(t, tagEntity, response)
		assert.Nil(t, err)
	})

	t.Run("UpdateTag", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		tagRepository := mocks.NewMockTagRepository(ctrl)
		tagRepository.EXPECT().SaveTag(tagEntity).Return(nil)

		data := &models.TagUpdate{
			Name: "0",
		}
		tagService := &tagServiceImpl{tagRepository: tagRepository}
		assert.Nil(t, tagService.UpdateTag(tagEntity, data))

		assert.Equal(t, data.Name, tagEntity.Name)
		assert.NotNil(t, tagEntity.Updated)
	})

	t.Run("DeleteTag", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		tagRepository := mocks.NewMockTagRepository(ctrl)
		tagRepository.EXPECT().RemoveTag(tagEntity).Return(nil)

		tagService := &tagServiceImpl{tagRepository: tagRepository}
		assert.Nil(t, tagService.DeleteTag(tagEntity))
	})
}
