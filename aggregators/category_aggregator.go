package aggregators

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type categoryAggregatorImpl struct {
	categoryModelFactory contracts.CategoryModelFactory
}

//NewCategoryAggregator
func NewCategoryAggregator(categoryModelFactory contracts.CategoryModelFactory) contracts.CategoryAggregator {
	return &categoryAggregatorImpl{categoryModelFactory}
}

//AggregateCategory
func (a *categoryAggregatorImpl) AggregateCategory(categoryEntity *entities.CategoryEntity) (category *models.Category) {
	category = a.categoryModelFactory.CreateCategory()
	category.Id = categoryEntity.Id
	category.Name = categoryEntity.Name
	category.Created = categoryEntity.Created

	return
}

//AggregateCategoryTree
func (a *categoryAggregatorImpl) AggregateCategoryTree(
	categoryEntityTree *entities.CategoryEntityTree,
) (category *models.CategoryTree) {
	categoryEntityTreeRoot := categoryEntityTree.Roots[0]
	category = a.aggregateCategoryTreeBranch(categoryEntityTreeRoot)
	return
}

func (a *categoryAggregatorImpl) aggregateCategoryTreeBranch(
	categoryEntityTreeBranch *entities.CategoryEntityTreeBranch,
) (categoryTree *models.CategoryTree) {
	category := a.AggregateCategory(categoryEntityTreeBranch.CategoryEntity)

	categoryTree = a.categoryModelFactory.CreateCategoryTree()
	categoryTree.Category = category

	var childCategories []*models.CategoryTree
	for _, categoryEntityTreeBranchChild := range categoryEntityTreeBranch.Children {
		childCategory := a.aggregateCategoryTreeBranch(categoryEntityTreeBranchChild)
		childCategories = append(childCategories, childCategory)
	}

	categoryTree.Categories = childCategories
	return
}

//AggregateCategories
func (a *categoryAggregatorImpl) AggregateCategories(
	categoryEntities []*entities.CategoryEntity,
) (categories []*models.Category) {
	categories = make([]*models.Category, len(categoryEntities))

	for k, categoryEntity := range categoryEntities {
		categories[k] = a.AggregateCategory(categoryEntity)
	}

	return
}

//AggregateCategoriesTree
func (a *categoryAggregatorImpl) AggregateCategoriesTree(
	categoryEntityTree *entities.CategoryEntityTree,
) (categoriesTree []*models.CategoryTree) {
	categoriesTree = make([]*models.CategoryTree, len(categoryEntityTree.Roots))

	for k, categoryEntityTreeBranch := range categoryEntityTree.Roots {
		categoriesTree[k] = a.aggregateCategoryTreeBranch(categoryEntityTreeBranch)
	}

	return
}

//AggregatePaginationResult
func (a *categoryAggregatorImpl) AggregatePaginationResult(
	entityPaginationResult *models.PaginationResult,
) (
	paginationResult *models.PaginationResult,
) {
	categoryEntities := entityPaginationResult.Data.([]*entities.CategoryEntity)
	categories := a.AggregateCategories(categoryEntities)
	return &models.PaginationResult{Total: entityPaginationResult.Total, Data: categories}
}
