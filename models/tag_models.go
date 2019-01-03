package models

import (
	"github.com/emelnychenko/go-press/common"
	"time"
)

type (
	TagId = common.ModelId

	Tag struct {
		Id      *TagId     `json:"id" xml:"id"`
		Name    string     `json:"name" xml:"name"`
		Created *time.Time `json:"created" xml:"created"`
		Updated *time.Time `json:"updated" xml:"updated"`
	}

	TagCreate struct {
		Name string `json:"name" xml:"name"`
	}

	TagUpdate struct {
		Name string `json:"name" xml:"name"`
	}

	TagPaginationQuery struct {
		*PaginationQuery
	}
)
