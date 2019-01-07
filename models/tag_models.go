package models

import (
	"time"
)

type (
	TagId = ModelId

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
