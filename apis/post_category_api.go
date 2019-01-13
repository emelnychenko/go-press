package apis

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	postCategoryApiImpl struct {
		eventDispatcher          contracts.EventDispatcher
		postCategoryEventFactory contracts.PostCategoryEventFactory
		postService              contracts.PostService
		categoryService          contracts.CategoryService
		postCategoryService      contracts.PostCategoryService
		categoryAggregator       contracts.CategoryAggregator
	}
)

//NewPostCategoryApi
func NewPostCategoryApi(
	eventDispatcher contracts.EventDispatcher,
	postCategoryEventFactory contracts.PostCategoryEventFactory,
	postService contracts.PostService,
	categoryService contracts.CategoryService,
	postCategoryService contracts.PostCategoryService,
	categoryAggregator contracts.CategoryAggregator,
) (postCategoryApi contracts.PostCategoryApi) {
	return &postCategoryApiImpl{
		eventDispatcher:          eventDispatcher,
		postCategoryEventFactory: postCategoryEventFactory,
		postService:              postService,
		categoryService:          categoryService,
		postCategoryService:      postCategoryService,
		categoryAggregator:       categoryAggregator,
	}
}

//ListPostCategories
func (a *postCategoryApiImpl) ListPostCategories(
	postId *models.PostId, categoryPaginationQuery *models.CategoryPaginationQuery,
) (
	paginationResult *models.PaginationResult, err errors.Error,
) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	entityPaginationResult, err := a.postCategoryService.ListPostCategories(postEntity, categoryPaginationQuery)

	if nil != err {
		return
	}

	paginationResult = a.categoryAggregator.AggregatePaginationResult(entityPaginationResult)
	return
}

//AddPostCategory
func (a *postCategoryApiImpl) AddPostCategory(postId *models.PostId, categoryId *models.CategoryId) (err errors.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	categoryEntity, err := a.categoryService.GetCategory(categoryId)

	if nil != err {
		return
	}

	err = a.postCategoryService.AddPostCategory(postEntity, categoryEntity)

	if nil != err {
		return
	}

	postCategoryEvent := a.postCategoryEventFactory.CreatePostCategoryAddedEvent(postEntity, categoryEntity)
	a.eventDispatcher.Dispatch(postCategoryEvent)

	return
}

//RemovePostCategory
func (a *postCategoryApiImpl) RemovePostCategory(postId *models.PostId, categoryId *models.CategoryId) (err errors.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	categoryEntity, err := a.categoryService.GetCategory(categoryId)

	if nil != err {
		return
	}

	err = a.postCategoryService.RemovePostCategory(postEntity, categoryEntity)

	if nil != err {
		return
	}

	postCategoryEvent := a.postCategoryEventFactory.CreatePostCategoryRemovedEvent(postEntity, categoryEntity)
	a.eventDispatcher.Dispatch(postCategoryEvent)

	return
}
