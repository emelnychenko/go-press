package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
	"github.com/labstack/echo"
	"net/http"
)

const PostId = "postId"

type PostController struct {
	postApi contracts.PostApi
}

func ParsePostIdFromContext(context echo.Context) (*models.PostId, error) {
	return common.ParseModelId(context.Param(PostId))
}

func NewPostController(postApi contracts.PostApi) (c *PostController) {
	return &PostController{postApi}
}

func (c *PostController) ListPosts(context echo.Context) error {
	posts, err := c.postApi.ListPosts()

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, posts)
}

func (c *PostController) GetPost(context echo.Context) error {
	postId, err := ParsePostIdFromContext(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	post, err2 := c.postApi.GetPost(postId)

	if err2 != nil {
		return context.JSON(err2.Code(), err2)
	}

	return context.JSON(http.StatusOK, post)
}

func (c *PostController) CreatePost(context echo.Context) error {
	data := new(models.PostCreate)

	if err := context.Bind(data); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	systemUser := models.NewSystemUser() // TODO: Replace stub
	post, err := c.postApi.CreatePost(systemUser, data)

	if err != nil {
		return context.JSON(err.Code(), err)
	}

	return context.JSON(http.StatusOK, post)
}

func (c *PostController) UpdatePost(context echo.Context) error {
	postId, err := ParsePostIdFromContext(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	data := new(models.PostUpdate)

	if err := context.Bind(data); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	if err2 := c.postApi.UpdatePost(postId, data); err2 != nil {
		return context.JSON(err2.Code(), err2)
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *PostController) DeletePost(context echo.Context) error {
	postId, err := ParsePostIdFromContext(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	if err := c.postApi.DeletePost(postId); err != nil {
		return context.JSON(err.Code(), err)
	}

	return context.JSON(http.StatusOK, nil)
}
