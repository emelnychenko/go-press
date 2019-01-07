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

	t.Run("GetTagXrefs", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		var tagXrefEntities []*entities.TagXrefEntity

		tagRepository := mocks.NewMockTagRepository(ctrl)
		tagRepository.EXPECT().GetTagXrefs(tagEntity).Return(tagXrefEntities, nil)

		tagService := &tagServiceImpl{tagRepository: tagRepository}
		results, err := tagService.GetTagXrefs(tagEntity)

		assert.Equal(t, tagXrefEntities, results)
		assert.Nil(t, err)
	})

	t.Run("GetTagObjectXrefs", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		var tagXrefEntities []*entities.TagXrefEntity

		tagRepository := mocks.NewMockTagRepository(ctrl)
		tagRepository.EXPECT().GetTagObjectXrefs(postEntity).Return(tagXrefEntities, nil)

		tagService := &tagServiceImpl{tagRepository: tagRepository}
		results, err := tagService.GetTagObjectXrefs(postEntity)

		assert.Equal(t, tagXrefEntities, results)
		assert.Nil(t, err)
	})

	t.Run("GetTagXref", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		postEntity := new(entities.PostEntity)
		tagXrefEntity := new(entities.TagXrefEntity)

		tagRepository := mocks.NewMockTagRepository(ctrl)
		tagRepository.EXPECT().GetTagXref(tagEntity, postEntity).Return(tagXrefEntity, nil)

		tagService := &tagServiceImpl{tagRepository: tagRepository}
		result, err := tagService.GetTagXref(tagEntity, postEntity)

		assert.Equal(t, tagXrefEntity, result)
		assert.Nil(t, err)
	})

	t.Run("CreateTagXref", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		postEntity := new(entities.PostEntity)
		tagXrefEntity := new(entities.TagXrefEntity)

		tagEntityFactory := mocks.NewMockTagEntityFactory(ctrl)
		tagEntityFactory.EXPECT().CreateTagXrefEntity(tagEntity, postEntity).Return(tagXrefEntity)

		tagRepository := mocks.NewMockTagRepository(ctrl)
		tagRepository.EXPECT().SaveTagXref(tagXrefEntity).Return(nil)

		tagService := &tagServiceImpl{
			tagRepository: tagRepository, tagEntityFactory: tagEntityFactory,
		}
		result, err := tagService.CreateTagXref(tagEntity, postEntity)

		assert.Equal(t, tagXrefEntity, result)
		assert.Nil(t, err)
	})

	t.Run("DeleteTagXref", func(t *testing.T) {
		tagXrefEntity := new(entities.TagXrefEntity)

		tagRepository := mocks.NewMockTagRepository(ctrl)
		tagRepository.EXPECT().RemoveTagXref(tagXrefEntity).Return(nil)

		tagService := &tagServiceImpl{tagRepository: tagRepository}
		err := tagService.DeleteTagXref(tagXrefEntity)

		assert.Nil(t, err)
	})

	t.Run("ListObjectTags", func(t *testing.T) {
		tagPaginationQuery := new(models.TagPaginationQuery)
		tagObject := mocks.NewMockObject(ctrl)

		var tagEntities *models.PaginationResult
		tagRepository := mocks.NewMockTagRepository(ctrl)
		tagRepository.EXPECT().ListObjectTags(tagObject, tagPaginationQuery).Return(tagEntities, nil)

		tagService := &tagServiceImpl{tagRepository: tagRepository}
		response, err := tagService.ListObjectTags(tagObject, tagPaginationQuery)

		assert.Equal(t, tagEntities, response)
		assert.Nil(t, err)
	})
}
