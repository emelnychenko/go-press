package models

import (
	"github.com/emelnychenko/go-press/common"
	"time"
)

type (
	CommentId = common.ModelId

	Comment struct {
		Id      *CommentId `json:"id" xml:"id"`
		Content string     `json:"content" xml:"content"`
		Created *time.Time `json:"created" xml:"created"`
		Updated *time.Time `json:"updated" xml:"updated"`
	}

	CommentCreate struct {
		Content string `json:"content" xml:"content"`
	}

	CommentUpdate struct {
		Content string `json:"content" xml:"content"`
	}

	CommentPaginationQuery struct {
		*PaginationQuery
	}
)
