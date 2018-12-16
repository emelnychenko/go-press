package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

func BindPostController(e *echo.Echo, c *PostController) {
	e.GET("/v0/posts", c.ListPosts)
	e.GET(fmt.Sprintf("/v0/post/:%s", PostId), c.GetPost)

	e.POST("/v0/posts", c.CreatePost)
	e.POST(fmt.Sprintf("/v0/post/:%s", PostId), c.UpdatePost)

	e.DELETE(fmt.Sprintf("/v0/post/:%s", PostId), c.DeletePost)
}
