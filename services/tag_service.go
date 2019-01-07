package services

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"time"
)

type (
	tagServiceImpl struct {
		tagEntityFactory contracts.TagEntityFactory
		tagRepository    contracts.TagRepository
	}
)

//NewTagService
func NewTagService(
	tagEntityFactory contracts.TagEntityFactory, tagRepository contracts.TagRepository,
) (tagService contracts.TagService) {
	return &tagServiceImpl{
		tagEntityFactory,
		tagRepository,
	}
}

//ListTags
func (s *tagServiceImpl) ListTags(tagPaginationQuery *models.TagPaginationQuery) (
	*models.PaginationResult, errors.Error,
) {
	return s.tagRepository.ListTags(tagPaginationQuery)
}

//GetTag
func (s *tagServiceImpl) GetTag(tagId *models.TagId) (*entities.TagEntity, errors.Error) {
	return s.tagRepository.GetTag(tagId)
}

//CreateTag
func (s *tagServiceImpl) CreateTag(data *models.TagCreate) (tagEntity *entities.TagEntity, err errors.Error) {
	tagEntity = s.tagEntityFactory.CreateTagEntity()
	tagEntity.Name = data.Name

	err = s.tagRepository.SaveTag(tagEntity)
	return
}

//UpdateTag
func (s *tagServiceImpl) UpdateTag(tagEntity *entities.TagEntity, data *models.TagUpdate) errors.Error {
	tagEntity.Name = data.Name

	updated := time.Now().UTC()
	tagEntity.Updated = &updated

	return s.tagRepository.SaveTag(tagEntity)
}

//DeleteTag
func (s *tagServiceImpl) DeleteTag(tagEntity *entities.TagEntity) errors.Error {
	return s.tagRepository.RemoveTag(tagEntity)
}

//GetTagXrefs
func (s *tagServiceImpl) GetTagXrefs(tagEntity *entities.TagEntity) (
	[]*entities.TagXrefEntity, errors.Error,
) {
	return s.tagRepository.GetTagXrefs(tagEntity)
}

//GetTagObjectXrefs
func (s *tagServiceImpl) GetTagObjectXrefs(tagObject models.Object) (
	[]*entities.TagXrefEntity, errors.Error,
) {
	return s.tagRepository.GetTagObjectXrefs(tagObject)
}

//GetTagXref
func (s *tagServiceImpl) GetTagXref(tagEntity *entities.TagEntity, tagObject models.Object) (
	*entities.TagXrefEntity, errors.Error,
) {
	return s.tagRepository.GetTagXref(tagEntity, tagObject)
}

//CreateTagXref
func (s *tagServiceImpl) CreateTagXref(tagEntity *entities.TagEntity, tagObject models.Object) (
	tagXrefEntity *entities.TagXrefEntity, err errors.Error,
) {
	tagXrefEntity = s.tagEntityFactory.CreateTagXrefEntity(tagEntity, tagObject)
	err = s.tagRepository.SaveTagXref(tagXrefEntity)
	return
}

//DeleteTagXref
func (s *tagServiceImpl) DeleteTagXref(tagXrefEntity *entities.TagXrefEntity) errors.Error {
	return s.tagRepository.RemoveTagXref(tagXrefEntity)
}

//ListObjectTags
func (s *tagServiceImpl) ListObjectTags(tagObject models.Object, tagPaginationQuery *models.TagPaginationQuery) (
	*models.PaginationResult, errors.Error,
) {
	return s.tagRepository.ListObjectTags(tagObject, tagPaginationQuery)
}
