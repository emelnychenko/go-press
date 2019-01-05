package aggregators

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type tagAggregatorImpl struct {
	tagModelFactory contracts.TagModelFactory
}

func NewTagAggregator(tagModelFactory contracts.TagModelFactory) contracts.TagAggregator {
	return &tagAggregatorImpl{tagModelFactory}
}

func (a *tagAggregatorImpl) AggregateTag(tagEntity *entities.TagEntity) (tag *models.Tag) {
	tag = a.tagModelFactory.CreateTag()
	tag.Id = tagEntity.Id
	tag.Name = tagEntity.Name
	tag.Created = tagEntity.Created

	return
}

func (a *tagAggregatorImpl) AggregateTags(tagEntities []*entities.TagEntity) (tags []*models.Tag) {
	tags = make([]*models.Tag, len(tagEntities))

	for k, tagEntity := range tagEntities {
		tags[k] = a.AggregateTag(tagEntity)
	}

	return
}

func (a *tagAggregatorImpl) AggregatePaginationResult(
	entityPaginationResult *models.PaginationResult,
) (
	paginationResult *models.PaginationResult,
) {
	tagEntities := entityPaginationResult.Data.([]*entities.TagEntity)
	tags := a.AggregateTags(tagEntities)
	return &models.PaginationResult{Total: entityPaginationResult.Total, Data: tags}
}
