package services

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	postTagServiceImpl struct {
		tagService contracts.TagService
	}
)

//NewPostTagService
func NewPostTagService(tagService contracts.TagService) contracts.PostTagService {
	return &postTagServiceImpl{tagService: tagService}
}

//ListPostTags
func (s *postTagServiceImpl) ListPostTags(
	postEntity *entities.PostEntity, tagPaginationQuery *models.TagPaginationQuery,
) (*models.PaginationResult, errors.Error) {
	return s.tagService.ListObjectTags(postEntity, tagPaginationQuery)
}

//AddPostTag
func (s *postTagServiceImpl) AddPostTag(
	postEntity *entities.PostEntity, tagEntity *entities.TagEntity,
) (err errors.Error) {
	_, err = s.tagService.CreateTagXref(tagEntity, postEntity)
	return
}

//RemovePostTag
func (s *postTagServiceImpl) RemovePostTag(
	postEntity *entities.PostEntity, tagEntity *entities.TagEntity,
) (err errors.Error) {
	tagXrefEntity, err := s.tagService.GetTagXref(tagEntity, postEntity)

	if nil != err {
		return
	}

	return s.tagService.DeleteTagXref(tagXrefEntity)
}
