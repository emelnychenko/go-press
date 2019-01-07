package jobs

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
)

type (
	categoryEdgesBuilderJobImpl struct {
		categoryRepository       contracts.CategoryRepository
		categoryNestedSetBuilder contracts.CategoryNestedSetBuilder
	}
)

//NewCategoryEdgesBuilderJob
func NewCategoryEdgesBuilderJob(
	categoryRepository contracts.CategoryRepository,
	categoryNestedSetBuilder contracts.CategoryNestedSetBuilder,
) contracts.CategoryEdgesBuilderJob {
	return &categoryEdgesBuilderJobImpl{
		categoryRepository,
		categoryNestedSetBuilder,
	}
}

//BuildCategoriesEdges
func (j *categoryEdgesBuilderJobImpl) BuildCategoriesEdges() (err errors.Error) {
	categoryEntities, err := j.categoryRepository.GetCategories()

	if nil != err {
		return
	}

	categoryEntityNestedSet, err := j.categoryNestedSetBuilder.BuildCategoryEntityNestedSet(categoryEntities)

	if nil != err {
		return
	}

	for _, categoryEntityNestedSetNode := range categoryEntityNestedSet.Nodes {
		if categoryEntityNestedSetNode.EdgesDifferent() {
			categoryEntityNestedSetNode.SetEntityEdges()
			err = j.categoryRepository.SaveCategory(categoryEntityNestedSetNode.CategoryEntity)

			if nil != err {
				return
			}
		}
	}

	return
}
