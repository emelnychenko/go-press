package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	categoryApiImpl struct {
		eventDispatcher      contracts.EventDispatcher
		categoryEventFactory contracts.CategoryEventFactory
		categoryService      contracts.CategoryService
		categoryAggregator   contracts.CategoryAggregator
	}
)

//NewCategoryApi
func NewCategoryApi(
	eventDispatcher contracts.EventDispatcher,
	categoryEventFactory contracts.CategoryEventFactory,
	categoryService contracts.CategoryService,
	categoryAggregator contracts.CategoryAggregator,
) (categoryApi contracts.CategoryApi) {
	return &categoryApiImpl{
		eventDispatcher,
		categoryEventFactory,
		categoryService,
		categoryAggregator,
	}
}

//ListCategories
func (a *categoryApiImpl) ListCategories(
	categoryPaginationQuery *models.CategoryPaginationQuery,
) (paginationResult *models.PaginationResult, err common.Error) {
	entityPaginationResult, err := a.categoryService.ListCategories(categoryPaginationQuery)

	if nil != err {
		return
	}

	paginationResult = a.categoryAggregator.AggregatePaginationResult(entityPaginationResult)
	return
}

//GetCategoriesTree
func (a *categoryApiImpl) GetCategoriesTree() (categoriesTree []*models.CategoryTree, err common.Error) {
	categoriesEntityTree, err := a.categoryService.GetCategoriesTree()

	if nil != err {
		return
	}

	categoriesTree = a.categoryAggregator.AggregateCategoriesTree(categoriesEntityTree)
	return
}

//GetCategory
func (a *categoryApiImpl) GetCategory(categoryId *models.CategoryId) (category *models.Category, err common.Error) {
	categoryEntity, err := a.categoryService.GetCategory(categoryId)

	if nil != err {
		return
	}

	category = a.categoryAggregator.AggregateCategory(categoryEntity)
	return
}

//GetCategoryTree
func (a *categoryApiImpl) GetCategoryTree(categoryId *models.CategoryId) (
	categoryTree *models.CategoryTree, err common.Error,
) {
	categoryEntityTree, err := a.categoryService.GetCategoryTree(categoryId)

	if nil != err {
		return
	}

	categoryTree = a.categoryAggregator.AggregateCategoryTree(categoryEntityTree)
	return
}

//CreateCategory
func (a *categoryApiImpl) CreateCategory(data *models.CategoryCreate) (category *models.Category, err common.Error) {
	categoryEntity, err := a.categoryService.CreateCategory(data)

	if nil != err {
		return
	}

	categoryCreatedEvent := a.categoryEventFactory.CreateCategoryCreatedEvent(categoryEntity)
	a.eventDispatcher.Dispatch(categoryCreatedEvent)

	category = a.categoryAggregator.AggregateCategory(categoryEntity)
	return
}

//UpdateCategory
func (a *categoryApiImpl) UpdateCategory(categoryId *models.CategoryId, data *models.CategoryUpdate) (err common.Error) {
	categoryService := a.categoryService
	categoryEntity, err := categoryService.GetCategory(categoryId)

	if nil != err {
		return
	}

	err = categoryService.UpdateCategory(categoryEntity, data)

	if nil != err {
		return
	}

	categoryUpdatedEvent := a.categoryEventFactory.CreateCategoryUpdatedEvent(categoryEntity)
	a.eventDispatcher.Dispatch(categoryUpdatedEvent)
	return
}

//ChangeCategoryParent
func (a *categoryApiImpl) ChangeCategoryParent(
	categoryId *models.CategoryId, parentCategoryId *models.CategoryId,
) (err common.Error) {
	categoryService := a.categoryService
	categoryEntity, err := categoryService.GetCategory(categoryId)

	if nil != err {
		return
	}

	parentCategoryEntity, err := categoryService.GetCategory(parentCategoryId)

	if nil != err {
		return
	}

	err = categoryService.ChangeCategoryParent(categoryEntity, parentCategoryEntity)

	if nil != err {
		return
	}

	categoryParentChangedEvent := a.categoryEventFactory.CreateCategoryParentChangedEvent(categoryEntity)
	a.eventDispatcher.Dispatch(categoryParentChangedEvent)
	return
}

//RemoveCategoryParent
func (a *categoryApiImpl) RemoveCategoryParent(categoryId *models.CategoryId) (err common.Error) {
	categoryService := a.categoryService
	categoryEntity, err := categoryService.GetCategory(categoryId)

	if nil != err {
		return
	}

	err = categoryService.RemoveCategoryParent(categoryEntity)

	if nil != err {
		return
	}

	categoryParentRemovedEvent := a.categoryEventFactory.CreateCategoryParentRemovedEvent(categoryEntity)
	a.eventDispatcher.Dispatch(categoryParentRemovedEvent)
	return
}

//DeleteCategory
func (a *categoryApiImpl) DeleteCategory(categoryId *models.CategoryId) (err common.Error) {
	categoryService := a.categoryService
	categoryEntity, err := categoryService.GetCategory(categoryId)

	if nil != err {
		return
	}

	err = categoryService.DeleteCategory(categoryEntity)

	if nil != err {
		return
	}

	categoryDeletedEvent := a.categoryEventFactory.CreateCategoryDeletedEvent(categoryEntity)
	a.eventDispatcher.Dispatch(categoryDeletedEvent)

	return
}
