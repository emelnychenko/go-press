package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"time"
)

type (
	tagServiceImpl struct {
		tagEntityFactory contracts.TagEntityFactory
		tagRepository    contracts.TagRepository
	}
)

func NewTagService(
	tagEntityFactory contracts.TagEntityFactory,
	tagRepository contracts.TagRepository,
) (tagService contracts.TagService) {
	return &tagServiceImpl{
		tagEntityFactory,
		tagRepository,
	}
}

func (s *tagServiceImpl) ListTags(
	tagPaginationQuery *models.TagPaginationQuery,
) (*models.PaginationResult, common.Error) {
	return s.tagRepository.ListTags(tagPaginationQuery)
}

func (s *tagServiceImpl) GetTag(tagId *models.TagId) (*entities.TagEntity, common.Error) {
	return s.tagRepository.GetTag(tagId)
}

func (s *tagServiceImpl) CreateTag(data *models.TagCreate) (tagEntity *entities.TagEntity, err common.Error) {
	tagEntity = s.tagEntityFactory.CreateTagEntity()
	tagEntity.Name = data.Name

	err = s.tagRepository.SaveTag(tagEntity)
	return
}

func (s *tagServiceImpl) UpdateTag(tagEntity *entities.TagEntity, data *models.TagUpdate) common.Error {
	tagEntity.Name = data.Name

	updated := time.Now().UTC()
	tagEntity.Updated = &updated

	return s.tagRepository.SaveTag(tagEntity)
}

func (s *tagServiceImpl) DeleteTag(tagEntity *entities.TagEntity) common.Error {
	return s.tagRepository.RemoveTag(tagEntity)
}
