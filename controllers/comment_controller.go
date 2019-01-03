package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
)

type commentControllerImpl struct {
	commentHttpHelper   contracts.CommentHttpHelper
	commentModelFactory contracts.CommentModelFactory
	commentApi          contracts.CommentApi
}

func NewCommentController(
	commentHttpHelper contracts.CommentHttpHelper,
	commentModelFactory contracts.CommentModelFactory,
	commentApi contracts.CommentApi,
) (commentController contracts.CommentController) {
	return &commentControllerImpl{
		commentHttpHelper,
		commentModelFactory,
		commentApi,
	}
}

func (c *commentControllerImpl) ListComments(httpContext contracts.HttpContext) (paginationResult interface{}, err common.Error) {
	commentPaginationQuery := c.commentModelFactory.CreateCommentPaginationQuery()

	if err = httpContext.BindModel(commentPaginationQuery.PaginationQuery); err != nil {
		return
	}

	if err = httpContext.BindModel(commentPaginationQuery); err != nil {
		return
	}

	paginationResult, err = c.commentApi.ListComments(commentPaginationQuery)
	return
}

func (c *commentControllerImpl) GetComment(httpContext contracts.HttpContext) (comment interface{}, err common.Error) {
	commentId, err := c.commentHttpHelper.ParseCommentId(httpContext)

	if err != nil {
		return
	}

	comment, err = c.commentApi.GetComment(commentId)
	return
}

func (c *commentControllerImpl) CreateComment(httpContext contracts.HttpContext) (comment interface{}, err common.Error) {
	data := c.commentModelFactory.CreateCommentCreate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	comment, err = c.commentApi.CreateComment(data)
	return
}

func (c *commentControllerImpl) UpdateComment(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	commentId, err := c.commentHttpHelper.ParseCommentId(httpContext)

	if err != nil {
		return
	}

	data := c.commentModelFactory.CreateCommentUpdate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	err = c.commentApi.UpdateComment(commentId, data)
	return
}

func (c *commentControllerImpl) DeleteComment(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	commentId, err := c.commentHttpHelper.ParseCommentId(httpContext)

	if err != nil {
		return
	}

	err = c.commentApi.DeleteComment(commentId)
	return
}
